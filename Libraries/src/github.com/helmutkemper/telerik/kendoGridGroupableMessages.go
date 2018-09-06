package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridGroupableMessages struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/groupable.messages.empty
	//
	// The text displayed in the grouping drop area. (default: "Drag a column header and drop it here to group by that
	// column")
	//
	//    Example - set the "empty" grouping message
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "productName" },
	//          { field: "category" }
	//        ],
	//        dataSource: [
	//          { productName: "Tea", category: "Beverages" },
	//          { productName: "Coffee", category: "Beverages" },
	//          { productName: "Ham", category: "Food" },
	//          { productName: "Bread", category: "Food" }
	//        ],
	//        groupable: {
	//          messages: {
	//            empty: "Drop columns here"
	//          }
	//        }
	//      });
	//</script>
	Empty string `jsObject:"empty"`

	*ToJavaScriptConverter
}

func (el *KendoGridGroupableMessages) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridGroupableMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
