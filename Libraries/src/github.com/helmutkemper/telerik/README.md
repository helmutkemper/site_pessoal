# Kendo-UI for Golang

telerik kendo ui for golang

https://godoc.org/github.com/helmutkemper/telerik

#### this package is still under development, so it still contains bugs.

For this example, I am have:
**DataSource: KendoDataSource** -> **Transport: KendoTransport** -> **Read: KendoRead** -> **Url: "/static/test/read"**. 

returns:

```json
{
    "Meta": {
        "Total": 2
    },
    "Objects": [
         {
             "id": 1,
             "ExposedPortsNumber": 80,
             "ExposedPortsProtocol": "TCP",
             "ExposedPortsShow": "80/TCP"
         },
         {
             "id": 2,
             "ExposedPortsNumber": 8080,
             "ExposedPortsProtocol": "TCP",
             "ExposedPortsShow": "8080/TCP"
         }
     ]
}
```
**DataSource: KendoDataSource** -> **Transport: KendoTransport** -> **Create: KendoCreate** -> **Url: "/static/test/create"**.

returns:

```json
{
    "Meta": {
        "Total": 1
    },
    "Objects": [
         {
             "id": 3,
             "ExposedPortsNumber": 90,
             "ExposedPortsProtocol": "UDP",
             "ExposedPortsShow": "90/UDP"
         }
     ]
}
```

Afther this, **DataSource: KendoDataSource** -> **Transport: KendoTransport** -> **Read: KendoRead** -> **Url: "/static/test/read"**. 

returns:

```json
{
    "Meta": {
        "Total": 3
    },
    "Objects": [
         {
             "id": 1,
             "ExposedPortsNumber": 80,
             "ExposedPortsProtocol": "TCP",
             "ExposedPortsShow": "80/TCP"
         },
         {
             "id": 2,
             "ExposedPortsNumber": 8080,
             "ExposedPortsProtocol": "TCP",
             "ExposedPortsShow": "8080/TCP"
         },
         {
             "id": 3,
             "ExposedPortsNumber": 90,
             "ExposedPortsProtocol": "UDP",
             "ExposedPortsShow": "90/UDP"
         }
     ]
}
```

Example:

