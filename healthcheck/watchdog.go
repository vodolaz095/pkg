package healthcheck

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/go-systemd/v22/daemon"
	"github.com/rs/zerolog/log"
)

// Pinger is interface StartWatchDog uses
type Pinger interface {
	Ping(context.Context) error
}

// Ready is used to nofity system is ready and check, if we need to perform systemd healthchecks
func Ready() (supported bool, err error) {
	return daemon.SdNotify(false, daemon.SdNotifyReady)
}

// SetStatus sets systemd unit status in human-readable format
// https://www.freedesktop.org/software/systemd/man/latest/sd_notify.html#STATUS=%E2%80%A6
func SetStatus(status string) (err error) {
	_, err = daemon.SdNotify(false, fmt.Sprintf("STATUS=%s", status))
	return
}

// SetReloading notify systemd that service is reloading
// https://www.freedesktop.org/software/systemd/man/latest/sd_notify.html#RELOADING=1
func SetReloading() (err error) {
	_, err = daemon.SdNotify(false, daemon.SdNotifyReloading)
	return
}

// SetStopping notify systemd that service is stopping
// https://www.freedesktop.org/software/systemd/man/latest/sd_notify.html#STOPPING=1
func SetStopping() (err error) {
	_, err = daemon.SdNotify(false, daemon.SdNotifyStopping)
	return
}

// Notify sends free form notification to systemd about service state change
func Notify(state string) (err error) {
	_, err = daemon.SdNotify(false, state)
	return
}

// StartWatchDog starts background process that notifies systemd if application is running properly
func StartWatchDog(mainCtx context.Context, pingers []Pinger) (err error) {
	var ok bool
	if len(pingers) == 0 {
		return fmt.Errorf("pingers are not set")
	}
	interval, err := daemon.SdWatchdogEnabled(false)
	if err != nil {
		return
	}
	if interval == 0 {
		log.Info().Msgf("Watchdog not enabled")
		return
	}
	ticker := time.NewTicker(interval / 2)
	for {
		select {
		case <-mainCtx.Done():
			ticker.Stop()
			return nil
		case t := <-ticker.C:
			ctx, cancel := context.WithDeadline(mainCtx, t.Add(interval/2))
			ok = true
			for i := range pingers {
				err = pingers[i].Ping(ctx)
				if err != nil {
					log.Error().Err(err).Msgf("error pinging %v: %s", i, err)
					ok = false
				}
			}
			cancel()

			if ok {
				_, err = daemon.SdNotify(false, daemon.SdNotifyWatchdog)
				if err != nil {
					log.Error().Err(err).Msgf("error sending watchdog notification: %s", err)
					continue
				}
				log.Trace().Msgf("Service is healthy!")
			} else {
				log.Warn().Msgf("Service is broken!")
			}
		}
	}
}
