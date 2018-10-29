package telerik

import "fmt"

func ExampleGetTemplate() {
	customTemplate := CustomTemplate{
		Id: []byte(`ExposePort`),
	}
	customTemplate.Footer = customTemplate.GetFooterTemplateButton([]byte(`Add`), []byte(`centerText`))
	customTemplate.NoData = []byte(`<div>No ports found to add. Please add a port and a protocol before continuing.</div>`)
	customTemplate.Template = []byte(`<span>Port: #= Value #<br>Usually used in: #= ImageName #</span>`)

	html := KendoUiMultiSelect{
		Html: HtmlElementFormSelect{
			Global: HtmlGlobalAttributes{
				Id: "multiSelect",
			},
		},
		AutoClose:      FALSE,
		AutoWidth:      TRUE,
		ClearButton:    FALSE,
		DataTextField:  "productName",
		DataValueField: "productId",
		Delay:          100,
		Filter:         FILTER_CONTAINS,
		// FixedGroupTemplate: "Fixed header: #: data #",
		FooterTemplate: JavaScript{
			Code: string(customTemplate.GetElementFooterTemplate()),
		},
		// GroupTemplate: "Group template: #: data #",
		NoDataTemplate: JavaScript{
			Code: string(customTemplate.GetElementNoDataTemplate()),
		},
		Placeholder: "Select...",
		// HeaderTemplate: "<div><h2>Fruits</h2></div>",
		ItemTemplate: JavaScript{
			Code: string(customTemplate.GetElementTemplate()),
		},
		// TagMode: TAG_MODE_MULTIPLE,
		Value: []map[string]interface{}{
			{"productName": "Item 1", "productId": "1"},
			{"productName": "Item 2", "productId": "2"},
			{"productName": "Item 3", "productId": "3"},
		},
	}

	fmt.Printf(`<!DOCTYPE html>
<html>
<head>
    <base href="https://demos.telerik.com/kendo-ui/grid/index">
    <style>html { font-size: 14px; font-family: Arial, Helvetica, sans-serif; }</style>
    <title></title>
    <link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.default-v2.min.css" />

    <script src="https://kendo.cdn.telerik.com/2018.1.221/js/jquery.min.js"></script>
    <script src="https://kendo.cdn.telerik.com/2018.1.221/js/kendo.all.min.js"></script>
    <script>
    %s
    $(document).ready(function () {
    %s
    });
    </script>
    %s
    %s
    %s

</head>
<body>

<div id="example">
    %s
</div>

<style type="text/css">
    .customer-photo {
        display: inline-block;
        width: 32px;
        height: 32px;
        border-radius: 50%%;
        background-size: 32px 35px;
        background-position: center center;
        vertical-align: middle;
        line-height: 32px;
        box-shadow: inset 0 0 1px #999, inset 0 0 10px rgba(0,0,0,.2);
        margin-left: 5px;
    }

    .customer-name {
        display: inline-block;
        vertical-align: middle;
        line-height: 32px;
        padding-left: 3px;
    }
</style>


</body>
</html>
`,
		customTemplate.GetFooterTemplateButtonOpenDialogJavaScript(),
		html.ToJavaScript(),
		customTemplate.GetTemplate(),
		customTemplate.GetFooterTemplate(),
		customTemplate.GetNoDataTemplate(),
		html.ToHtml(),
	)

	// Output:
	//
}

