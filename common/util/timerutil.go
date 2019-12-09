package util

import (
	"time"
)

func SingleAtTime(f func(), t time.Time) {
	now := time.Now()
	timer := time.NewTimer(t.Sub(now))
	go func() {
		<-timer.C
		f()
	}()
}

func SingleDelayed(f func(), d time.Duration) {
	timer := time.NewTimer(d)
	go func() {
		<-timer.C
		f()
	}()
}

func ScheduleAtTime(f func(), hour int) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), hour, 0, 0, 0, next.Location())
			timer := time.NewTimer(next.Sub(now))
			<-timer.C
		}
	}()
}

func ScheduleDelayed(f func(), d time.Duration) {
	ticker := time.NewTicker(d)
	go func() {
		for range ticker.C {
			f()
		}
	}()
}
