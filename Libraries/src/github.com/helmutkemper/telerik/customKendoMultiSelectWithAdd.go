package telerik

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
)

type CustomKendoMultiSelectWithAdd struct {
	DockerElementName        string // docker
	ConfigElementName        string // Host
	ConfigItemName           string // EnvVar
	FormName                 string // env
	FormHelpText             string // List of environment variable to set in the container
	FormLabelText            string // Environment Variable
	DialogWindowTitle        string // Environments vars from container
	DialogWindowTextAddClose string // Add and close
	DialogWindowTextAdd      string // Add
	DialogWindowTextClose    string // Close
	DialogErrorTitle         string // Configuration error
	DialogErrorText          string // Please, select a valid command.
	FooterTemplateButtonText string // Add new enviromment var
	TemplateSelectText       string // Environment var:
	NoDataTemplateText       string // No environments vars found to add. Please add an environment var before continuing.
}

func (el *CustomKendoMultiSelectWithAdd) toTemplate(data string) string {
	var output bytes.Buffer
	htmlTemplate := template.Must(template.New("").Funcs(templateFuncMap).Parse(data))
	err := htmlTemplate.Execute(&output, el)
	if err != nil {
		fmt.Println(err.Error())
	}
	return output.String()
}

func (el *CustomKendoMultiSelectWithAdd) ToGlobalVars() string {
	return el.toTemplate(el.getTemplateFromGlobalVars())
}

func (el *CustomKendoMultiSelectWithAdd) ToInitializationVars() string {
	return el.toTemplate(el.getTemplateFromInitialization())
}

func (el *CustomKendoMultiSelectWithAdd) ToFunctions(typeTemplate string) string {
	return el.toTemplate(el.getTemplateFromFunctions(typeTemplate))
}

func (el *CustomKendoMultiSelectWithAdd) ToKendoTemplates(typeTemplate string) string {
	var ret string

	if typeTemplate == "String" {
		ret = el.toTemplate(el.getTemplateFromKendoTemplateForStringInput())
	} else if typeTemplate == "MapStringString" {
		ret = el.toTemplate(el.getTemplateFromKendoTemplateForMapStringStringInput())
	}

	expReg := regexp.MustCompile(`(>\s*)(")`)
	ret = expReg.ReplaceAllString(ret, "$1")

	expReg = regexp.MustCompile(`(")(\s*<)`)
	ret = expReg.ReplaceAllString(ret, "$2")
	//  $1     $2   3    4   5   6   7      8
	expReg = regexp.MustCompile(`(<span>)(.*?:)(")( #= )(")(.*?)(")( #</span>)`)
	ret = expReg.ReplaceAllString(ret, "$1$2$4$6$8")

	return ret
}

func (el *CustomKendoMultiSelectWithAdd) ToDiv() string {
	return el.toTemplate(el.getTemplateFromDivContent())
}

func (el *CustomKendoMultiSelectWithAdd) ToForm() string {
	return el.toTemplate(el.getTemplateFromElementFormContent())
}

func (el *CustomKendoMultiSelectWithAdd) getTemplateFromGlobalVars() string {
	return `// Custom kendo multi select with add - automatically generated code - start
var {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsCount = -1;
var {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsIdToAdd = -1;
var {{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef;
var {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Ref;
var {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}DialogWindowRef;
// Custom kendo multi select with add - automatically generated code - end`
}

func (el *CustomKendoMultiSelectWithAdd) getTemplateFromInitialization() string {
	return `// Custom kendo multi select with add - automatically generated code - start
{{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef = $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput");
{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}DialogWindowRef = $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}");
$("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}kendoMultiSelect").kendoMultiSelect({
  filter: "contains",
  placeholder: "",
  itemTemplate: kendo.template($("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Template").html()),
  noDataTemplate: kendo.template($("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}NoDataTemplate").html()),
  footerTemplate: kendo.template($("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}FooterTemplate").html()),
  dataTextField: "{{.ConfigItemName}}",
  dataValueField: "Id",
  dataSource: new kendo.data.DataSource({
    schema: {
	  model: {
	    id: "Id",
	    fields: {
		  Id: { type: "number" },
  		  {{.ConfigItemName}}: { type: "string" }
	    }
	  }
    }
  })
});
{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Ref = $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}kendoMultiSelect").data("kendoMultiSelect");
// Custom kendo multi select with add - automatically generated code - end`
}

