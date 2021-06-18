package processor

import (
	"authorizer/entity"
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ProcessorSuite struct {
	suite.Suite
	*require.Assertions

	ctrl        *gomock.Controller
	handlerMock *MockHandler
}

func TestProcessorSuite(t *testing.T) {
	suite.Run(t, new(ProcessorSuite))
}

func (s *ProcessorSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())
	s.handlerMock = NewMockHandler(s.ctrl)
}

func (s *ProcessorSuite) ShutDown() {
	s.ctrl.Finish()
}

func (s *ProcessorSuite) TestSuccessProcessor() {
	inJson := `{"account": {"active-card": true, "available-limit": 100}}
				{"transaction": {"merchant": "Burger King", "amount": 20, "time": "2019-02-13T10:00:00.000Z"}}`

	s.handlerMock.EXPECT().Handle(gomock.Any()).Return(Output{
		Account: &entity.Account{
			ActiveCard:     true,
			AvaliableLimit: 100,
		},
		Violations: []string{},
	}, nil).Return(Output{
		Account: &entity.Account{
			ActiveCard:     true,
			AvaliableLimit: 100,
		},
		Violations: []string{},
	}, nil).Times(2)

	processor := NewProcessor(bytes.NewBuffer([]byte(inJson)), bytes.NewBuffer(nil), s.handlerMock)
	err := processor.Process()

	s.NoError(err)

}

func (s *ProcessorSuite) TestProcessorEmptyInput() {

	processor := NewProcessor(bytes.NewBuffer(nil), bytes.NewBuffer(nil), s.handlerMock)
	err := processor.Process()

	s.NoError(err)
}

func (s *ProcessorSuite) TestProcessorInvalidInputDecodeError() {
	inJson := `error`

	processor := NewProcessor(bytes.NewBuffer([]byte(inJson)), bytes.NewBuffer(nil), s.handlerMock)
	err := processor.Process()

	s.Error(err, "error while reading the input transaction file")
}

func (s *ProcessorSuite) TestProcessorHandleError() {
	inJson := `{}`

	s.handlerMock.EXPECT().Handle(gomock.Any()).Return(Output{}, errors.New("invalid operation type"))

	processor := NewProcessor(bytes.NewBuffer([]byte(inJson)), bytes.NewBuffer(nil), s.handlerMock)
	err := processor.Process()

	s.Error(err, "error while handling the transaction operation")
}