```golang
  el := HtmlElementDiv{
    Global: HtmlGlobalAttributes{
      Id: "spanCreateTemplateExposedPortsAddNewPort",
      Class: "k-content",
      Style: "width: 300px !important;",
    },
    Content: Content{

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigHostName",
            Content: Content{
              "Host Name",
            },
            Global: HtmlGlobalAttributes{
              Style: "width: 200px important!;",
            },
          },

          &HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigHostName",
              Class: "k-textbox",
            },
            Name: "HostName",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigDomainName",
            Content: Content{
              "Domain Name",
            },
          },

          &HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigDomainName",
              Class: "k-textbox",
            },
            Name: "DomainName",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigUser",
            Content: Content{
              "User",
            },
          },

          &HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigUser",
              Class: "k-textbox",
            },
            Name: "User",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigAttachStdIn",
            Content: Content{
              "Attach Std In",
            },
          },

          &KendoUiDropDownList{
            Html: HtmlInputText{
              Global: HtmlGlobalAttributes{
                Id: "ConfigAttachStdIn",
              },
              Name: "AttachStdIn",
            },
            DataValueField: "key",
            DataTextField: "value",

            DataSource: []map[string]interface{}{
              {
                "key": -1,
                "value": "Default",
              },
              {
                "key": 0,
                "value": "No",
              },
              {
                "key": 1,
                "value": "Yes",
              },
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigAttachStdOut",
            Content: Content{
              "Attach Std Out",
            },
          },

          &KendoUiMobileSwitch{
            Html: HtmlInputCheckBox{
              Global: HtmlGlobalAttributes{
                Id: "ConfigAttachStdOut",
              },
              Name: "AttachStdOut",
            },
            OnLabel: "Yes",
            OffLabel: "No",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigAttachStdErr",
            Content: Content{
              "Host Attach Std Err",
            },
          },

          &KendoUiMobileSwitch{
            Html: HtmlInputCheckBox{
              Global: HtmlGlobalAttributes{
                Id: "ConfigAttachStdErr",
              },
              Name: "AttachStdErr",
            },
            OnLabel: "Yes",
            OffLabel: "No",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigExposedPorts",
            Content: Content{
              "Exposed Ports",
            },
          },

          &KendoUiMultiSelect{
            Html: HtmlElementFormSelect{
              Global: HtmlGlobalAttributes{
                Id: "ConfigExposedPorts",
              },
              Name: "ExposedPorts",
            },
            ClearButton: FALSE,
            DataValueField: "id",
            DataTextField: "ExposedPortsShow",
            DataSource: KendoDataSource{
              //Type: KENDO_TYPE_DATA_JSON,
              Transport: KendoTransport{
                Read: KendoRead{
                  Url: "/static/test/read",
                  Type: HTML_METHOD_GET,
                  DataType: KENDO_TYPE_DATA_JSON_JSON,
                },
                Create: KendoCreate{
                  Url: "/static/test/create",
                  Type: HTML_METHOD_POST,
                  DataType: KENDO_TYPE_DATA_JSON_JSON,
                },
              },
              Schema: KendoSchema{
                Data:  "Objects",
                Total: "Total",
                Model: KendoDataModel{
                  Id: "id",
                  Fields: map[string]KendoField{
                    "id": {
                      Type: JAVASCRIPT_NUMBER,
                    },
                    "ExposedPortsNumber": {
                      Type: JAVASCRIPT_NUMBER,
                    },
                    "ExposedPortsProtocol": {
                      Type: JAVASCRIPT_STRING,
                    },
                    "ExposedPortsShow": {
                      Type: JAVASCRIPT_STRING,
                    },
                  },
                },
              },
              //PageSize: 10,
              ServerPaging: TRUE,
            },
            Dialog: KendoUiDialog{
              Html: HtmlElementDiv{
                Global: HtmlGlobalAttributes{
                  Id: getAutoId(),
                },
              },
              Title: "Add new exposed port.",
              Content: Content{

                // regra, o form valida automaticamente
                &HtmlElementDiv{
                  Global: HtmlGlobalAttributes{
                    Id: "ConfigExposedPortsDialogContent",
                    Class: "k-content",
                  },
                  Content: Content{

                    &HtmlElementDiv{
                      Content: Content{

                        &HtmlElementFormLabel{
                          For: "ExposedPortsNumber",
                          Content: Content{
                            "Port number",
                          },
                        },

                        &KendoUiNumericTextBox{
                          Html: HtmlInputNumber{
                            Name:         "ExposedPortsNumber",
                            PlaceHolder:  "",
                            AutoComplete: FALSE,
                            Required:     TRUE,
                            // Pattern: "[^=]*",
                            Global: HtmlGlobalAttributes{
                              Id:    "ExposedPortsNumber",
                              Class: "oneThirdSize",
                              Extra: map[string]interface{}{
                                "validationMessage": "Enter a {0}",
                              },
                            },
                          },
                          Format: "#",
                        },
                      },
                    },

                    &HtmlElementDiv{
                      Content: Content{

                        &HtmlElementFormLabel{
                          For: "ExposedPortsProtocol",
                          Content: Content{
                            "Port protocol",
                          },
                        },

                        &KendoUiComboBox{
                          Html: HtmlElementFormSelect{
                            Global: HtmlGlobalAttributes{
                              Id:    "ExposedPortsProtocol",
                              Class: "oneThirdSize",
                              Data:  map[string]string{"required-msg": "Select start time"},
                            },
                            Name:     "ExposedPortsProtocol",
                            Required: TRUE,
                            Options: []HtmlOptions{
                              {
                                Label: "Please, select one",
                                Key:   "",
                              },
                              {
                                Label: "TCP",
                                Key:   "TCP",
                              },
                              {
                                Label: "UDP",
                                Key:   "UDP",
                              },
                            },
                          },
                        },
                      },
                    },

                    &HtmlInputHidden{
                      Global: HtmlGlobalAttributes{
                        Id: "ExposedPortsShow",
                      },
                      Name: "ExposedPortsShow",
                    },

                  },
                },
              },
              Visible: FALSE,
              Width: 400,
              Actions: []KendoActions{
                {
                  Primary: FALSE,
                  Text:    "Close",
                },
                {
                  Primary: FALSE,
                  Text:    "Add",
                  ButtonType: BUTTON_TYPE_ADD,
                },
                {
                  Primary: TRUE,
                  Text:    "Add and close",
                  ButtonType: BUTTON_TYPE_ADD_AND_CLOSE,
                },
              },
            },
            NoDataTemplate: HtmlElementScript{
              Global: HtmlGlobalAttributes{
                Id: getAutoId(),
              },
              Type: SCRIPT_TYPE_KENDO_TEMPLATE,
              Content: Content{

                &HtmlElementDiv{
                  Content: Content{
                    "No data found. Do you want to add new item?",
                  },
                },

                "<br>",
                "<br>",

                &HtmlElementFormButton{
                  ButtonType: BUTTON_TYPE_ADD_IN_TEMPLATE,
                  Global: HtmlGlobalAttributes{
                    Id: getAutoId(),
                    Class: "k-button",
                  },
                  Content: Content{
                    "Add new item",
                  },
                },

              },
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigTry",
            Content: Content{
              "Try",
            },
          },

          &KendoUiMobileSwitch{
            Html: HtmlInputCheckBox{
              Global: HtmlGlobalAttributes{
                Id: "ConfigTry",
              },
              Name: "Try",
            },
            OnLabel: "Yes",
            OffLabel: "No",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigOpenStdIn",
            Content: Content{
              "Open Std In",
            },
          },

          &KendoUiMobileSwitch{
            Html: HtmlInputCheckBox{
              Global: HtmlGlobalAttributes{
                Id: "ConfigOpenStdIn",
              },
              Name: "OpenStdIn",
            },
            OnLabel: "Yes",
            OffLabel: "No",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigStdInOnce",
            Content: Content{
              "Std In Once",
            },
          },

          &KendoUiMobileSwitch{
            Html: HtmlInputCheckBox{
              Global: HtmlGlobalAttributes{
                Id: "ConfigStdInOnce",
              },
              Name: "StdInOnce",
            },
            OnLabel: "Yes",
            OffLabel: "No",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigEnv",
            Content: Content{
              "Env",
            },
          },

          &KendoUiMultiSelect{
            Html: HtmlElementFormSelect{
              Global: HtmlGlobalAttributes{
                Id: "ConfigEnv",
              },
              Name: "Env",
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigCmd",
            Content: Content{
              "Cmd",
            },
          },

          &KendoUiMultiSelect{
            Html: HtmlElementFormSelect{
              Global: HtmlGlobalAttributes{
                Id: "ConfigCmd",
              },
              Name: "Cmd",
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigHealthCheck",
            Content: Content{
              "Health Check",
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigArgsEscaped",
            Content: Content{
              "Args Escaped",
            },
          },

          &KendoUiMobileSwitch{
            Html: HtmlInputCheckBox{
              Global: HtmlGlobalAttributes{
                Id: "ConfigArgsEscaped",
              },
              Name: "ArgsEscaped",
            },
            OnLabel: "Yes",
            OffLabel: "No",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigImage",
            Content: Content{
              "Image",
            },
          },

          &KendoUiAutoComplete{
            Html: HtmlInputText{
              Global: HtmlGlobalAttributes{
                Id: "ConfigHostImage",
                Class: "k-textbox",
              },
              Name: "Image",
            },
            DataTextField: "ExposedPortsShow",
            ClearButton: FALSE,
            DataSource: KendoDataSource{
              //VarName: "testDataSource",
              //Type: KENDO_TYPE_DATA_JSON,
              Transport: KendoTransport{
                Read: KendoRead{
                  Url: "/static/test/read",
                  Type: HTML_METHOD_GET,
                  DataType: KENDO_TYPE_DATA_JSON_JSON,
                },
                Create: KendoCreate{
                  Url: "/static/test/create",
                  Type: HTML_METHOD_POST,
                  DataType: KENDO_TYPE_DATA_JSON_JSON,
                },
              },
              Schema: KendoSchema{
                Data:  "Objects",
                Total: "Total",
                Model: KendoDataModel{
                  Id: "id",
                  Fields: map[string]KendoField{
                    "id": {
                      Type: JAVASCRIPT_NUMBER,
                    },
                    "ExposedPortsNumber": {
                      Type: JAVASCRIPT_NUMBER,
                    },
                    "ExposedPortsProtocol": {
                      Type: JAVASCRIPT_STRING,
                    },
                    "ExposedPortsShow": {
                      Type: JAVASCRIPT_STRING,
                    },
                  },
                },
              },
              //PageSize: 10,
              ServerPaging: TRUE,
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigVolumes",
            Content: Content{
              "Volumes",
            },
          },

          &KendoUiMultiSelect{
            Html: HtmlElementFormSelect{
              Global: HtmlGlobalAttributes{
                Id: "ConfigVolumes",
              },
              Name: "Volumes",
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigWorkingDir",
            Content: Content{
              "Working Dir",
            },
          },

          &HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigWorkingDir",
              Class: "k-textbox",
            },
            Name: "WorkingDir",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigEntryPoint",
            Content: Content{
              "EntryPoint",
            },
          },

          &KendoUiMultiSelect{
            Html: HtmlElementFormSelect{
              Global: HtmlGlobalAttributes{
                Id: "ConfigEntryPoint",
              },
              Name: "EntryPoint",
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigNetworkDisabled",
            Content: Content{
              "Network Disabled",
            },
          },

          &KendoUiMobileSwitch{
            Html: HtmlInputCheckBox{
              Global: HtmlGlobalAttributes{
                Id: "ConfigNetworkDisabled",
              },
              Name: "NetworkDisabled",
            },
            OnLabel: "Yes",
            OffLabel: "No",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigMacAddress",
            Content: Content{
              "Mac Address",
            },
          },

          &HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigMacAddress",
              Class: "k-textbox",
            },
            Name: "MacAddress",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigOnBuild",
            Content: Content{
              "On Build",
            },
          },

          &KendoUiMultiSelect{
            Html: HtmlElementFormSelect{
              Global: HtmlGlobalAttributes{
                Id: "ConfigOnBuild",
              },
              Name: "OnBuild",
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigLabels",
            Content: Content{
              "Labels",
            },
          },

          &KendoUiMultiSelect{
            Html: HtmlElementFormSelect{
              Global: HtmlGlobalAttributes{
                Id: "ConfigLabels",
              },
              Name: "Labels",
            },
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigStopSignal",
            Content: Content{
              "Stop Signal",
            },
          },

          &HtmlInputText{
            Global: HtmlGlobalAttributes{
              Id: "ConfigStopSignal",
              Class: "k-textbox",
            },
            Name: "StopSignal",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigStopTimeout",
            Content: Content{
              "Stop Timeout",
            },
          },

          &KendoUiNumericTextBox{
            Html: HtmlInputNumber{
              Global: HtmlGlobalAttributes{
                Id: "ConfigStopTimeout",
              },
              Name: "StopTimeout",
            },
            Format: "#s",
          },

        },
      },

      &HtmlElementDiv{
        Content: Content{

          &HtmlElementFormLabel{
            For: "ConfigShell",
            Content: Content{
              "Shell",
            },
          },

          &KendoUiMultiSelect{
            Html: HtmlElementFormSelect{
              Global: HtmlGlobalAttributes{
                Id: "ConfigShell",
              },
              Name: "Shell",
            },
          },

        },
      },

    },
  }
  
  fmt.Printf( `<!DOCTYPE html>
