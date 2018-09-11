<html>
<head>
  <link rel="stylesheet" type="text/css" href="/static/template/datatable/datatables.min.css"/>
  <script type="text/javascript" src="/static/template/datatable/datatables.min.js"></script>
  <style>
    .display{

    }
    .content{
      margin: auto;
      width: 800px;
    }
  </style>
  <script>
    $(document).ready( function () {
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

      // http://localhost:8888/getData/runLogAll
      $('#table_id').DataTable({
        responsive: true,
        ajax: {
          url: '/getData/runLogAll',
          cache: false,
          dataType: 'json',
          dataSrc: 'Objects'
        },
        columns: [
          {
            data: 'codeName',
            type: 'string',
            title: 'Code Name'
          },
          {
            data: 'elapsed_ms',
            type: 'numeric',
            title: 'Exe. Time (ms)'
          },
          {
            data: 'start',
            render: function ( data, type, row ) {
              var options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };
              return new Date( data ).toLocaleDateString("en-US", options);
            },
            type: 'numeric',
            title: 'Start'
          },
          {
            data: 'end',
            render: function ( data, type, row ) {
              var options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };
              return new Date( data ).toLocaleDateString("en-US", options);
            },
            type: 'numeric',
            title: 'End'
          }
        ]
      });
    });
  </script>
</head>
<body>
<div>
  <div class="content"><table id="table_id" class="display"></table></div>
</div>
</body>
</html>