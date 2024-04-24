package config

type MySQL struct {
	WagersDB struct {
		Master string
		Slave  string
	}
	GameDB struct {
		Master string
		Slave  string
	}
}
