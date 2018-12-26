package main

import (
	"encoding/json"
	"errors"
	"fmt"
	rpx "github.com/helmutkemper/SimpleReverseProxy"
	js "github.com/helmutkemper/javascript"
	"github.com/helmutkemper/mgo/bson"
	"github.com/helmutkemper/mongodb"
	"github.com/helmutkemper/safeType"
	log "github.com/helmutkemper/seelog"
	"github.com/helmutkemper/telerik"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

var JsQueryListOutput safeType.SafeMapStringInterface
var JsUpdateQueryListOutput safeType.SafeMapStringInterface

func addRouteToGetData(name, collection string, query bson.M, regExp string, selectOut bson.M) {

	JsQueryListOutput.Set(name, map[string]interface{}{
		"collection": collection,
		"query":      query,
		"regexp":     regExp,
		"select":     selectOut,
	})
}

func addRouteToUpdateData(name, collection string, query bson.M, regExp string, selectOut bson.M) {

	JsUpdateQueryListOutput.Set(name, map[string]interface{}{
		"collection": collection,
		"query":      query,
		"regexp":     regExp,
		"select":     selectOut,
	})
}

func runJsCodeOnLoad() {

	var codeName = "onLoad"
	var query = mongodb.QueryStt{
		Limit: 1,
		Query: bson.M{
			"name": codeName,
		},
	}

	var db = mongodb.DbStt{}
	db.Init("main")
	db.Collection("bigData_javascriptCode")
	db.Find(&query)

	var jsData = db.DataGetAll()
	if len(jsData) == 0 {
		log.Error("runJsCodeOnLoad() not found code 'onLoad' in collection 'bigData_javascriptCode'")
		return
	}

	if jsData[0]["code"].(string) == "" {
		log.Error("runJsCodeOnLoad() found code 'onLoad', but, content is empty")
		return
	}

	err := js.RunJavaScriptCode(codeName, jsData[0]["code"].(string))
	if err != nil {
		log.Errorf("JavaScript error: %s", err.Error())
	}
}

func initCodeJsOnDb(force bool) {
	var query = mongodb.QueryStt{
		Query: bson.M{},
	}
	var db = mongodb.DbStt{}
	db.Init("main")
	db.Collection("bigData_javascriptCode")
	db.Find(&query)

	var jsData = db.DataGetAll()
	var passOnLoad = false
	for _, codeData := range jsData {
		if codeData["name"] == "onLoad" {
			passOnLoad = true
		}
	}

	if passOnLoad == true && force == true {
		query.Query = bson.M{"name": "onLoad"}
		db.Remove(&query)
	}

	if passOnLoad == false || force == true {
		dataToInsert := map[string]interface{}{
			"name": "onLoad",
			"code": "",
		}
		db.Insert(dataToInsert)
	}

}

func reloadOnLoad(w rpx.ProxyResponseWriter, r *rpx.ProxyRequest) {
	js.CronSupport.Stop()
	runJsCodeOnLoad()
	js.CronSupport.Start()

	output := rpx.JSonOutStt{}
	output.ToOutput(1, nil, []int{}, w)
}

func resetJsCode(w rpx.ProxyResponseWriter, r *rpx.ProxyRequest) {
	js.CronSupport.Stop()
	initCodeJsOnDb(true)
	runJsCodeOnLoad()
	js.CronSupport.Start()

	output := rpx.JSonOutStt{}
	output.ToOutput(1, nil, []int{}, w)
}

func runJsCode(w rpx.ProxyResponseWriter, r *rpx.ProxyRequest) {

	output := rpx.JSonOutStt{}

	var err error
	var query = mongodb.QueryStt{
		Limit: 1,
		Query: bson.M{
			"name": r.ExpRegMatches["codeName"],
		},
	}

	var db = mongodb.DbStt{}
	db.Init("main")
	db.Collection("bigData_javascriptCode")
	db.Find(&query)

	var jsData = db.DataGetAll()

	if len(jsData) == 0 {
		output.ToOutput(1, errors.New(r.ExpRegMatches["codeName"]+": code not found"), []int{}, w)
		return
	}

	if len(jsData) == 0 {
		output.ToOutput(1, errors.New(r.ExpRegMatches["codeName"]+": code not found"), []int{}, w)
		return
	}

	if jsData[0]["code"] == nil {
		output.ToOutput(1, errors.New(r.ExpRegMatches["codeName"]+": code found, but, content is empty"), []int{}, w)
		return
	}

	err = js.RunJavaScriptCode(r.ExpRegMatches["codeName"], jsData[0]["code"].(string))

	if err != nil {
		output.ToOutput(1, err.Error(), []int{}, w)
		return
	}

	output.ToOutput(1, nil, []int{}, w)
}

func getData(w rpx.ProxyResponseWriter, r *rpx.ProxyRequest) {

	output := rpx.JSonOutStt{}

	jsQueryListData := JsQueryListOutput.GetKey(r.ExpRegMatches["queryName"])
	if jsQueryListData != nil {
		if jsQueryListData.(map[string]interface{})["regexp"].(string) != "" {
			matched, err := regexp.MatchString(jsQueryListData.(map[string]interface{})["regexp"].(string), r.R.URL.Path)
			if err != nil {
				log.Criticalf("getData.regexp.MatchString.error: %v", err.Error())
				output.ToOutput(1, err, []int{}, w)
				return
			}
			if matched == true {
				re := regexp.MustCompile(jsQueryListData.(map[string]interface{})["regexp"].(string))
				for k, v := range re.SubexpNames() {
					if k == 0 || v == "" {
						continue
					}

					r.ExpRegMatches[v] = re.ReplaceAllString(r.R.URL.Path, `${`+v+`}`)
				}
			} else {
				output.ToOutput(1, errors.New("page not found"), []int{}, w)
				return
			}
		}
	}

	var db = mongodb.DbStt{}
	db.Init("main")
	db.Collection(jsQueryListData.(map[string]interface{})["collection"].(string))

	var query = mongodb.QueryStt{
		Query: db.ParserQuery(jsQueryListData.(map[string]interface{})["query"], r.ExpRegMatches),
	}

	if len(jsQueryListData.(map[string]interface{})["select"].(bson.M)) != 0 {
		query.Select = jsQueryListData.(map[string]interface{})["select"].(bson.M)
	}

	db.Find(&query)

	var jsData = db.DataGetAll()

	if len(jsData) == 0 {
		output.ToOutput(1, errors.New(r.ExpRegMatches["codeName"]+": code not found"), []int{}, w)
		return
	}

	out := db.DataGetAll()

	output.ToOutput(len(out), nil, out, w)
}

func updateData(w rpx.ProxyResponseWriter, r *rpx.ProxyRequest) {

	output := rpx.JSonOutStt{}

	raw, err := ioutil.ReadAll(r.R.Body)
	if err != nil {
		log.Criticalf("updateData.ioutil.ReadAll.error: %v\n", err.Error())
	}
	rawMap := make(map[string]interface{})

	err = json.Unmarshal(raw, &rawMap)
	if err != nil {
		log.Criticalf("updateData.json.Unmarshal.error: %v\n", err.Error())
		output.ToOutput(1, err, []int{}, w)
		return
	}

	jsQueryListData := JsQueryListOutput.GetKey(r.ExpRegMatches["queryName"])

	if jsQueryListData.(map[string]interface{})["regexp"].(string) != "" {
		matched, err := regexp.MatchString(jsQueryListData.(map[string]interface{})["regexp"].(string), r.R.URL.Path)
		if err != nil {
			log.Criticalf("updateData.regexp.MatchString.error: %v\n", err.Error())
			output.ToOutput(1, err, []int{}, w)
			return
		}
		if matched == true {
			re := regexp.MustCompile(jsQueryListData.(map[string]interface{})["regexp"].(string))
			for k, v := range re.SubexpNames() {
				if k == 0 || v == "" {
					continue
				}

				r.ExpRegMatches[v] = re.ReplaceAllString(r.R.URL.Path, `${`+v+`}`)
			}
		} else {
			output.ToOutput(1, errors.New("page not found"), []int{}, w)
			return
		}
	}

	var db = mongodb.DbStt{}
	db.Init("main")
	db.Collection(jsQueryListData.(map[string]interface{})["collection"].(string))

	var query = mongodb.QueryStt{
		Query:  db.ParserQuery(jsQueryListData.(map[string]interface{})["query"], r.ExpRegMatches),
		Update: rawMap,
	}

	if len(jsQueryListData.(map[string]interface{})["select"].(bson.M)) != 0 {
		query.Select = jsQueryListData.(map[string]interface{})["select"].(bson.M)
	}

	if db.Count(&query) == 0 {
		db.Insert(query.Update)
	} else {
		db.Update(&query)
	}

	db.Find(&query)

	var jsData = db.DataGetAll()

	if len(jsData) == 0 {
		output.ToOutput(1, errors.New(r.ExpRegMatches["codeName"]+": code not found"), []int{}, w)
		return
	}

	out := db.DataGetAll()

	output.ToOutput(len(out), nil, out, w)
}

type Article struct {
	Title   string
	Content string
}

type telerikJs struct {
	ScriptTemplate string
	VarGlobal      string
	OnLoadCode     string
	HtmlSupport    string
}

type Page struct {
	TemplateName string

	Telerik  telerikJs
	Title    string
	Articles []Article
}

func tplFunc(w rpx.ProxyResponseWriter, r *rpx.ProxyRequest) {
	data := map[string]interface{}{
		"PageTitle":     "gOsm",
		"MainTitle":     "Documentation",
		"Title":         "gOsm server",
		"Text":          "<p>É claro que a complexidade dos estudos efetuados talvez venha a ressaltar a relatividade das condições financeiras e administrativas exigidas. O incentivo ao avanço tecnológico, assim como a constante divulgação das informações garante a contribuição de um grupo importante na determinação das diretrizes de desenvolvimento para o futuro. Por outro lado, a contínua expansão de nossa atividade cumpre um papel essencial na formulação de todos os recursos funcionais envolvidos. No mundo atual, a estrutura atual da organização ainda não demonstrou convincentemente que vai participar na mudança das posturas dos órgãos dirigentes com relação às suas atribuições.</p><p>O que temos que ter sempre em mente é que a crescente influência da mídia acarreta um processo de reformulação e modernização das novas proposições. A prática cotidiana prova que o fenômeno da Internet oferece uma interessante oportunidade para verificação dos paradigmas corporativos. Todas estas questões, devidamente ponderadas, levantam dúvidas sobre se o surgimento do comércio virtual causa impacto indireto na reavaliação do sistema de formação de quadros que corresponde às necessidades. Caros amigos, o julgamento imparcial das eventualidades obstaculiza a apreciação da importância dos conhecimentos estratégicos para atingir a excelência.</p><p>Acima de tudo, é fundamental ressaltar que a percepção das dificuldades estende o alcance e a importância das direções preferenciais no sentido do progresso. A certificação de metodologias que nos auxiliam a lidar com o acompanhamento das preferências de consumo pode nos levar a considerar a reestruturação das formas de ação. Não obstante, a consolidação das estruturas maximiza as possibilidades por conta do fluxo de informações. Todavia, o novo modelo estrutural aqui preconizado prepara-nos para enfrentar situações atípicas decorrentes dos procedimentos normalmente adotados.</p><p>Percebemos, cada vez mais, que o comprometimento entre as equipes auxilia a preparação e a composição das condições inegavelmente apropriadas. Do mesmo modo, a adoção de políticas descentralizadoras estimula a padronização do processo de comunicação como um todo. Ainda assim, existem dúvidas a respeito de como a mobilidade dos capitais internacionais faz parte de um processo de gerenciamento do sistema de participação geral. Gostaria de enfatizar que o aumento do diálogo entre os diferentes setores produtivos agrega valor ao estabelecimento dos índices pretendidos.</p><p>É importante questionar o quanto a revolução dos costumes facilita a criação da gestão inovadora da qual fazemos parte. A nível organizacional, a necessidade de renovação processual promove a alavancagem dos modos de operação convencionais. O empenho em analisar o consenso sobre a necessidade de qualificação nos obriga à análise dos métodos utilizados na avaliação de resultados. Pensando mais a longo prazo, a valorização de fatores subjetivos exige a precisão e a definição de alternativas às soluções ortodoxas.</p><p>Evidentemente, a consulta aos diversos militantes não pode mais se dissociar dos níveis de motivação departamental. Podemos já vislumbrar o modo pelo qual o início da atividade geral de formação de atitudes possibilita uma melhor visão global do orçamento setorial. Nunca é demais lembrar o peso e o significado destes problemas, uma vez que a hegemonia do ambiente político desafia a capacidade de equalização do levantamento das variáveis envolvidas.</p>",
		"CopyrightYear": "2018",
		"CopyrightLink": "",
		"CopyrightName": "Helmut Kemper",
		"RightMenu": []map[string]interface{}{
			{
				"Link":   "#",
				"Text":   "Introduction",
				"Active": true,
				"Icon":   "fe fe-activity",
			},
			{
				"Link":   "#",
				"Text":   "Alerts",
				"Active": false,
				"Icon":   "fe fe-airplay",
			},
			{
				"Link":   "#",
				"Text":   "Demos",
				"Active": false,
				"Icon":   "fe fe-airplay",
			},
		},
		"MainMenu": []map[string]interface{}{
			{
				"Icon":  "fe fe-activity",
				"Label": "Menu 1",
				"SubMenu": []map[string]interface{}{
					{
						"Link":  "#",
						"Label": "um",
					},
					{
						"Link":  "#",
						"Label": "Dois",
					},
					{
						"Link":  "#",
						"Label": "Tres",
					},
				},
			},
			{
				"Icon":  "fe fe-airplay",
				"Label": "Menu 2",
				"SubMenu": []map[string]interface{}{
					{
						"Link":  "#",
						"Label": "quatro",
					},
					{
						"Link":  "#",
						"Label": "cinco",
					},
					{
						"Link":  "#",
						"Label": "seis",
					},
				},
			},
			{
				"Icon":  "fe fe-alert-circle",
				"Label": "Menu 3",
				"SubMenu": []map[string]interface{}{
					{
						"Link":  "#",
						"Label": "sete",
					},
					{
						"Link":  "#",
						"Label": "oito",
					},
					{
						"Link":  "#",
						"Label": "nove",
					},
				},
			},
		},
	}

	var fns = template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
		"nLast": func(x int, a interface{}) bool {
			return x != reflect.ValueOf(a).Len()-1
		},
		"htmlSafe": func(str string) template.HTML {
			return template.HTML(str)
		},
		"js": func(str string) template.JS {
			return template.JS(str)
		},
	}

	tmpl := template.Must(
		template.New("index").Funcs(fns).ParseFiles(
			"./static/site_original_template/docs/index.tmpl",
			"./static/site_original_template/docs/index/mainMenu.tmpl",
			"./static/site_original_template/docs/index/searchForm.tmpl",
			"./static/site_original_template/docs/index/userMenu.tmpl",
			"./static/site_original_template/docs/index/bell.tmpl",
			"./static/site_original_template/docs/index/admin.tmpl",
			"./static/site_original_template/docs/index/rightMenu.tmpl",
		),
	)

	err := tmpl.ExecuteTemplate(w, "index", &data)
	if err != nil {
		fmt.Printf("template.error: %v", err.Error())
	}
}

