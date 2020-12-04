package model

// Config ...
type Config struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`
	Debug   bool   `json:"debug"`

	DB   *Postgres `json:"db"`
	Flip *Flip     `json:"flip"`
}

// Postgres ...
type Postgres struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Pass         string `json:"pass"`
	Name         string `json:"name"`
	MaxOpenConn  int    `json:"max_open_conn"`
	MaxIdleCount int    `json:"max_idle_count"`
}

// Flip ...
type Flip struct {
	Host          string `json:"host"`
	Authorization string `json:"authorization"`
}
