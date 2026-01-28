<!---
   Copyright Axis Communications AB
   For a full list of individual contributors, please see the commit history.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
--->

<img src="./images/logo.png" alt="Eiffel Events SDK – Go" width="350"/>

[![Go Reference](https://pkg.go.dev/badge/github.com/eiffel-community/eiffelevents-sdk-go.svg)](https://pkg.go.dev/github.com/eiffel-community/eiffelevents-sdk-go)
[![Sandbox badge](https://img.shields.io/badge/Stage-Sandbox-yellow)](https://github.com/eiffel-community/community/blob/master/PROJECT_LIFECYCLE.md#stage-sandbox)

# Eiffel Events SDK – Go
This repository contains data types, constants, and functions for working with Eiffel events in the [Go](https://golang.org/) language, including marshaling to and from JSON. Its goal is to make it easy to create and process Eiffel events in Go.

The module declares a Go struct for every major version of each event type.
These structs are generated from the JSON schemas and named as in the examples
below.

## Creating new events

The struct types used to represent Eiffel events are named after the event
types without the "Eiffel" prefix and "Event" suffix, and with a version suffix.
For non-experimental event versions (1.0.0 and up) the version suffix is the
event's major version (i.e. each major version gets its own struct) while for
experimental event versions (0.x.y) every single version gets its own struct
(because every version is allowed to contain backwards incompatible changes).

The following example shows two methods of creating events, with and without
a factory.

```go
package main

import (
	"fmt"
	"time"

	"github.com/eiffel-community/eiffelevents-sdk-go"
)

func main() {
	// Manual initialization of all struct members.
	var event1 eiffelevents.CompositionDefinedV3
	event1.Meta.Type = "EiffelCompositionDefinedEvent"
	event1.Meta.Version = "3.2.0"
	event1.Meta.ID = "87dac043-2e1b-41c5-833a-712833f2a613"
	event1.Meta.Time = time.Now().UnixMilli()
	event1.Data.Name = "my-composition"
	fmt.Println(event1.String())

	// Equivalent example using the factory that pre-populates all
	// required meta members (picking the most recent event version in
	// the chosen major version). Note that the factory returns
	// a struct pointer.
	event2, err := eiffelevents.NewCompositionDefinedV3()
	if err != nil {
		panic(err)
	}
	event2.Data.Name = "my-composition"
	fmt.Println(event2.String())
}
```

The example below shows how modifier functions can be passed to factories
to populate the newly created events with additional fields. In trivial
cases modifiers are superfluous and the caller can just set the desired fields
after obtaining the event from the factory, but apart from being a compact
representation modifiers can be used with any event type. You can also use them
to create custom factories that apply a preconfigured set of modifiers.

```go
package main

import (
	"fmt"

	"github.com/eiffel-community/eiffelevents-sdk-go"
)

func main() {
	// Create an event with modifiers that select a particular
	// version of the event and makes sure meta.source.host is
	// populated with the name of the current host.
	event1, err := eiffelevents.NewCompositionDefinedV3(
		eiffelevents.WithVersion("3.1.0"),
		eiffelevents.WithAutoSourceHost(),
	)
	if err != nil {
		panic(err)
	}
	event1.Data.Name = "my-composition"
	fmt.Println(event1.String())

	// Create a custom factory with the chosen modifiers.
	newComposition := func() (*eiffelevents.CompositionDefinedV3, error) {
		return eiffelevents.NewCompositionDefinedV3(
			eiffelevents.WithVersion("3.1.0"),
			eiffelevents.WithAutoSourceHost(),
		)
	}

	// Create a new event using the custom factory.
	event2, err := newComposition()
	if err != nil {
		panic(err)
	}
	event2.Data.Name = "my-composition"
	fmt.Println(event2.String())
}
```

## Preferring events from a particular Eiffel edition

Each Eiffel edition has a subpackage containing version-less struct type
aliases and factories for creating events with the correct version for
the chosen edition. This removes much of the need to litter the code with
"V3" etc suffixes.

```go
package main

import (
	"fmt"

	"github.com/eiffel-community/eiffelevents-sdk-go/editions/lyon"
)

func main() {
	event, err := eiffelevents.NewCompositionDefined()
	if err != nil {
		panic(err)
	}
	// The event struct has the type eiffelevents.CompositionDefined,
	// which is an alias for the parent package's CompositionDefinedV3.
	// The event version is set to 3.2.0. By instead importing the paris
	// subpackage the event version would've been set to 3.1.0.

	event.Data.Name = "my-composition"
	fmt.Println(event.String())
}
```

## Unmarshaling event JSON strings into Go structs

To unmarshal a JSON string into one of the structs defined in this package use
the UnmarshalAny function and use e.g. a type switch to access the event members:

```go
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/eiffel-community/eiffelevents-sdk-go"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	anyEvent, err := eiffelevents.UnmarshalAny(input)
	if err != nil {
		panic(err)
	}

	switch event := anyEvent.(type) {
	case *eiffelevents.CompositionDefinedV3:
		fmt.Printf("Received %s composition\n", event.Data.Name)
	default:
		fmt.Printf("This event I don't know much about: %s\n", event)
	}
}
```

If you have a compound JSON structure containing e.g. an array of event
objects you can declare its type to be []*eiffelevents.Any. After unmarshaling
the data you can use a type switch to process the events:

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/eiffel-community/eiffelevents-sdk-go"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var apiResponse struct {
		Events []*eiffelevents.Any `json:"events"`
	}

	err = json.Unmarshal(input, &apiResponse)
	if err != nil {
		panic(err)
	}

	for _, anyEvent := range apiResponse.Events {
		switch event := anyEvent.Get().(type) {
		case *eiffelevents.CompositionDefinedV3:
			fmt.Printf("Received %s composition\n", event.Data.Name)
		default:
			fmt.Printf("This event I don't know much about: %s\n", event)
		}
	}
}
```

## Validating events

Eiffel events are defined by their schemas, and publishers are expected to
send events that conform to those schemas. The validator subpackage can assist
with that task as well as other user-defined validation tasks. Validation is
done via a validator.Set instance, where one or more implementations of
validator.Validator inspect an event in the configured order. To ease the
configuration burden, validator.DefaultSet returns a reasonably configured
validator.Set instance that's ready to be used. See the documentation of
the validator subpackage for details.

## Signing events and verifying signatures

The SDK supports cryptographic signing of (typically) outbound events as
well as verification of the signature of inbound events. The signing is
done according to the standard method, with the signature and metadata under
the `meta.security` field. See the documentation of the signature subpackage
for details and code examples, but also the cmd/eiffelsignature
subpackage from which you can build a standalone CLI executable for
signing events and verifying the signatures of signed events.

## Code of Conduct and Contributing
To get involved, please see [Code of Conduct](https://github.com/eiffel-community/.github/blob/master/CODE_OF_CONDUCT.md) and [contribution guidelines](https://github.com/eiffel-community/.github/blob/master/CONTRIBUTING.md).

Note that these files are located in the .github repository. See [this](https://docs.github.com/en/github/building-a-strong-community/creating-a-default-community-health-file) page for further details regarding default community health files.

## About this repository
The contents of this repository are licensed under the [Apache License 2.0](./LICENSE).

## About Eiffel
This repository forms part of the Eiffel Community. Eiffel is a protocol for technology agnostic machine-to-machine communication in continuous integration and delivery pipelines, aimed at securing scalability, flexibility and traceability. Eiffel is based on the concept of decentralized real time messaging, both to drive the continuous integration and delivery system and to document it.

Visit [Eiffel Community](https://eiffel-community.github.io) to get started and get involved.
