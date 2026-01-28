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
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrUsage = errors.New("invalid command line usage")

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	subcommand := os.Args[1]

	var err error
	switch subcommand {
	case "sign":
		err = signCmd(os.Args[2:], os.Stdin, os.Stdout)
	case "verify":
		err = verifyCmd(context.Background(), os.Args[2:], os.Stdin)
	default:
		fmt.Fprintf(os.Stderr, "Unknown subcommand: %s\n", subcommand)
		fmt.Fprintln(os.Stderr)
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if errors.Is(err, ErrUsage) {
			fmt.Fprintln(os.Stderr)
			usage()
			os.Exit(2)
		}
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <command> <args>\n", filepath.Base(os.Args[0]))
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Commands:")
	fmt.Fprintln(os.Stderr, "  sign <private key PEM file> <author identity DN> <algorithm>")
	fmt.Fprintln(os.Stderr, "  verify <public key directory>")
}
