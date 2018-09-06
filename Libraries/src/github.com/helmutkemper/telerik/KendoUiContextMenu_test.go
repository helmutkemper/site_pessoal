package telerik

import "fmt"

func ExampleKendoUiContextMenu_ToHtml() {
	html := KendoUiContextMenu{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "context-menu",
			},
		},
		AlignToAnchor: TRUE,
		Animation: KendoAnimation{
			Open: KendoOpen{
				Duration: 500,
				Effects:  EFFECT_EXPAND_IN,
			},
			Close: KendoClose{
				Duration: 500,
				Effects:  EFFECT_EXPAND_OUT,
			},
		},
		AppendTo:         "#container",
		CloseOnClick:     TRUE,
		CopyAnchorStyles: TRUE,
		DataSource: []map[string]interface{}{
			{
				"text":     "Item 1",
				"cssClass": "myClass",
				"url":      "http://www.kendoui.com",
			},
			{
				"text":    "<b>Item 2</b>",
				"encoded": false,
				"content": "text",
			},
			{
				"text":     "Item 3",
				"imageUrl": "http://www.kendoui.com/test.jpg",
				"items": []map[string]interface{}{
					{
						"text": "Sub Item 1",
					},
					{
						"text": "Sub Item 2",
					},
				},
			},
			{
				"text":           "Item 4",
				"spriteCssClass": "imageClass3",
			},
		},
		Direction:      CONTEXT_MENU_DIRECTION_LEFT,
		Filter:         ".box",
		HoverDelay:     200,
		Orientation:    ORIENTATION_HORIZONTAL,
		PopupCollision: FALSE,
		ShowOn:         MOUSE_EVENT_CLICK,
		Target:         "#target",
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#context-menu").kendoContextMenu({alignToAnchor: true,animation: { close: { effects: "expand:out",duration: 500,},open: { effects: "expand:in",duration: 500,},},appendTo: "#container",closeOnClick: true,copyAnchorStyles: true,dataSource: [{"url": "http://www.kendoui.com","text": "Item 1","cssClass": "myClass",},{"text": "<b>Item 2</b>","encoded": false,"content": "text",},{"text": "Item 3","imageUrl": "http://www.kendoui.com/test.jpg","items": [{"text": "Sub Item 1",},{"text": "Sub Item 2",},],},{"text": "Item 4","spriteCssClass": "imageClass3",},],direction: "left",filter: ".box",hoverDelay: 200,orientation: "horizontal",popupCollision: false,showOn: "click",target: "#target",});
}
