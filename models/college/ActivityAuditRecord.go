package college

import "time"

type ActivityAuditRecord struct {
	RecordId        uint
	ActivityName    string
	StartTime       time.Time
	EndTime         time.Time
	ActivityAddress string
	Organizer       string
	ApplicationTime time.Time
	Status          uint
}
