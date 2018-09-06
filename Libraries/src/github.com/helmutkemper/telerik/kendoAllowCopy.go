package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridAllowCopy struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/allowcopy#allowCopy.delimeter
	//
	// Changes the delimeter between the items on the same row. Use this option if you want to change the default TSV
	// format to CSV - set the delimeter to comma ','. (default: "\t")
	//
	//    Example - change the clipboard format from default TSV to CSV
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        selectable: "multiple cell",
	//        allowCopy: {
	//          delimeter: ",",
	//        },
	//        columns: [
	//          { field: "productName" },
	//          { field: "category" }
	//        ],
	//        dataSource: [
	//          { productName: "Tea", category: "Beverages" },
	//          { productName: "Coffee", category: "Beverages" },
	//          { productName: "Ham", category: "Food" },
	//          { productName: "Bread", category: "Food" }
	//        ]
	//      });
	//    </script>
	Delimeter string `jsObject:"delimeter"`

	*ToJavaScriptConverter
}

func (el *KendoGridAllowCopy) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridAllowCopy.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
