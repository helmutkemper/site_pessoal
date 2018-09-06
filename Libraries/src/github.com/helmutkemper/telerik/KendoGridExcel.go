package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridExcel struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/excel.allpages
	//
	// If set to true the grid will export all pages of data. By default the grid exports only the current page.
	//
	// Important:
	//
	// If the grid is bound to remote data and allPages is set to true it will request all data items from the remote
	// service. Be careful if you have a lot of data.
	//
	//    Example - export all pages of data
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        toolbar: ["excel"],
	//        excel: {
	//          allPages: true
	//        },
	//        dataSource: {
	//          transport: {
	//            read: {
	//              url: "https://demos.telerik.com/kendo-ui/service/products",
	//              dataType: "jsonp"
	//            }
	//          },
	//          pageSize: 10
	//        },
	//        pageable: true
	//      });
	//    </script>
	AllPages Boolean `jsObject:"allpages"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/excel.filename
	//
	// Specifies the file name of the exported Excel file.
	//
	//    Example - set the default Excel file name
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        toolbar: ["excel"],
	//        excel: {
	//          fileName: "Products.xlsx"
	//        },
	//        dataSource: {
	//          transport: {
	//            read: {
	//              url: "https://demos.telerik.com/kendo-ui/service/products",
	//              dataType: "jsonp"
	//            }
	//          },
	//          pageSize: 10
	//        },
	//        pageable: true
	//      });
	//    </script>
	FileName string `jsObject:"fileName"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/excel.filterable
	//
	// Enables or disables column filtering in the Excel file. Not to be mistaken with the grid filtering feature.
	//
	//    Example - enable filtering in the output Excel file
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        toolbar: ["excel"],
	//        excel: {
	//          filterable: false
	//        },
	//        dataSource: {
	//          transport: {
	//            read: {
	//              url: "https://demos.telerik.com/kendo-ui/service/products",
	//              dataType: "jsonp"
	//            }
	//          },
	//          pageSize: 10
	//        },
	//        pageable: true
	//      });
	//    </script>
	Filterable Boolean `jsObject:"filterable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/excel.forceproxy
	//
	// If set to true, the content will be forwarded to proxyURL even if the browser supports saving files locally.
	// (default: false)
	//
	ForceProxy Boolean `jsObject:"forceproxy"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/excel.proxyurl
	//
	// The URL of the server side proxy which will stream the file to the end user. (default: null)
	//
	// A proxy will be used when the browser isn't capable of saving files locally. Such browsers are IE version 9 and
	// lower and Safari.
	//
	// The developer is responsible for implementing the server-side proxy.
	//
	// The proxy will receive a POST request with the following parameters in the request body:
	//
	// > contentType: The MIME type of the file
	// > base64: The base-64 encoded file content
	// > fileName: The file name, as requested by the caller.
	//
	// The proxy should return the decoded file with the "Content-Disposition" header set to attachment;
	// filename="<fileName.xslx>".
	//
	//    Example - set the server proxy URL
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        toolbar: ["excel"],
	//        excel: {
	//          proxyURL: "/save"
	//        },
	//        dataSource: {
	//          transport: {
	//            read: {
	//              url: "https://demos.telerik.com/kendo-ui/service/products",
	//              dataType: "jsonp"
	//            }
	//          },
	//          pageSize: 10
	//        },
	//        pageable: true
	//      });
	//    </script>
	ProxyURL string `jsObject:"proxyURL"`

	*ToJavaScriptConverter
}

func (el *KendoGridExcel) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridExcel.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
