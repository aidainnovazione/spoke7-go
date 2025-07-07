package postgres

type PostgresConf struct {
	Host         string
	Port         uint16
	DatabaseName string
	Username     string
	Password     string
	SSLMode      bool
	SSLRootCert  string
	BatchSize    int
}
