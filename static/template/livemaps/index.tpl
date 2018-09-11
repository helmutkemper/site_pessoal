{{ define "index" }}<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" href="/static/template/kendo-ui/styles/kendo.common-material.min.css" />
  <link rel="stylesheet" href="/static/template/kendo-ui/styles/kendo.material.min.css" />
  <link rel="stylesheet" href="/static/template/kendo-ui/styles/kendo.material.mobile.min.css" />

  <script src="/static/template/kendo-ui/js/jquery.min.js"></script>
  <script src="/static/template/kendo-ui/js/kendo.all.min.js"></script>

  <link rel="stylesheet" href="/static/template/fonts/fontawesome/v5.0.13/css/all.css" integrity="sha384-DNOHZ68U8hZfKXOrtjWvjxusGo9WQnrNx2sqG0tfsghAvtVlRW3tvkXWZh58N9jp" crossorigin="anonymous">
  <script src="/static/template/ace/ace.js" type="text/javascript" charset="utf-8"></script>
  <style>

    div.container {
      width: 1400px;
      /*border: 1px solid gray;*/
      margin: auto;
    }

    header, footer {
      padding: 0.5em;
      color: white;
      background-color: black;
      clear: left;
      text-align: center;
    }

    nav {
      float: left;
      max-width: 160px;
      margin: 0;
      padding: 1em;
    }

    nav ul {
      list-style-type: none;
      padding: 0;
    }

    nav ul a {
      text-decoration: none;
    }

    article {
      margin-left: 170px;
      border-left: 1px solid #eeeeee;
      padding: 1em;
      overflow: hidden;
      min-height: 800px;
    }

    #aceEditor{
      width: 1300px;
      height: 600px;
      border: 1px solid #eeeeee;
    }

  </style>
  {{htmlSafe .Telerik.ScriptTemplate}}
  <script>
    {{js .Telerik.VarGlobal}}
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

      $.ajax({
        url: '/getData/editCode/' + $.QueryString['name'],
        dataType: 'json',
        method: 'GET',
        success: function( data, textStatus, jqXHR ){
          setValueById( 'id:codeName', data.Objects[0].name );
          setValueById( 'id:aceEditor', data.Objects[0].code );

          $('#submitButton').click(function(){
            var obj = {
              name: getValueById( 'id:codeName' ),
              code: getValueById( 'id:aceEditor' ),
            };

            $.ajax({
              url: '/updateData/editCode/' + $.QueryString['name'],
              dataType: 'json',
              method: 'POST',
              data: JSON.stringify(obj),
              success: function( data, textStatus, jqXHR ) {
                setValueById('id:codeName', data.Objects[0].name);
                setValueById('id:aceEditor', data.Objects[0].code);
              }
            });

          });
        }
      });

      {{js .Telerik.OnLoadCode}}
    });
  </script>
</head>
<body>

<div class="container">

  <header>
    <h1>{{.Title}}</h1>
  </header>

  <nav>
    <ul>
      <li><a href="/admin/reloadConfig">Reload onLoad JS</a></li>
      <li><a href="#">Paris</a></li>
      <li><a href="#">Tokyo</a></li>
    </ul>
  </nav>

  <article>
    {{range .Articles}}
      <h1>{{.Title}}</h1>
      {{htmlSafe .Content}}
    {{end}}
  </article>

  <footer>Copyright &copy; Ahgora Sistemas 2018</footer>

</div>

</body>
</html>{{js .Telerik.HtmlSupport}}{{ end }}