package framework

import "fmt"

type UserQueue interface {
	Contains(mu MessageUser) bool
	Head() *MessageUser
	Length() int
	Pop() *MessageUser
	Push(mu MessageUser)
	Remove(mu MessageUser) bool
}

type SimpleUserQueue struct {
	queue []MessageUser
}

func NewSimpleUserQueue(users []MessageUser) SimpleUserQueue {
	return SimpleUserQueue{queue: users}
}

func (s *SimpleUserQueue) Contains(mu MessageUser) bool {
	for _, u := range s.queue {
		if mu.ID == u.ID {
			return true
		}
	}

	return false
}

func (s *SimpleUserQueue) Head() *MessageUser {
	if len(s.queue) == 0 {
		return nil
	}

	return &s.queue[0]
}

func (s *SimpleUserQueue) Length() int {
	return len(s.queue)
}

func (s *SimpleUserQueue) Pop() *MessageUser {
	if len(s.queue) == 0 {
		return nil
	}

	out := s.queue[0]
	s.queue = s.queue[1:]
	return &out
}

func (s *SimpleUserQueue) Push(mu MessageUser) {
	s.queue = append(s.queue, mu)
}

func (s *SimpleUserQueue) Remove(mu MessageUser) bool {
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

func (s SimpleUserQueue) String() string {
	return fmt.Sprintf("%s", s.queue)
}
