package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridPageableMessages struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.display
	//
	// The pager info text. Uses kendo.format. (default: "{0} - {1} of {2} items")
	//
	// Contains three placeholders:
	//
	// > {0} - the first data item index
	// > {1} - the last data item index
	// > {2} - the total number of data items
	//
	//    Example - set the "display" pager message
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
	//        pageable: {
	//          pageSize: 2,
	//          messages: {
	//            display: "Showing {0}-{1} from {2} data items"
	//          }
	//        }
	//      });
	//    </script>
	Display string `jsObject:"display"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.empty
	//
	// The text displayed when the grid is empty. (default: "No items to display")
	//
	//    Example - set the "empty" pager message
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "productName" },
	//          { field: "category" }
	//        ],
	//        dataSource: [],
	//        pageable: {
	//          pageSize: 2,
	//          messages: {
	//            empty: "No data"
	//          }
	//        }
	//      });
	//    </script>
	Empty string `jsObject:"empty"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.page
	//
	// The label displayed before the pager input. (default: "Page")
	//
	//    Example - set the label before the pager input
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
	//        pageable: {
	//          pageSize: 2,
	//          input: true,
	//          messages: {
	//            page: "Enter page"
	//          }
	//        }
	//      });
	//    </script>
	Page string `jsObject:"page"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.of
	//
	// The label displayed before the pager input. Uses kendo.format.
	// [ https://docs.telerik.com/kendo-ui/api/javascript/kendo/methods/format ]
	// Contains one optional placeholder {0} which represents the total number of pages.
	// (default: "of {0}")
	//
	//    Example - set the label after the pager input
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
	//        pageable: {
	//          pageSize: 2,
	//          input: true,
	//          messages: {
	//            of: "from {0}"
	//          }
	//        }
	//      });
	//    </script>
	Of string `jsObject:"of"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.itemsPerPage
	//
	// The label displayed after the page size DropDownList. (default: "items per page")
	//
	//    Example - set the label after the page size DropDownList
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
	//        pageable: {
	//          pageSize: 2,
	//          pageSizes: true,
	//          messages: {
	//            itemsPerPage: "data items per page"
	//          }
	//        }
	//      });
	//    </script>
	ItemsPerPage string `jsObject:"itemsPerPage"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.first
	//
	//  The tooltip of the button which goes to the first page. (default: "Go to the first page")
	//
	//    Example - set the Tooltip of the first page button
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
	//        pageable: {
	//          pageSize: 2,
	//          messages: {
	//            first: "First page"
	//          }
	//        }
	//      });
	//    </script>
	First string `jsObject:"first"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.last
	//
	// The tooltip of the button which goes to the last page. (default: "Go to the last page")
	//
	//    Example - set the Tooltip of the last page button
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
	//        pageable: {
	//          pageSize: 2,
	//          messages: {
	//            last: "Last page"
	//          }
	//        }
	//      });
	//    </script>
	Last string `jsObject:"last"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.next
	//
	// The Tooltip of the button which goes to the next page. (default: "Go to the next page")
	//
	//    Example - set the Tooltip of the next page button
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
	//        pageable: {
	//          pageSize: 2,
	//          messages: {
	//            next: "Next page"
	//          }
	//        }
	//      });
	//    </script>
	Next string `jsObject:"next"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.previous
	//
	// The Tooltip of the button which goes to the previous page. (default: "Go to the previous page")
	//
	//    Example - set the Tooltip of the previous page button
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
	//        pageable: {
	//          pageSize: 2,
	//          messages: {
	//            previous: "Previous page"
	//          }
	//        }
	//      });
	//    </script>
	Previous string `jsObject:"previous"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.refresh
	//
	// The Tooltip of the refresh button. (default: "Refresh")
	//
	//    Example - set the Tooltip of the refresh button
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
	//        pageable: {
	//          pageSize: 2,
	//          refresh: true,
	//          messages: {
	//            refresh: "Refresh the grid"
	//          }
	//        }
	//      });
	//    </script>
	Refresh string `jsObject:"refresh"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages#pageable.messages.morePages
	//
	// The Tooltip of the ellipsis ("...") button, which appears when the number of pages is greater than the buttonCount.
	// (default: "More pages")
	//
	//    Example - set the Tooltip of the ellipsis button
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
	//        pageable: {
	//          pageSize: 1,
	//          buttonCount: 2,
	//          refresh: true,
	//          messages: {
	//            morePages: "More pages"
	//          }
	//        }
	//      });
	//    </script>
	MorePages string `jsObject:"morePages"`

	*ToJavaScriptConverter
}

func (el *KendoGridPageableMessages) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridPageableMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
