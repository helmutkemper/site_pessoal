package telerik

import "fmt"

func ExampleKendoUiColorPicker_ToHtml_Input() {
	html := KendoUiColorPicker{
		Html: HtmlInputText{
			Global: HtmlGlobalAttributes{
				Id: "picker",
			},
		},
		Buttons:     TRUE,
		ClearButton: TRUE,
		Columns:     3,
		TileSize: KendoTileSize{
			Width:  32,
			Height: 32,
		},
		Messages: KendoColorMessages{
			Apply:  "update",
			Cancel: "Discard",
		},
		Palette:  "websafe",
		Opacity:  TRUE,
		Preview:  TRUE,
		ToolIcon: "k-foreColor",
		Value:    "#b72bba",
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#picker").kendoColorPicker({buttons: true,clearButton: true,columns: 3,tileSize: { width: 32,height: 32,},messages: { apply: "update",cancel: "Discard",},palette: "websafe",opacity: true,preview: true,toolIcon: "k-foreColor",value: "#b72bba",});
}

func ExampleKendoUiColorPicker_ToHtml_Div() {
	html := KendoUiColorPicker{
		Html: HtmlInputText{
			Global: HtmlGlobalAttributes{
				Id: "picker",
			},
		},
		Buttons:     TRUE,
		ClearButton: TRUE,
		Columns:     3,
		TileSize: KendoTileSize{
			Width:  32,
			Height: 32,
		},
		Messages: KendoColorMessages{
			Apply:  "update",
			Cancel: "Discard",
		},
		Palette:  "websafe",
		Opacity:  TRUE,
		Preview:  TRUE,
		ToolIcon: "k-foreColor",
		Value:    "#b72bba",
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#picker").kendoColorPicker({buttons: true,clearButton: true,columns: 3,tileSize: { width: 32,height: 32,},messages: { apply: "update",cancel: "Discard",},palette: "websafe",opacity: true,preview: true,toolIcon: "k-foreColor",value: "#b72bba",});
}
