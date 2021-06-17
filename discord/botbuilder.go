package discord

func (d *DiscordBuilder) Open() error {
	return d.session.Open()
}

func (d *DiscordBuilder) Close() error {
	return d.session.Close()
}
