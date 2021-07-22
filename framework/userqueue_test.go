package framework

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserQueueSuite struct{ suite.Suite }

func (suite UserQueueSuite) TestHeadPop() {
	q := UserQueue{}

	// empty queue
	actual := q.Head()
	suite.Require().Nil(actual)

	actual = q.Pop()
	suite.Require().Nil(actual)

	// nonempty queue
	au := MessageUser{ID: "amethyst#4422", Username: "amethyst"}
	ou := MessageUser{ID: "osh#1219", Username: "osh"}
	q = NewUserQueue([]MessageUser{au, ou})

	actual = q.Head()
	suite.Require().Equal(&au, actual)

	actual = q.Pop()
	suite.Require().Equal(&au, actual)
	suite.Require().Equal(1, q.Length())

	actual = q.Head()
	suite.Require().Equal(&ou, actual)
}

func (suite UserQueueSuite) TestPushRemove() {
	q := UserQueue{}
	mu := MessageUser{ID: "amethyst#4422", Username: "amethyst"}

	q.Push(mu)
	actual := q.Head()
	suite.Require().Equal(&mu, actual)

	q.Remove(mu)
	actual = q.Head()
	suite.Require().Nil(actual)
}

func TestUserQueue(t *testing.T) {
	suite.Run(t, new(UserQueueSuite))
}
