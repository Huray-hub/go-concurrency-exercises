//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed

// Beginner level
// func HandleRequest(process func(), u *User) bool {
//     var tick <-chan time.Time
// 	if !u.IsPremium {
//         tick = time.Tick(10 * time.Second)
// 	}
//
// 	done := make(chan bool)
// 	go func() {
// 		process()
// 		done <- true
// 	}()
//
// 	select {
// 	case <-tick:
// 		return false
// 	case <-done:
// 		return true
// 	}
// }

// Advanced level
func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	}

	if u.TimeUsed >= 10 {
		return false
	}

	done := make(chan bool)
	tick := time.Tick(time.Duration(time.Second))
	go func(procCh chan<- bool) {
		process()
		procCh <- true
	}(done)

	for {
		select {
		case <-tick:
			if u.TimeUsed++; u.TimeUsed >= 10 {
				return false
			}
		case <-done:
			return true
		}
	}
}

func main() {
	RunMockServer()
}
