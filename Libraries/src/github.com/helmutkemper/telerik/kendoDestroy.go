package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoDestroy struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.destroy#transport.destroy.cache

	  If set to false, the request result will not be cached by the browser. Setting cache to false will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.

	  Refer to the jQuery.ajax documentation for further information.

	  Example - enable request caching
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      destroy: {
	        / * omitted for brevity * /
	        cache: true
	      }
	    }
	  });
	  </script>
	*/
	Cache Boolean `jsObject:"cache"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.destroy#transport.destroy.contentType

	  The content-type HTTP header sent to the server. The default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON. Refer to the jQuery.ajax documentation for further information.

	  Example - set the content type
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      destroy: {
	        / * omitted for brevity * /
	        contentType: "application/json"
	      }
	    }
	  });
	  </script>
	*/
	ContentType string `jsObject:"contentType"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.destroy#transport.destroy.data

	  Additional parameters which are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for sorting, filtering, paging, and grouping.

	  Refer to the jQuery.ajax documentation for further information.

	  Example - send additional parameters as an object
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      destroy: {
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
	      destroy: {
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
	Data interface{} `jsObject:"contentType" jsType:"*JavaScript,[]map[string]interface{}"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.destroy#transport.destroy.dataType

	  The type of result expected from the server. Commonly used values are "json" and "jsonp".

	  Refer to the jQuery.ajax documentation for further information.

	  Example - set the data type to JSON
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      destroy: {
	        / * omitted for brevity * /
	        dataType: "json"
	      }
	    }
	  });
	  </script>
	*/
	DataType KendoTypeDataJSon `jsObject:"dataType"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.destroy#transport.destroy.type

	  The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
	  The type option is ignored if dataType is set to "jsonp". JSONP always uses GET requests.

	  Refer to the jQuery.ajax documentation for further information.

	  Example
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      destroy: {
	        / * omitted for brevity * /
	        type: "POST"
	      }
	    }
	  });
	  </script>
	*/
	Type HtmlMethod `jsObject:"type"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.destroy#transport.destroy.url

	  The URL to which the request is sent.

	  If set to function, the data source will invoke it and use the result as the URL.

	  Example - specify the URL as a string
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp"
	      },
	      destroy: {
	        url: "https://demos.telerik.com/kendo-ui/service/products/destroy",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      },
	      parameterMap: function(data, type) {
	        if (type == "destroy") {
	          return { models: kendo.stringify(data.models) }
	        }
	      }
	    },
	    batch: true,
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  dataSource.fetch(function() {
	    var products = dataSource.data();
	    dataSource.remove(products[0]);
	    dataSource.sync();
	  });
	  </script>

	  Example - specify the URL as a function
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp"
	      },
	      destroy: {
	        url: function (options) {
	          return "https://demos.telerik.com/kendo-ui/service/products/destroy"
	        },
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      },
	      parameterMap: function(data, type) {
	        if (type == "destroy") {
	          return { models: kendo.stringify(data.models) }
	        }
	      }
	    },
	    batch: true,
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  dataSource.fetch(function() {
	    var products = dataSource.data();
	    dataSource.remove(products[0]);
	    dataSource.sync();
	  });
	  </script>
	*/
	Url interface{} `jsObject:"url" jsType:"*JavaScript,string"`

	*ToJavaScriptConverter
}

func (el *KendoDestroy) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoDestroy.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
