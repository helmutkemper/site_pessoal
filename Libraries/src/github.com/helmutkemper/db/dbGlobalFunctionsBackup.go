package db

import (
	log "github.com/helmutkemper/seelog"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	MONGODB_BACKUP_RESTORE MongoBackupCommand = iota
	MONGODB_BACKUP_DUMP
)

type MongoBackupCommand int

var mongoBackupCommands = [...]string{
	"mongorestore",
	"mongodump",
}

func (e MongoBackupCommand) String() string {
	return mongoBackupCommands[e]
}

type MongoBackupData struct {

	// --host <hostname><:port>, -h <hostname><:port>
	// Default: localhost:27017
	//
	// Specifies a resolvable hostname for the mongod to which to connect. By default, the mongorestore attempts to
	// connect to a MongoDB instance running on the localhost on port number 27017.
	//
	// To connect to a replica set, specify the replSetName and a seed list of set members, as in the following:
	//
	//   <replSetName>/<hostname1><:port>,<hostname2><:port>,<...>
	//
	// You can always connect directly to a single MongoDB instance by specifying the host and port number directly.
	//
	// Changed in version 3.0.0: If you use IPv6 and use the <address>:<port> format, you must enclose the portion of an
	// address and port combination in brackets (e.g. [<address>]).
	Host string

	// --port <port>¶
	// Default: 27017
	//
	// Specifies the TCP port on which the MongoDB instance listens for client connections.
	Port string

	// --ssl
	// New in version 2.6.
	//
	// Enables connection to a mongod or mongos that has TLS/SSL support enabled.
	//
	// Changed in version 3.0: Most MongoDB distributions include support for TLS/SSL. See Configure mongod and mongos for
	// TLS/SSL and TLS/SSL Configuration for Clients for more information about TLS/SSL and MongoDB.
	//
	// Changed in version 3.4: If --sslCAFile is not specified when connecting to an TLS/SSL-enabled server, the
	// system-wide CA certificate store will be used.
	Ssl bool

	// --sslCAFile <filename>¶
	// New in version 2.6.
	//
	// Specifies the .pem file that contains the root certificate chain from the Certificate Authority. Specify the file
	// name of the .pem file using relative or absolute paths.
	//
	// Changed in version 3.0: Most MongoDB distributions include support for TLS/SSL. See Configure mongod and mongos for
	// TLS/SSL and TLS/SSL Configuration for Clients for more information about TLS/SSL and MongoDB.
	//
	// Changed in version 3.4: If --sslCAFile is not specified when connecting to an TLS/SSL-enabled server, the
	// system-wide CA certificate store will be used.
	//
	// WARNING
	//   Version 3.2 and earlier: For SSL connections (--ssl) to mongod and mongos, if the mongorestore runs without the
	//   --sslCAFile, mongorestore will not attempt to validate the server certificates. This creates a vulnerability to
	//   expired mongod and mongos certificates as well as to foreign processes posing as valid mongod or mongos
	//   instances.
	//   Ensure that you always specify the CA file to validate the server certificates in cases where intrusion is a
	//   possibility.
	SslCAFile string

	// --sslPEMKeyFile <filename>
	// New in version 2.6.
	//
	// Specifies the .pem file that contains both the TLS/SSL certificate and key. Specify the file name of the .pem file
	// using relative or absolute paths.
	//
	// This option is required when using the --ssl option to connect to a mongod or mongos that has CAFile enabled
	// without allowConnectionsWithoutCertificates.
	//
	// Changed in version 3.0: Most MongoDB distributions include support for TLS/SSL. See Configure mongod and mongos for
	// TLS/SSL and TLS/SSL Configuration for Clients for more information about TLS/SSL and MongoDB.
	//
	// Changed in version 3.4: If --sslCAFile is not specified when connecting to an TLS/SSL-enabled server,
	// the system-wide CA certificate store will be used.
	SslPEMKeyFile string

	// --sslPEMKeyPassword <value>
	// New in version 2.6.
	//
	// Specifies the password to de-crypt the certificate-key file (i.e. --sslPEMKeyFile). Use the --sslPEMKeyPassword
	// option only if the certificate-key file is encrypted. In all cases, the mongorestore will redact the password from
	// all logging and reporting output.
	//
	// If the private key in the PEM file is encrypted and you do not specify the --sslPEMKeyPassword option, the
	// mongorestore will prompt for a passphrase. See SSL Certificate Passphrase.
	//
	// Changed in version 3.0: Most MongoDB distributions include support for TLS/SSL. See Configure mongod and mongos for
	// TLS/SSL and TLS/SSL Configuration for Clients for more information about TLS/SSL and MongoDB.
	//
	// Changed in version 3.4: If --sslCAFile is not specified when connecting to an TLS/SSL-enabled server, the
	// system-wide CA certificate store will be used.
	SslPEMKeyPassword string

	// --sslCRLFile <filename>
	// New in version 2.6.
	//
	// Specifies the .pem file that contains the Certificate Revocation List. Specify the file name of the .pem file using
	// relative or absolute paths.
	//
	// Changed in version 3.0: Most MongoDB distributions include support for TLS/SSL. See Configure mongod and mongos for
	// TLS/SSL and TLS/SSL Configuration for Clients for more information about TLS/SSL and MongoDB.
	//
	// Changed in version 3.4: If --sslCAFile is not specified when connecting to an TLS/SSL-enabled server, the
	// system-wide CA certificate store will be used.
	SslCRLFile string

	// --sslAllowInvalidCertificates
	// New in version 2.6.
	//
	// Bypasses the validation checks for server certificates and allows the use of invalid certificates. When using the
	// allowInvalidCertificates setting, MongoDB logs as a warning the use of the invalid certificate.
	//
	// Changed in version 3.0: Most MongoDB distributions include support for TLS/SSL. See Configure mongod and mongos for
	// TLS/SSL and TLS/SSL Configuration for Clients for more information about TLS/SSL and MongoDB.
	//
	// Changed in version 3.4: If --sslCAFile is not specified when connecting to an TLS/SSL-enabled server, the
	// system-wide CA certificate store will be used.
	SslAllowInvalidCertificates bool

	// --sslAllowInvalidHostnames
	// New in version 3.0.
	//
	// Disables the validation of the hostnames in TLS/SSL certificates. Allows mongorestore to connect to MongoDB
	// instances even if the hostname in their certificates do not match the specified hostname.
	//
	// Changed in version 3.0: Most MongoDB distributions include support for TLS/SSL. See Configure mongod and mongos for
	// TLS/SSL and TLS/SSL Configuration for Clients for more information about TLS/SSL and MongoDB.
	//
	// Changed in version 3.4: If --sslCAFile is not specified when connecting to an TLS/SSL-enabled server, the
	// system-wide CA certificate store will be used.
	SslAllowInvalidHostNames bool

	// --sslFIPSMode
	// New in version 2.6.
	//
	// Directs the mongorestore to use the FIPS mode of the installed OpenSSL library. Your system must have a FIPS
	// compliant OpenSSL library to use the --sslFIPSMode option.
	//
	//   NOTE
	//   FIPS-compatible SSL is available only in MongoDB Enterprise. See Configure MongoDB for FIPS for more information.
	SslFIPSMode bool

	// --username <username>, -u <username>
	// Specifies a username with which to authenticate to a MongoDB database that uses authentication. Use in conjunction
	// with the --password and --authenticationDatabase options.
	Username string

	// --password <password>, -p <password>
	// Specifies a password with which to authenticate to a MongoDB database that uses authentication. Use in conjunction
	// with the --username and --authenticationDatabase options.
	//
	// Changed in version 3.0.0: If you do not specify an argument for --password, mongorestore returns an error.
	//
	// Changed in version 3.0.2: If you wish mongorestore to prompt the user for the password, pass the --username option
	// without --password or specify an empty string as the --password value, as in --password "" .
	Password string

	// --authenticationDatabase <dbname>
	// Specifies the database in which the user is created. See Authentication Database.
	AuthenticationDatabase string

	// --authenticationMechanism <name>
	// Default: SCRAM-SHA-1
	//
	// Changed in version 2.6: Added support for the PLAIN and MONGODB-X509 authentication mechanisms.
	//
	// Changed in version 3.0: Added support for the SCRAM-SHA-1 authentication mechanism. Changed default mechanism to
	// SCRAM-SHA-1.
	//
	// Specifies the authentication mechanism the mongorestore instance uses to authenticate to the mongod or mongos.
	//
	// |-------------------|---------------------------------------------------------------------------------------------|
	// | Value             | Description                                                                                 |
	// |-------------------|---------------------------------------------------------------------------------------------|
	// | SCRAM-SHA-1       | RFC 5802 standard Salted Challenge Response Authentication Mechanism using the SHA1 hash    |
	// |                   | function.                                                                                   |
	// | MONGODB-CR        | MongoDB challenge/response authentication.                                                  |
	// | MONGODB-X509      | MongoDB TLS/SSL certificate authentication.                                                 |
	// | GSSAPI (Kerberos) | External authentication using Kerberos. This mechanism is available only in MongoDB         |
	// |                   | Enterprise.                                                                                 |
	// | PLAIN (LDAP SASL) | External authentication using LDAP. You can also use PLAIN for authenticating in-database   |
	// |                   | users. PLAIN transmits passwords in plain text. This mechanism is available only in MongoDB |
	// |                   | Enterprise.                                                                                 |
	// |-------------------|---------------------------------------------------------------------------------------------|
	// fixme criar constantes
	AuthenticationMechanism string

	// --gssapiServiceName
	// New in version 2.6.
	//
	// Specify the name of the service using GSSAPI/Kerberos. Only required if the service does not use the default name
	// of mongodb.
	//
	// This option is available only in MongoDB Enterprise.
	GssapiServiceName bool

	// --gssapiHostName
	// New in version 2.6.
	//
	// Specify the hostname of a service using GSSAPI/Kerberos. Only required if the hostname of a machine does not match the hostname resolved by DNS.
	//
	// This option is available only in MongoDB Enterprise.
	GssapiHostName bool

	// --db <database>, -d <database>
	// Specifies a database for mongorestore to restore data into. If the database does not exist, mongorestore creates the database. If you do not specify a <db>, mongorestore creates new databases that correspond to the databases where data originated and data may be overwritten. Use this option to restore data into a MongoDB instance that already has data.
	//
	//   --db does not control which BSON files mongorestore restores. You must use the mongorestore path option to limit that restored data.
	Db string

	// --collection <collection>, -c <collection>
	// Specifies a single collection for mongorestore to restore. If you do not specify --collection, mongorestore takes the collection name from the input filename. If the input file has an extension, MongoDB omits the extension of the file from the collection name.
	Collection string

	// --nsExclude <namespace pattern>
	// New in version 3.4.
	//
	// Excludes the specified namespaces from the restore operation.
	//
	// --nsExclude accepts a namespace pattern as its argument. The namespace pattern permits --nsExclude to refer to any namespace that matches the specified pattern. mongorestore matches the smallest valid occurence of the namespace pattern.
	//
	// Use asterisks (*) as wild cards. Escape all literal asterisks and backslashes with a backslash. Restore Collections Using Wild Cards provides an example of using asterisks as wild cards.
	NsExclude string

	// --nsInclude <namespace pattern>
	// New in version 3.4.
	//
	// Includes only the specified namespaces in the restore operation. By enabling you to specify multiple collections to restore, --nsInclude offers a superset of the functionality of the --collection option.
	//
	// --nsInclude accepts a namespace pattern as its argument. The namespace pattern permits --nsInclude to refer to any namespace that matches the specified pattern. mongorestore matches the smallest valid occurence of the namespace pattern.
	//
	// Use asterisks (*) as wild cards. Escape all literal asterisks and backslashes with a backslash. Restore Collections Using Wild Cards provides an example of using asterisks as wild cards.
	NsInclude string

	// --nsFrom <namespace pattern>
	// New in version 3.4.
	//
	// Use with --nsTo to rename a namespace during the restore operation. --nsFrom specifies the collection in the dump file, while --nsTo specifies the name that should be used in the restored database.
	//
	// --nsFrom accepts a namespace pattern as its argument. The namespace pattern permits --nsFrom to refer to any namespace that matches the specified pattern. mongorestore matches the smallest valid occurence of the namespace pattern.
	//
	// For simple replacements, use asterisks (*) as wild cards. Escape all literal asterisks and backslashes with a backslash. Replacements correspond linearly to matches: each asterisk in --nsFrom must correspond to an asterisk in --nsTo, and the first asterisk in --nsFrom matches the first asterisk in nsTo.
	//
	// For more complex replacements, use dollar signs to delimit a “wild card” variable to use in the replacement. Change Collections’ Namespaces during Restore provides an example of complex replacements with dollar sign-delimited wild cards.
	//
	// Unlike replacements with asterisks, replacements with dollar sign-delimited wild cards do not need to be linear.
	NsFrom string

	// --nsTo <namespace pattern>
	// New in version 3.4.
	//
	// Use with --nsFrom to rename a namespace during the restore operation. --nsTo specifies the new collection name to use in the restored database, while --nsFrom specifies the name in the dump file.
	//
	// --nsTo accepts a namespace pattern as its argument. The namespace pattern permits --nsTo to refer to any namespace that matches the specified pattern. mongorestore matches the smallest valid occurence of the namespace pattern.
	//
	// For simple replacements, use asterisks (*) as wild cards. Escape all literal asterisks and backslashes with a backslash. Replacements correspond linearly to matches: each asterisk in --nsFrom must correspond to an asterisk in --nsTo, and the first asterisk in --nsFrom matches the first asterisk in nsTo.
	//
	// For more complex replacements, use dollar signs to delimit a “wild card” variable to use in the replacement. Change Collections’ Namespaces during Restore provides an example of complex replacements with dollar sign-delimited wild cards.
	//
	// Unlike replacements with asterisks, replacements with dollar sign-delimited wild cards do not need to be linear.
	NsTo string

	// --objcheck
	// Forces mongorestore to validate all requests from clients upon receipt to ensure that clients never insert invalid documents into the database. For objects with a high degree of sub-document nesting, --objcheck can have a small impact on performance.
	//
	// Changed in version 2.4: MongoDB enables --objcheck by default, to prevent any client from inserting malformed or invalid BSON into a MongoDB database.
	Objcheck bool

	// --drop
	// Before restoring the collections from the dumped backup, drops the collections from the target database. --drop does not drop collections that are not in the backup.
	//
	// When the restore includes the admin database, mongorestore with --drop removes all user credentials and replaces them with the users defined in the dump file. Therefore, in systems with authorization enabled, mongorestore must be able to authenticate to an existing user and to a user defined in the dump file. If mongorestore can’t authenticate to a user defined in the dump file, the restoration process will fail, leaving an empty database.
	Drop bool

	// --dryRun
	// New in version 3.4.
	//
	// Runs mongorestore without actually importing any data, returning the mongorestore summary information. Use with --verbose to produce more detailed summary information.
	DryRun bool

	// --oplogReplay
	// After restoring the database dump, replays the oplog entries from the oplog.bson file located in the top level of the dump directory. When used in conjunction with mongodump --oplog, mongorestore --oplogReplay restores the database to the point-in-time backup captured with the mongodump --oplog command. For an example of --oplogReplay, see Restore Point in Time Oplog Backup.
	//
	// mongorestore --oplogReplay replays any valid oplog.bson file found in the top level of the dump directory. That is, if you have a bson file that contains valid oplog entries, you can name the file oplog.bson and move it to the top level of the dump directory for mongorestore --oplogReplay to replay.
	//
	// SEE ALSO
	//   mongorestore Required Access
	OplogReplay bool

	//fixme: timestamp
	// --oplogLimit <timestamp>
	// Prevents mongorestore from applying oplog entries with timestamp newer than or equal to <timestamp>. Specify <timestamp> values in the form of <time_t>:<ordinal>, where <time_t> is the seconds since the UNIX epoch, and <ordinal> represents a counter of operations in the oplog that occurred in the specified second.
	//
	// You must use --oplogLimit in conjunction with the --oplogReplay option.
	OplogLimit string

	// --oplogFile <path>
	// New in version 3.4.
	//
	// Specifies the path to the oplog file containing oplog data for the restore.
	OplogFile string

	// --keepIndexVersion
	// Prevents mongorestore from upgrading the index to the latest version during the restoration process.
	KeepIndexVersion bool

	// --noIndexRestore
	// Prevents mongorestore from restoring and building indexes as specified in the corresponding mongodump output.
	NoIndexRestore bool

	// --noOptionsRestore
	// Prevents mongorestore from setting the collection options, such as those specified by the collMod database command, on restored collections.
	NoOptionsRestore bool

	// --restoreDbUsersAndRoles
	// Restore user and role definitions for the given database. See system.roles Collection and system.users Collection for more information.
	RestoreDbUsersAndRoles bool

	// --writeConcern <document>
	// Default: majority
	//
	// Specifies the write concern for each write operation that mongorestore writes to the target database.
	//
	// Specify the write concern as a document with w options.
	WriteConcern string

	// --maintainInsertionOrder
	// Default: False
	//
	// If specified, mongorestore inserts the documents in the order of their appearance in the input source, otherwise mongorestore may perform the insertions in an arbitrary order.
	MaintainInsertionOrder bool

	// --numParallelCollections int, -j int
	// Default: 4
	//
	// Number of collections mongorestore should restore in parallel.
	//
	// If you specify -j when restoring a single collection, -j maps to the --numInsertionWorkersPerCollection option rather than --numParallelCollections.
	NumParallelCollections int64

	// --numInsertionWorkersPerCollection int
	// Default: 1
	//
	// New in version 3.0.0.
	//
	// Specifies the number of insertion workers to run concurrently per collection.
	//
	// For large imports, increasing the number of insertion workers may increase the speed of the import.
	NumInsertionWorkersPerCollection int64

	// --stopOnError
	// New in version 3.0.
	//
	// Forces mongorestore to halt the restore when it encounters an error.
	StopOnError bool

	// --bypassDocumentValidation
	// Enables mongorestore to bypass document validation during the operation. This lets you insert documents that do not meet the validation requirements.
	//
	// New in version 3.2.1.
	BypassDocumentValidation bool

	// --gzip
	// New in version 3.2.
	//
	// Restores from compressed files or data stream created by mongodump --archive
	//
	// To restore from a dump directory that contains compressed files, run mongorestore with the new --gzip option.
	//
	// To restore from a compressed archive file, run mongorestore with the --gzip option in conjunction with the --archive option.
	//
	// <path>
	// The final argument of the mongorestore command is a directory path. This argument specifies the location of the database dump from which to restore.
	//
	// You cannot specify both the <path> argument and the --dir option, which also specifies the dump directory, to mongorestore.
	GZip bool

	// --archive <file or null>
	// New in version 3.2.
	//
	// Restores from an archive file or from the standard input (stdin).
	//
	// To restore from an archive file, run mongorestore with the --archive option and the archive filename.
	//
	// To restore from the standard input, run mongorestore with the archive option but omit the filename.
	//
	//   NOTE
	//     You cannot use the --archive option with the --dir option.
	//     mongorestore still supports the positional - parameter to restore a single collection from the standard input.
	Archive string

	// --out <path>, -o <path>
	// Specifies the directory where mongodump will write BSON files for the dumped databases. By default, mongodump saves
	// output files in a directory named dump in the current working directory.
	//
	// To send the database dump to standard output, specify “-” instead of a path. Write to standard output if you want
	// process the output before saving it, such as to use gzip to compress the dump. When writing standard output,
	// mongodump does not write the metadata that writes in a <dbname>.metadata.json file when writing to files directly.
	//
	// You cannot use the --archive option with the --out option.
	Out string

	// --dir string
	// Specifies the dump directory.
	//
	// You cannot specify both the --dir option and the <path> argument, which also specifies the dump directory, to
	// mongorestore.
	//
	//   You cannot use the --archive option with the --dir option.
	Dir string
}

