package data_registry

import (
	"fmt"

	ent "local/auth_example/api/entities"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db sqlx.DB

// Open the DB connection and cache the query strings.
// Should be called once when the app starts.
func InitDB(connStr string) error {

	db_con, err := sqlx.Open("postgres", connStr)

	db = *db_con

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}

// Make sure the DB is accessible.
// Return an error if it's not
func PingDB() error {
	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}

// Upsert a user into the DB. Returns the user id.
func UpsertUser(user ent.User) error {

	fmt.Println(user)

	_, err := db.NamedExec(
		"insert into users (user_id, creation_time, provider, email) values (:user_id, :creation_time, :provider, :email) on conflict (email, provider) do nothing;",
		user,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
