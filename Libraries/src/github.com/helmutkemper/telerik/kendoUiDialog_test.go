package telerik

import "fmt"

func ExampleKendoUiDialog_String() {
	element := KendoUiDialog{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "dialog",
			},
		},
		Modal:   TRUE,
		Visible: TRUE,
		Title:   "Configuration error",
		Content: "Please, set a valid environment var.",
		Width:   400,
		Actions: []KendoActions{
			{
				Action: JavaScript{
					Code: "function(){ setTimeout( function(){ containerConfigurationEnvVarNameRef.focus(); }, 1000); }",
				},
				Primary: TRUE,
				Text:    "OK",
			},
		},
	}

	fmt.Printf("%s\n", element.ToHtml())
	fmt.Printf("<script>%s</script>\n", element.ToJavaScript())

	// Output:
	// <div  id="dialog" ></div>
	// <script>$("#dialog").kendoDialog({ actions: [{ text: "OK",action: function(){ setTimeout( function(){ containerConfigurationEnvVarNameRef.focus(); }, 1000); },primary: true, },],content: "Please, set a valid environment var.",modal: true,title: "Configuration error",visible: true,width: 400, });</script>
}
