package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiNumericTextBox struct {
	Html HtmlInputNumber `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/culture#culture

	  Specifies the culture info used by the widget.

	  Example - specify German culture internationalization
	  <!--
	    TODO: Add the kendo.culture.de-DE.min.js file as it is required!

	    Here is a sample script tag:
	    <script src="http://kendo.cdn.telerik.com/{kendo version}/js/cultures/kendo.culture.de-DE.min.js"></script>

	    For more information check this help topic:
	    http://docs.telerik.com/kendo-ui/framework/globalization/overview
	  -->

	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	    culture: "de-DE"
	  });
	  </script>
	*/
	Culture string `jsObject:"culture"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/decimals#decimals

	  Specifies the number precision applied to the widget value and when the NumericTextBox is focused. If not set, the
	  precision defined by the current culture is used. If the user enters a number with a greater precision than is
	  currently configured, the widget value will be rounded. For example, if decimals is 2 and the user inputs 12.346, the
	  value will become 12.35. If the user inputs 12.99, the value will become 13.00.

	  Compare with the format property.

	  Example
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      decimals: 1
	  });
	  </script>
	*/
	Decimals int `jsObject:"decimals"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/downarrowtext#downArrowText

	  Specifies the text of the tooltip on the down arrow. (default: "Decrease value")

	  Example
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      downArrowText: "Less"
	  });
	  </script>
	*/
	DownArrowText string `jsObject:"downArrowText"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/factor#factor

	  Specifies the factor by which the value is multiplied. The obtained result is used as edit value. So, if 15 as string
	  is entered in the NumericTextBox and the factor value is set to 100 the visual value will be 1500. On blur the visual
	  value will be divided by 100 thus scaling the widget value to the original proportion. (default: 1)

	  Example
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	     format: "p0",
	     factor: 100,
	     min: 0,
	     max: 1,
	     step: 0.01
	  });
	  </script>
	*/
	Factor int `jsObject:"factor"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/format#format

	  @see https://docs.telerik.com/kendo-ui/framework/globalization/numberformatting

	  Specifies the number format used when the widget is not focused. Any valid number format is allowed.

	  Compare with the decimals property. (default: "n")

	  Example
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	     format: "c0"
	  });
	  </script>
	*/
	Format string `jsObject:"format"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/max#max

	  Specifies the largest value the user can enter.

	  Example - specify max option
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      max: 100
	  });
	  </script>

	  Example - specify max option as a HTML attribute
	  <input id="numerictextbox" max="100" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox();
	  </script>
	*/
	Max int `jsObject:"max"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/min#min

	  Specifies the smallest value the user can enter.
	  Example - specify min option
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      min: -100
	  });
	  </script>

	  Example - specify min option as a HTML attribute
	  <input id="numerictextbox" min="-100" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox();
	  </script>
	*/
	Min int `jsObject:"min"`

	/*
	  @sse https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/placeholder#placeholder

	  The hint displayed by the widget when it is empty. Not set by default.

	  Example
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      placeholder: "Select A Value"
	  });
	  </script>
	*/
	PlaceHolder string `jsObject:"placeholder"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/restrictdecimals#restrictDecimals

	  Specifies whether the decimals length should be restricted during typing. The length of the fraction is defined by the
	  decimals value.

	  Example
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      decimals: 3,
	      restrictDecimals: true
	  });
	  </script>
	*/
	RestrictDecimals Boolean `jsObject:"restrictDecimals"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/round#round

	  Specifies whether the value should be rounded or truncated. The length of the fraction is defined by the decimals
	  value. (default: true)

	  Example
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      round: false
	  });
	  </script>
	*/
	Round Boolean `jsObject:"round"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/step#step

	  Specifies the value used to increment or decrement widget value. (default: true)

	  Example - specify step option
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      step: 0.1
	  });
	  </script>

	  Example - specify step option as a HTML attribute
	  <input id="numerictextbox" step="0.1" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox();
	  </script>
	*/
	Step int `jsObject:"step"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/uparrowtext#upArrowText

	  Specifies the text of the tooltip on the up arrow. (default: "Increase value")

	  Example
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      upArrowText: "More"
	  });
	  </script>
	*/
	UpArrowText string `jsObject:"upArrowText"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/value#value

	  Specifies the value of the NumericTextBox widget.

	  Example - specify value option
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      value: 10
	  });
	  </script>

	  Example - specify value option as a HTML attribute
	  <input id="numerictextbox" value="10" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox();
	  </script>
	*/
	Value int `jsObject:"value"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/numerictextbox/configuration/spinners#spinners

	  Specifies whether the up and down spin buttons should be rendered (default: true)

	  Example - hide spin buttons
	  <input id="numerictextbox" />
	  <script>
	  $("#numerictextbox").kendoNumericTextBox({
	      spinners: false
	  });
	  </script>
	*/
	Spinners Boolean `jsObject:"spinners"`

	*ToJavaScriptConverter
}

func (el *KendoUiNumericTextBox) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoNumericTextBox.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoNumericTextBox({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiNumericTextBox) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiNumericTextBox) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
