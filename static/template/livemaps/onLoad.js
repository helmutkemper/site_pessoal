(function() {
  addRouteToGetData( 'criticalErrors', 'bigData_ahpiCriticalErrors', {}, '', {} );
  addRouteToGetData( 'runLogAll', 'bigData_javascriptRunLog', {}, '', {} );
  addRouteToGetData( 'runLogCode', 'bigData_javascriptRunLog', {"codeName": "#name#"}, '^.*?/(?P<name>[a-zA-Z0-9_]+)$', {} );
  addRouteToGetData( 'processLog', 'bigData_ahpiProcessLog', {"empresa": "#company#"}, '^.*?/(?P<company>a[0-9]+)$', {} );
  addRouteToGetData( 'processLogAll', 'bigData_ahpiProcessLog', {}, '', {} );
  addRouteToGetData( 'processCompany', 'bigData_ahpiDistinctCompany', {}, '', {} );
  addRouteToGetData( 'codeList', 'bigData_javascriptCode', {}, '', { 'name': 1 } );

  addRouteToGetData( 'editCode', 'bigData_javascriptCode', {"name": "#name#"}, '^.*?/(?P<name>[a-zA-Z0-9_]+)$', {} );

  var jwt = getJWT("a961585", "livemaps@ahgora.com.br", "12zetaSC");
  print(jwt);

  var body = getRoute( "https://www.ahgora.com.br/login/jwt/a961585", "GET", {"Authorization": "basic bGl2ZW1hcHNAYWhnb3JhLmNvbS5icjoxMnpldGFTQw=="} );
  print(body);

  //cronGetRoute( '@every 20s', 'https://www.ahgora.com.br/login/jwt/a961585', 'GET', {'Authorization': 'basic bGl2ZW1hcHNAYWhnb3JhLmNvbS5icjoxMnpldGFTQw=='} )
}).call(this);