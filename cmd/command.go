package cmd

type command struct {
	cmd     func(*Context)
	helpMsg string
}

func (c command) RunWithContext(ctx *Context) {
	c.cmd(ctx)
}
