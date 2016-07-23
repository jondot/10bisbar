package main

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

type CLISuite struct{}

var _ = Suite(&CLISuite{})

func (s *CLISuite) TestFoobar(c *C) {
	//should fail - fix me!
	//c.Check(42, Equals, 0)
}

func TestDaysLeftEndMonthSundayToThursday(t *testing.T) {

	d := time.Date(2016, 07, 31, 0, 0, 0, 0, time.UTC)//Friday
	israelWorkingDayDefault := true
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestDaysLeftEndMonthMondayToFriday(t *testing.T) {

	d := time.Date(2016, 07, 31, 0, 0, 0, 0, time.UTC)//Friday
	israelWorkingDayDefault := false
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 0 {
		t.Error("Expected 0, got ", v)
	}
}

func TestDaysLeftStartOfMonthSundayToThursday(t *testing.T) {
	d := time.Date(2016, 07, 3, 0, 0, 0, 0, time.UTC)//Sunday
	israelWorkingDayDefault := true
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 21 {
		t.Error("Expected 20, got ", v)
	}
}

func TestDaysLeftStartOfMonthMondayToFriday(t *testing.T) {
	d := time.Date(2016, 07, 3, 0, 0, 0, 0, time.UTC)//Sunday
	israelWorkingDayDefault := false
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 20 {
		t.Error("Expected 20, got ", v)
	}
}

func TestDaysLeftStartOfWeekSundayToThursday(t *testing.T) {
	d := time.Date(2016, 07, 24, 0, 0, 0, 0, time.UTC)//Sunday
	israelWorkingDayDefault := true
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 6 {
		t.Error("Expected 6, got ", v)
	}
}


func TestDaysLeftStartOfWeekMondayToFriday(t *testing.T) {
	d := time.Date(2016, 07, 24, 0, 0, 0, 0, time.UTC)//Sunday
	israelWorkingDayDefault := false
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}


func TestDaysLeftEndOfWeekSundayToThursday(t *testing.T) {
	d := time.Date(2016, 07, 14, 0, 0, 0, 0, time.UTC)//Sunday
	israelWorkingDayDefault := true
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 12 {
		t.Error("Expected 12, got ", v)
	}
}


func TestDaysLeftEndOfWeekMondayToFriday(t *testing.T) {
	d := time.Date(2016, 07, 14, 0, 0, 0, 0, time.UTC)//Sunday
	israelWorkingDayDefault := false
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 12 {
		t.Error("Expected 12, got ", v)
	}
}
func TestDaysLeftStartMiddleOfMonthSundayToThursday(t *testing.T) {
	d := time.Date(2016, 07, 12, 0, 0, 0, 0, time.UTC)//Sunday
	israelWorkingDayDefault := true
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 14 {
		t.Error("Expected 14, got ", v)
	}
}
func TestDaysLeftStartMiddleOfMonthMondayToFriday(t *testing.T) {
	d := time.Date(2016, 07, 12, 0, 0, 0, 0, time.UTC)//Sunday
	israelWorkingDayDefault := false
	v := daysLeft(d, israelWorkingDayDefault)
	if v != 14 {
		t.Error("Expected 14, got ", v)
	}
}

func TestIsWeekDaySundayToThursday(t *testing.T) {
	d := time.Date(2016, 07, 24, 0, 0, 0, 0, time.UTC)//Sunday
	v := isWeekDay(d,true)
	if v != true {
		t.Error("Expected true, got ", v)
	}
}


func TestIsWeekDayMondayToFriday1(t *testing.T) {
	d := time.Date(2016, 07, 29, 0, 0, 0, 0, time.UTC)//Friday
	v := isWeekDay(d,true)
	if v != false {
		t.Error("Expected false, got ", v)
	}
}

func TestIsWeekDayMondayToFriday2(t *testing.T) {
	d := time.Date(2016, 07, 29, 0, 0, 0, 0, time.UTC)//Friday
	v := isWeekDay(d,false)
	if v != true {
		t.Error("Expected ture, got ", v)
	}
}

func TestIsWeekDayMiddleOfWeek(t *testing.T) {
	d := time.Date(2016, 07, 26, 0, 0, 0, 0, time.UTC)//Tuesday
	v1 := isWeekDay(d,false)
	v2 := isWeekDay(d,false)
	if v1 != true && v2 != true {
		t.Error("Expected ture, got ", v1,v2)
	}


}

