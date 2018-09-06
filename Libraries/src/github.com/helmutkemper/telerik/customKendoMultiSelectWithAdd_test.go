package telerik

import "fmt"

var multiSelectWithAdd CustomKendoMultiSelectWithAdd

func ExampleCustomKendoMultiSelectWithAdd_ToDiv() {
	fmt.Print(multiSelectWithAdd.ToDiv())

	// Output:
	//
}

func ExampleCustomKendoMultiSelectWithAdd_ToForm() {
	fmt.Print(multiSelectWithAdd.ToForm())

	// Output:
	//
}

func ExampleCustomKendoMultiSelectWithAdd_ToFunctions() {
	fmt.Print(multiSelectWithAdd.ToFunctions("String"))

	// Output:
	//
}

func ExampleCustomKendoMultiSelectWithAdd_ToGlobalVars() {
	fmt.Print(multiSelectWithAdd.ToGlobalVars())

	// Output:
	//
}

func ExampleCustomKendoMultiSelectWithAdd_ToInitializationVars() {
	fmt.Print(multiSelectWithAdd.ToInitializationVars())

	// Output:
	//
}

func ExampleCustomKendoMultiSelectWithAdd_ToKendoTemplates() {
	fmt.Print(multiSelectWithAdd.ToKendoTemplates("String"))

	// Output:
	//
}

func init() {
	multiSelectWithAdd = CustomKendoMultiSelectWithAdd{
		DockerElementName:        "docker",
		ConfigElementName:        "Host",
		ConfigItemName:           "OnBuild",
		FormName:                 "onbuild",
		FormHelpText:             "ONBUILD metadata that were defined on the image Dockerfile",
		FormLabelText:            "On Build",
		DialogWindowTitle:        "On Build metadata from container",
		DialogWindowTextAddClose: "Add and close",
		DialogWindowTextAdd:      "Add",
		DialogWindowTextClose:    "Close",
		DialogErrorTitle:         "Configuration error",
		DialogErrorText:          "Please, select a valid on build metadata.",
		FooterTemplateButtonText: "Add new on build metadata",
		TemplateSelectText:       "on build metadata:",
		NoDataTemplateText:       "No on build metadata found to add. Please add a on build metadata before continuing.",
	}
}
