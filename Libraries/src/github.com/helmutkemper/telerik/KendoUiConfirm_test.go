package telerik

import "fmt"

func ExampleKendoUiConfirm_ToHtml() {
	html := KendoUiConfirm{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "confirm",
			},
		},
		Messages: KendoConfirmMessages{
			OkText: "Ok",
			Cancel: "Cancel",
		},
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#confirm").kendoConfirm({messages: { okText: "Ok",cancel: "Cancel",},});
}
