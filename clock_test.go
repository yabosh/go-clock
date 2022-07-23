package clock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_set_fixed_clock(t *testing.T) {
	defer ResetClock()

	// Given that the clock has been set to a fixed value
	SetFixedClock()

	// When the clock value is retrieved
	now1 := Now()
	time.Sleep(10 * time.Millisecond)
	now2 := Now()

	// It should be the fixed value
	assert.Equal(t, FixedBaseline, now1)
	assert.Equal(t, FixedBaseline, now2)
}

func Test_reset_fixed_clock(t *testing.T) {
	defer ResetClock()

	// Given that the clock has been set to a fixed value
	SetFixedClock()

	// When the clock value is retrieved
	now1 := Now()
	time.Sleep(10 * time.Millisecond)
	now2 := Now()
	ResetClock()
	now3 := Now()

	// It should be the fixed value
	assert.Equal(t, FixedBaseline, now1)
	assert.Equal(t, FixedBaseline, now2)
	assert.NotEqual(t, FixedBaseline, now3)
}

func Test_max(t *testing.T) {
	SetFixedClock()
	defer ResetClock()

	// Given two different times
	time1 := Now()
	time2 := Now()

	// When Max() is executed
	// It should return the greater of the two dates
	assert.Equal(t, time2, AppClock.Max(time1, time2))
	assert.Equal(t, time2, AppClock.Max(time2, time1))
}

func Test_advance_clock(t *testing.T) {
	SetFixedClock()
	defer ResetClock()

	// Given a fixed clock
	time1 := Now()

	// When the clock is advanced
	AdvanceFixed(0, 0, 1)

	// It should increment by the specified amount
	assert.True(t, Now().After(time1))
}

func Test_tick(t *testing.T) {
	SetFixedClock()
	defer ResetClock()

	// Given a fixed clock
	time1 := Now()

	// When the Tick() runs
	Tick()

	// It should increment by one second
	assert.True(t, Now().After(time1))
}
