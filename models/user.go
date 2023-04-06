package models

import "database/sql"

type NullShortUser struct {
	ID        sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
}
