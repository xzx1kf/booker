package util

import (
	"strconv"
	"time"
)

// The timeslot of the court can be calculated given that we know the first timeslot
// for each court. If we take the time in minutes remaining after the court time
// is subtracted from the first timeslot time and divide it by 40 minutes we get
// the timeslot for that court. This value then needs to be adjusted depending
// on which court is being booked. Court 1 requires zero adjustment. Each subsequent
// court requires 20 timeslots adding on.
func Timeslot(court, date string) string {
	startTimes := [6]string{"00:00", "09:10", "09:10", "09:00", "09:00", "09:05"}

	layout := "15:04"

	// court
	c, _ := strconv.Atoi(court)
	courtTime, _ := time.Parse(layout, date[11:])
	startTime, _ := time.Parse(layout, startTimes[c])

	// Remainder minutes
	r := courtTime.Sub(startTime)

	// Calculate the slot for the court. 1 - 20
	courtSlot := (int(r.Minutes()) + 40) / 40

	// Calculate the slot for the day, this value is dependant
	// on the court selected 1 - 100.
	daySlot := courtSlot + ((c - 1) * 20)

	return strconv.Itoa(daySlot)
}
