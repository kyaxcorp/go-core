package gor

import (
	"github.com/KyaXTeam/go-core/v2/core/logger/model"
	"time"
)

type ReRunOnRecoverOptions struct {
	// After how much time it should run again after recovery! Duration it's measured in nanoseconds -> 100000000 = 100 ms or 0.1 seconds
	RunAfterDuration time.Duration `yaml:"run_after_duration" mapstructure:"run_after_duration" default:"100000000"`
	// How many panics it should be allowed to happen... -1 is infinite!
	MaxNrOfPanics int `yaml:"max_nr_of_panics" mapstructure:"max_nr_of_panics" default:"-1"`
}

// Config -> options
type Config struct {
	// Logger
	Logger *model.Logger

	OnRun func(instance *GInstance)
	// AutoRecover -> from panic!
	// It's a boolean
	AutoRecover interface{} `yaml:"auto_recover" mapstructure:"auto_recover" default:"yes"`
	// If the function OnRun had a panic, and it AutoRecovered, should it rerun itself again?!
	// If setting no, it will not rerun itself,
	// If Setting yes, it will rerun itself,
	// but be very carefully, it can run infinitely if you have infinite panics!
	// It's a boolean
	ReRunOnRecover interface{} `yaml:"rerun_on_recover" mapstructure:"rerun_on_recover" default:"no"`
	// This is the recovery options related to ReRunOnRecover
	ReRunOnRecoverOptions ReRunOnRecoverOptions `yaml:"rerun_on_recover_options" mapstructure:"rerun_on_recover_options"`

	// Add how many times it should run! default it's 1, -1 it's infinite, any number of times!
	RunTimes int `yaml:"run_times" mapstructure:"run_times" default:"1"`
	// AutoRun -> It's a Boolean, Should it run on creation!?
	AutoRun interface{} `yaml:"auto_run" mapstructure:"auto_run" default:"yes"`
	// when the execution finished!
	OnRunFinished func(instance *GInstance)

	OnExiting func(instance *GInstance)
	// OnAllFinished -> when everything has being finished as planned
	OnAllFinished func(instance *GInstance)
	// TODO: should we add here logger configuration? like level or other things?

}