func ExampleIdea() {

	dataSource := &KendoDataSource{
		//VarName: "exposedPortsDataSource",
		Schema: KendoSchema{
			Model: KendoDataModel{
				Id: "Id",
				Fields: map[string]KendoField{
					"Id": {
						Type: JAVASCRIPT_NUMBER,
					},
					"Value": {
						Type: JAVASCRIPT_STRING,
					},
					"ImageName": {
						Type: JAVASCRIPT_STRING,
					},
				},
			},
		},
	}

	dialogWindow := &HtmlElementScript{
		Global: HtmlGlobalAttributes{
			Id: "containerCreateTemplateExposedPortsAddNewPort",
		},
		Type: SCRIPT_TYPE_KENDO_TEMPLATE,
		Content: Content{

			&HtmlElementDiv{
				Global: HtmlGlobalAttributes{
					Id: "spanCreateTemplateExposedPortsAddNewPort",
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
											},
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
						},
					},
				},
			},
		},
	}

	content := HtmlContent{
		Content: Content{

			dataSource,

			&HtmlElementScript{
				Global: HtmlGlobalAttributes{
					Id: "containerHostExposedPortsFooterTemplate",
				},
				Type: SCRIPT_TYPE_KENDO_TEMPLATE,
				Content: Content{

					&HtmlElementFormButton{
						Global: HtmlGlobalAttributes{
							Id:      "buttonHostExposedPortsFooterTemplate",
							Class:   "k-button k-primary centerText",
							OnClick: "containerHostExposedPortsAddNewPort()",
						},
						Content: Content{
							"Add new port and protocol",
						},
					},
				},
			},

			&HtmlElementScript{
				Type: SCRIPT_TYPE_JAVASCRIPT,
				Content: Content{

					`function containerHostExposedPortsAddNewPort(){ $('#containerHostExposedPortsAddNewPort').data('kendoDialog').open(); }`,
				},
			},

			&HtmlElementScript{
				Global: HtmlGlobalAttributes{
					Id: "containerHostExposedPortsTemplate",
				},
				Type: SCRIPT_TYPE_KENDO_TEMPLATE,
				Content: Content{

					&HtmlElementSpan{
						Content: Content{
							"containerHostExposedPortsTemplate",
						},
					},
				},
			},

			dialogWindow,

			&HtmlElementScript{
				Global: HtmlGlobalAttributes{
					Id: "containerHostExposedPortsNoDataTemplate",
				},
				Type: SCRIPT_TYPE_KENDO_TEMPLATE,
				Content: Content{

					&HtmlElementDiv{
						Global: HtmlGlobalAttributes{
							Id: "divExposedPortsFooterTemplate",
						},
						Content: Content{
							"No ports found to add. Please add a port and a protocol before continuing.",
						},
					},
				},
			},

			&KendoUiMultiSelect{
				Html: HtmlElementFormSelect{
					Global: HtmlGlobalAttributes{
						Id: "containerHostExposedPorts",
					},
				},
				Filter:      FILTER_CONTAINS,
				Placeholder: "",
				ItemTemplate: JavaScript{
					Code: "kendo.template($('#containerHostExposedPortsTemplate').html())",
				},
				NoDataTemplate: JavaScript{
					Code: "kendo.template($('#containerHostExposedPortsNoDataTemplate').html())",
				},
				FooterTemplate: JavaScript{
					Code: "kendo.template($('#containerHostExposedPortsFooterTemplate').html())",
				},
				DataTextField:  "Value",
				DataValueField: "Id",
				DataSource:     dataSource,
			},

			&KendoUiDialog{
				Html: HtmlElementDiv{
					Global: HtmlGlobalAttributes{
						Id: "containerHostExposedPortsAddNewPort",
					},
				},
				Modal:   TRUE,
				Visible: FALSE,
				Title:   "Expose port from container",
				Width:   400,
				Content: JavaScript{
					Code: string(dialogWindow.ToKendoTemplate()),
				},
				Actions: []KendoActions{
					{
						Text: "--Close",
					},
					{
						/*Action:  JavaScript{
						  Code: `function(input){
						          if(!$('#spanCreateTemplateExposedPortsAddNewPort').kendoValidator().data('kendoValidator').validate())
						          {
						            return false;
						          }

						          dataSource.add({
						            Id: dataSource.total(),
						            Value: $('#ExposedPortsNumber').val() + '/' + $('#ExposedPortsProtocol').val(),
						            ImageName: 'fixme: containerConfigurationImageNameRef.text()'
						          });

						          dataSource.one('requestEnd', function(args) {
						            if (args.type !== 'create') {
						              return;
						            }

						            dataSource.one('sync', function() {
						              //$('#containerHostExposedPorts').value($('#containerHostExposedPorts').value().concat([containerHostExposedPortsItemsIdToAdd]));
						            });

						            $('#ExposedPortsProtocol').data('kendoComboBox').val('')
						            $('#ExposedPortsNumber').data('kendoNumericTextBox').value('');
						          });

						          dataSource.sync();
						          return false;
						        }`,
						},*/
						Text: "--Add",
					},
					{
						/*Action:  JavaScript{
						  Code: "function(input){ if(!$('#spanCreateTemplateExposedPortsAddNewPort').kendoValidator().data('kendoValidator').validate()){ return false; } }",
						},*/
						Primary: TRUE,
						Text:    "--Add and close",
					},
				},
				EventOpen: JavaScript{
					Code: "function(){ " + string(dialogWindow.Content.ToJavaScript()) + "}",
				},
			},
		},
	}

	fmt.Printf("%s\n\n", content.ToJavaScript())
	fmt.Printf("%s", content.ToHtml())

	// Output:
	//
}
func ExampleSoUmTest() {
	el := HtmlElementDiv{
		Global: HtmlGlobalAttributes{
			Id:    "spanCreateTemplateExposedPortsAddNewPort",
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
							Id:    "ConfigHostName",
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
							Id:    "ConfigDomainName",
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
							Id:    "ConfigUser",
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
						DataTextField:  "value",

						DataSource: []map[string]interface{}{
							{
								"key":   -1,
								"value": "Default",
							},
							{
								"key":   0,
								"value": "No",
							},
							{
								"key":   1,
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
						OnLabel:  "Yes",
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
						OnLabel:  "Yes",
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
						ClearButton:    FALSE,
						DataValueField: "id",
						DataTextField:  "ExposedPortsShow",
						DataSource: &KendoDataSource{
							//Type: KENDO_TYPE_DATA_JSON,
							Transport: KendoTransport{
								Read: KendoRead{
									Url:      "/static/test/read",
									Type:     HTML_METHOD_GET,
									DataType: KENDO_TYPE_DATA_JSON_JSON,
								},
								Create: KendoCreate{
									Url:      "/static/test/create",
									Type:     HTML_METHOD_POST,
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
									Id: GetAutoId(),
								},
							},
							Title: "Add new exposed port.",
							Content: Content{

								// regra, o form valida automaticamente
								&HtmlElementDiv{
									Global: HtmlGlobalAttributes{
										Id:    "ConfigExposedPortsDialogContent",
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
							Width:   400,
							Actions: []KendoActions{
								{
									Primary: FALSE,
									Text:    "Close",
								},
								{
									Primary:    FALSE,
									Text:       "Add",
									ButtonType: BUTTON_TYPE_ADD,
								},
								{
									Primary:    TRUE,
									Text:       "Add and close",
									ButtonType: BUTTON_TYPE_ADD_AND_CLOSE,
								},
							},
						},
						NoDataTemplate: HtmlElementScript{
							Global: HtmlGlobalAttributes{
								Id: GetAutoId(),
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
										Id:    GetAutoId(),
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
						OnLabel:  "Yes",
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
						OnLabel:  "Yes",
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
						OnLabel:  "Yes",
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
						OnLabel:  "Yes",
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
								Id:    "ConfigHostImage",
								Class: "k-textbox",
							},
							Name: "Image",
						},
						DataTextField: "ExposedPortsShow",
						ClearButton:   FALSE,
						DataSource: &KendoDataSource{
							//VarName: "testDataSource",
							//Type: KENDO_TYPE_DATA_JSON,
							Transport: KendoTransport{
								Read: KendoRead{
									Url:      "/static/test/read",
									Type:     HTML_METHOD_GET,
									DataType: KENDO_TYPE_DATA_JSON_JSON,
								},
								Create: KendoCreate{
									Url:      "/static/test/create",
									Type:     HTML_METHOD_POST,
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
							Id:    "ConfigWorkingDir",
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
						OnLabel:  "Yes",
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
							Id:    "ConfigMacAddress",
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
							Id:    "ConfigStopSignal",
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

	/*
	  desejado:
	  {
	    "Docker": {
	      "AttachStdin": true,
	      "AttachStdout": true,
	      "AttachStderr": true,
	      "ExposedPorts": {
	        "27018/tcp": {}
	      },
	      "OpenStdin": true,
	      "Cmd": [
	        "bash"
	      ],
	      "Image": "mongo:latest"
	    },
	    "Host": {
	      "PortBindings": {
	        "27017/tcp": [
	          {
	            "HostPort": "27018"
	          }
	        ]
	      }
	    },
	    "Name": "mongo_test"
	  }

	  formato:
	  {
	    "root": [
	      "Docker", "AttachStdin"
	    ],
	    "types": [
	      {
	        "name": "Docker",
	        "type": "object",
	        "attributes": [ "ExposedPorts" ]
	      },
	      {
	        "name": "ExposedPortKey",
	        "type": "concat",
	        "attributes": [ "id:ExposedPortNumber", "/", "id:ExposedPortProtocol" ]
	      },
	      {
	        "name": "ExposedPortValue",
	        "type": "object"
	      },
	      {
	        "name": "AttachStdin",
	        "type": "boolean",
	        "souce": {
	          "value": "id:AttachStdin"
	        }
	      },
	      {
	        "name": "ExposedPorts",
	        "type": "object",
	        "source": {
	          "key": "name:ExposedPortKey",
	          "value": "name:ExposedPortValue"
	        }
	      }
	    }
	  ]
	*/

	fmt.Printf(`<!DOCTYPE html>
<html>
<head>
    <!--base href="http://localhost:8888/static/js/kendo-ui/"-->
    <style>
      html { 
        font-size: 14px; 
        font-family: Arial, Helvetica, sans-serif; 
      }
    </style>
    <title></title>
    <link rel="stylesheet" href="./styles/kendo.common-material.min.css" />
    <link rel="stylesheet" href="./styles/kendo.material.min.css" />
    <link rel="stylesheet" href="./styles/kendo.material.mobile.min.css" />
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
    <script src="./js/jquery.min.js"></script>
    <script src="./js/kendo.all.min.js"></script>
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
%s`, el.Content.MakeJsScript(), el.Content.MakeJsObject(), el.Content.ToJavaScript(), el.ToHtml(), el.ToHtmlSupport())

	// Output:
	//
}
