package sender

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/err/define"
)

var ErrConstructorNameEmpty = define.Err(100, "name is empty")
var ErrConstructorDBConnect = define.Err(101, "failed to connect to the db")
