package message

type Messager interface {
	Replyer
	DirectMessager
}

type Replyer interface {
	Reply(content string) error
}

type DirectMessager interface {
	DirectMessage(recipientID, content string) error
}

type User struct {
	ID       string `bson:"id"`
	Username string `bson:"username"`
}
