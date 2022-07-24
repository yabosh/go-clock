# go-clock
Wrapper around time.Time that allows the clock to be changed to support unit testing.

It can be difficult to test scenarios that involve the passage of time.  In some cases, inserting simple delays in test cases can be adequate provided the time-intervals are small.  For instance, to test a scenario where a resource times out after a few seconds, it would be possible to simply wait for a few seconds so the resource can time out.  This is generally not desirable because it can increase the overall runtime of a test suite but it is possible.  Other scenarios require waiting significantly longer intervals of time or waiting for a specific time event such as the change from one month to the next.  It is possible to change the system clock to verify this type of scenario but that approach is clumsy and very difficult to automate.  

go-clock is designed to simplify the process of creating temporal unit tests.  It is a wrapper around the go `time.Time` module that allows the time returned to be modified at runtime.  Using go-clock is simple.  Replace all calls that access `time.Time.Now()` with calls to `go-clock.clock.Now()`.  

Example usage:

```go
// Retrieve using time.Time
t := time.Now()

// Retrieve using go-clock.  The value returned by clock.Now() can be varied during unit tests to show the passage of time.
t = clock.Now()
```

## Testing with go-clock
go-clock can be used to simulate the passage of time in a unit test without the need to insert pauses.  This allows for the creation of temporal tests without slowing down execution of a test suite.  The following is an example of a unit test that simulates time passage.

###Function to be tested

```go
func myfunc() {

}
```

### Unit test for myfunc

```go

  // Configure go-clock to 'freeze' time.  This resets the clock back to Jan 1, 1970.  Repeated calls to clock.Now() will return
  // the same value and time appears to have frozen.
  clock.SetFixedClock()
  
  
```
