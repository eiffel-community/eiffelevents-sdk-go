package signature

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthorIdentity_UnmarshalText(t *testing.T) {
	var data struct {
		Identity *AuthorIdentity `json:"identity"`
	}
	require.NoError(t, json.Unmarshal([]byte(`{"identity": "CN=foo,DC=example,DC=com"}`), &data))
	assert.Equal(t, "CN=foo,DC=example,DC=com", data.Identity.String())
}
