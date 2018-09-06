package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridNoRecords struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/norecords#noRecords.template

	  The template [ https://docs.telerik.com/kendo-ui/api/javascript/kendo/methods/template ] which is rendered when current view contains no records.
	*/
	//
	//    Example - customize the noRecords message
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      pageable: true,
	//      noRecords: {
	//        template: "No data available on current page. Current page is: #=this.dataSource.page()#"
	//      },
	//      dataSource: {
	//        data: [{name: "John", age: 29}],
	//        page: 2,
	//        pageSize: 10
	//      }
	//    });
	//    </script>
	//
	Template interface{} `jsObject:"template" jsType:"JavaScript,string"`

	*ToJavaScriptConverter
}

func (el *KendoGridNoRecords) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridNoRecords.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
