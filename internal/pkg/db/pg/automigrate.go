package pg

import (
	"redditclone/internal/domain/user"
	"redditclone/internal/pkg/session"
)

func (db *DB) AutoMigrateAll() {
	db.DB().AutoMigrate(
		&user.User{},
		&session.Session{},
	)
}
