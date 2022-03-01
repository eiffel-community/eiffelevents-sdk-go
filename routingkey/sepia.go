// Package routingkey contains functions for generating AMQP routing keys (topics) based on Eiffel events.
package routingkey

import (
	"fmt"

	"github.com/eiffel-community/eiffelevents-sdk-go"
)

// Sepia returns the AMQP routing key that the Sepia standard recommends
// (https://eiffel-community.github.io/eiffel-sepia/rabbitmq-message-broker.html).
// The family and tag strings are optional and may be empty. If they're empty
// a short replacement string will be used in their place.
func Sepia(event eiffelevents.MetaTeller, family string, tag string) string {
	return fmt.Sprintf("eiffel.%s.%s.%s.%s", valueOrDefault(family), event.Type(), valueOrDefault(tag), valueOrDefault(event.DomainID()))
}

func valueOrDefault(value string) string {
	if value != "" {
		return value
	}
	return "_"
}
