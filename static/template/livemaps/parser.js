(function() {

  /**
   * @var include em funções pré definidas para a coleção 'untreatedData'
   */
  var ud = require('untreatedData');

  /**
   * @var include em funções predefindas para coleções não críticas
   */
  var newData = require('db');
  // define o nome da coleção para salvar os dados
  newData.collection('bigData_parserByJs');

  // calcula a quantidade de itens encontrado
  var count = ud.count({ query: {} });

  // valores muito elevados podem fazer a execução do javascript ficar lenta
  var limit = 100;

  // recebe os documentos da coleção
  var data;

  // laço para pegar todos os documentos
  for( var i = 0, l = Math.ceil( count/limit ); i < l; i += 1 ){

    // find na coleção 'untreatedData'
    ud.find({
      'query': {},
      'limit': limit,
      'skip': i * limit
    });

    // pega todos os dados da coleção
    data = ud.dataGetAll();

    // laço para separar cada documento da resposta
    for( var dataKey in data ){

      // cada linha é um documento do banco
      var line = data[dataKey];

      // dados dos empregados
      var employeeMatricula = line.employee.matricula;
      var employeeNome      = line.employee.nome;
      var employeePis       = line.employee.pis;
      var employeeId        = line.employee._id;

      // demais dados
      var daysInfos         = line.daysInfos;
      var shifts            = line.shifts;
      var company           = line.company;
      var totals            = line.totals;

      // laço para cada linha do daysInfos contido no documento do banco de dados
      for( var daysInfosCount = 0, daysInfosLength = daysInfos.length; daysInfosCount < daysInfosLength; daysInfosCount += 1 ){

        // separa um day info
        var daysInfosElement  = daysInfos[daysInfosCount];

        // separa os elementos dos indicadores
        var daysInfosSumTotal       = daysInfosElement.sumTotal;
        var daysInfosPunches        = daysInfosElement.punches;
        var daysInfosJornada        = daysInfosElement.jornada;
        var daysInfosDia            = daysInfosElement.dia;
        var daysInfosIsAbsentAllDay = daysInfosElement.isAbsentAllDay;

        // monta o dado a ser salvo por day info
        var dataToInsert = {
          company:                 company,
          daysInfosSumTotal:       daysInfosSumTotal,
          daysInfosPunches:        daysInfosPunches,
          daysInfosJornada:        daysInfosJornada,
          daysInfosDia:            daysInfosDia,
          daysInfosIsAbsentAllDay: daysInfosIsAbsentAllDay,
          //daysInfos:               daysInfosElement,

          totals:                  totals,

          employeeMatricula:       employeeMatricula,
          employeeNome:            employeeNome,
          employeePis:             employeePis,
          employeeId:              employeeId,
        };

        // insere no banco de dados
        newData.insert( dataToInsert );

        // testa e aborta em caso de erro
        var error = db.getLastError();
        if( error !== '' ){
          print('error: ' + error);
          return;
        }
      }
    }
  }
}).call(this);