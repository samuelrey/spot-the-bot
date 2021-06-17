package framework

type MessageBuilder interface {
	ReplyBuilder
	DirectMessageBuilder
}

// TODO use only one message builder interface
type V2MessageBuilder interface {
	V2ReplyBuilder
	DirectMessageBuilder
}

type V2ReplyBuilder interface {
	Reply(channelID, content string) error
}

type ReplyBuilder interface {
	Reply(content string) error
}

type DirectMessageBuilder interface {
	DirectMessage(recipientID, content string) error
}

type BotBuilder interface {
	Open() error
	Close() error
}
