package framework

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CommandRegistrySuite struct{ suite.Suite }

func (suite CommandRegistrySuite) TestFormatHelpMessageNoCommand() {
	cr := NewCommandRegistry()

	helpMsg := cr.FormatHelpMessage()

	suite.Require().Empty(helpMsg)
}

func (suite CommandRegistrySuite) TestFormatHelpMessageWithCommand() {
	cr := NewCommandRegistry()
	cr.Register("rocktheboat", nil, "don't rock the boat baby")

	helpMsg := cr.FormatHelpMessage()

	suite.Require().NotEmpty(helpMsg)
}

func TestCommandRegistry(t *testing.T) {
	suite.Run(t, new(CommandRegistrySuite))
}