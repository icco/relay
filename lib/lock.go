package lib

import (
	"database/sql"
	"time"

	"cirello.io/pglock"
)

// GetLockClient returns a lock client with default settings.
func GetLockClient(db *sql.DB) (*pglock.Client, error) {
	return pglock.New(db,
		pglock.WithLeaseDuration(3*time.Second),
		pglock.WithHeartbeatFrequency(1*time.Second),
	)
}
