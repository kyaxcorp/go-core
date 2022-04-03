package record

import "github.com/KyaXTeam/go-core/v2/core/helpers/err/define"

func (r *Record) Exists() (bool, error) {
	if r.LoadData() {
		// load data procedure is ok!
		return r.isRecordExists.Get(), nil
	} else {
		// failed to load data!
		return false, define.Err(0, "load data procedure failed")
	}
}
