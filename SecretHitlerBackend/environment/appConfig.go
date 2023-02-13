package environment

import "database/sql"

type AppConfig struct {
	DB *sql.DB
}
