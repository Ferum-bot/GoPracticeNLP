package nlp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTokenize(t *testing.T) {
	text := "Who's on the first?"

	actual := Tokenize(text)
	expected := []string{"who", "s", "on", "the", "first"}

	require.Equal(t, expected, actual)
}
