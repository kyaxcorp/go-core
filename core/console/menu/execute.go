package menu

// Execute executes the root Command.
func (m *Menu) Execute() error {
	m.init()
	//return m.RootCmd.Execute()
	// We are using cancel context here!
	return m.RootCmd.ExecuteContext(m.ctx.Context())
}
