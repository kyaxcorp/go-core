package gor

func (g *GInstance) Stop() error {
	// Check if Run function is not running...
	// Stop can run only after Run has finished!
	if g.isRunFuncRunning.Get() {
		return ErrRunFunctionAlreadyRunning
	}

	// Check if it's running Stop function
	if g.isStopFuncRunning.IfFalseSetTrue() {
		return ErrStopFunctionAlreadyRunning
	}

	defer func() {
		// Set that is not running anymore!
		g.isStopFuncRunning.False()
	}()

	// Call the cancel context
	if g.IsRunning() && g.ctx != nil {
		g.ctx.Cancel()
	}

	return nil
}
