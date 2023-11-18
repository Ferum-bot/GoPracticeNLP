package nlp

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
	"testing"
)

const tomlFileName = "testdata/tokenize_cases.toml"

type TomlCases struct {
	Cases []TomlCase
}

type TomlCase struct {
	Text   string
	Tokens []string
}

var testCases = []struct {
	inputText      string
	expectedOutput []string
}{
	{
		inputText:      "Who's on the first?",
		expectedOutput: []string{"who", "on", "the", "first"},
	},
	{
		inputText:      "",
		expectedOutput: nil,
	},
}

func TestTokenizeTable(t *testing.T) {
	for number, testCase := range testCases {
		testName := fmt.Sprintf("Test(%d)", number)

		t.Run(testName, func(t *testing.T) {
			actualOutput := Tokenize(testCase.inputText)

			require.Equal(t, testCase.expectedOutput, actualOutput)
		})
	}
}

func TestTokenize(t *testing.T) {
	text := "Who's on the first?"

	actual := Tokenize(text)
	expected := []string{"who", "on", "the", "first"}

	require.Equal(t, expected, actual)
}

func TestTokenizeTomlTable(t *testing.T) {
	var testCases TomlCases
	_, err := toml.DecodeFile(tomlFileName, &testCases)
	require.NoError(t, err)

	for testNumber, testCase := range testCases.Cases {
		testName := fmt.Sprintf("TomlTest(%d)", testNumber)
		actualOutput := Tokenize(testCase.Text)

		t.Run(testName, func(t *testing.T) {
			require.Equal(t, testCase.Tokens, actualOutput)
		})
	}
}
