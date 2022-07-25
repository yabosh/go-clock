# go-clock

Wrapper around time.Time that allows the clock to be changed to support unit testing.

It can be difficult to test scenarios that involve the passage of time.  In some cases, inserting simple delays in test cases can be adequate provided the time-intervals are small.  For instance, to test a scenario where a resource times out after a few seconds, it would be possible to simply wait for a few seconds so the resource can time out.  This is generally not desirable because it can increase the overall runtime of a test suite but it is possible.  Other scenarios require waiting significantly longer intervals of time or waiting for a specific time event such as the change from one month to the next.  It is possible to change the system clock to verify this type of scenario but that approach is clumsy and very difficult to automate.  

go-clock is designed to simplify the process of creating temporal unit tests.  It is a wrapper around the go `time.Time` module that allows the time returned to be manipulated at runtime without changing the system clock.  

## Usage
There are two basic steps required to use go-clock:

* Replace calls to `time.Now()` with `clock.Now()`
* Manipulate clock during unit testing.

### Replace calls to time.Now()
The value returned by `time.Now()` is always based on the underlying system clock.  *go-clock* provides a means of returning a system-time that can be changed at runtime, primarily during unit testing.  In order for *go-clock* to work, any code that retrieves the time should use `clock.Now()` instead of `time.Time`.  `clock.Now()` returns an instance of `time.Time` just like `time.Now()` does and in most cases, `clock.Now()` is simply calls `time.Now()`.

Example:

```go
// At runtime, t1 and t2 return the value of the system clock.

// Retrieve using time.Time
t1 := time.Now()

// Retrieve using go-clock.  The value returned by clock.Now() can be varied during unit tests without changing the system clock.
t2 = clock.Now()
```

## Testing with go-clock
Any code that uses `clock.Now()` can be 
go-clock can be used to simulate the passage of time in a unit test without the need to insert pauses.  This allows for the creation of temporal tests without slowing down execution of a test suite.  The following is an example of a unit test that simulates time passage.

Common Test Scenarios
---
* Assert that a specific time value is available
* Wait for a specific time to pass to signal some sort of activity

### Asserting Specific Time Values
There are instances where it may be necessary to validate that a data structure has been populated with the expected time.  This can be difficult when using the system clock but simple with *go-clock*.

```go
	// Configure the Now() function to always return the same time
  // In this case, always return: 1970-01-01T00:00:00+0000  (UTC)
	clock.SetFixedClock()

	// Perform some action that uses the current time
	mystruct := getSomeData()

  // Verify that the call to getSomeData updated the Timestamp value as expected
  assert.Equal(t, "1970-01-01T00:00:00+0000", mystruct.Timestamp.Format(time.RFC3339))
  
```

### Wait for passage of time
Assume there is a function that waits until a specific point in time to take some action.  *go-clock* allows the fixed clock to be advanced to any point to simulate the passage of time.

```go
	// Configure the Now() function to always return the same time
  // In this case, always return: 1970-01-01T00:00:00+0000  (UTC)
	clock.SetFixedClock()
  
  // Perform some action that waits a specific point in time
  mystruct.RunAfter("1970-01-02T00:00:00+0000", myfunc) 
  
  assert.False(t, mystruct.HasExecuted)
  
  // Advance the clock by 1 hour and 1 minute
  clock.AdvanceFixed(1, 1, 0)
  
  assert.True(t, mystruct.HasExecuted)
  
  // Reset the clock to return the system time
  clock.ResetNow()
```

## go-clock functions

`SetFixedClock()` - Sets the clock to a fixed point in time, the beginning of the Unix epoch.  The value returned by clock.Now() will not change unless another clock function changes it.

`SetNowTime()` - Sets the clock to the specific time pointed to by a time.Time value. The value returned by clock.Now() will not change unless another clock function changes it.

`ResetClock()` - Restores the clock to return the system time.

`AdvanceFixed()` - Increments the value of the clock by a specific amount of time defined in hours, minute, and seconds.
