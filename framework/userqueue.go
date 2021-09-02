package framework

import "fmt"

type UserQueue struct {
	queue []MessageUser
}

func NewUserQueue(users []MessageUser) UserQueue {
	return UserQueue{queue: users}
}

func (s *UserQueue) Contains(mu MessageUser) bool {
	for _, u := range s.queue {
		if mu.ID == u.ID {
			return true
		}
	}

	return false
}

func (s *UserQueue) Head() *MessageUser {
	if len(s.queue) == 0 {
		return nil
	}

	return &s.queue[0]
}

func (s *UserQueue) Length() int {
	return len(s.queue)
}

func (s *UserQueue) Pop() *MessageUser {
	if len(s.queue) == 0 {
		return nil
	}

	out := s.queue[0]
	s.queue = s.queue[1:]
	return &out
}

func (s *UserQueue) Push(mu MessageUser) {
	s.queue = append(s.queue, mu)
}

func (s *UserQueue) Remove(mu MessageUser) bool {
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
