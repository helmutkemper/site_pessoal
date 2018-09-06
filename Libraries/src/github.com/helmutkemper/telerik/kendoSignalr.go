package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoSignalr struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.client

	  Specifies the client-side CRUD methods of the SignalR hub.
	*/
	Client KendoClient `jsObject:"client"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.hub

	  The SignalR hub object returned by the createHubProxy method. The hub option is mandatory.
	*/
	Hub JavaScript `jsObject:"hub"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.promise

	  The promise returned by the start method of the SignalR connection. The promise option is mandatory.
	*/
	Promise JavaScript `jsObject:"promise"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.server

	  Specifies the server-side CRUD methods of the SignalR hub.
	*/
	Server KendoServer `jsObject:"server"`

	*ToJavaScriptConverter
}

func (el *KendoSignalr) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoSignalr.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
