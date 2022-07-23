package clock

import (
	"time"
)

/*

   Provide basic time functions to allow time to be mocked

*/

// FixedBaseline is the date and time that is used to initialize the
// baseline clock before it is advanced.
var FixedBaseline = time.Date(1970, 01, 01, 00, 00, 00, 00, time.UTC)

// Clock represents an abstraction of time.Now() that can be modified
// to provide different values for 'now'
type Clock struct {
	nowFunc func() time.Time
}

// AppClock provides a global replacement for time.Now
var AppClock = &Clock{
	nowFunc: time.Now,
}

// Now retrieves the current time from the current implementation of AppClock
func Now() time.Time {
	return AppClock.Now()
}

// Now retrieves the current time from the clock defined for AppClock.
// The source of the clock used can be changed to provide a deterministic clock
func (c *Clock) Now() time.Time {
	return c.nowFunc()
}

// SetNow changes the source for Now() to a function that returns an instance
// of time.Time.
func (c *Clock) SetNow(newNow func() time.Time) {
	c.nowFunc = newNow
}

// ResetNow changes the clock source to time.Now
func (c *Clock) ResetNow() {
	c.nowFunc = time.Now
}

// Max returns the latest date from the two provided
func (c *Clock) Max(t1 time.Time, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t1
	}

	return t2
}

// Used to set a deterministic value for the time.Now() function when used by
// filter-related operations.
var fixedClock time.Time

// SetFixedClock will set the AppClock to a fixed time.  Repeated calls to AppClock.Now() will
// always return the fixed value, it will not increment with the passage of time.
func SetFixedClock() {
	// Unix epoch timestamp 1
	fixedClock = FixedBaseline
	AppClock.SetNow(func() time.Time {
		return fixedClock
	})
}

// ResetClock will set the AppClock to time.Now().
func ResetClock() {
	AppClock.ResetNow()
}

// AdvanceFixed will advance the fixed clock forward the specified number of hours, minutes, and seconds.
// Once the clock is advanced it will return the same time for subsequent calls until it is advanced again
func AdvanceFixed(hours int, minutes int, seconds int) {
	fixedClock = fixedClock.Add(time.Duration(hours) * time.Hour)
	fixedClock = fixedClock.Add(time.Duration(minutes) * time.Minute)
	fixedClock = fixedClock.Add(time.Duration(seconds) * time.Second)
}

// Tick advances the fixed clock one second.  Used to simplify simulating
// the passage of time when testing.
func Tick() {
	AdvanceFixed(0, 0, 1)
}

// AdvanceHours will advance the fixed clock by the specified number of hours
//
// This is a convience function that can be used to make unit tests more readable.
func AdvanceHours(delta int) {
	fixedClock = fixedClock.Add(time.Duration(delta) * time.Hour)
}

// AdvanceMinutes will advance the fixed clock by the specified number of minutes
//
// This is a convience function that can be used to make unit tests more readable.
func AdvanceMinutes(delta int) {
	fixedClock = fixedClock.Add(time.Duration(delta) * time.Minute)
}

// AdvanceSeconds will advance the fixed clock by the specified number of seconds
//
// This is a convience function that can be used to make unit tests more readable.
func AdvanceSeconds(delta int) {
	fixedClock = fixedClock.Add(time.Duration(delta) * time.Second)
}
