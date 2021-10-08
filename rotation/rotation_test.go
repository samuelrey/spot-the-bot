package rotation

import (
	"testing"

	"github.com/samuelrey/spot-the-bot/message"
	"github.com/stretchr/testify/suite"
)

type RotationSuite struct{ suite.Suite }

func (suite RotationSuite) TestHeadPop() {
	q := Rotation{}

	// empty queue
	actual := q.Head()
	suite.Require().Nil(actual)

	actual = q.Pop()
	suite.Require().Nil(actual)

	// nonempty queue
	au := message.User{ID: "amethyst#4422", Username: "amethyst"}
	ou := message.User{ID: "osh#1219", Username: "osh"}
	q = NewRotation([]message.User{au, ou})

	actual = q.Head()
	suite.Require().Equal(&au, actual)

	actual = q.Pop()
	suite.Require().Equal(&au, actual)
	suite.Require().Equal(1, q.Length())

	actual = q.Head()
	suite.Require().Equal(&ou, actual)
}

func (suite RotationSuite) TestPushRemove() {
	q := Rotation{}
	mu := message.User{ID: "amethyst#4422", Username: "amethyst"}

	q.Push(mu)
	actual := q.Head()
	suite.Require().Equal(&mu, actual)

	q.Leave(mu)
	actual = q.Head()
	suite.Require().Nil(actual)
}

func TestRotation(t *testing.T) {
	suite.Run(t, new(RotationSuite))
}
