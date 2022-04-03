package dbresolver

import "github.com/KyaXTeam/go-core/v2/core/clients/db/driver"

func (dr *DBResolver) SetMainConfig(config driver.Config) {
	dr.mainConfig = config
}
