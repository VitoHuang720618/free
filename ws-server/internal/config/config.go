package config

type Config struct {
	// rest.RestConf
	Name string
	Host string
	Port string

	// db
	DB MySQL

	Redis struct {
		Host     string
		Password string
		Db       int
	}

	Etcd struct {
		Key   string
		Hosts []string
	}
}
