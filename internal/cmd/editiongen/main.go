// Copyright Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"text/template"
	"time"

	"github.com/Masterminds/semver"

	"github.com/eiffel-community/eiffelevents-sdk-go"
	"github.com/eiffel-community/eiffelevents-sdk-go/internal/codetemplate"
)

// editionTags maps names of edition-specific Go packages to the corresponding
// tags in the protocol git.
var editionTags = map[string]string{
	"agen":     "edition-agen",
	"agen1":    "edition-agen-1",
	"arica":    "edition-arica",
	"bordeaux": "edition-bordeaux",
	"lyon":     "edition-lyon",
	"orizaba":  "edition-orizaba",
	"paris":    "edition-paris",
	"toulouse": "edition-toulouse",
}

//go:embed templates/eventfile.tmpl
var eventFileTemplate string

// createEditionDefinitions creates the Go file(s) for a particular Eiffel
// edition given a map containing the edition's events and their versions.
func createEditionDefinitions(packageName string, outputRootDir string, eventVersions map[string]*semver.Version) error {
	// The map we get as input isn't entirely suitable for text/template,
	// so transform it into a sorted list of structs that's convenient
	// to use from the template.
	type eventTypeInfo struct {
		EventType           string
		Version             *semver.Version
		StructName          string
		VersionedStructName string
	}
	var events []eventTypeInfo
	for eventType, version := range eventVersions {
		events = append(events, eventTypeInfo{
			EventType:           eventType,
			Version:             version,
			StructName:          eiffelevents.StructName(eventType, version),
			VersionedStructName: eiffelevents.VersionedStructName(eventType, version),
		})
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].EventType < events[j].EventType
	})

	outputDir := filepath.Join(outputRootDir, packageName)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	ct := codetemplate.New(filepath.Join(outputDir, "events.go"))
	if err := ct.ExpandTemplate(eventFileTemplate, events, template.FuncMap{}); err != nil {
		return err
	}
	return ct.Close()
}

// getLatestEvents scans an Eiffel protocol Git repository for event schemas
// at the commit pointed to by the given tag and returns a map with the most
// recent version of each encountered event.
func getLatestEvents(ctx context.Context, dir string, tagName string) (map[string]*semver.Version, error) {
	gitCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(gitCtx, "git", "ls-tree", "-r", "--name-only", tagName, "schemas")
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running %q to list available schemas: %w", cmd.String(), err)
	}

	latestEventVersions := make(map[string]*semver.Version)
	schemaFileRegexp := regexp.MustCompile(`^schemas/(Eiffel[^/]+Event)/([^/]+)\.json$`)
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		// Listing all files in the git and matching their paths against
		// a regexp obviously isn't very efficient, but it's fast enough
		// for this use case.
		matches := schemaFileRegexp.FindStringSubmatch(scanner.Text())
		if matches == nil {
			continue
		}
		eventType := matches[1]
		versionString := matches[2]
		version, err := semver.NewVersion(versionString)
		if err != nil {
			return nil, fmt.Errorf("%s: Error parsing version %q: %s", scanner.Text(), versionString, err)
		}
		if currentLatest, exists := latestEventVersions[eventType]; !exists || version.GreaterThan(currentLatest) {
			latestEventVersions[eventType] = version
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading output from %q: %w", cmd.String(), err)
	}
	return latestEventVersions, nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s SCHEMA_REPO ROOT_OUTPUT_DIR", filepath.Base(os.Args[0]))
	}

	for editionName, editionTag := range editionTags {
		latestEventVersions, err := getLatestEvents(context.Background(), os.Args[1], editionTag)
		if err != nil {
			log.Fatalf("%s: %s", filepath.Base(os.Args[0]), err)
		}

		if err = createEditionDefinitions(editionName, filepath.Join(os.Args[2], "editions"), latestEventVersions); err != nil {
			log.Fatalf("%s: %s", filepath.Base(os.Args[0]), err)
		}
	}
}
