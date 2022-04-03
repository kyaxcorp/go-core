package gor

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/err/define"
)

const ErrCodeAlreadyRunning = 100
const ErrCodeRunTimeIsZero = 101
const ErrCodeRunTimesIsLowerThanOne = 102
const ErrCodeOnRunIsNotAFunction = 103
const ErrCodeFailedToSetDefaultValuesForConfig = 104
const ErrCodeReRunOnRecoverMaxNrOfPanicsIsLowerThanOne = 105
const ErrCodeRunFunctionAlreadyRunning = 106
const ErrCodeStopFunctionAlreadyRunning = 107

var ErrAlreadyRunning = define.Err(ErrCodeAlreadyRunning, "already running...")
var ErrRunTimeIsZero = define.Err(ErrCodeRunTimeIsZero, "'run times' is 0")
var ErrRunTimesIsLowerThanOne = define.Err(ErrCodeRunTimesIsLowerThanOne, "'run times' is lower than -1")
var ErrOnRunIsNotAFunction = define.Err(ErrCodeOnRunIsNotAFunction, "'OnRun' if not a function")
var ErrFailedToSetDefaultValuesForConfig = define.Err(ErrCodeFailedToSetDefaultValuesForConfig, "failed to set default values for config")
var ErrReRunOnRecoverMaxNrOfPanicsIsLowerThanOne = define.Err(ErrCodeReRunOnRecoverMaxNrOfPanicsIsLowerThanOne, "'rerun on recover max nr of panics' times is lower than -1")
var ErrRunFunctionAlreadyRunning = define.Err(ErrCodeRunFunctionAlreadyRunning, "Run function already running...")
var ErrStopFunctionAlreadyRunning = define.Err(ErrCodeStopFunctionAlreadyRunning, "Stop function already running...")
