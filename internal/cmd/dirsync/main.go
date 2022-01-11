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
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	programName := filepath.Base(os.Args[0])

	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s SRC_DIR DEST_DIR", programName)
	}

	src := os.Args[1]
	dest := os.Args[2]
	// As a safety measure since we're allowing files to be deleted
	// the destination directory must be relative.
	if filepath.IsAbs(dest) {
		log.Fatalf("%s: destination directory path must be relative: %s", programName, dest)
	}

	gitCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(gitCtx, "rsync", "-a", "--delete", src, dest)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("%s: error running %q to list available schemas: %s", programName, cmd.String(), err)
	}
}
