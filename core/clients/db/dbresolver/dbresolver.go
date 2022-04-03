package dbresolver

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/_context"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_bool"
	"github.com/KyaXTeam/go-core/v2/core/helpers/sync/_int"
	"gorm.io/gorm"
	"sync"
)

const (
	Write Operation = "write"
	Read  Operation = "read"
)

func (dr *DBResolver) Name() string {
	return "gorm:db_resolver"
}

// Initialize this is how gorm initializes our code...
func (dr *DBResolver) Initialize(db *gorm.DB) error {
	dr.DB = db
	// Set the context
	dr.ctx = db.Statement.Context

	dr.nrOfActiveMasters = _int.NewVal(0)
	dr.nrOfInactiveMasters = _int.NewVal(0)
	//dr.isMonitoringActive = _bool.NewVal(false)
	dr.isMonitoringActive = _bool.NewValContext(false, dr.ctx)

	dr.activeMastersLock = &sync.RWMutex{}
	// Register callbacks
	dr.registerCallbacks(db)
	// Compile
	return dr.compile()
}

func (dr *DBResolver) IsTerminating() bool {
	// We simply create a temporary context to not be blocked by using directly
	tmpContext := _context.WithCancel(dr.ctx)
	if tmpContext.IsDone() {
		return true
	}
	return false
}
