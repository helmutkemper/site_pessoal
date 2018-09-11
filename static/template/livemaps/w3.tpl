{{ define "w3" }}<!DOCTYPE html>
<html>
<title></title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="/static/template/livemaps/style.css">

<link rel="stylesheet" href="/static/template/kendo-ui/styles/kendo.common-material.min.css" />
<link rel="stylesheet" href="/static/template/kendo-ui/styles/kendo.material.min.css" />
<link rel="stylesheet" href="/static/template/kendo-ui/styles/kendo.material.mobile.min.css" />

<script src="/static/template/kendo-ui/js/jquery.min.js"></script>
<script src="/static/template/kendo-ui/js/kendo.all.min.js"></script>

<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.0.13/css/all.css" integrity="sha384-DNOHZ68U8hZfKXOrtjWvjxusGo9WQnrNx2sqG0tfsghAvtVlRW3tvkXWZh58N9jp" crossorigin="anonymous">
<script src="/static/template/ace/ace.js" type="text/javascript" charset="utf-8"></script>

<style>
  body {
    background-color: coral;
  }

  header, footer {
    padding: 0.1em;
    color: white;
    background-color: #101010;
    clear: left;
    text-align: center;
  }

  article {
    height: 100%;
    min-height: 800px;
  }
  .title {
    padding-left: 50px;
  }

  .k-content{
    width: 100%;
    height: 100%;
  }

  .k-textbox{
    width: 100%;
  }

  #editor{
    //width: 1300px;
    height: 600px;
    border: 1px solid #eeeeee;
  }
</style>

{{htmlSafe .Telerik.ScriptTemplate}}
<script>
  {{js .Telerik.VarGlobal}}
  var notificationWidget;
  (function($) {
    $.QueryString = (function(paramsArray) {
      let params = {};

      for (let i = 0; i < paramsArray.length; ++i)
      {
        let param = paramsArray[i]
        .split('=', 2);

        if (param.length !== 2)
          continue;

        params[param[0]] = decodeURIComponent(param[1].replace(/\+/g, " "));
      }

      return params;
    })(window.location.search.substr(1).split('&'))
  })(jQuery);

  $(document).ready(function () {

    notificationWidget = $("#notification").kendoNotification().data("kendoNotification");
    $.ajax({
      url: '/getData/editCode/' + $.QueryString['name'],
      dataType: 'json',
      method: 'GET',
      success: function( data, textStatus, jqXHR ){
        if( data.Objects.length != 0 ) {
          setValueById('id:codeName', data.Objects[0].name);
          setValueById('id:editor', data.Objects[0].code);
        }

        $('#submitButton').click(function(){
          var obj = {
            name: getValueById( 'id:codeName' ),
            code: getValueById( 'id:editor' ),
          };

          $.ajax({
            url: '/updateData/editCode/' + $.QueryString['name'],
            dataType: 'json',
            method: 'POST',
            data: JSON.stringify(obj),
            success: function(){ notificationWidget.show("Arquivo salvo", "success"); },
            error: function(){ notificationWidget.show("Error ao salvar o arquivo", "error"); },
            /*success: function( data, textStatus, jqXHR ) {
              setValueById('id:codeName', data.Objects[0].name);
              setValueById('id:editor', data.Objects[0].code);
            }*/
          });

        });
      }
    });

  {{js .Telerik.OnLoadCode}}
  });
</script>

<body style="background:#e6e6e6;color:black;">

<div class="w3-sidebar w3-light-grey w3-card-4 w3-animate-left" style="width:200px" id="mySidebar">
  <div class="w3-bar w3-dark-grey">
    <span class="w3-bar-item w3-padding-16">Content</span>
    <button onclick="w3_close()"
            class="w3-bar-item w3-button w3-right w3-padding-16" title="close Sidebar">&times;</button>
  </div>
  <div class="w3-bar-block">
    <a class="w3-bar-item w3-button w3-green" href="javascript:void(0)">Home</a>
    <a class="w3-bar-item w3-button" href="/static/template/livemaps/processLogCompany.tpl?company=a961585">Process Log Company</a>
    <a class="w3-bar-item w3-button" href="javascript:void(0)">Contact</a>
    <a class="w3-bar-item w3-button" href="javascript:void(0)">Support</a>
  </div>
</div>

<div id="main" style="margin-left:200px">
  <header>
    <h1>{{.Title}}</h1>
  </header>
  <div class="w3-container w3-display-container">
    <span title="open Sidebar" style="display:none" id="openNav" class="w3-button w3-transparent w3-display-topleft w3-xlarge" onclick="w3_open()">&#9776;</span>
    <article>
    {{range .Articles}}
      <h1 class="title">{{.Title}}</h1>
    {{htmlSafe .Content}}
    {{end}}
    </article>

    <footer>Copyright &copy; Ahgora Sistemas 2018</footer>
  </div>

</div>

<script>
  function w3_open() {
    document.getElementById("main").style.marginLeft = "180px";
    document.getElementById("mySidebar").style.width = "180px";
    document.getElementById("mySidebar").style.display = "block";
    document.getElementById("openNav").style.display = 'none';
  }
  function w3_close() {
    document.getElementById("main").style.marginLeft = "0";
    document.getElementById("mySidebar").style.display = "none";
    document.getElementById("openNav").style.display = "inline-block";
  }
</script>
<span id="notification"></span>
</body>
</html>{{js .Telerik.HtmlSupport}}{{ end }}