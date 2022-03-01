package routingkey

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rooteiffelevents "github.com/eiffel-community/eiffelevents-sdk-go"
	eiffelevents "github.com/eiffel-community/eiffelevents-sdk-go/editions/lyon"
)

func ExampleSepia() {
	event, _ := eiffelevents.NewCompositionDefined()
	fmt.Println(Sepia(event, "", "random-tag"))
	// Output: eiffel._.EiffelCompositionDefinedEvent.random-tag._
}

func TestSepia(t *testing.T) {
	testcases := []struct {
		name     string
		domainID string
		family   string
		tag      string
		expected string
	}{
		{
			"Minimal example with no extras",
			"",
			"",
			"",
			"eiffel._.EiffelCompositionDefinedEvent._._",
		},
		{
			"With domain ID",
			"example",
			"",
			"",
			"eiffel._.EiffelCompositionDefinedEvent._.example",
		},
		{
			"With family",
			"",
			"example",
			"",
			"eiffel.example.EiffelCompositionDefinedEvent._._",
		},
		{
			"With tag",
			"",
			"",
			"example",
			"eiffel._.EiffelCompositionDefinedEvent.example._",
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := eiffelevents.NewCompositionDefined(rooteiffelevents.WithSourceDomainID(tc.domainID))
			require.NoError(t, err)
			assert.Equal(t, tc.expected, Sepia(event, tc.family, tc.tag))
		})
	}
}
