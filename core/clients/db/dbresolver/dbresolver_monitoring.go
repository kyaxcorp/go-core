package dbresolver

import (
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/clients/db/codes"
	"github.com/kyaxcorp/go-core/core/helpers/gor"
	"github.com/rs/zerolog"
	"time"
)

/*
This process will monitor the status of the resolvers
Based on their status it will set the global resolver for processing...

*/
func (dr *DBResolver) startResolversMonitoring() error {
	info := func() *zerolog.Event {
		return dr.LInfoF("resolvers_monitoring")
	}
	_error := func() *zerolog.Event {
		return dr.LErrorF("resolvers_monitoring")
	}

	info().Msg("starting resolvers monitoring...")

	resolverMonitoring, _err := gor.GO(dr.ctx, gor.Config{
		Logger: dr.Logger,
		OnRun: func(instance *gor.GInstance) {
			info().Msg(color.LightGreen.Render("resolvers monitoring started..."))

			for {
				nrOfMasters := len(dr.masters)

				if instance.IsTerminating() {
					break
				}

				// The Priority it's set by their order!
				var activeMasters []*resolver
				var inactiveMasters []*resolver
				for _, r := range dr.masters {
					r.isMonitoringActive.WaitUntilTrue()

					if r.CanProcessWriteOp() {
						// Check if started...
						activeMasters = append(activeMasters, r)
					} else {
						inactiveMasters = append(inactiveMasters, r)
					}
				}

				dr.activeMastersLock.Lock()
				dr.nrOfActiveMasters.Set(len(activeMasters))
				dr.nrOfInactiveMasters.Set(len(inactiveMasters))
				dr.activeMasters = activeMasters
				dr.inactiveMasters = inactiveMasters
				dr.activeMastersLock.Unlock()

				// Set as active!
				dr.isMonitoringActive.True()

				info().
					Int("nr_of_masters", nrOfMasters).
					Int("nr_of_active_masters", dr.nrOfActiveMasters.Get()).
					Int("nr_of_inactive_masters", dr.nrOfInactiveMasters.Get()).
					Msg(color.LightYellow.Render("available masters"))

				if instance.IsTerminating() {
					break
				}
				time.Sleep(time.Second)
			}
		},
	})

	if _err != nil {
		// Error... return this error!
		_error().Err(codes.ErrFailedToGenerateResolverMonitoringGoroutine).Msg("")
		return codes.ErrFailedToGenerateResolverMonitoringGoroutine
	}
	// Set the connection revival here...
	dr.resolversMonitoring = resolverMonitoring
	_err = dr.resolversMonitoring.Run()
	if _err != nil {
		_error().Err(codes.ErrFailedToRunResolverMonitoringGoroutine).Msg("")
		return codes.ErrFailedToRunResolverMonitoringGoroutine
	}

	return nil
}
