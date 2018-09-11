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
          url: '/getData/codeList',
          cache: false,
          dataType: 'json',
          dataSrc: 'Objects'
        },
        columns: [
          {
            data: 'name',
            render: function ( data, type, row ) {
              return '/customJs/' + data;
            },
            type: 'string',
            title: 'Code Name'
          },
          {
            data: '_id',
            render: function ( data, type, row ) {
              return '<a href="/edit/js/' + data + '">edit</a>';
            },
            type: 'string',
            title: 'Edit'
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