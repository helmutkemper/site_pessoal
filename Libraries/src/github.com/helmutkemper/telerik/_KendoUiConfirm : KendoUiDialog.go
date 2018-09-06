package telerik

type KendoUiConfirm struct {
	HtmlId string

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/confirm#configuration-messages

	  Defines the text of the labels that are shown within the confirm dialog. Used primarily for localization.

	  Example
	   <div id="confirm"></div>
	   <script>
	   $("#confirm").kendoConfirm({
	     messages:{
	       okText: "OK",
	       cancel: "No"
	     }
	   }).data("kendoConfirm").open();
	   </script>
	*/

	Messages *KendoMessages
}

func (el *KendoUiConfirm) IsSet() bool {
	return el != nil
}
