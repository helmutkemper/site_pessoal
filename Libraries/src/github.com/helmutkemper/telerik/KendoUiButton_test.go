package telerik

import "fmt"

func ExampleKendoUiButton_ToHtml() {
	html := KendoUiButton{
		Html: HtmlElementFormButton{
			Global: HtmlGlobalAttributes{
				Id: "button",
			},
			Name: "send",
			Content: []interface{}{
				"Send",
			},
		},
	}

	fmt.Printf("%s", html.ToHtml())

	// Output:
	// <button id="button" name="send">Send</button>
}

func ExampleKendoUiButton_ToJavaScript_1() {
	html := KendoUiButton{
		Html: HtmlElementFormButton{
			Global: HtmlGlobalAttributes{
				Id:    "button",
				Class: "blue",
			},
			Name: "send",
			Content: []interface{}{
				"Send",
			},
		},
		SpriteCssClass: "k-icon netherlandsFlag",
		Enable:         FALSE,
		Icon:           "filter",
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#button").kendoButton({enable: false,icon: "filter",spriteCssClass: "k-icon netherlandsFlag",});
}

func ExampleKendoUiButton_ToJavaScript_2() {
	html := KendoUiButton{
		Html: HtmlElementFormButton{
			Global: HtmlGlobalAttributes{
				Id:    "button",
				Class: "blue",
			},
			Name: "send",
			Content: []interface{}{
				"Send",
			},
		},
		SpriteCssClass: "k-icon netherlandsFlag",
		Enable:         FALSE,
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#button").kendoButton({enable: false,spriteCssClass: "k-icon netherlandsFlag",});
}

func ExampleKendoUiButton_ToJavaScript_3() {
	html := KendoUiButton{
		Html: HtmlElementFormButton{
			Global: HtmlGlobalAttributes{
				Id: "button",
			},
			Name: "send",
			Content: []interface{}{
				"Send",
			},
		},
		Enable:   FALSE,
		ImageUrl: "../content/shared/icons/sports/snowboarding.png",
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#button").kendoButton({enable: false,imageUrl: "../content/shared/icons/sports/snowboarding.png",});
}