func (el *CustomKendoMultiSelectWithAdd) getTemplateFromFunctions(typeTemplate string) string {
	if typeTemplate == "String" {
		return `// Custom kendo multi select with add - automatically generated code - start
function {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}(){
  {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}DialogWindowRef.kendoDialog({
	  title: "{{.DialogWindowTitle}}",
  	content: kendo.template($("#{{.DockerElementName}}CreateTemplate{{.ConfigItemName}}AddNew{{.ConfigItemName}}").html()),
	  visible: false,
  	modal: true,
	  open: function(){
	    {{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef = $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput");

	    setTimeout( function(){ {{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef.focus(); }, 1000)
	  },
    actions: [
	    {
		    text: "{{.DialogWindowTextClose}}"
  	  },
	    {
		    text: "{{.DialogWindowTextAdd}}",
		    action: function(e){
		      {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}Function();
  		    return false;
	  	  }
	    },
  	  {
	  	  text: "{{.DialogWindowTextAddClose}}",
		    action: function(e){
		      {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}Function();
  		  },
	  	  primary: true
	    }
	  ]
  });
  {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}DialogWindowRef.data("kendoDialog").open();
}

function {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}Function(){
  if( {{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef.val() == "" ){
	  $("#dialog").kendoDialog({
	    modal: true,
  	  visible: true,
	    title: "{{.DialogErrorTitle}}",
	    content: "{{.DialogErrorText}}",
  	  width: 400,
	    actions: [
		    { text: "OK", primary: true, action: function(){ setTimeout( function(){ {{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef.focus(); }, 1000); } }
	    ]
  	});
	  return;
  }

  let dataSource = {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Ref.dataSource;
  {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsCount += 1;
  {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsIdToAdd = {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsCount;

  let dataInDataSource = dataSource.data();
  let pass = true;

  for (var i in dataInDataSource) {
	  if( isNaN( i ) ){
	    break
  	}
	  i = parseInt(i);
	  if (dataInDataSource[i].{{.ConfigItemName}} === $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput").val()) {
	    {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsIdToAdd = dataInDataSource[i].Id;
  	  pass = false;
	    break;
	  }
  }

  if( pass === true ) {
	  dataSource.add({
	    Id: {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsCount,
  	  {{.ConfigItemName}}: $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput").val()
	  });
  }

  dataSource.one("requestEnd", function(args) {
  	if (args.type !== "create") {
	    return;
  	}

	  dataSource.one("sync", function() {
	    {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Ref.value({{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Ref.value().concat([{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsIdToAdd]));
  	});

	  $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput").val("");
  });

  dataSource.sync();
}
// Custom kendo multi select with add - automatically generated code - end`
	} else if typeTemplate == "MapStringString" {
		return `// Custom kendo multi select with add - automatically generated code - start
function {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}(){
  {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}DialogWindowRef.kendoDialog({
	  title: "{{.DialogWindowTitle}}",
  	content: kendo.template($("#{{.DockerElementName}}CreateTemplate{{.ConfigItemName}}AddNew{{.ConfigItemName}}").html()),
	  visible: false,
  	modal: true,
	  open: function(){
	    {{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef = $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput");

	    setTimeout( function(){ {{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef.focus(); }, 1000)
	  },
    actions: [
	    {
		    text: "{{.DialogWindowTextClose}}"
  	  },
	    {
		    text: "{{.DialogWindowTextAdd}}",
		    action: function(e){
		      {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}Function();
  		    return false;
	  	  }
	    },
  	  {
	  	  text: "{{.DialogWindowTextAddClose}}",
		    action: function(e){
		      {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}Function();
  		  },
	  	  primary: true
	    }
	  ]
  });
  {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}DialogWindowRef.data("kendoDialog").open();
}

function {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}Function(){
  if( {{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef.val() == "" ){
	  $("#dialog").kendoDialog({
	    modal: true,
  	  visible: true,
	    title: "{{.DialogErrorTitle}}",
	    content: "{{.DialogErrorText}}",
  	  width: 400,
	    actions: [
		    { text: "OK", primary: true, action: function(){ setTimeout( function(){ {{.DockerElementName}}Configuration{{.ConfigItemName}}NameRef.focus(); }, 1000); } }
	    ]
  	});
	  return;
  }

  let dataSource = {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Ref.dataSource;
  {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsCount += 1;
  {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsIdToAdd = {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsCount;

  let dataInDataSource = dataSource.data();
  let pass = true;

  for (var i in dataInDataSource) {
	  if( isNaN( i ) ){
	    break
  	}
	  i = parseInt(i);
	  if (dataInDataSource[i].{{.ConfigItemName}} === $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput").val()) {
	    {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsIdToAdd = dataInDataSource[i].Id;
  	  pass = false;
	    break;
	  }
  }

  if( pass === true ) {
	  dataSource.add({
	    Id: {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsCount,
  	  {{.ConfigItemName}}: $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput").val()
	  });
  }

  dataSource.one("requestEnd", function(args) {
  	if (args.type !== "create") {
	    return;
  	}

	  dataSource.one("sync", function() {
	    {{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Ref.value({{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Ref.value().concat([{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ItemsIdToAdd]));
  	});

	  $("#{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput").val("");
  });

  dataSource.sync();
}
// Custom kendo multi select with add - automatically generated code - end`
	}

	return ``
}

