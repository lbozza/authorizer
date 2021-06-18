package main

import (
	"authorizer/processor"
	"authorizer/usecase/authorizer"
	"os"
)

func main() {

	in := os.Stdin
	out := os.Stdout

	processor := processor.NewProcessor(in, out, authorizer.NewAuthorizeHandler())
	processor.Process()
}
