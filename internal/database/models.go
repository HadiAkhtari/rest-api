package database

import "database/sql"

type Models struct {
	Users     UserModel
	Events    EventsModel
	Attendees AttendeeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:     UserModel{db},
		Events:    EventsModel{db},
		Attendees: AttendeeModel{db},
	}
}
