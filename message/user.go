package message

type User struct {
	ID       string
	Username string
}

func (u User) String() string {
	return u.Username
}
