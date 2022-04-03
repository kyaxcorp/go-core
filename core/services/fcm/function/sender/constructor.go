package sender

import "github.com/KyaXTeam/go-core/v2/core/clients/db"

func New(sender *Sender) (*Sender, error) {
	sender.stopSender = make(chan bool)
	// Check the Database if everything is ok with the tables!
	if sender.Name == "" {
		// TODO: raise an error!
		return sender, ErrConstructorNameEmpty
	}

	_db, err := db.GetDefaultClient()
	// Create necessary tables!
	if err != nil {
		return sender, ErrConstructorDBConnect
	}

	_db.AutoMigrate(
		&messageQueue{},
		&recipients{},
		&inTimeStatistics{},
	)

	return sender, nil
}
