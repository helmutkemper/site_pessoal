package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoTransport struct {
	/*
	  The configuration used when the data source destroys data items. Those are items removed from the data source via the remove method.
	  The data source uses jQuery.ajax to make an HTTP request to the remote service. The value configured via transport.destroy is passed to jQuery.ajax. This means that you can set all options supported by jQuery.ajax via transport.destroy except the success and error callback functions which are used by the transport.

	  If the value of transport.destroy is a function, the data source invokes that function instead of jQuery.ajax.

	  If the value of transport.destroy is a string, the data source uses this string as the URL of the remote service.
	  All transport actions (read, update, create, destroy) must be defined in the same way, that is, as functions or as objects. Mixing the different configuration alternatives is not possible.

	  Example - set the destroy remote service
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp"
	      },
	      // make JSONP request to https://demos.telerik.com/kendo-ui/service/products/destroy
	      destroy: {
	        url: "https://demos.telerik.com/kendo-ui/service/products/destroy",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      },
	      parameterMap: function(data, type) {
	        if (type == "destroy") {
	          // send the destroyed data items as the "models" service parameter encoded in JSON
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
	    // remove the first data item
	    dataSource.remove(products[0]);
	    // send the destroyed data item to the remote service
	    dataSource.sync();
	  });
	  </script>

	  Example - set destroy as a function
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: function(options) {
	        $.ajax({
	          url: "https://demos.telerik.com/kendo-ui/service/products",
	          dataType: "jsonp",
	          success: function(result) {
	            options.success(result);
	          }
	        });
	      },
	      destroy: function (options) {
	        // make JSONP request to https://demos.telerik.com/kendo-ui/service/products/destroy
	        $.ajax({
	          url: "https://demos.telerik.com/kendo-ui/service/products/destroy",
	          dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	          // send the destroyed data items as the "models" service parameter encoded in JSON
	          data: {
	            models: kendo.stringify(options.data.models)
	          },
	          success: function(result) {
	            // notify the data source that the request succeeded
	            options.success(result);
	          },
	          error: function(result) {
	            // notify the data source that the request failed
	            options.error(result);
	          }
	        });
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
	Destroy KendoDestroy `jsObject:"destroy" jsType:"*KendoDestroy,*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.parametermap#transport.parameterMap

	  The function which converts the request parameters to a format suitable for the remote service. By default, the data source sends the parameters using jQuery's conventions.

	  The parameterMap method is often used to encode the parameters in JSON format.
	  The parameterMap function will not be called when using custom functions for the read, update, create, and destroy operations.
	  If a transport.read.data function is used together with parameterMap, do not forget to preserve the result from the data function that will be received in the parameterMap arguments. An example is provided below. Generally, the parameterMap function is designed to transform the request payload, not add new parameters to it.

	  pseudo transport: { read: { url: "my-data-service-url", data: function () { return { foo: 1 }; } }, parameterMap: function (data, type) { // if type is "read", then data is { foo: 1 }, we also want to add { "bar": 2 } return kendo.stringify($.extend({ "bar": 2 }, data)); } }

	  Parameters
	  data Object
	  The parameters which will be sent to the remote service. The value specified in the data field of the transport settings (create, read, update or destroy) is included as well. If batch is set to false, the fields of the changed data items are also included.

	  data.aggregate Array
	  The current aggregate configuration as set via the aggregate option. Available if the serverAggregates option is set to true and the data source makes a "read" request.

	  data.group Array
	  The current grouping configuration as set via the group option. Available if the serverGrouping option is set to true and the data source makes a "read" request.

	  data.filter Object
	  The current filter configuration as set via the filter option. Available if the serverFiltering option is set to true and the data source makes a "read" request.

	  data.models Array
	  All changed data items. Available if there are any data item changes and the batch option is set to true.

	  data.page Number
	  The current page. Available if the serverPaging option is set to true and the data source makes a "read" request.

	  data.pageSize Number
	  The current page size as set via the pageSize option. Available if the serverPaging option is set to true and the data source makes a "read" request.

	  data.skip Number
	  The number of data items to skip. Available if the serverPaging option is set to true and the data source makes a "read" request.

	  data.sort Array
	  The current sort configuration as set via the sort option. Available if the serverSorting option is set to true and the data source makes a "read" request.

	  data.take Number
	  The number of data items to return (the same as data.pageSize). Available if the serverPaging option is set to true and the data source makes a "read" request.

	  type String
	  The type of the request which the data source makes.

	  The supported values are:

	  "create"
	  "read"
	  "update"
	  "destroy"
	  Returns
	  Object the request parameters converted to a format required by the remote service.

	  Example - convert data source request parameters
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: {
	        url: "https://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders?$format=json",
	        dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	        jsonp: "$callback",
	        cache: true
	      },
	      parameterMap: function(data, type) {
	        if (type == "read") {
	          // send take as "$top" and skip as "$skip"
	          return {
	            $top: data.take,
	            $skip: data.skip
	          }
	        }
	      }
	    },
	    schema: {
	      data: "d"
	    },
	    pageSize: 20,
	    serverPaging: true // enable serverPaging so take and skip are sent as request parameters
	  });
	  dataSource.fetch(function() {
	    console.log(dataSource.view().length); // displays "20"
	  });
	  </script>

	  Example - send request parameters as JSON
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      create: {
	        url: "https://demos.telerik.com/kendo-ui/service/products/create",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      },
	      parameterMap: function(data, type) {
	        return kendo.stringify(data);
	      }
	    },
	    batch: true,
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  dataSource.add( { ProductName: "New Product" });
	  dataSource.sync();
	  </script>
	*/
	ParameterMap JavaScript `jsObject:"parameterMap"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.create#transport.create

	  The configuration used when the data source saves newly created data items. Those are items added to the data source via the add or insert methods.

	  The data source uses jQuery.ajax to make a HTTP request to the remote service. The value configured via transport.create is passed to jQuery.ajax. This means that you can set all options supported by jQuery.ajax via transport.create except the success and error callback functions which are used by the transport.

	  If the value of transport.create is a function, the data source invokes that function instead of jQuery.ajax. Check the jQuery documentation for more details on the provided argument.

	  If the value of transport.create is a string, the data source uses this string as the URL of the remote service.

	  The remote service must return the inserted data items and the data item field configured as the id must be set. For example, if the id of the data item is ProductID, the "create" server response must be [{ "ProductID": 79, "AnotherProperties": "value"}] including the ID and the other properties of the data items.
	  All transport actions (read, update, create, destroy) must be defined in the same way, that is, as functions or as objects. Mixing the different configuration alternatives is not possible.

	  Example - set the create remote service
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      // make JSONP request to https://demos.telerik.com/kendo-ui/service/products/create
	      create: {
	        url: "https://demos.telerik.com/kendo-ui/service/products/create",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      },
	      parameterMap: function(data, type) {
	        if (type == "create") {
	          // send the created data items as the "models" service parameter encoded in JSON
	          return { models: kendo.stringify(data.models) };
	        }
	      }
	    },
	    batch: true,
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  // create a new data item
	  dataSource.add( { ProductName: "New Product" });
	  // save the created data item
	  dataSource.sync(); // server response is [{"ProductID":78,"ProductName":"New Product","UnitPrice":0,"UnitsInStock":0,"Discontinued":false}]
	  </script>

	  Example - set create as a function
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: function(options) {
	        / * implementation omitted for brevity * /
	      },
	      create: function(options) {
	        // make JSONP request to https://demos.telerik.com/kendo-ui/service/products/create
	        $.ajax({
	          url: "https://demos.telerik.com/kendo-ui/service/products/create",
	          dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	          // send the created data items as the "models" service parameter encoded in JSON
	          data: {
	            models: kendo.stringify(options.data.models)
	          },
	          success: function(result) {
	            // notify the data source that the request succeeded
	            options.success(result);
	          },
	          error: function(result) {
	            // notify the data source that the request failed
	            options.error(result);
	          }
	        });
	      }
	    },
	    batch: true,
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  dataSource.add( { ProductName: "New Product" });
	  dataSource.sync();
	  </script>
	*/
	Create interface{} `jsObject:"create" jsType:"*KendoCreate,*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.push

	  The function invoked during transport initialization which sets up push notifications. The data source will call this function only once and provide callbacks which will handle push notifications (data pushed from the server).

	  Parameters
	  callbacks Object
	  An object containing callbacks for notifying the data source of push notifications.

	  callbacks.pushCreate Function
	  Function that should be invoked to notify the data source about newly created data items that are pushed from the server. Accepts a single argument - the object pushed from the server which should follow the schema.data configuration.

	  callbacks.pushDestroy Function
	  Function that should be invoked to notify the data source about destroyed data items that are pushed from the server. Accepts a single argument - the object pushed from the server which should follow the schema.data configuration.

	  callbacks.pushUpdate Function
	  Function that should be invoked to notify the data source about updated data items that are pushed from the server. Accepts a single argument - the object pushed from the server which should follow the schema.data configuration.

	  Example
	  <script src="http://ajax.aspnetcdn.com/ajax/signalr/jquery.signalr-1.1.3.min.js"></script>
	  <script>
	  var hubUrl = "https://demos.telerik.com/kendo-ui/service/signalr/hubs";
	  var connection = $.hubConnection(hubUrl, { useDefaultPath: false});
	  var hub = connection.createHubProxy("productHub");
	  var hubStart = connection.start({ jsonp: true });
	  var dataSource = new kendo.data.DataSource({
	  transport: {
	    push: function(callbacks) {
	      hub.on("create", function(result) {
	        console.log("push create");
	        callbacks.pushCreate(result);
	      });
	      hub.on("update", function(result) {
	        console.log("push update");
	        callbacks.pushUpdate(result);
	      });
	      hub.on("destroy", function(result) {
	        console.log("push destroy");
	        callbacks.pushDestroy(result);
	      });
	    }
	  },
	  schema: {
	    model: {
	      id: "ID",
	      fields: {
	        "ID": { editable: false },
	        "CreatedAt": { type: "date" },
	        "UnitPrice": { type: "number" }
	      }
	    }
	  }
	  });
	  </script>
	*/
	Push JavaScript `jsObject:"push"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.read#transport.read

	  The configuration used when the data source loads data items from a remote service.
	  The data source uses jQuery.ajax to make a HTTP request to the remote service. The value configured via transport.read is passed to jQuery.ajax. This means that you can set all options supported by jQuery.ajax via transport.read except the success and error callback functions which are used by the transport.

	  If the value of transport.read is a function, the data source invokes that function instead of jQuery.ajax.

	  If the value of transport.read is a string, the data source uses this string as the URL of the remote service.
	  All transport actions (read, update, create, destroy) must be defined in the same way, that is, as functions or as objects. Mixing the different configuration alternatives is not possible.

	  Example - set the read remote service
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      // make JSONP request to https://demos.telerik.com/kendo-ui/service/products
	      read: {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      }
	    }
	  });
	  dataSource.fetch(function() {
	    console.log(dataSource.view().length); // displays "77"
	  });
	  </script>

	  Example - send additional parameters to the remote service
	  <input value="2" id="search" />
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: {
	        url: "https://demos.telerik.com/kendo-ui/service/products/read",
	        dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	        data: function() {
	            return {
	                skip: 0,
	                take: $("#search").val() // send the value of the #search input to the remote service
	            };
	        }
	      }
	    }
	  });
	  dataSource.fetch();
	  </script>

	  Example - set read as a function
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: function(options) {
	        // make JSONP request to https://demos.telerik.com/kendo-ui/service/products
	        $.ajax({
	          url: "https://demos.telerik.com/kendo-ui/service/products",
	          dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	          success: function(result) {
	            // notify the data source that the request succeeded
	            options.success(result);
	          },
	          error: function(result) {
	            // notify the data source that the request failed
	            options.error(result);
	          }
	        });
	      }
	    }
	  });
	  dataSource.fetch(function() {
	    console.log(dataSource.view().length); // displays "77"
	  });
	  </script>
	*/
	Read KendoRead `jsObject:"read"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.signalr#transport.signalr

	  The configuration used when type is set to "signalr". Configures the SignalR settings - hub, connection promise, server, and client hub methods.

	  Live demo available at demos.telerik.com/kendo-ui.

	  It is recommended to get familiar with the SignalR JavaScript API.

	  Example
	  <script src="http://ajax.aspnetcdn.com/ajax/signalr/jquery.signalr-1.1.3.min.js"></script>
	  <script>
	      var hubUrl = "https://demos.telerik.com/kendo-ui/service/signalr/hubs";
	      var connection = $.hubConnection(hubUrl, { useDefaultPath: false});
	      var hub = connection.createHubProxy("productHub");
	      var hubStart = connection.start({ jsonp: true });

	      var dataSource = new kendo.data.DataSource({
	          type: "signalr",
	          schema: {
	              model: {
	                  id: "ID",
	                  fields: {
	                      "ID": { editable: false, nullable: true },
	                      "CreatedAt": { type: "date" },
	                      "UnitPrice": { type: "number" }
	                  }
	              }
	          },
	          transport: {
	              signalr: {
	                  promise: hubStart,
	                  hub: hub,
	                  server: {
	                      read: "read",
	                      update: "update",
	                      destroy: "destroy",
	                      create: "create"
	                  },
	                  client: {
	                      read: "read",
	                      update: "update",
	                      destroy: "destroy",
	                      create: "create"
	                  }
	              }
	          }
	      });
	  </script>
	*/
	Signalr KendoSignalr `jsObject:"signalr"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.submit#transport.submit

	  A function that will handle create, update and delete operations in a single batch when custom transport is used, i.e. the transport.read is defined as a function.

	  The transport.create, transport.update and transport.delete operations will not be executed in this case.

	  This function will only be invoked when the DataSource is in batch mode.

	  Parameters
	  e.data Object
	  An object containing the created (e.data.created), updated (e.data.updated), and destroyed (e.data.destroyed) items.

	  e.success Function
	  A callback that should be called for each operation with two parameters - items and operation. See example below.

	  e.error Function
	  A callback that should be called in case of failure of any of the operations.

	  Example - declare transport submit function
	  <script>
	    var dataSource = new kendo.data.DataSource({
	      transport: {
	        read:  function(options){
	          $.ajax({
	            url: "https://demos.telerik.com/kendo-ui/service/products",
	            dataType: "jsonp",
	            success: function(result) {
	              options.success(result);
	            },
	            error: function(result) {
	              options.error(result);
	            }
	          });
	        },
	        submit: function(e) {
	          var data = e.data;
	          console.log(data);

	          // send batch update to desired URL, then notify success/error
	          e.success(data.updated,"update");
	          e.success(data.created,"create");
	          e.success(data.destroyed,"destroy");
	          e.error(null, "customerror", "custom error");
	        }
	      },
	      batch: true,
	      pageSize: 20,
	      schema: {
	        model: {
	          id: "ProductID",
	          fields: {
	            ProductID: { editable: false, nullable: true },
	            ProductName: { validation: { required: true } },
	            UnitPrice: { type: "number", validation: { required: true, min: 1} },
	            Discontinued: { type: "boolean" },
	            UnitsInStock: { type: "number", validation: { min: 0, required: true } }
	          }
	        }
	      }
	    });
	    dataSource.read().then(function(){
	      var productOne = dataSource.at(1),
	      productTwo = dataSource.at(2);
	      productOne.set("UnitPrice",42);
	      productTwo.set("UnitPrice",42);
	      dataSource.sync();
	    });
	  </script>
	*/
	Submit JavaScript `jsObject:"submit"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport.update

	  The configuration used when the data source saves updated data items. Those are data items whose fields have been updated.

	  The data source uses jQuery.ajax to make a HTTP request to the remote service. The value configured via transport.update is passed to jQuery.ajax. This means that you can set all options supported by jQuery.ajax via transport.update except the success and error callback functions which are used by the transport.

	  If the value of transport.update is a function, the data source invokes that function instead of jQuery.ajax.

	  If the value of transport.update is a string, the data source uses this string as the URL of the remote service.

	  All transport actions (read, update, create, destroy) must be defined in the same way, that is, as functions or as objects. Mixing the different configuration alternatives is not possible.

	  Example - specify update as a string
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
	    dataSource.sync(); //makes request to https://demos.telerik.com/kendo-ui/service/products/update
	  });
	  </script>

	  Example - specify update as a function
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: function(options) {
	        / * implementation omitted for brevity * /
	      },
	      update: function(options) {
	        // make JSONP request to https://demos.telerik.com/kendo-ui/service/products/update
	        $.ajax({
	          url: "https://demos.telerik.com/kendo-ui/service/products/update",
	          dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	          // send the updated data items as the "models" service parameter encoded in JSON
	          data: {
	            models: kendo.stringify(options.data.models)
	          },
	          success: function(result) {
	            // notify the data source that the request succeeded
	            options.success(result);
	          },
	          error: function(result) {
	            // notify the data source that the request failed
	            options.error(result);
	          }
	        });
	      }
	    },
	    batch: true,
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  dataSource.fetch(function() {
	    var product = dataSource.at(0);
	    product.set("UnitPrice", 20);
	    dataSource.sync(); //makes request to https://demos.telerik.com/kendo-ui/service/products/update
	  });
	  </script>
	*/
	Update KendoUpdate `jsObject:"update"`
	*ToJavaScriptConverter
}

func (el *KendoTransport) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoTransport.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
