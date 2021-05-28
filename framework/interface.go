package framework

type Replyer interface {
	Reply(content string) error
}

func Reply(replyer Replyer, content string) error {
	return replyer.Reply(content)
}
