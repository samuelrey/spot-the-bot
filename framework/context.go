package framework

type Context struct {
	Messager
	EnrolledUsers *[]User
	Actor         User
	ChannelID     string
}
