package database

import "database/sql"

type Models struct {
	Users     UserModel
	Events    EventsModel
	Attendees AttendeeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:     UserModel{DB: db},
		Events:    EventsModel{DB: db},
		Attendees: AttendeeModel{DB: db},
	}
}
