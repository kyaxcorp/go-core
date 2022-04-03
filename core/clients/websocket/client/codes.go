package client

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/err/define"
)

var WriterError = define.Err(100, "writer error")
var ReaderError = define.Err(200, "reader error")
var ConnectionFailedToHost = define.Err(300, "connection failed to host")
var NoConnectionsDefined = define.Err(400, "no connections have being defined")
