package cmd

import (
	"fmt"

	"github.com/samuelrey/spot-the-bot/framework"
)

type SimpleUserQueue struct {
	queue []framework.MessageUser
}

func (s *SimpleUserQueue) Head() *framework.MessageUser {
	if len(s.queue) == 0 {
		return nil
	}

	return &s.queue[0]
}

func (s *SimpleUserQueue) Length() int {
	return len(s.queue)
}

func (s *SimpleUserQueue) Pop() *framework.MessageUser {
	if len(s.queue) == 0 {
		return nil
	}

	out := s.queue[0]
	s.queue = s.queue[1:]
	return &out
}

func (s *SimpleUserQueue) Push(mu framework.MessageUser) {
	s.queue = append(s.queue, mu)
}

func (s *SimpleUserQueue) Remove(mu framework.MessageUser) {
	found := -1
	for i, user := range s.queue {
		if mu.ID == user.ID {
			found = i
			break
		}
	}

	if found != -1 {
		s.queue = append(s.queue[:found], s.queue[found+1:]...)
	}
}

func (s SimpleUserQueue) String() string {
	return fmt.Sprintf("%s", s.queue)
}