<html>
<head>
    <style>
      html { 
        font-size: 14px; 
        font-family: Arial, Helvetica, sans-serif; 
      }
    </style>
    <title></title>
    <link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.common-material.min.css" />
    <link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.material.min.css" />
    <link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.material.mobile.min.css" />
    <style>
      .k-content {
        width: 500px;
      }
      input[type=text], .k-multiselect-wrap, .k-textbox, .k-numerictextbox, .k-combobox {
        width: 100%%;
        padding: 5px 5px;
        margin: 10px 0;
        display: inline-block;
        border: 1px solid #ccc;
        border-radius: 4px;
        box-sizing: border-box;
      }

      input[type=text], .k-multiselect-wrap, .k-textbox, .k-numerictextbox, .k-combobox, .k-dropdown {
        width: 100%%;
        padding: 5px 5px;
        margin: 10px 0;
        display: inline-block;
        border: 1px solid #ccc;
        border-radius: 4px;
        box-sizing: border-box;
    }

    .k-multiselect.k-header, .k-switch.k-header, .k-dropdown {
      border-color: #ffffff;
      background: #ffffff;
    }

    .k-multiselect.k-header.k-state-focused, .k-dropdown {
      background-color: #ffffff;
    }

      .k-switch, .km-switch {
        margin: 5px 0;
        display: block;
      }
    </style>
    <script src="https://kendo.cdn.telerik.com/2018.1.221/js/jquery.min.js"></script>
    <script src="https://kendo.cdn.telerik.com/2018.1.221/js/kendo.all.min.js"></script>
    %s
    <script>
      %s
      $(document).ready(function () {
        %s
      });

    function getData(){
      var ConfigExposedPorts = [];

      var data = ConfigExposedPortsDataSource.data();
      var value = getFormValue( 'id:ExposedPorts' );

      for( var kValue in value ){
        if( !value.hasOwnProperty( kValue ) ){
          continue;
        }

        for( var kData in data ){
          if( !data.hasOwnProperty( kData ) ){
            continue;
          }

          if( data[ kData ].id === value[ kValue ] ){
            ConfigExposedPorts.push( data[ kData ] );
          }

        }
      }

      console.log( 'ConfigHostName:', getFormValue( 'id:HostName' ) );
      console.log( 'data to send:', ConfigExposedPorts );
      console.log( 'data to send:', ConfigExposedPorts );
    }
    </script>
</head>
<body>
  <div id="example">
    %s
  </div>
</body>
</html>
%s`, el.Content.MakeJsScript(), el.Content.MakeJsObject(), el.Content.ToJavaScript(), el.ToHtml(), el.ToHtmlSupport() )
```


if you're wondering, why do it this way? the answer is simple, a complex, handmade form took a few hours to get ready. The same form done this way took about ten minutes to get ready.

Briefly I put the complete example with Bank CRUD for mongodb
