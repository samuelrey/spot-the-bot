package framework

type UserQueue interface {
	Head() *MessageUser
	Pop() *MessageUser
	Push(mu MessageUser) UserQueue
	Remove(mu MessageUser) UserQueue
	String() string
}
