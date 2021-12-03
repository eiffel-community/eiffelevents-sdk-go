package eiffelevents

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleModifier() {
	// Define a factory function that encapsulates a WithDomainID modifier
	compositionFactory := func() (*CompositionDefinedV3, error) {
		return NewCompositionDefinedV3(WithSourceDomainID("example.com"))
	}

	// Use the newly defined factory to create a customized new event
	myComposition, _ := compositionFactory()
	fmt.Println(myComposition.Meta.Source.DomainID)

	// Output: example.com
}

func TestWithSourceDomainID(t *testing.T) {
	event, err := NewCompositionDefinedV3(WithSourceDomainID("example.com"))
	assert.NoError(t, err)
	assert.Equal(t, "example.com", event.Meta.Source.DomainID)
}

func TestWithSourceHost(t *testing.T) {
	event, err := NewCompositionDefinedV3(WithSourceHost("hostname.example.com"))
	assert.NoError(t, err)
	assert.Equal(t, "hostname.example.com", event.Meta.Source.Host)
}

func TestWithSourceName(t *testing.T) {
	event, err := NewCompositionDefinedV3(WithSourceName("My Application"))
	assert.NoError(t, err)
	assert.Equal(t, "My Application", event.Meta.Source.Name)
}

func TestWithSourceURI(t *testing.T) {
	event, err := NewCompositionDefinedV3(WithSourceURI("http://www.example.com"))
	assert.NoError(t, err)
	assert.Equal(t, "http://www.example.com", event.Meta.Source.URI)
}

func TestWithVersion(t *testing.T) {
	newestEventVersion, err := NewCompositionDefinedV3()
	assert.NoError(t, err)
	customEventVersion, err := NewCompositionDefinedV3(WithVersion("3.1.0"))
	assert.NoError(t, err)

	// Make sure we're not accidentally getting a pass because the default
	// version happens to coincide with the one we're request (which shouldn't
	// happen unless there's a bug elsewhere).
	assert.NotEqual(t, "3.1.0", newestEventVersion.Meta.Version)
	assert.Equal(t, "3.1.0", customEventVersion.Meta.Version)
}
