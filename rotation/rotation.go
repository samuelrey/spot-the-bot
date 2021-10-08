package rotation

import (
	"errors"
	"fmt"

	"github.com/samuelrey/spot-the-bot/message"
)

type Rotation struct {
	queue []message.User
}

func NewRotation(users []message.User) Rotation {
	return Rotation{queue: users}
}

func (s *Rotation) Join(mu message.User) error {
	if s.contains(mu) {
		return errors.New("duplicate user")
	}

	s.push(mu)
	return nil
}

func (s *Rotation) Rotate() (*message.User, error) {
	if s.Length() == 0 {
		return nil, errors.New("rotation is empty")
	}

	head := s.pop()
	s.push(*head)
	return s.Head(), nil
}

func (s *Rotation) Leave(mu message.User) bool {
	found := -1
	for i, user := range s.queue {
		if mu.ID == user.ID {
			found = i
			break
		}
	}

	if found != -1 {
		s.queue = append(s.queue[:found], s.queue[found+1:]...)
		return true
	}

	return false
}

func (s *Rotation) contains(mu message.User) bool {
	for _, u := range s.queue {
		if mu.ID == u.ID {
			return true
		}
	}

	return false
}

func (s *Rotation) Head() *message.User {
	if len(s.queue) == 0 {
		return nil
	}

	return &s.queue[0]
}

func (s *Rotation) Length() int {
	return len(s.queue)
}

func (s *Rotation) pop() *message.User {
	if len(s.queue) == 0 {
		return nil
	}

	out := s.queue[0]
	s.queue = s.queue[1:]
	return &out
}

func (s *Rotation) push(mu message.User) {
	s.queue = append(s.queue, mu)
}

func (s Rotation) String() string {
	return fmt.Sprintf("%s", s.queue)
}
