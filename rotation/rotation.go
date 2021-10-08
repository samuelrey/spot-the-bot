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
	return s.Current(), nil
}

func (s *Rotation) Leave(mu message.User) bool {
	i := s.indexOf(mu)
	if i != -1 {
		s.queue = append(s.queue[:i], s.queue[i+1:]...)
		return true
	}

	return false
}

func (s *Rotation) contains(mu message.User) bool {
	for _, user := range s.queue {
		if mu.ID == user.ID {
			return true
		}
	}

	return false
}

func (s *Rotation) indexOf(mu message.User) int {
	for i, user := range s.queue {
		if mu.ID == user.ID {
			return i
		}
	}

	return -1
}

func (s *Rotation) Current() *message.User {
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
