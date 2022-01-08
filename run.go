package sentry

// Run starts a sentry instance
func (instance *SentryInstance) Run() chan error {
	err := make(chan error, 1)
	go func() {
		err <- instance.server.ListenAndServe()
	}()
	return err
}
