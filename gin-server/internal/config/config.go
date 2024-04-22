package config

type Config struct {
	// rest.RestConf
	Name string
	Host string
	Port string

	WagersDB struct {
		Master string
		Slave  string
	}

	GameDB struct {
		Master string
		Slave  string
	}

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
