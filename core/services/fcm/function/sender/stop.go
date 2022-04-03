package sender

func (s *Sender) Stop() {
	s.stopSender <- true
}
