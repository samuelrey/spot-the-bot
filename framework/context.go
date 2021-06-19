package framework

type Context struct {
	Messager
	EnrolledUsers *[]User
	Actor         User
	ChannelID     string
}

type User struct {
	ID       string
	Username string
}

func (u User) String() string {
	return u.Username
}
