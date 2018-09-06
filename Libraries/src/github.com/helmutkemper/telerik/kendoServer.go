package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoServer struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.server.create

	  Specifies the name of the server-side method of the SignalR hub responsible for creating data items.
	*/
	Create string `jsObject:"create"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.server.destroy

	  Specifies the name of the server-side method of the SignalR hub responsible for destroying data items.
	*/
	Destroy string `jsObject:"destroy"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.server.read

	  Specifies the name of the server-side method of the SignalR hub responsible for reading data items.
	*/
	Read string `jsObject:"read"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.server.update

	  Specifies the name of the server-side method of the SignalR hub responsible for updating data items.
	*/
	Update string `jsObject:"update"`

	*ToJavaScriptConverter
}

func (el *KendoServer) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoServer.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