func memory(w rpx.ProxyResponseWriter, r *rpx.ProxyRequest) {
	var mem runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&mem)

	output := rpx.JSonOutStt{}
	output.ToOutput(1, nil, mem, w)
}

func formJs(w rpx.ProxyResponseWriter, r *rpx.ProxyRequest) {

	page := Page{
		TemplateName: "w3",
		Title:        "live maps",
		Articles:     []Article{},
	}

	var err error

	var fns = template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
		"nLast": func(x int, a interface{}) bool {
			return x != reflect.ValueOf(a).Len()-1
		},
		"htmlSafe": func(str string) template.HTML {
			return template.HTML(str)
		},
		"js": func(str string) template.JS {
			return template.JS(str)
		},
	}

	el := telerik.HtmlElementDiv{
		Global: telerik.HtmlGlobalAttributes{
			Id:    "spanCreateTemplateExposedPortsAddNewPort",
			Class: "k-content",
			//Style: "width: 300px !important;",
		},
		Content: telerik.Content{

			&telerik.HtmlElementDiv{
				Content: telerik.Content{

					&telerik.HtmlElementFormLabel{
						For: "codeName",
						Content: telerik.Content{
							"Code Name",
						},
						Global: telerik.HtmlGlobalAttributes{
							//Style: "width: 200px important!;",
						},
					},

					&telerik.HtmlInputText{
						Global: telerik.HtmlGlobalAttributes{
							Id:    "codeName",
							Class: "k-textbox",
						},
						Name: "codeName",
					},
				},
			},

			&telerik.HtmlElementDiv{
				Content: telerik.Content{

					&telerik.HtmlElementFormLabel{
						For: "ConfigHostName",
						Content: telerik.Content{
							"Code",
						},
						Global: telerik.HtmlGlobalAttributes{
							//Style: "width: 200px important!;",
						},
					},

					&telerik.AceEditor{
						Html: telerik.HtmlElementDiv{
							Name: "editor",
							Global: telerik.HtmlGlobalAttributes{
								Id:    "editor",
								Class: "",
							},
						},
						FontSize: 16,
						Mode:     telerik.ACE_MODE_JAVASCRIPT,
						Theme:    telerik.ACE_THEME_BRIGHT_CHROME,
						WarpMode: telerik.FALSE,
						TabSize:  4,
						TabSoft:  telerik.TRUE,
					},
				},
			},

			&telerik.HtmlElementDiv{
				Content: telerik.Content{

					&telerik.KendoUiButton{
						Html: telerik.HtmlElementFormButton{
							Global: telerik.HtmlGlobalAttributes{
								Id: "submitButton",
							},
							Content: telerik.Content{
								"Save",
							},
						},
					},
				},
			},
		},
	}

	page.Articles = make([]Article, 1)

	page.Telerik.ScriptTemplate = string(el.Content.MakeJsScript())
	page.Telerik.VarGlobal = string(el.Content.MakeJsObject())
	page.Telerik.OnLoadCode = string(el.Content.ToJavaScript())
	page.Telerik.HtmlSupport = string(el.ToHtmlSupport())

	page.Articles[0].Title = "Edit"
	page.Articles[0].Content = string(el.ToHtml())

	fileList := make([]string, 0)
	err = filepath.Walk("./static/template/livemaps/", func(path string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			if strings.HasSuffix(path, ".tpl") || strings.HasSuffix(path, ".htm") || strings.HasSuffix(path, ".html") {
				fileList = append(fileList, path)
			}
		}
		return nil
	})

	tpl := template.Must(template.New(page.TemplateName).Funcs(fns).ParseFiles(fileList...))
	err = tpl.Execute(w, page)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}

}

