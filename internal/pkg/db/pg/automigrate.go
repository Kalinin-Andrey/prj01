package pg

import (
	"carizza/internal/domain/user"
	"carizza/internal/pkg/session"
)

func (db *DB) AutoMigrateAll() {
	db.DB().AutoMigrate(
		&user.User{},
		&session.Session{},
	)
}
