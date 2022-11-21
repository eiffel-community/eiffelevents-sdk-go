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
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/semver"
	"github.com/tidwall/gjson"

	"github.com/eiffel-community/eiffelevents-sdk-go"
	"github.com/eiffel-community/eiffelevents-sdk-go/internal/codetemplate"
)

type tableEntry struct {
	Filename   string
	StructName string
}

//go:embed templates/roundtriptable.tmpl
var tableFileTemplate string

func generateExampleTable(exampleDir string, output *codetemplate.OutputFile) error {
	filenames, err := filepath.Glob(filepath.Join(exampleDir, "*/*.json"))
	if err != nil {
		return err
	}

	var table []tableEntry
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		eventExample, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		// Extract the event type and version from the JSON payload.
		schemaSelectors := gjson.GetManyBytes(eventExample, "meta.type", "meta.version")
		metaType := schemaSelectors[0].String()
		if metaType == "" {
			return fmt.Errorf("unable to extract meta.type from input file %s", filename)
		}
		metaVersion := schemaSelectors[1].String()
		if metaVersion == "" {
			return fmt.Errorf("unable to extract meta.version from input file %s", filename)
		}
		v, err := semver.NewVersion(metaVersion)
		if err != nil {
			return fmt.Errorf("error parsing %q version: %w", metaVersion, err)
		}
		table = append(table, tableEntry{filename, eiffelevents.VersionedStructName(metaType, v)})
	}

	return output.ExpandTemplate(tableFileTemplate, table, template.FuncMap{})
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s EXAMPLE_DIR OUTPUT_FILE", filepath.Base(os.Args[0]))
	}
	output := codetemplate.New(os.Args[2])
	if err := generateExampleTable(os.Args[1], output); err != nil {
		log.Fatalf("Error generating code: %s", err)
	}
	if err := output.Close(); err != nil {
		log.Fatalf("Error writing final output file: %s", err)
	}
}
