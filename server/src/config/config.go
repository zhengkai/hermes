package config

// config
var (
	Prod bool
	Dir  string

	MySQL     = `user:pass@/dbname`
	Port      = `:30023`
	StaticDir = `/www/hermes/static`
)