func main() {

	var err error

	var db mongodb.MongoDb
	err, _ = db.Connect("127.0.0.1:27017", "gOsm", "main")
	if err != nil {
		log.Criticalf("db.Connect.error: %v", err.Error())
		os.Exit(1)
	}

	//js.ExternalFunc.Set( "makeGrid", makeGrid )
	js.ExternalFunc.Set("addRouteToGetData", addRouteToGetData)
	js.ExternalFunc.Set("addRouteToUpdateData", addRouteToUpdateData)

	// ********************************************************
	// ********************************************************
	// ********************************************************
	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "run javascript code",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "",
			},
			Path: rpx.ProxyPath{
				Path:   "/admin",
				Method: "GET",
				ExpReg: "^/admin/reloadConfig$",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: reloadOnLoad,
			},
		},
	)
	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "run javascript code",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "",
			},
			Path: rpx.ProxyPath{
				Path:   "/customJs",
				Method: "GET",
				ExpReg: "^/customJs/(?P<codeName>[a-zA-Z0-9_]+)$",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: runJsCode,
			},
		},
	)
	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "run query",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "",
			},
			Path: rpx.ProxyPath{
				Path:   "/getData",
				Method: "GET",
				ExpReg: "^/getData/(?P<queryName>[a-zA-Z0-9_]+).*",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: getData,
			},
		},
	)
	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "update query",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "",
			},
			Path: rpx.ProxyPath{
				Path:   "/updateData",
				Method: "POST",
				ExpReg: "^/updateData/(?P<queryName>[a-zA-Z0-9_]+).*",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: updateData,
			},
		},
	)

	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "api",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "",
			},
			Path: rpx.ProxyPath{
				Path:   "/editJs",
				Method: "GET",
				ExpReg: "^/editJs$",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: formJs,
			},
		},
	)
	// ********************************************************
	// ********************************************************
	// ********************************************************

	// Cria a roda para código JavaScript
	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "reset javascript code",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
			},
			Path: rpx.ProxyPath{
				Path:   "/admin",
				Method: "GET",
				ExpReg: "^/admin/resetJsCode$",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: resetJsCode,
			},
		},
	)
	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "run javascript code",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "bigdata-go-ahgora.apps.ahgoracloud.com.br",
			},
			Path: rpx.ProxyPath{
				Path:   "/admin",
				Method: "GET",
				ExpReg: "^/admin/reloadConfig$",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: reloadOnLoad,
			},
		},
	)
	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "run javascript code",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "bigdata-go-ahgora.apps.ahgoracloud.com.br",
			},
			Path: rpx.ProxyPath{
				Path:   "/customJs",
				Method: "GET",
				ExpReg: "^/customJs/(?P<codeName>[a-zA-Z0-9_]+)$",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: runJsCode,
			},
		},
	)
	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "run query",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "bigdata-go-ahgora.apps.ahgoracloud.com.br",
			},
			Path: rpx.ProxyPath{
				Path:   "/getData",
				Method: "GET",
				ExpReg: "^/getData/(?P<queryName>[a-zA-Z0-9_]+).*",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: getData,
			},
		},
	)
	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "update query",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "bigdata-go-ahgora.apps.ahgoracloud.com.br",
			},
			Path: rpx.ProxyPath{
				Path:   "/updateData",
				Method: "POST",
				ExpReg: "^/updateData/(?P<queryName>[a-zA-Z0-9_]+).*",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: updateData,
			},
		},
	)

	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "api",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "bigdata-go-ahgora.apps.ahgoracloud.com.br",
			},
			Path: rpx.ProxyPath{
				Path:   "/test",
				Method: "GET",
				ExpReg: "^/test$",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: formJs,
			},
		},
	)

	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "memory",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "",
			},
			Path: rpx.ProxyPath{
				Path:   "/memory",
				Method: "GET",
				ExpReg: "^/memory$",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: memory,
			},
		},
	)

	rpx.ProxyRootConfig.AddRouteFromFuncStt(
		rpx.ProxyRoute{
			Name: "memory",
			Domain: rpx.ProxyDomain{
				NotFoundHandle: rpx.ProxyRootConfig.ProxyNotFound,
				ErrorHandle:    rpx.ProxyRootConfig.ProxyError,
				Host:           "",
			},
			Path: rpx.ProxyPath{
				Path:   "/template",
				Method: "GET",
				ExpReg: "^/template$",
			},
			ProxyEnable: false,
			Handle: rpx.ProxyHandle{
				Handle: tplFunc,
			},
		},
	)

	rpx.ProxyRootConfig.ListenConfig(":8888")
	rpx.ProxyRootConfig.Prepare()

	//js.CronSupport.Start()

	if _, err = os.Stat("./static"); os.IsNotExist(err) {
		err = os.Mkdir("./static", 0755)
		if err != nil {
			log.Errorf("main.os.Mkdir.static.error: %v", err.Error())
		}
	}

	log.Infof("listen and server: %v", rpx.ProxyRootConfig.ListenAndServe)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", rpx.ProxyFunc)
	err = http.ListenAndServe(rpx.ProxyRootConfig.ListenAndServe, nil)
	if err != nil {
		log.Errorf("main.http.ListenAndServe.error: %v", err.Error())
	}
}
