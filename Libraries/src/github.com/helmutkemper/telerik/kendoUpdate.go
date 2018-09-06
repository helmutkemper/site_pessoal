package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUpdate struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.update#transport.update.cache

	  If set to false, the request result will not be cached by the browser. Setting cache to false will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.

	  Refer to the jQuery.ajax documentation for further information.

	  Example - enable request caching
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      update: {
	        / * omitted for brevity * /
	        cache: true
	      }
	    }
	  });
	  </script>
	*/
	Cache Boolean `jsObject:"cache"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.update#transport.update.contentType

	  The content-type HTTP header sent to the server. Default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON. Refer to the jQuery.ajax documentation for further information.

	  Example - set content type
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      update: {
	        / * omitted for brevity * /
	        contentType: "application/json"
	      }
	    }
	  });
	  </script>
	*/
	ContentType string `jsObject:"contentType"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.update#transport.update.data

	  Additional parameters which are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for sorting, filtering, paging, and grouping.

	  Refer to the jQuery.ajax documentation for further information.

	  Example - send additional parameters as an object
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      update: {
	        / * omitted for brevity * /
	        data: {
	          name: "Jane Doe",
	          age: 30
	        }
	      }
	    }
	  });
	  </script>

	  Example - send additional parameters by returning them from a function
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      update: {
	        / * omitted for brevity * /
	        data: function() {
	          return {
	            name: "Jane Doe",
	            age: 30
	          }
	        }
	      }
	    }
	  });
	  </script>
	*/
	Data interface{} `jsObject:"data" jsType:"[]map[string]interface{},*JavaScript"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.create#transport.create.dataType

	  The type of result expected from the server. Commonly used values are "json" and "jsonp".

	  Refer to the jQuery.ajax documentation for further information.

	  Example - set the data type to JSON
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      update: {
	        / * omitted for brevity * /
	        dataType: "json"
	      }
	    }
	  });
	  </script>
	*/
	DataType KendoTypeDataJSon `jsObject:"dataType"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.update#transport.update.type

	  The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
	  The type option is ignored if dataType is set to "jsonp". JSONP always uses GET requests.

	  Refer to the jQuery.ajax documentation for further information.

	  Example - set the HTTP verb of the request
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      update: {
	        / * omitted for brevity * /
	        type: "POST"
	      }
	    }
	  });
	  </script>
	*/
	Type HtmlMethod `jsObject:"type"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.update#transport.update.url

	  The URL to which the request is sent.

	  If set to function, the data source will invoke it and use the result as the URL.

	  Example - specify URL as a string
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read:  {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      },
	      update: {
	        url: "https://demos.telerik.com/kendo-ui/service/products/update",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      }
	    },
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  dataSource.fetch(function() {
	    var product = dataSource.at(0);
	    product.set("UnitPrice", 20);
	    dataSource.sync();
	  });
	  </script>

	  Example - specify URL as a function
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read:  {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      },
	      update: {
	        url: function(options) {
	          return "https://demos.telerik.com/kendo-ui/service/products/update"
	        },
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      }
	    },
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  dataSource.fetch(function() {
	    var product = dataSource.at(0);
	    product.set("UnitPrice", 20);
	    dataSource.sync();
	  });
	  </script>
	*/
	Url interface{} `jsObject:"url" jsType:"*JavaScript,string"`

	*ToJavaScriptConverter
}

func (el *KendoUpdate) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoUpdate.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
