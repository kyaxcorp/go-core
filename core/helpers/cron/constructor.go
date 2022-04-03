package cron

import c "github.com/robfig/cron/v3"

func New() *Cron {
	return &Cron{
		C: c.New(),
	}
}
