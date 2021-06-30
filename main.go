package main

import (
	"authorizer/processor"
	"authorizer/usecase/authorizer"
	"fmt"
	"os"
)

func main() {

	in := os.Stdin
	out := os.Stdout

	processor := processor.NewProcessor(in, out, authorizer.NewAuthorizeHandler())
	err := processor.Process()

	if err != nil {
		fmt.Printf(err.Error())
	}
}
