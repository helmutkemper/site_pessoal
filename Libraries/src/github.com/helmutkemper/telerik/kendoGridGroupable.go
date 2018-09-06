package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridGroupable struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/groupable.showfooter

	  When enabled the group footer rows will remain visible when the corresponding group is collapsed. (default: false)
	*/
	//
	//    Example - show footer when groups are collapsed
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName", groupFooterTemplate: "this is a footer" },
	//        { field: "category" }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: "Beverages" },
	//        { productName: "Beer", category: "Beverages" },
	//        { productName: "Cheese", category: "Food" },
	//      ],
	//      groupable: {
	//        showFooter: true
	//      }
	//    });
	//    </script>
	//
	ShowFooter Boolean `jsObject:"showFooter"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/groupable.enabled

	  When set to false grouping is considered disabled. (default: true)
	*/
	//
	//    Example - enable grouping
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category" }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: "Beverages" },
	//        { productName: "Coffee", category: "Beverages" },
	//        { productName: "Ham", category: "Food" },
	//        { productName: "Bread", category: "Food" }
	//      ],
	//      groupable: {
	//        enabled: false
	//      }
	//    });
	//    </script>
	//
	Enabled Boolean `jsObject:"enabled"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/groupable.messages

	  The text messages displayed during grouping.
	*/
	Messages KendoGridGroupableMessages `jsObject:"messages"`

	*ToJavaScriptConverter
}

func (el *KendoGridGroupable) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridGroupable.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
