package framework

type UserQueue interface {
	Head() *MessageUser
	Length() int
	Pop() *MessageUser
	Push(mu MessageUser)
	Remove(mu MessageUser)
}