func (el *MongoBackupData) String() (string, error) {
	var outputLStr string = ""
	var filePathLStr string
	var errLErr error

	if el.Host != "" && el.Port != "" {
		outputLStr += " --host=\"" + el.Host + ":" + el.Port + "\""
	} else if el.Host != "" && el.Port == "" {
		outputLStr += " --host=\"" + el.Host + "\""
	} else if el.Host == "" && el.Port != "" {
		outputLStr += " --host=\":" + el.Port + "\""
	}

	if el.SslCAFile != "" {
		filePathLStr, errLErr = filepath.Abs(el.SslCAFile)
		if errLErr != nil {
			return "", errLErr
		}
		outputLStr += " --sslCAFile=\"" + filePathLStr + "\""
	}

	if el.SslPEMKeyFile != "" {
		filePathLStr, errLErr = filepath.Abs(el.SslPEMKeyFile)
		if errLErr != nil {
			return "", errLErr
		}
		outputLStr += " --sslPEMKeyFile=\"" + filePathLStr + "\""
	}

	if el.SslPEMKeyPassword != "" {
		outputLStr += " --sslPEMKeyPassword=\"" + el.SslPEMKeyPassword + "\""
	}

	if el.SslCRLFile != "" {
		filePathLStr, errLErr = filepath.Abs(el.SslCRLFile)
		if errLErr != nil {
			return "", errLErr
		}
		outputLStr += " --sslCRLFile=\"" + filePathLStr + "\""
	}

	if el.Username != "" {
		outputLStr += " --username=\"" + el.Username + "\""
	}

	if el.Password != "" {
		outputLStr += " --password=\"" + el.Password + "\""
	}

	if el.AuthenticationDatabase != "" {
		outputLStr += " --authenticationDatabase=\"" + el.AuthenticationDatabase + "\""
	}

	if el.AuthenticationMechanism != "" {
		outputLStr += " --authenticationMechanism=\"" + el.AuthenticationMechanism + "\""
	}

	if el.Db != "" {
		outputLStr += " --db=\"" + el.Db + "\""
	}

	if el.Collection != "" {
		outputLStr += " --collection=\"" + el.Collection + "\""
	}

	if el.NsExclude != "" {
		outputLStr += " --nsExclude=\"" + el.NsExclude + "\""
	}

	if el.NsInclude != "" {
		outputLStr += " --nsInclude=\"" + el.NsInclude + "\""
	}

	if el.NsFrom != "" {
		outputLStr += " --nsFrom=\"" + el.NsFrom + "\""
	}

	if el.NsTo != "" {
		outputLStr += " --nsTo=\"" + el.NsTo + "\""
	}

	if el.OplogLimit != "" {
		outputLStr += " --oplogLimit=\"" + el.OplogLimit + "\""
	}

	if el.OplogFile != "" {
		filePathLStr, errLErr = filepath.Abs(el.OplogFile)
		if errLErr != nil {
			return "", errLErr
		}
		outputLStr += " --oplogFile=\"" + filePathLStr + "\""
	}

	if el.WriteConcern != "" {
		outputLStr += " --writeConcern=\"" + el.WriteConcern + "\""
	}

	if el.NumParallelCollections != 0 {
		outputLStr += " --numParallelCollections=" + strconv.FormatInt(el.NumParallelCollections, 10)
	}

	if el.NumInsertionWorkersPerCollection != 0 {
		outputLStr += " --numInsertionWorkersPerCollection=" + strconv.FormatInt(el.NumInsertionWorkersPerCollection, 10)
	}

	if el.Archive != "" {
		filePathLStr, errLErr = filepath.Abs(el.Archive)
		if errLErr != nil {
			return "", errLErr
		}
		outputLStr += " --archive=\"" + filePathLStr + "\""
	}

	if el.Dir != "" {
		filePathLStr, errLErr = filepath.Abs(el.Dir)
		if errLErr != nil {
			return "", errLErr
		}
		outputLStr += " --dir=\"" + filePathLStr + "\""
	}

	if el.Out != "" {
		filePathLStr, errLErr = filepath.Abs(el.Out)
		if errLErr != nil {
			return "", errLErr
		}
		outputLStr += " --out=\"" + filePathLStr + "\""
	}

	if el.Objcheck == true {
		outputLStr += " --objcheck=true"
	}

	if el.Ssl == true {
		outputLStr += " --ssl=true"
	}

	if el.SslAllowInvalidCertificates == true {
		outputLStr += " --sslAllowInvalidCertificates=true"
	}

	if el.SslAllowInvalidHostNames == true {
		outputLStr += " --sslAllowInvalidHostNames=true"
	}

	if el.SslFIPSMode == true {
		outputLStr += " --sslFIPSMode=true"
	}

	if el.GssapiServiceName == true {
		outputLStr += " --gssapiServiceName=true"
	}

	if el.GssapiHostName == true {
		outputLStr += " --gssapiHostName=true"
	}

	if el.Drop == true {
		outputLStr += " --drop=true"
	}

	if el.DryRun == true {
		outputLStr += " --dryRun=true"
	}

	if el.OplogReplay == true {
		outputLStr += " --oplogReplay=true"
	}

	if el.KeepIndexVersion == true {
		outputLStr += " --keepIndexVersion=true"
	}

	if el.NoIndexRestore == true {
		outputLStr += " --noIndexRestore=true"
	}

	if el.NoOptionsRestore == true {
		outputLStr += " --noOptionsRestore=true"
	}

	if el.RestoreDbUsersAndRoles == true {
		outputLStr += " --restoreDbUsersAndRoles=true"
	}

	if el.MaintainInsertionOrder == true {
		outputLStr += " --maintainInsertionOrder=true"
	}

	if el.StopOnError == true {
		outputLStr += " --stopOnError=true"
	}

	if el.BypassDocumentValidation == true {
		outputLStr += " --bypassDocumentValidation=true"
	}

	if el.GZip == true {
		outputLStr += " --gzip"
	}

	return outputLStr, nil
}

