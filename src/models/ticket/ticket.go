package ticket

import (
	"time"
)

// Ticket is ticket
type Ticket struct {
	Slot               int    // Number of slot
	RegistrationNumber string // Car registration number
	CreatedAt          int    // Create time
	UpdatedAt          int    // Update time
	Status             int    // Ticket status. 1 = Active, 2 = Deleted, 0 = Inactive
	// ID                 int // Ticket ID
	// CreatedAt          time.Time // Create time
}

// var ID = 0
var Slot = 0

// Create new ticket
func CreateNew(registrationNumber string) *Ticket {
	// ID = ID + 1
	Slot = Slot + 1

	return &Ticket{
		Slot:               Slot,
		RegistrationNumber: registrationNumber,
		CreatedAt:          int(time.Now().Unix()), // 1558429930
		UpdatedAt:          0,
		Status:             1,
		// ID:                 ID,
		// CreatedAt:          time.Now(), // 2019-05-21 16:04:37.891132 +0700 WIB m=+0.000355928
	}
}
