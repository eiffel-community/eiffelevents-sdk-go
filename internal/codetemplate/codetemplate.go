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

// Package codetemplate provides a small wrapper type that makes it more
// convenient to generate correctly formatted Go source files.
package codetemplate

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/google/renameio"
)

// OutputFile buffers file contents, either written directly via its io.Writer
// interface or via text/template expansion, and writes the data to the chosen
// destination file. The data is initially written to a temporary file and
// renamed into place to replace the destination file, resulting in an atomic
// replacement.
type OutputFile struct {
	Filename string
	buf      bytes.Buffer
}

func New(filename string) *OutputFile {
	return &OutputFile{Filename: filename}
}

// ExpandTemplate parses a text/template and executes it with the provided input data.
func (of *OutputFile) ExpandTemplate(text string, data interface{}, funcs template.FuncMap) error {
	templ, err := template.New("(template name unused)").Funcs(funcs).Parse(text)
	if err != nil {
		return fmt.Errorf("parse error: %w", err)
	}
	if err := templ.Execute(&of.buf, data); err != nil {
		return fmt.Errorf("template expansion error: %w", err)
	}
	return nil
}

// Close creates a temporary file, writes the buffered file contents to it, and
// renames it into place.
func (of *OutputFile) Close() error {
	output, err := renameio.TempFile("", of.Filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = output.Cleanup()
	}()

	formatted, err := format.Source(of.buf.Bytes())
	if err != nil {
		return fmt.Errorf("unable to format source code: %w", err)
	}
	if _, err := output.Write(formatted); err != nil {
		return err
	}
	return output.CloseAtomicallyReplace()
}

func (of *OutputFile) Write(b []byte) (int, error) {
	return of.buf.Write(b)
}
