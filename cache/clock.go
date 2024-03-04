/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache

import (
	"time"
)

type PassiveClock interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

type Clock interface {
	PassiveClock

	// After returns the channel of a new Timer.
	// This method does not allow to free/GC the backing timer before it fires. Use
	// NewTimer instead.
	After(d time.Duration) <-chan time.Time

	// NewTimer returns a new Timer.
	NewTimer(d time.Duration) Timer

	// Sleep sleeps for the provided duration d.
	Sleep(d time.Duration)

	// Tick returns the channel of a new Ticker.
	// This method does not allow to free/GC the backing ticker. Use
	// NewTicker from WithTicker instead.
	Tick(d time.Duration) <-chan time.Time
}

// Ticker defines the Ticker interface.
type Ticker interface {
	C() <-chan time.Time
	Stop()
}

// Timer allows for injecting fake or real timers into code that
// needs to do arbitrary things based on time.
type Timer interface {
	C() <-chan time.Time
	Stop() bool
	Reset(d time.Duration) bool
}

// RealClock really calls time.Now()
type RealClock struct{}

// Now returns the current time.
func (RealClock) Now() time.Time {
	return time.Now()
}

// Since returns time since the specified timestamp.
func (RealClock) Since(ts time.Time) time.Duration {
	return time.Since(ts)
}

// After is the same as time.After(d).
// This method does not allow to free/GC the backing timer before it fires. Use
// NewTimer instead.
func (RealClock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

// NewTimer is the same as time.NewTimer(d)
func (RealClock) NewTimer(d time.Duration) Timer {
	return &realTimer{
		timer: time.NewTimer(d),
	}
}

// AfterFunc is the same as time.AfterFunc(d, f).
func (RealClock) AfterFunc(d time.Duration, f func()) Timer {
	return &realTimer{
		timer: time.AfterFunc(d, f),
	}
}

// Tick is the same as time.Tick(d)
// This method does not allow to free/GC the backing ticker. Use
// NewTicker instead.
func (RealClock) Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}

// NewTicker returns a new Ticker.
func (RealClock) NewTicker(d time.Duration) Ticker {
	return &realTicker{
		ticker: time.NewTicker(d),
	}
}

// Sleep is the same as time.Sleep(d)
// Consider making the sleep interruptible by using 'select' on a context channel and a timer channel.
func (RealClock) Sleep(d time.Duration) {
	time.Sleep(d)
}

var _ = Timer(&realTimer{})

// realTimer is backed by an actual time.Timer.
type realTimer struct {
	timer *time.Timer
}

// C returns the underlying timer's channel.
func (r *realTimer) C() <-chan time.Time {
	return r.timer.C
}

// Stop calls Stop() on the underlying timer.
func (r *realTimer) Stop() bool {
	return r.timer.Stop()
}

// Reset calls Reset() on the underlying timer.
func (r *realTimer) Reset(d time.Duration) bool {
	return r.timer.Reset(d)
}

type realTicker struct {
	ticker *time.Ticker
}

func (r *realTicker) C() <-chan time.Time {
	return r.ticker.C
}

func (r *realTicker) Stop() {
	r.ticker.Stop()
}
