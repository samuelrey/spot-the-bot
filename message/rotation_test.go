package message

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RotationSuite struct{ suite.Suite }

func (suite RotationSuite) TestJoin() {
	q := Rotation{}
	mu := User{ID: "amethyst#4422", Username: "amethyst"}

	err := q.Join(mu)

	suite.Require().Nil(err)
	suite.Require().Equal(1, q.Length())
	suite.Require().True(q.contains(mu))

	err = q.Join(mu)

	suite.Require().NotNil(err)
	suite.Require().Equal(1, q.Length())
	suite.Require().True(q.contains(mu))
}

func (suite RotationSuite) TestRotate() {
	q := Rotation{}
	m1 := User{ID: "amethyst#4422", Username: "amethyst"}
	m2 := User{ID: "osh#1219", Username: "osh"}

	user, err := q.Rotate()

	suite.Require().Nil(user)
	suite.Require().NotNil(err)

	_ = q.Join(m1)
	_ = q.Join(m2)
	user, err = q.Rotate()

	suite.Require().Nil(err)
	suite.Require().NotNil(user)
	suite.Require().Equal(m2, *user)
}

func (suite RotationSuite) TestLeave() {
	q := Rotation{}
	mu := User{ID: "amethyst#4422", Username: "amethyst"}

	succeed := q.Leave(mu)

	suite.Require().False(succeed)
	suite.Require().Equal(0, q.Length())
	suite.Require().False(q.contains(mu))

	_ = q.Join(mu)
	succeed = q.Leave(mu)

	suite.Require().True(succeed)
	suite.Require().Equal(0, q.Length())
	suite.Require().False(q.contains(mu))
}

func TestRotation(t *testing.T) {
	suite.Run(t, new(RotationSuite))
}
