package framework

type MessageConnecter interface {
	Open() error
	Close() error
}

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
