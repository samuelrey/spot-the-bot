package framework

type MessageBuilder interface {
	ReplyBuilder
	DirectMessageBuilder
}

type ReplyBuilder interface {
	Reply(content string) error
}

type DirectMessageBuilder interface {
	DirectMessage(recipientID, content string) error
}
