package telerik

import "fmt"

func ExampleKendoUiColorPalette_ToJavaScript() {
	color := KendoUiColorPalette{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "palette",
			},
		},
		Columns: 4,
		TileSize: KendoTileSize{
			Width:  34,
			Height: 19,
		},
		Palette: []string{
			"#f0d0c9", "#e2a293", "#d4735e", "#65281a",
			"#eddfda", "#dcc0b6", "#cba092", "#7b4b3a",
			"#fcecd5", "#f9d9ab", "#f6c781", "#c87d0e",
			"#e1dca5", "#d0c974", "#a29a36", "#514d1b",
			"#c6d9f0", "#8db3e2", "#548dd4", "#17365d",
		},
	}

	fmt.Printf(`<!DOCTYPE html>
<html>
<head>
<base href="https://demos.telerik.com/kendo-ui/colorpicker/index">
<style>html { font-size: 14px; font-family: Arial, Helvetica, sans-serif; }</style>
<title></title>
<link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.default-v2.min.css" />
<script src="https://kendo.cdn.telerik.com/2018.1.221/js/jquery.min.js"></script>
<script src="https://kendo.cdn.telerik.com/2018.1.221/js/kendo.all.min.js"></script>
</head>
<body>
<div id="example">
<div class="demo-section hidden-on-narrow k-content wide">
<div id="background">
%s
</div>
</div>
</div>
<script>
%s
</script>
</body>
</html>
`, color.ToHtml(), color.ToJavaScript())

	//Output:
	//<!DOCTYPE html>
	//<html>
	//<head>
	//<base href="https://demos.telerik.com/kendo-ui/colorpicker/index">
	//<style>html { font-size: 14px; font-family: Arial, Helvetica, sans-serif; }</style>
	//<title></title>
	//<link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.default-v2.min.css" />
	//<script src="https://kendo.cdn.telerik.com/2018.1.221/js/jquery.min.js"></script>
	//<script src="https://kendo.cdn.telerik.com/2018.1.221/js/kendo.all.min.js"></script>
	//</head>
	//<body>
	//<div id="example">
	//<div class="demo-section hidden-on-narrow k-content wide">
	//<div id="background">
	//<div id="palette"></div>
	//</div>
	//</div>
	//</div>
	//<script>
	//$("#palette").kendoColorPalette({palette: ["#f0d0c9","#e2a293","#d4735e","#65281a","#eddfda","#dcc0b6","#cba092","#7b4b3a","#fcecd5","#f9d9ab","#f6c781","#c87d0e","#e1dca5","#d0c974","#a29a36","#514d1b","#c6d9f0","#8db3e2","#548dd4","#17365d",],columns: 4,tileSize: { width: 34,height: 19,},});
	//</script>
	//</body>
	//</html>
}
