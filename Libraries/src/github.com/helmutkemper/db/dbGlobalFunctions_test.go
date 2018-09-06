package db

import (
	"fmt"
	"github.com/helmutkemper/gOsm/consts"
	"os"
	"path/filepath"
)

func ExampleCreateBackup() {
	// Conecta ao banco de dados
	Connect(consts.DB_SERVER_TEST, consts.DB_DATABASE_TEST)

	// Faz um índice 2d
	err := Index2DMake(consts.DB_OSM_FILE_NODES_COLLECTIONS, "loc")
	if err != nil {
		return
	}

	// Cria uma pasta para o backup
	if _, err = os.Stat("./backup"); os.IsNotExist(err) {
		err = os.Mkdir("./backup", 0777)
		if err != nil {
			return
		}
	}

	// Define host como "127.0.0.1", db como sendo "gOsm_test" e o caminho do arquivo a ser salvo
	var dataCommand MongoBackupData = MongoBackupData{
		Host:    consts.DB_SERVER_TEST,
		Db:      consts.DB_DATABASE_TEST,
		Archive: "./backup/mongo_backup.bson",
	}

	// Faz o backup
	type backupStt struct {
		*DbStt `bson:"-"`
	}
	var bk backupStt = backupStt{}
	err = bk.DbStt.BackupMake(&dataCommand)

	// Apaga o banco de dados
	DropDatabase(consts.DB_DATABASE_TEST)

	// Recupera o backup
	err = bk.DbStt.BackupRestore(&dataCommand)

	// Apaga todo o conteúdo da pasta de backup
	d, err := os.Open("./backup")
	if err != nil {
		return
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, name := range names {
		// Imprime o nome do arquivo de backup antes de apagar
		fmt.Printf("Backup file: %v\n", names)
		err = os.RemoveAll(filepath.Join("./backup", name))
		if err != nil {
			return
		}
	}

	// Procura pelo índice 2d gravado no backup
	index, err := IndexList(consts.DB_OSM_FILE_NODES_COLLECTIONS)

	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Index name: %v\n", index[0].Name)
	fmt.Printf("Index bits: %v\n", index[0].Bits)
	fmt.Printf("Index key: %v\n", index[0].Key)

	fmt.Printf("Index name: %v\n", index[1].Name)
	fmt.Printf("Index bits: %v\n", index[1].Bits)
	fmt.Printf("Index key: %v\n", index[1].Key)

	// Apaga o banco de dados após o teste.
	DropDatabase(consts.DB_DATABASE_TEST)

	// Desconecta.
	Disconnect()

	// Output:
	// Backup file: [mongo_backup.bson]
	// Error: <nil>
	// Index name: _id_
	// Index bits: 0
	// Index key: [_id]
	// Index name: loc
	// Index bits: 26
	// Index key: [$2d:loc]
}

func ExampleTestConnection() {
	// Exemplo:
	// Testa a conexão com o banco de dados e retorna um erro.
	//
	// Descrição:
	// Caso o banco de dados esteja conectado, o teste ira retornar 'nil', caso contrário, a mensagem de erro do banco.

	// Testa a conexão do banco.
	fmt.Printf("%v\n", TestConnection())

	// Output:
	// error: db is not connected to 'gOsm_test' on address '127.0.0.1'
}

func ExampleConnect() {
	// Exemplo:
	// Testa a conexão com o banco de dados e retorna ok
	//
	// Descrição:
	// Caso o banco de dados esteja conectado, o teste ira retornar 'nil', caso contrário, a mensagem de erro do banco.

	// Faz a conexão do banco;
	Connect(consts.DB_SERVER_TEST, consts.DB_DATABASE_TEST)

	// Testa a conexão do banco;
	fmt.Printf("%v\n", TestConnection())

	// Desconecta.
	Disconnect()

	// Output:
	// <nil>
}

func ExampleIndex2DMake() {
	// Exemplo:
	// Monta um index para pontos em 2D, ou seja, um array [x,y].
	//
	// Descrição:
	// Para se trabalhar com mapas em 2D, o MongoDB tem algumas funções prontas e estas funções esperam por dados na forma
	// de ponto cartesiano, (x,y), por isto, o mongo usa um array no formato [longitude,latitude], respectivamente, [x,y].
	// No caso de um grande volume de dados e consultas constantes, é necessário se adicionar um índice 2D de 26bits para
	// melhorar o desempenho.

	// Faz a conexão do banco;
	Connect(consts.DB_SERVER_TEST, consts.DB_DATABASE_TEST)

	// Cria um índice 2D na chave de nome 'loc' e assume que o conteúdo da mesma é um array [x,y]
	Index2DMake(consts.DB_OSM_FILE_NODES_COLLECTIONS, "loc")

	// Lista todos os índices do banco.
	index, err := IndexList(consts.DB_OSM_FILE_NODES_COLLECTIONS)

	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Index name: %v\n", index[0].Name)
	fmt.Printf("Index bits: %v\n", index[0].Bits)
	fmt.Printf("Index key: %v\n", index[0].Key)

	fmt.Printf("Index name: %v\n", index[1].Name)
	fmt.Printf("Index bits: %v\n", index[1].Bits)
	fmt.Printf("Index key: %v\n", index[1].Key)

	// Apaga o banco de dados após o teste.
	DropDatabase(consts.DB_DATABASE_TEST)

	// Desconecta.
	Disconnect()

	// Output:
	// Error: <nil>
	// Index name: _id_
	// Index bits: 0
	// Index key: [_id]
	// Index name: loc
	// Index bits: 26
	// Index key: [$2d:loc]
}

func ExampleIndexUniqueMake() {
	// Exemplo:
	// Monta um índice único para o MongoDB.
	//
	// Descrição:
	// Em certas ocasiões, é necessário impedir que um dado seja repetido no banco, e a maneira mais segura e rápida de
	// fazer isto é criar um índice único.

	// Faz a conexão do banco;
	Connect(consts.DB_SERVER_TEST, consts.DB_DATABASE_TEST)

	// Cria um índice único na chave de nome 'keyName'
	IndexUniqueMake(consts.DB_OSM_FILE_NODES_COLLECTIONS, "keyName")

	// Lista todos os índices do banco.
	index, err := IndexList(consts.DB_OSM_FILE_NODES_COLLECTIONS)

	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Index name: %v\n", index[0].Name)
	fmt.Printf("Index bits: %v\n", index[0].Bits)
	fmt.Printf("Index key: %v\n", index[0].Key)

	fmt.Printf("Index name: %v\n", index[1].Name)
	fmt.Printf("Index unique: %v\n", index[1].Unique)
	fmt.Printf("Index key: %v\n", index[1].Key)

	// Apaga o banco de dados após o teste.
	DropDatabase(consts.DB_DATABASE_TEST)

	// Desconecta.
	Disconnect()

	// Output:
	// Error: <nil>
	// Index name: _id_
	// Index bits: 0
	// Index key: [_id]
	// Index name: keyName_1
	// Index unique: true
	// Index key: [keyName]
}

func ExampleIndexDropByName() {
	// Exemplo:
	// Apaga um índice por nome
	//
	// Descrição:
	// Indices deixam as consultas do banco mais rápidas, mas, cobram um preço caro no espaço de ocupado no disco do
	// servidor, por isto, em algumas situações, é necessário apagar um índice.

	// Faz a conexão do banco;
	Connect(consts.DB_SERVER_TEST, consts.DB_DATABASE_TEST)

	// Cria um índice 2D na chave de nome 'loc' e assume que o conteúdo da mesma é um array [x,y]
	Index2DMake(consts.DB_OSM_FILE_NODES_COLLECTIONS, "loc")

	// Lista todos os índices do banco.
	index, err := IndexList(consts.DB_OSM_FILE_NODES_COLLECTIONS)

	fmt.Printf("Error: %v\n", err)

	fmt.Print("After make index 2D\n")
	for _, idx := range index {
		fmt.Printf("Index name: %v\n", idx.Name)
		fmt.Printf("Index bits: %v\n", idx.Bits)
		fmt.Printf("Index key: %v\n", idx.Key)
	}

	// Apaga o índice de nome 'loc'
	IndexDropByName(consts.DB_OSM_FILE_NODES_COLLECTIONS, "loc")

	// Lista todos os índices do banco.
	index, err = IndexList(consts.DB_OSM_FILE_NODES_COLLECTIONS)

	fmt.Print("After drop index 2D\n")
	for _, idx := range index {
		fmt.Printf("Index name: %v\n", idx.Name)
		fmt.Printf("Index bits: %v\n", idx.Bits)
		fmt.Printf("Index key: %v\n", idx.Key)
	}
	fmt.Print("\n")

	// Apaga o banco de dados após o teste.
	DropDatabase(consts.DB_DATABASE_TEST)

	// Desconecta.
	Disconnect()

	// Output:
	// Error: <nil>
	// After make index 2D
	// Index name: _id_
	// Index bits: 0
	// Index key: [_id]
	// Index name: loc
	// Index bits: 26
	// Index key: [$2d:loc]
	// After drop index 2D
	// Index name: _id_
	// Index bits: 0
	// Index key: [_id]
}

func ExampleListCollections() {
	Connect(consts.DB_SERVER_TEST, consts.DB_DATABASE_TEST)

	err := Index2DMake(consts.DB_OSM_FILE_NODES_COLLECTIONS, "loc")

	fmt.Printf("Error: %v\n", err)

	collections, err := ListCollections()

	fmt.Printf("Error: %v\n", err)

	fmt.Printf("Collections: %v\n", collections)

	// Drop database after test
	DropDatabase(consts.DB_DATABASE_TEST)

	// Mongo main disconnect
	Disconnect()

	// Output:
	// Error: <nil>
	// Error: <nil>
	// Collections: [nodes]
}
