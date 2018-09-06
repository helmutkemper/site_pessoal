package telerik

import "fmt"

func ExampleKendoUiMenu_ToHtml_1() {
	el := KendoUiMenu{
		Html: HtmlElementUl{
			Global: HtmlGlobalAttributes{
				Id: "menu",
			},
		},
		DataSource: []MenuObject{
			{
				Text:     "Item 1",
				CssClass: "myClass",                // Add custom CSS class to the item, optional, added 2012 Q3 SP1.
				Url:      "http://www.kendoui.com", // Link URL if navigation is needed, optional.
				Attr: map[string]interface{}{
					"custom": "value", // Add attributes with specified values
					"other":  "value",
				},
			},
			{
				Text:      "<b>Item 2</b>",
				AllowHtml: FALSE,  // Allows use of HTML for item text
				Content:   "text", // content within an ite
			},
			{
				Text: "Item 3",
				ImageAttr: MenuImageAttr{ // Add additional image attributes
					Alt:    "Image",
					Height: 25,
					Width:  25,
				},
				ImageUrl: "image.png", // Item image URL, optional.
				Items: []MenuObject{
					{
						Text: "Sub item 1",
					},
					{
						Text: "Sub item 2",
					},
				},
			},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s\n", el.ToJavaScript())

	// Output:
	//
}

func ExampleKendoUiMenu_ToHtml_2() {
	el := KendoUiMenu{
		Html: HtmlElementUl{
			Global: HtmlGlobalAttributes{
				Id: "menu",
			},
			Content: Content{

				&HtmlElementLi{
					Content: Content{
						"Item 1",

						&HtmlElementUl{
							Content: Content{

								&HtmlElementLi{
									Content: Content{
										"Sub Item 1",
									},
								},

								&HtmlElementLi{
									Content: Content{
										"Sub Item 2",
									},
								},

								&HtmlElementLi{
									Content: Content{
										"Sub Item 3",
									},
								},
							},
						},
					},
				},

				&HtmlElementLi{
					Content: Content{
						"Item 2",

						&HtmlElementUl{
							Content: Content{

								&HtmlElementLi{
									Content: Content{
										"Sub Item 1",
									},
								},

								&HtmlElementLi{
									Content: Content{
										"Sub Item 2",
									},
								},

								&HtmlElementLi{
									Content: Content{
										"Sub Item 3",
									},
								},
							},
						},
					},
				},
			},
		},

		Direction:    MENU_DIRECTION_BOTTOM_LEFT,
		CloseOnClick: TRUE,
		HoverDelay:   200,
		OpenOnClick: MenuOpenOnClick{
			RootMenuItems: TRUE,
			SubMenuItems:  TRUE,
		},
		Orientation:    MENU_ORIENTATION_HORIZONTAL,
		PopupCollision: MENU_COLLISION_FIT,
		Scrollable: MenuScrollable{
			Distance: 200,
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s\n", el.ToJavaScript())

	// Output:
	//
}
