package college

import "time"

type JAuditRecord struct {
	ID          uint
	StudentName string
	StudentID   string
	Status      int
	UpdatedAt   time.Time
}
