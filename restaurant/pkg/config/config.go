package config

type Config struct {
	Restaurant
	DB
	Customer
}

type Restaurant struct {
	HTTPAddr string `env:"GATEWAY_HTTP_ADDR" envDefault:":13103"`
	GRPCAddr string `env:"GATEWAY_GRPC_ADDR" envDefault:":13104"`
}

type DB struct {
	PgPort   string `env:"PG_PORT" envDefault:"5454"`
	PgHost   string `env:"PG_HOST" envDefault:"0.0.0.0"`
	PgDBName string `env:"PG_DB_NAME" envDefault:"db"`
	PgUser   string `env:"PG_USER" envDefault:"db"`
	PgPwd    string `env:"PG_PWD" envDefault:"db"`
}

type Customer struct {
	GRPCAddr string `env:"GATEWAY_GRPC_ADDR" envDefault:":13102"`
}
