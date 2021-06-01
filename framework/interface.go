package framework

type Messager interface {
	Replyer
	DirectMessager
}

type Replyer interface {
	Reply(content string) error
}

func Reply(replyer Replyer, content string) error {
	return replyer.Reply(content)
}

type DirectMessager interface {
	DirectMessage(recipient, content string) error
}

func DirectMessage(
	directMessager DirectMessager,
	recipient string,
	content string,
) error {
	return directMessager.DirectMessage(recipient, content)
}
