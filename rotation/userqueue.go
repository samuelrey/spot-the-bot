package rotation

import (
	"fmt"

	"github.com/samuelrey/spot-the-bot/message"
)

type UserQueue struct {
	queue []message.User
}

func NewUserQueue(users []message.User) UserQueue {
	return UserQueue{queue: users}
}

func (s *UserQueue) Contains(mu message.User) bool {
	for _, u := range s.queue {
		if mu.ID == u.ID {
			return true
		}
	}

	return false
}

func (s *UserQueue) Head() *message.User {
	if len(s.queue) == 0 {
		return nil
	}

	return &s.queue[0]
}

func (s *UserQueue) Length() int {
	return len(s.queue)
}

func (s *UserQueue) Pop() *message.User {
	if len(s.queue) == 0 {
		return nil
	}

	out := s.queue[0]
	s.queue = s.queue[1:]
	return &out
}

func (s *UserQueue) Push(mu message.User) {
	s.queue = append(s.queue, mu)
}

func (s *UserQueue) Remove(mu message.User) bool {
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

func (s UserQueue) String() string {
	return fmt.Sprintf("%s", s.queue)
}
