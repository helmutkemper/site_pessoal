package db

import log "github.com/helmutkemper/seelog"

func init() {
	logger, _ := log.LoggerFromConfigAsString(`<seelog levels="off"></seelog>`)
	log.ReplaceLogger(logger)
}
