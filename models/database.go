package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/moeabdol/birdpedia-golang/utils"
)

var store *Store

// ConnectDatabase function
func ConnectDatabase() {
	dbSource := fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=%s",
		utils.Config.DBDialect,
		utils.Config.DBUser,
		utils.Config.DBPassword,
		utils.Config.DBHost,
		utils.Config.DBPort,
		utils.Config.DBName,
		utils.Config.DBSslmode,
	)
	conn, err := sql.Open(utils.Config.DBDialect, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	store = New(conn)
}