func (el *CustomKendoMultiSelectWithAdd) getTemplateFromKendoTemplateForStringInput() string {
	return `<script id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}FooterTemplate" type="text/x-kendo-tmpl">
  <button class="k-button k-primary centerText" onclick="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}()">{{.FooterTemplateButtonText}}</button>
</script>
<script id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Template" type="text/x-kendo-tmpl">
  <span>{{.TemplateSelectText}} #= {{.ConfigItemName}} #</span>
</script>
<script id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}NoDataTemplate" type="text/x-kendo-tmpl">
  <div>
    {{.NoDataTemplateText}}
  </div>
</script>
<script type="text/x-kendo-template" id="{{.DockerElementName}}CreateTemplate{{.ConfigItemName}}AddNew{{.ConfigItemName}}">
  <span>
    <input type="text" id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}TextInput" name="{{.FormName}}" class="k-textbox" placeholder="" required validationMessage="Enter a {0}" autocomplete="off" />
  </span>
</script>`
}

func (el *CustomKendoMultiSelectWithAdd) getTemplateFromKendoTemplateForMapStringStringInput() string {
	return `<script id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}FooterTemplate" type="text/x-kendo-tmpl">
  <button class="k-button k-primary centerText" onclick="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}()">{{.FooterTemplateButtonText}}</button>
</script>
<script id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}Template" type="text/x-kendo-tmpl">
  <span>{{.TemplateSelectText}} #= {{.ConfigItemName}} #</span>
</script>
<script id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}NoDataTemplate" type="text/x-kendo-tmpl">
  <div>
    {{.NoDataTemplateText}}
  </div>
</script>
<script type="text/x-kendo-template" id="{{.DockerElementName}}CreateTemplate{{.ConfigItemName}}AddNew{{.ConfigItemName}}">
  <span>
    <input type="text" id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}KeyTextInput" name="{{.FormName}}Key" class="k-textbox" placeholder="key" required validationMessage="Enter a {0}" autocomplete="off" />
    <input type="text" id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}ValueTextInput" name="{{.FormName}}Value" class="k-textbox" placeholder="value" required validationMessage="Enter a {0}" autocomplete="off" />
  </span>
</script>`
}

func (el *CustomKendoMultiSelectWithAdd) getTemplateFromDivContent() string {
	return `<!-- Custom kendo multi select with add - automatically generated code - start -->
<div id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}AddNew{{.ConfigItemName}}"></div>
<!-- Custom kendo multi select with add - automatically generated code - end -->`
}

func (el *CustomKendoMultiSelectWithAdd) getTemplateFromElementFormContent() string {
	return `<!-- Custom kendo multi select with add - automatically generated code - start -->
<span>
  <label for="{{.FormName}}" title="{{.FormHelpText}}">{{.FormLabelText}}</label>
  <select id="{{.DockerElementName}}{{.ConfigElementName}}{{.ConfigItemName}}kendoMultiSelect" name="{{.FormName}}" autocomplete="off"></select>
</span>
<!-- Custom kendo multi select with add - automatically generated code - end -->`
}
