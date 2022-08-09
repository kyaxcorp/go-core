package vars

import "github.com/kyaxcorp/go-core/core/logger/model"

// ApplicationLogger -> This is the app logger which handles all logs writing to a single file
var ApplicationLogger *model.Logger

// CoreLogger -> this is the first logger which is been created...
// it's more for debugging lib things
var CoreLogger *model.Logger
