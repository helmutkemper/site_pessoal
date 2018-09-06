package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

/*
@see https://docs.telerik.com/kendo-ui/api/javascript/kendo
fixme: imcompleto
*/
type Kendo struct {
	Id string `jsObject:"-"`

	/*
	  Compiles a template to a function that builds HTML. Useful when a template will be used several times. Templates offer
	  way of creating HTML chunks. Options such as HTML encoding and compilation for optimal performance are available.

	  Example - Basic template
	  <span id="output"></span>
	  <script>
	    var template = kendo.template("Hello, #= firstName # #= lastName #");
	    var data = { firstName: "John", lastName: "Doe" };
	    $("#output").html(template(data));
	  </script>

	  Example - encode HTML
	  <span id="output"></span>
	  <script>
	    var template = kendo.template("HTML tags are encoded: #: html #");
	    var data = { html: "<strong>lorem ipsum</strong>" };
	    $("#output").html(template(data));
	  </script>

	  Example - use JavaScript in the template
	  <span id="output"></span>
	  <script>
	    var template = kendo.template("#if (foo) {# foo is true #}#");
	    var data = { foo: true };
	    $("#output").html(template(data));
	  </script>

	  Example - escape sharp ("#") symbols in an inline template
	  <span id="output"></span>
	  <script>
	    var template = kendo.template("<a href='\\#'>link</a>");
	    $("#output").html(template({}));
	  </script>

	  Example - escape sharp ("#") symbols in a template defined in a script element
	  <span id="output"></span>
	  <script type="text/x-kendo-template" id="template">
	    <a href="\#">link</a>
	  </script>
	  <script>
	    var template = kendo.template($("#template").html());
	    $("#output").html(template({}));
	  </script>

	  ReturnsFunction the compiled template as a JavaScript function. When called this function will return the generated
	  HTML string.

	  Parameters:

	    template (string): The template that will be compiled.
	    options (object) [optional]: Template compilation options.
	    options
	      paramName (string)  - default: "data": The name of the parameter used by the generated function. Useful when
	      useWithBlock is set to false.

	        Example:

	        var template = kendo.template("<strong>#: d.name #</strong>", { paramName: "d", useWithBlock: false });

	      useWithBlock (boolean) - default: true: Wraps the generated code in a with block. This allows the usage of
	      unqualified fields in the template. Disabling the with block will improve the performance of the template.

	        Example:

	        // Note that "data." is used to qualify the field
	        var template = kendo.template("<strong>#: data.name #</strong>", { useWithBlock: false });

	        console.log(template({ name: "John Doe" })); // outputs "<strong>John Doe</strong>"
	*/
	Template interface{} `jsObject:"template" jsType:"*JavaScript,string"`

	*ToJavaScriptConverter
}

func (el *Kendo) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Id == "" {
		el.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendo.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Id + `").kendo({`))
	ret.Write(data)
	ret.Write([]byte(`});`))

	return ret.Bytes()
}
