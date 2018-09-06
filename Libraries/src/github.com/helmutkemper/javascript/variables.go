package javascript

import (
	"crypto/tls"
	"github.com/helmutkemper/safeType"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DisableCompression:  true,
		TLSHandshakeTimeout: 30 * time.Second,
	},
	Timeout: 120 * time.Minute,
}

/*
  Example of use

  js.ExternalFunc.Set( "makeGrid", makeGrid )
  js.ExternalFunc.Set( "addRouteToGetData", addRouteToGetData )
  js.ExternalFunc.Set( "addRouteToUpdateData", addRouteToUpdateData )
  js.ExternalFunc.Set( "getJWT", getJWT )
  js.ExternalFunc.Set( "getRoute", getRoute )
  js.ExternalFunc.Set( "getProcessLogStarted", GetProcessLogStarted )
  js.ExternalFunc.Set( "getProcessLogEnd", GetProcessLogEnd )
*/
var ExternalFunc safeType.SafeMapStringInterface
