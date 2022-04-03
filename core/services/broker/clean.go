package broker

func (b *Broker) Clean() {
	b.createPipes() // It's recreating the pipes...meaning that all existing will be deleted
}
