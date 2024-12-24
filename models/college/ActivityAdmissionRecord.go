package college

import "time"

type ActivityAdmissionRecord struct {
	RecordId        uint
	ActivityName    string
	StudentName     string
	StudentId       string
	ApplicationTime time.Time
	Status          int
}
