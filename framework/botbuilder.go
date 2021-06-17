package framework

type BotBuilder interface {
	Open() error
	Close() error
}
