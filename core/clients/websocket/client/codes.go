package client

import (
	"github.com/kyaxcorp/go-core/core/helpers/errors2/define"
)

var WriterError = define.Err(100, "writer error")
var ReaderError = define.Err(200, "reader error")
var ConnectionFailedToHost = define.Err(300, "connection failed to host")
var NoConnectionsDefined = define.Err(400, "no connections have being defined")
