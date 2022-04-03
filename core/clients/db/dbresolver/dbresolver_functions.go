package dbresolver

import "github.com/kyaxcorp/go-core/core/clients/db/driver"

func (dr *DBResolver) SetMainConfig(config driver.Config) {
	dr.mainConfig = config
}
