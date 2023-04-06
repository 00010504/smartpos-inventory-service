package models

import "database/sql"

type NullShortShop struct {
	ID   sql.NullString
	Name sql.NullString
}
