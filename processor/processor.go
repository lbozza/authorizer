package processor

import (
	"authorizer/entity"
	"encoding/json"
	"fmt"
	"io"
)

type Input struct {
	Account     *entity.Account     `json:"account"`
	Transaction *entity.Transaction `json:"transaction"`
}

type Output struct {
	Account    *entity.Account `json:"account"`
	Violations []string        `json:"violations"`
}

//go:generate go run github.com/golang/mock/mockgen -package=processor -self_package=processor -destination=./handler_mock.go . Handler
type Handler interface {
	Handle(Input) (Output, error)
}

type Processor struct {
	stdin   *json.Decoder
	stdout  *json.Encoder
	handler Handler
}

func NewProcessor(stdin io.Reader, stdout io.Writer, handler Handler) *Processor {
	return &Processor{
		stdin:   json.NewDecoder(stdin),
		stdout:  json.NewEncoder(stdout),
		handler: handler,
	}
}

func (p *Processor) Process() error {

	for {
		var input Input
		if err := p.stdin.Decode(&input); err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error while reading the input transaction file: %w", err)
		}

		operation, err := p.handler.Handle(input)
		if err != nil {
			return fmt.Errorf("error while handling the transaction operation: %w", err)
		}

		if err := p.stdout.Encode(operation); err != nil {
			return fmt.Errorf("error while trying to write the output response :%w", err)
		}
	}

	return nil
}
