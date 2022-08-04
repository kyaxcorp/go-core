package sender

import "time"

type Sender struct {
	FcmSenderId  string
	FcmServerKey string
	Name         string // The name should be unique, because based on this, it will use specific DB Table Name!

	// TODO: also see Web Push Configuration
	// TODO: also see IOS Configuration

	stopSender chan bool
}

// TODO: the DB table name should be different than usual... it should have some kind of prefix
// TODO: the messages should be sent in a order but not randomly! This is important!
// The messages are being deleted after they are being sent!
type messageQueue struct {
	ID uint64 `gorm:"primaryKey"` // -> it's primary key! But when it gets to max value we can reset it... or if we don't have
	// anything in the DB, we also can reset it!
	Payload   string // JSONB -> this is the payload that we will be sending to the recipients
	CreatedAt uint64 `gorm:"autoCreateTime:nano"` // When it has being created!
	UpdatedAt uint64 `gorm:"autoUpdateTime:nano"` // When last time has being modified

	//
	LastRetry         time.Time `gorm:"index"` // This is the time when the process started to send the messages
	RetryAfterSeconds uint64    `gorm:"index"` // This is the time after when the Check should retry the sending?!
	LockedBy          uint64    `gorm:"index"` // By whom it's being locked (by what process/goroutine?)
	LockedWhen        time.Time `gorm:"index"` // When it's being locked
	LockedTTL         uint64    `gorm:"index"` // For how long it's being locked?!
}

// The recipients are being deleted after the
// Maybe in time we will delete the unnecessary data..!?
// If the token is invalid, the Recipient will be deleted from here!
// Also a callback should be called if there is an invalid recipient!
type recipients struct {
	MessageId          uint64    `gorm:"primaryKey"` // primary key
	To                 string    `gorm:"primaryKey"` // primary key This is the token, to which device ID should be sent the message! -> it's a very large string... and it can be costly
	Status             bool      // Sent or Failed
	SentTime           time.Time // When the messages has being sent successfully
	IsInvalidRecipient bool      // True/False
}

type inTimeStatistics struct {
	Time                 time.Time `gorm:"primaryKey"` // PRIMARY KEY: we should store here:  YEAR MONTH DAY HOUR
	SentMessages         uint64
	TotalRecipients      uint64
	TotalSuccessMessages uint64
	TotalFailedMessages  uint64
}
