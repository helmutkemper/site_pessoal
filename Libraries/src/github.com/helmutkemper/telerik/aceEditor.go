package telerik

import (
	"bytes"
	"strconv"
)

type AceEditor struct {
	Html                HtmlElementDiv `jsObject:"-"`
	Theme               AceTheme
	Mode                AceMode
	Content             string
	TabSize             int
	TabSoft             Boolean
	FontSize            int
	WarpMode            Boolean
	HighlightActiveLine Boolean
	ShowPrintMargin     Boolean
	ReadOnly            Boolean
}

func (el *AceEditor) ToJavaScript() []byte {
	var ret bytes.Buffer

	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	ret.Write([]byte(` var `))
	ret.WriteString(el.Html.Global.Id)
	ret.Write([]byte(` = `))
	ret.Write([]byte(`ace.edit("` + el.Html.Global.Id + `");`))

	if el.Theme != 0 {
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`.setTheme("` + el.Theme.String() + `");`))
	}

	if el.Mode != 0 {
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`.session.setMode("` + el.Mode.String() + `");`))
	}

	if len(el.Content) != 0 {
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`.setValue("` + el.Content + `");`))
	}

	if el.TabSoft != 0 {
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`.getSession().setUseSoftTabs(` + el.TabSoft.String() + `);`))
	}

	if el.TabSize != 0 {
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`.getSession().setTabSize(` + strconv.Itoa(el.TabSize) + `);`))
	}

	if el.FontSize != 0 {
		ret.Write([]byte(`document.getElementById("`))
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`").style.fontSize="` + strconv.Itoa(el.FontSize) + `px";`))
	}

	if el.WarpMode != 0 {
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`.getSession().setUseWrapMode(` + el.WarpMode.String() + `);`))
	}

	if el.HighlightActiveLine != 0 {
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`.setHighlightActiveLine(` + el.HighlightActiveLine.String() + `);`))
	}

	if el.ShowPrintMargin != 0 {
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`.setShowPrintMargin(` + el.ShowPrintMargin.String() + `);`))
	}

	if el.ReadOnly != 0 {
		ret.WriteString(el.Html.Global.Id)
		ret.Write([]byte(`.setReadOnly(` + el.ReadOnly.String() + `);`))
	}

	ret.Write([]byte{0x0A})

	ret.WriteString(`var useWebWorker = window.location.search.toLowerCase().indexOf('noworker') == -1;`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`var editor = ace.edit("` + el.Html.Global.Id + `"); // fixme: id do elemento aqui`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`editor.getSession().setUseWorker(useWebWorker);`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`editor.$blockScrolling = Infinity;//prevents ace from logging annoying warnings`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`//#endregion`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`ace.config.loadModule('ace/ext/tern', function () {`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`editor.setOptions({`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`/**`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* Either 'true'' or 'false'' or to enable with custom options pass object that`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* has options for tern server: http://ternjs.net/doc/manual.html#server_api`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* If 'true'', then default options will be used`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`*/`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`enableTern: {`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`/* http://ternjs.net/doc/manual.html#option_defs */`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`defs: ['browser', 'ecma5', 'db'],`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`/* http://ternjs.net/doc/manual.html#plugins */`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`plugins: {`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`doc_comment: {`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`fullDocs: true`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`}`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`},`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`/**`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* (default is true) If web worker is used for tern server.`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* This is recommended as it offers better performance, but prevents this from working in a local html file due to browser security restrictions`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`*/`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`useWorker: useWebWorker,`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`/* if your editor supports switching between different files (such as tabbed interface) then tern can do this when jump to defnition of function in another file is called, but you must tell tern what to execute in order to jump to the specified file */`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`switchToDoc: function (name, start) {`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`//console.log('switchToDoc called but not defined. name=' + name + '; start=', start);`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`},`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`/**`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* if passed, this function will be called once ternServer is started.`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* This is needed when useWorker=false because the tern source files are loaded asynchronously before the server is started.`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`*/`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`startedCb: function () {`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`//once tern is enabled, it can be accessed via editor.ternServer`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`//console.log('editor.ternServer:', editor.ternServer);`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`},`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`},`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`/**`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* when using tern, it takes over Ace's built in snippets support.`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* this setting affects all modes when using tern, not just javascript.`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`*/`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`enableSnippets: true,`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`/**`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* when using tern, Ace's basic text auto completion is enabled still by deafult.`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* This settings affects all modes when using tern, not just javascript.`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`* For javascript mode the basic auto completion will be added to completion results if tern fails to find completions or if you double tab the hotkey for get completion (default is ctrl+space, so hit ctrl+space twice rapidly to include basic text completions in the result)`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`*/`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`enableBasicAutocompletion: true,`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`});`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`});`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`//#region not relevant to tern (custom beautify plugin) and demo loading`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`ace.config.loadModule('ace/ext/html_beautify', function (beautify) {`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`editor.setOptions({`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`// beautify when closing bracket typed in javascript or css mode`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`autoBeautify: true,`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`// this enables the plugin to work with hotkeys (ctrl+b to beautify)`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`htmlBeautify: true,`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`});`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`//modify beautify options as needed:`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`window.beautifyOptions = beautify.options;`)
	ret.Write([]byte{0x0A})
	ret.WriteString(`});`)
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}

func (el *AceEditor) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *AceEditor) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *AceEditor) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
