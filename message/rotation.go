package message

import (
	"errors"
	"time"
)

type Rotation struct {
	ServerID  string    `bson:"server_id"`
	Queue     []User    `bson:"queue"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func NewRotation(users []User, serverID string) Rotation {
	return Rotation{Queue: users, ServerID: serverID}
}

func (s *Rotation) Join(mu User) error {
	if s.contains(mu) {
		return errors.New("duplicate user")
	}

	s.push(mu)
	return nil
}

func (s *Rotation) Rotate() (*User, error) {
	if s.Length() == 0 {
		return nil, errors.New("rotation is empty")
	}

	head := s.pop()
	s.push(*head)
	return s.Current(), nil
}

func (s *Rotation) Leave(mu User) bool {
	i := s.indexOf(mu)
	if i != -1 {
		s.Queue = append(s.Queue[:i], s.Queue[i+1:]...)
		return true
	}

	return false
}

func (s *Rotation) Current() *User {
	if len(s.Queue) == 0 {
		return nil
	}

	return &s.Queue[0]
}

func (s *Rotation) Length() int {
	return len(s.Queue)
}

func (s *Rotation) contains(mu User) bool {
	for _, user := range s.Queue {
		if mu.ID == user.ID {
			return true
		}
	}

	return false
}

func (s *Rotation) indexOf(mu User) int {
	for i, user := range s.Queue {
		if mu.ID == user.ID {
			return i
		}
	}

	return -1
}

func (s *Rotation) pop() *User {
	if len(s.Queue) == 0 {
		return nil
	}

	out := s.Queue[0]
	s.Queue = s.Queue[1:]
	return &out
}

func (s *Rotation) push(mu User) {
	s.Queue = append(s.Queue, mu)
}
