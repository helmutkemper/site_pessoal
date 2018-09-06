package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
	"time"
)

type KendoGridPdf struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.avoidlinks

	  A flag indicating whether to produce actual hyperlinks in the exported PDF file.

	  It's also possible to pass a CSS selector as argument. All matching links will be ignored. (default: false)

	  Important:

	  Available in versions 2015.3.1020 and later

	*/
	//
	//    Example - skip hyperlinks
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            creator: "John Doe",
	//            avoidLinks: true
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        columns: [
	//          { field: "ProductName",
	//            template: "<a href='producs/#= ProductID #/'>#= ProductName #</a>" }
	//        ],
	//        pageable: true
	//    });
	//    </script>
	//
	AvoidLinks interface{} `jsObject:"avoidLinks" jsType:"Boolean,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.scale

	  A scale factor. In many cases, text size on screen will be too big for print, so you can use this option to scale down
	  the output in PDF. See the documentation in drawDOM. (default: 1)

	*/
	Scale int `jsObject:"scale"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.proxyurl

	  The URL of the server side proxy which will stream the file to the end user. (default: null)

	  A proxy will be used when the browser is not capable of saving files locally, for example, Internet Explorer 9 and
	  Safari.

	  The developer is responsible for implementing the server-side proxy.

	  The proxy will receive a POST request with the following parameters in the request body:

	  > contentType: The MIME type of the file
	  > base64: The base-64 encoded file content
	  > fileName: The file name, as requested by the caller.

	  The proxy should return the decoded file with the "Content-Disposition" header set to `attachment;
	  filename="<fileName.pdf>"´.

	*/
	//
	//    Example - set the server proxy URL
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            proxyURL: "/save"
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	ProxyURL string `jsObject:"proxyURL"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.allpages

	  Exports all grid pages, starting from the first one. (default: false)

	  Important:

	  Chrome is known to crash when generating very large PDF-s.  A solution to this is to include the Pako
	  [ http://nodeca.github.io/pako/ ] library, which is bundled with Kendo as `pako_deflate.min.js´.  Simply loading this
	  library with a `<script>´ tag will enable compression in PDF, e.g.:
	  `<script src="http://kendo.cdn.telerik.com/2018.2.516/js/pako_deflate.min.js"></script>´
	  The allPages export is not supported when virtual scrolling is enabled.

	*/
	//
	//    Example - export all pages
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      toolbar: ["pdf"],
	//      columns: [
	//        { field: "name" }
	//      ],
	//      dataSource: {
	//        data: [{ name: "Jane Doe"},
	//               { name: "John Doe"},
	//               { name: "Tim Doe"},
	//               { name: "Alice Doe"}],
	//        pageSize: 2
	//      },
	//      pdf: {
	//        allPages: true
	//      }
	//    });
	//    var grid = $("#grid").data("kendoGrid");
	//    grid.saveAsPDF();
	//    </script>
	//
	AllPages Boolean `jsObject:"allPages"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.filename

	  Specifies the file name of the exported PDF file. (default: "Export.pdf")

	*/
	//
	//    Example - set the default PDF file name
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            fileName: "Products.pdf"
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	FileName string `jsObject:"fileName"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.creator

	  The creator of the PDF document. (default: "Kendo UI PDF Generator")

	*/
	//
	//    Example - set the creator
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            creator: "John Doe"
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	Creator string `jsObject:"creator"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.papersize

	  Specifies the paper size of the PDF document. (default: "auto")

	  The default "auto" means paper size is determined by content.

	  Important:

	  The size of the content in pixels will match the size of the output in points (1 pixel = 1/72 inch).

	  Supported values:

	  > A predefined size: "A4", "A3" etc
	  > An array of two numbers specifying the width and height in points (1pt = 1/72in)
	  > An array of two strings specifying the width and height in units. Supported units are "mm", "cm", "in" and "pt".

	*/
	//
	//    Example - set custom paper size
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            paperSize: ["20mm", "20mm"]
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	PaperSize interface{} `jsObject:"paperSize" jsType:"string,[]string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.date

	  The date when the PDF document is created. Defaults to `new Date()´.

	*/
	//
	//    Example - set the date
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            date: new Date("2014/10/10")
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	Date time.Time `jsObject:"date"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.forceproxy

	  If set to true, the content will be forwarded to proxyURL even if the browser supports saving files locally.
	  (default: false)

	*/
	ForceProxy Boolean `jsObject:"forceProxy"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.template

	  A piece of HTML to be included in each page. Can be used to display headers and footers. See the documentation in
	  drawDOM. (default: null)

	  Available template variables include:

	  pageNum
	  totalPages

	  Important:

	  Using a template requires setting paper size

	*/
	Template KendoPdfTemplate `jsObject:"template"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.title

	  Sets the title of the PDF file. (default: null)
	*/
	//
	//    Example - set the title
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            title: "Products"
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	Title string `jsObject:"title"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.keywords

	  Specifies the keywords of the exported PDF file. (default: null)
	*/
	//
	//    Example - set the keywords
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            keywords: "northwind products"
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	Keywords []string `jsObject:"keywords"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.landscape

	  Set to `true´ to reverse the paper dimensions if needed such that width is the larger edge. (default: false)

	*/
	//
	//    Example - enable landscape mode
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            landscape: true
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	Landscape Boolean `jsObject:"landscape"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.subject

	  Sets the subject of the PDF file. (default: null)
	*/
	//
	//    Example - set the subject
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            subject: "Products"
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	Subject string `jsObject:"subject"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.margin

	  Specifies the margins of the page (numbers or strings with units). Supported
	  units are "mm", "cm", "in" and "pt" (default).
	*/
	//
	//    Example - set the margins
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            margin: {
	//                left: 10,
	//                right: "10pt",
	//                top: "10mm",
	//                bottom: "1in"
	//            }
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	Margin KendoPdfMargin `jsObject:"margin"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.repeatheaders

	  Set this to true to repeat the grid headers on each page. (default: false)
	*/
	RepeatHeaders Boolean `jsObject:"repeatHeaders"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.author

	  The author of the PDF document. (default: null)
	*/
	//
	//    Example - set the author
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            author: "John Doe"
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	Author string `jsObject:"author"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf.proxytarget

	  A name or keyword indicating where to display the document returned from the proxy.

	  If you want to display the document in a new window or iframe,
	  the proxy should set the "Content-Disposition" header to `inline; filename="<fileName.pdf>"´. (default: "_self")
	*/
	//
	//    Example - open the generated document in a new window
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        toolbar: ["pdf"],
	//        pdf: {
	//            forceProxy: true,
	//            proxyURL: "/save",
	//            proxyTarget: "_blank"
	//        },
	//        dataSource: {
	//            transport: {
	//                read: {
	//                    url: "https://demos.telerik.com/kendo-ui/service/products",
	//                    dataType: "jsonp"
	//                }
	//            },
	//            pageSize: 10
	//        },
	//        pageable: true
	//    });
	//    </script>
	//
	ProxyTarget interface{} `jsObject:"proxyTarget" jsType:"string,HtmlTarget"`

	*ToJavaScriptConverter
}

func (el *KendoGridPdf) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridPdf.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
