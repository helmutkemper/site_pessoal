package telerik

import "fmt"

func ExampleFrameworkForm_ToForm() {
	f := FrameworkForm{
		Template: `<ul>{{range $k, $v := .}}<li>{{$v.Label | safeHTML}}{{$v.Input | safeHTML}}</li>{{end}}</ul>`,
		Form: HtmlElementForm{
			Action: "./cadastrar",
			Method: "get",
		},
		Elements: []FrameworkInput{
			{
				Id:                "nome",
				Label:             "Nome:",
				ValidationMessage: "Por favor, digite o seu nome completo.",
				Title:             "Digite o seu nome",
				PlaceHolder:       "Nome completo",
				Input:             HtmlInputText{},
			},
			{
				Id:                "endereco",
				Label:             "Endereço:",
				ValidationMessage: "Por favor, digite seu nome completo.",
				PlaceHolder:       "Endereço",
				Title:             "Digite o seu endereço",
				Input:             KendoUiAutoComplete{
					//DataSource: []string{"Albania","Andorra","Armenia","Austria","Azerbaijan","Belarus","Belgium","Bosnia & Herzegovina","Bulgaria","Croatia","Cyprus","Czech Republic","Denmark","Estonia","Finland","France","Georgia","Germany","Greece","Hungary","Iceland","Ireland","Italy","Kosovo","Latvia","Liechtenstein","Lithuania","Luxembourg","Macedonia","Malta","Moldova","Monaco","Montenegro","Netherlands","Norway","Poland","Portugal","Romania","Russia","San Marino","Serbia","Slovakia","Slovenia","Spain","Sweden","Switzerland","Turkey","Ukraine","United Kingdom","Vatican City"},
				},
			},
		},
	}

	fmt.Printf("%v\n", f.ToForm())

	// Output:
	// <form  action="./cadastrar"  method="get" ><ul><li><label  class=" required"  title="Digite o seu nome"  for="nome" >Nome:</label><input  class="k-textbox"  id="nome" validationMessage="Por favor, digite o seu nome completo."  type="text"  name="nome"  placeholder="Nome completo"  required ></li><li><label  class=" required"  title="Digite o seu endereço"  for="endereco" >Endereço:</label><input  class="k-textbox"  id="endereco" validationMessage="Por favor, digite seu nome completo."  type="search"  name="endereco"  placeholder="Endereço"  required ></li></ul></form>
}

func ExampleFrameworkForm_ToScript() {
	f := FrameworkForm{
		Template: `<ul>{{range $k, $v := .}}<li>{{$v.Label | safeHTML}}{{$v.Input | safeHTML}}</li>{{end}}</ul>`,
		Form: HtmlElementForm{
			Action: "./cadastrar",
			Method: "get",
		},
		Elements: []FrameworkInput{
			{
				Id:                "nome",
				Label:             "Nome:",
				ValidationMessage: "Por favor, digite o seu nome completo.",
				Title:             "Digite o seu nome",
				PlaceHolder:       "Nome completo",
				Input:             HtmlInputText{},
			},
			{
				Id:                "endereco",
				Label:             "Endereço:",
				ValidationMessage: "Por favor, digite seu nome completo.",
				PlaceHolder:       "Endereço",
				Title:             "Digite o seu endereço",
				Input:             KendoUiAutoComplete{
					//DataSource: []string{"Albania","Andorra","Armenia","Austria","Azerbaijan","Belarus","Belgium","Bosnia & Herzegovina","Bulgaria","Croatia","Cyprus","Czech Republic","Denmark","Estonia","Finland","France","Georgia","Germany","Greece","Hungary","Iceland","Ireland","Italy","Kosovo","Latvia","Liechtenstein","Lithuania","Luxembourg","Macedonia","Malta","Moldova","Monaco","Montenegro","Netherlands","Norway","Poland","Portugal","Romania","Russia","San Marino","Serbia","Slovakia","Slovenia","Spain","Sweden","Switzerland","Turkey","Ukraine","United Kingdom","Vatican City"},
				},
			},
		},
	}

	fmt.Printf("<script>%v</script>\n", f.ToScript())

	// Output:
	// <script>$("#endereco").kendoAutoComplete({ dataSource: ["Albania","Andorra","Armenia","Austria","Azerbaijan","Belarus","Belgium","Bosnia & Herzegovina","Bulgaria","Croatia","Cyprus","Czech Republic","Denmark","Estonia","Finland","France","Georgia","Germany","Greece","Hungary","Iceland","Ireland","Italy","Kosovo","Latvia","Liechtenstein","Lithuania","Luxembourg","Macedonia","Malta","Moldova","Monaco","Montenegro","Netherlands","Norway","Poland","Portugal","Romania","Russia","San Marino","Serbia","Slovakia","Slovenia","Spain","Sweden","Switzerland","Turkey","Ukraine","United Kingdom","Vatican City"], });</script>
}
