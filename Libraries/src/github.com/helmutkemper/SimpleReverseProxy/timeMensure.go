package SimpleReverseProxy

import (
	log "github.com/helmutkemper/seelog"
	"time"
)

func timeMeasure(start time.Time, name string) {
	elapsed := NetworkTime.Since(start)
	log.Infof("%s: %s", name, elapsed)
}
