package cmd

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RegistrySuite struct{ suite.Suite }

func (suite RegistrySuite) TestFormatHelpMessageNoCommand() {
	cr := NewRegistry()

	helpMsg := cr.FormatHelpMessage()

	suite.Require().Empty(helpMsg)
}

func (suite RegistrySuite) TestFormatHelpMessageWithCommand() {
	cr := NewRegistry()
	cr.Register("rocktheboat", nil, "don't rock the boat baby")

	helpMsg := cr.FormatHelpMessage()

	suite.Require().NotEmpty(helpMsg)
}

func TestCommandRegistry(t *testing.T) {
	suite.Run(t, new(RegistrySuite))
}
