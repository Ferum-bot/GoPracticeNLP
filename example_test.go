package nlp_test

import (
	"fmt"
	nlp "github.com/Ferum-Bot/GoPracticeNLP"
)

func ExampleTokenize() {
	text := "Who is on first?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	//Output:
	//[who i on first]
}
