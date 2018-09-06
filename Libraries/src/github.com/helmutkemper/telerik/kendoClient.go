package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoClient struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.client.create

	  Specifies the name of the client-side method of the SignalR hub responsible for creating data items.
	*/
	Create string `jsObject:"create"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.client.destroy

	  Specifies the name of the client-side method of the SignalR hub responsible for destroying data items.
	*/
	Destroy string `jsObject:"destroy"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.client.read

	  Specifies the name of the client-side method of the SignalR hub responsible for reading data items.
	*/
	Read string `jsObject:"read"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr.client.update

	  Specifies the name of the client-side method of the SignalR hub responsible for updating data items.
	*/
	Update string `jsObject:"update"`

	*ToJavaScriptConverter
}

func (el *KendoClient) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoClient.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