func (el *MongoBackupData) ToCommand() ([]string, error) {
	var outputLStr []string = make([]string, 0)
	var filePathLStr string
	var errLErr error

	if el.Host != "" && el.Port != "" {
		outputLStr = append(outputLStr, "--host=\""+el.Host+":"+el.Port+"\"")
	} else if el.Host != "" && el.Port == "" {
		outputLStr = append(outputLStr, "--host=\""+el.Host+"\"")
	} else if el.Host == "" && el.Port != "" {
		outputLStr = append(outputLStr, "--host=\":"+el.Port+"\"")
	}

	if el.SslCAFile != "" {
		filePathLStr, errLErr = filepath.Abs(el.SslCAFile)
		if errLErr != nil {
			return []string{}, errLErr
		}
		outputLStr = append(outputLStr, "--sslCAFile=\""+filePathLStr+"\"")
	}

	if el.SslPEMKeyFile != "" {
		filePathLStr, errLErr = filepath.Abs(el.SslPEMKeyFile)
		if errLErr != nil {
			return []string{}, errLErr
		}
		outputLStr = append(outputLStr, "--sslPEMKeyFile=\""+filePathLStr+"\"")
	}

	if el.SslPEMKeyPassword != "" {
		outputLStr = append(outputLStr, "--sslPEMKeyPassword=\""+el.SslPEMKeyPassword+"\"")
	}

	if el.SslCRLFile != "" {
		filePathLStr, errLErr = filepath.Abs(el.SslPEMKeyFile)
		if errLErr != nil {
			return []string{}, errLErr
		}
		outputLStr = append(outputLStr, "--sslCRLFile=\""+el.SslCRLFile+"\"")
	}

	if el.Username != "" {
		outputLStr = append(outputLStr, "--username=\""+el.Username+"\"")
	}

	if el.Password != "" {
		outputLStr = append(outputLStr, "--password=\""+el.Password+"\"")
	}

	if el.AuthenticationDatabase != "" {
		outputLStr = append(outputLStr, "--authenticationDatabase=\""+el.AuthenticationDatabase+"\"")
	}

	if el.AuthenticationMechanism != "" {
		outputLStr = append(outputLStr, "--authenticationMechanism=\""+el.AuthenticationMechanism+"\"")
	}

	if el.Db != "" {
		outputLStr = append(outputLStr, "--db=\""+el.Db+"\"")
	}

	if el.Collection != "" {
		outputLStr = append(outputLStr, "--collection=\""+el.Collection+"\"")
	}

	if el.NsExclude != "" {
		outputLStr = append(outputLStr, "--nsExclude=\""+el.NsExclude+"\"")
	}

	if el.NsInclude != "" {
		outputLStr = append(outputLStr, "--nsInclude=\""+el.NsInclude+"\"")
	}

	if el.NsFrom != "" {
		outputLStr = append(outputLStr, "--nsFrom=\""+el.NsFrom+"\"")
	}

	if el.NsTo != "" {
		outputLStr = append(outputLStr, "--nsTo=\""+el.NsTo+"\"")
	}

	if el.OplogLimit != "" {
		outputLStr = append(outputLStr, "--oplogLimit=\""+el.OplogLimit+"\"")
	}

	if el.OplogFile != "" {
		filePathLStr, errLErr = filepath.Abs(el.OplogFile)
		if errLErr != nil {
			return []string{}, errLErr
		}
		outputLStr = append(outputLStr, "--oplogFile=\""+filePathLStr+"\"")
	}

	if el.WriteConcern != "" {
		outputLStr = append(outputLStr, "--writeConcern=\""+el.WriteConcern+"\"")
	}

	if el.NumParallelCollections != 0 {
		outputLStr = append(outputLStr, "--numParallelCollections="+strconv.FormatInt(el.NumParallelCollections, 10))
	}

	if el.NumInsertionWorkersPerCollection != 0 {
		outputLStr = append(outputLStr, "--numInsertionWorkersPerCollection="+strconv.FormatInt(el.NumInsertionWorkersPerCollection, 10))
	}

	if el.Archive != "" {
		filePathLStr, errLErr = filepath.Abs(el.Archive)
		if errLErr != nil {
			return []string{}, errLErr
		}
		outputLStr = append(outputLStr, "--archive=\""+filePathLStr+"\"")
	}

	if el.Out != "" {
		filePathLStr, errLErr = filepath.Abs(el.Out)
		if errLErr != nil {
			return []string{}, errLErr
		}
		outputLStr = append(outputLStr, "--out=\""+filePathLStr+"\"")
	}

	if el.Dir != "" {
		filePathLStr, errLErr = filepath.Abs(el.Dir)
		if errLErr != nil {
			return []string{}, errLErr
		}
		outputLStr = append(outputLStr, "--dir=\""+filePathLStr+"\"")
	}

	if el.Objcheck == true {
		outputLStr = append(outputLStr, "--objcheck=true")
	}

	if el.Ssl == true {
		outputLStr = append(outputLStr, "--ssl=true")
	}

	if el.SslAllowInvalidCertificates == true {
		outputLStr = append(outputLStr, "--sslAllowInvalidCertificates=true")
	}

	if el.SslAllowInvalidHostNames == true {
		outputLStr = append(outputLStr, "--sslAllowInvalidHostNames=true")
	}

	if el.SslFIPSMode == true {
		outputLStr = append(outputLStr, "--sslFIPSMode=true")
	}

	if el.GssapiServiceName == true {
		outputLStr = append(outputLStr, "--gssapiServiceName=true")
	}

	if el.GssapiHostName == true {
		outputLStr = append(outputLStr, "--gssapiHostName=true")
	}

	if el.Drop == true {
		outputLStr = append(outputLStr, "--drop=true")
	}

	if el.DryRun == true {
		outputLStr = append(outputLStr, "--dryRun=true")
	}

	if el.OplogReplay == true {
		outputLStr = append(outputLStr, "--oplogReplay=true")
	}

	if el.KeepIndexVersion == true {
		outputLStr = append(outputLStr, "--keepIndexVersion=true")
	}

	if el.NoIndexRestore == true {
		outputLStr = append(outputLStr, "--noIndexRestore=true")
	}

	if el.NoOptionsRestore == true {
		outputLStr = append(outputLStr, "--noOptionsRestore=true")
	}

	if el.RestoreDbUsersAndRoles == true {
		outputLStr = append(outputLStr, "--restoreDbUsersAndRoles=true")
	}

	if el.MaintainInsertionOrder == true {
		outputLStr = append(outputLStr, "--maintainInsertionOrder=true")
	}

	if el.StopOnError == true {
		outputLStr = append(outputLStr, "--stopOnError=true")
	}

	if el.BypassDocumentValidation == true {
		outputLStr = append(outputLStr, "--bypassDocumentValidation=true")
	}

	if el.GZip == true {
		outputLStr = append(outputLStr, "--gzip")
	}

	return outputLStr, nil
}

func CreateBackup(command MongoBackupCommand, data *MongoBackupData) error {
	var errLErr error
	var commandNameLStr string
	var CommandDataLAStr []string
	//defer pWriter.Close()

	commandNameLStr, errLErr = exec.LookPath(command.String())
	if errLErr != nil {
		return errLErr
	}

	if data.Host == "" {
		data.Host = addressMStr
	}

	if data.Db == "" && command == MONGODB_BACKUP_DUMP {
		data.Db = nameDbMStr
	}

	if command == MONGODB_BACKUP_DUMP && data.Dir != "" {
		data.Out = data.Dir
		data.Dir = ""
	}

	if command == MONGODB_BACKUP_RESTORE && data.Out != "" {
		data.Dir = data.Out
		data.Out = ""
	}

	CommandDataLAStr, errLErr = data.ToCommand()

	cmd := exec.Command(commandNameLStr, CommandDataLAStr...)
	//cmd.Stdout = pWriter
	//cmd.Stderr = os.Stderr
	log.Infof("CMD: $ %s %s", commandNameLStr, strings.Join(cmd.Args, " "))
	errLErr = cmd.Run()
	if errLErr != nil {
		return errLErr
	}
	return nil
}
