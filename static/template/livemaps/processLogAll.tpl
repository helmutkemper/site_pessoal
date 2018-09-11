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
      function Employee ( name, position, salary, office ) {
        this.name = name;
        this.position = position;
        this.salary = salary;
        this._office = office;

        this.office = function () {
          return this._office;
        }
      };
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

        // http://localhost:8888/static/template/livemaps/processLogAll.tpl
        $('#table_id').DataTable({
          responsive: true,
          ajax: {
            url: '/getData/processLogAll/',
            cache: false,
            dataType: 'json',
            dataSrc: 'Objects'
          },
          columns: [
            {
              data: '__created',
              render: function ( data, type, row ) {
                var options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };
                return new Date( data ).toLocaleDateString("en-US", options);
              },
              type: 'string',
              title: 'Created'
            },
            {
              data: 'month',
              type: 'numeric',
              title: 'Process Month'
            },
            {
              data: 'year',
              type: 'numeric',
              title: 'Process Year'
            },
            {
              data: 'empresa',
              type: 'string',
              title: 'Company'
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