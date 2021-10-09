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
	actual := q.Current()
	suite.Require().Nil(actual)

	actual = q.pop()
	suite.Require().Nil(actual)

	// nonempty queue
	au := message.User{ID: "amethyst#4422", Username: "amethyst"}
	ou := message.User{ID: "osh#1219", Username: "osh"}
	q = NewRotation([]message.User{au, ou})

	actual = q.Current()
	suite.Require().Equal(&au, actual)

	actual = q.pop()
	suite.Require().Equal(&au, actual)
	suite.Require().Equal(1, q.Length())

	actual = q.Current()
	suite.Require().Equal(&ou, actual)
}

func (suite RotationSuite) TestPushRemove() {
	q := Rotation{}
	mu := message.User{ID: "amethyst#4422", Username: "amethyst"}

	q.push(mu)
	actual := q.Current()
	suite.Require().Equal(&mu, actual)

	q.Leave(mu)
	actual = q.Current()
	suite.Require().Nil(actual)
}

func (suite RotationSuite) TestJoin() {
	q := Rotation{}
	mu := message.User{ID: "amethyst#4422", Username: "amethyst"}

	err := q.Join(mu)

	suite.Require().Nil(err)
	suite.Require().Equal(1, q.Length())
	suite.Require().Equal(true, q.contains(mu))

	err = q.Join(mu)

	suite.Require().NotNil(err)
	suite.Require().Equal(1, q.Length())
	suite.Require().Equal(true, q.contains(mu))
}

func TestRotation(t *testing.T) {
	suite.Run(t, new(RotationSuite))
}
