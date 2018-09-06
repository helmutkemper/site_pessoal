package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridPageable struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.previousnext

	  If set to `true´ the pager will display buttons for going to the first, previous, next and last pages. By default
	  those buttons are displayed. (default: true)

	*/
	//
	//    Example - hide the first, previous, next, and last buttons
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
	//      pageable: {
	//        pageSize: 2,
	//        previousNext: false
	//      }
	//    });
	//    </script>
	//
	PreviousNext Boolean `jsObject:"previousNext"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.numeric

	  If set to `true´ the pager will display buttons for navigating to specific pages. By default those buttons are
	  displayed.

	  Using `pageable.numeric´ and <a href="pageable.input">pageable.input</a> at the same time is not recommended.
	  (default: true)

	*/
	//
	//    Example - hide the numeric pager buttons
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
	//      pageable: {
	//        pageSize: 2,
	//        numeric: false
	//      }
	//    });
	//    </script>
	//
	Numeric Boolean `jsObject:"numeric"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.refresh

	  If set to `true´ the pager will display the refresh button. Clicking the refresh button will refresh the grid. By
	  default the refresh button is not displayed. (default: false)

	*/
	//
	//    Example - show the refresh button
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
	//      pageable: {
	//        pageSize: 2,
	//        refresh: true
	//      }
	//    });
	//    </script>
	//
	Refresh Boolean `jsObject:"refresh"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.alwaysvisible

	  By default the grid will show the pager even when total amount of items in the DataSource is less than the pageSize.
	  (default: true)

	  If set to `false´ the grid will toggle the pager visibility as follows:

	  > when the total amount of items initially set in the DataSource is less than the pageSize number the pager will be
	  hidden.

	  > when the total amount of items initially set in the DataSource is greater than or equal to the pageSize number the
	  pager will be shown.

	  > when the total amount of items in the DataSource becomes less than the pageSize number (after delete, filter
	  operation or pageSize change) the pager will be hidden.

	  > when the total amount of items in the DataSource becomes greater than or equal to the pageSize number (after an
	  insert, filter operation or pageSize change) the pager will be shown.

	  Introduced in the Kendo UI 2017 R3 release.

	*/
	//
	//    Example - hide the pager if total items are less than pageSize
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
	//      pageable: {
	//        pageSize: 5,
	//        alwaysVisible: false
	//      }
	//    });
	//    </script>
	//
	AlwaysVisible Boolean `jsObject:"alwaysVisible"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.pagesizes

	  If set to `true´ the pager will display a drop-down which allows the user to pick a page size.
	  By default the page size drop-down is not displayed.

	  Can be set to an array of predefined page sizes to override the default list.
	  A special `all´ value is supported. It sets the page size to the total number of records.

	  If a `pageSize´ setting is provided for the data source then this value will be selected initially. (default: false)

	*/
	//
	//    Example - show the page size DropDownList
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "productName" },
	//          { field: "category" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { productName: "Tea", category: "Beverages" },
	//            { productName: "Coffee", category: "Beverages" },
	//            { productName: "Water", category: "Beverages" },
	//            { productName: "Juice", category: "Beverages" },
	//            { productName: "Decaffeinated Coffee", category: "Beverages" },
	//            { productName: "Iced Tea", category: "Beverages" },
	//            { productName: "Ham", category: "Food" },
	//            { productName: "Bread", category: "Food" },
	//            { productName: "Eggs", category: "Food" },
	//            { productName: "Bacon", category: "Food" },
	//            { productName: "Chips", category: "Food" },
	//            { productName: "Fish", category: "Food" }
	//          ],
	//          pageSize: 4
	//        },
	//        pageable: {
	//          pageSizes: true
	//        }
	//      });
	//    </script>
	//
	//
	//    Example - specify the page sizes as array
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category" }
	//      ],
	//      dataSource: {
	//        data: [
	//        { productName: "Tea", category: "Beverages" },
	//        { productName: "Coffee", category: "Beverages" },
	//        { productName: "Ham", category: "Food" },
	//        { productName: "Bread", category: "Food" }
	//        ],
	//        pageSize: 1
	//      },
	//      pageable: {
	//        pageSizes: [2, 3, 4, "all"],
	//        numeric: false
	//      }
	//    });
	//    </script>
	//
	PageSizes interface{} `jsObject:"pageSizes" jsType:"Boolean,[]interface{}"` //fixme: sizes?????????

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.messages

	  The text messages displayed in pager. Use this option to customize or localize the pager messages.
	*/
	Messages KendoGridPageableMessages `jsObject:"messages"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.buttoncount

	  The maximum number of buttons displayed in the numeric pager. The pager will display ellipsis (...) if there are more pages than the specified number. (default: 10)
	*/
	//
	//    Example - set pager button count
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
	//      pageable: {
	//        pageSize: 2,
	//        buttonCount: 1
	//      }
	//    });
	//    </script>
	//
	ButtonCount int `jsObject:"buttonCount"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.input

	  If set to `true´ the pager will display an input element which allows the user to type a specific page number. By default the page input is not displayed.

	  Using `pageable.input´ and <a href="pageable.numeric">pageable.numeric</a> at the same time is not recommended. (default: false)
	*/
	//
	//    Example - show the pager input
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
	//      pageable: {
	//        pageSize: 2,
	//        input: true
	//      }
	//    });
	//    </script>
	//
	Input Boolean `jsObject:"input"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.pagesize

	  The number of data items which will be displayed in the grid. This setting will not work if the Grid is assigned an already existing Kendo UI DataSource instance.
	*/
	//
	//    Example - set page size
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
	//      pageable: {
	//        pageSize: 2
	//      }
	//    });
	//    </script>
	//
	PageSize int `jsObject:"pageSize"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable.info

	   (default: true)
	*/
	Info Boolean `jsObject:"info"`

	*ToJavaScriptConverter
}

func (el *KendoGridPageable) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridPageable.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
