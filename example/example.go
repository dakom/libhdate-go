package main

import (
	"fmt"
	"math"
	"time"

	"github.com/dakom/libhdate-go/libhdate"
)

func main() {
	hDate := &libhdate.HebDate{}
	now := time.Now()

	// Set the Date
	libhdate.HdateSetGoTime(hDate, now)
	libhdate.HdateSetGdate(hDate, 25, 3, 2016)

	// get holydays
	holyday := libhdate.HdateGetHolyday(hDate, false)
	omer := libhdate.HdateGetOmerDay(hDate)
	reading := libhdate.HdateGetParasha(hDate, false)

	// HDATE_STRING_INT     0
	// HDATE_STRING_DOW     1
	// HDATE_STRING_PARASHA 2
	// HDATE_STRING_HMONTH  3
	// HDATE_STRING_GMONTH  4
	// HDATE_STRING_HOLIDAY 5
	// HDATE_STRING_OMER    6
	// HDATE STRING_SHORT   1
	// HDATE_STRING_LONG    0
	// HDATE_STRING_HEBREW  1
	// HDATE_STRING_LOCAL   0
	holydayName := libhdate.HdateString(libhdate.HDATE_STRING_HOLIDAY, holyday, false, false)

	readingName := libhdate.HdateString(libhdate.HDATE_STRING_PARASHA, reading, false, false)

	// get times
	latitude := 31.8903
	longitude := 35.0104

	times := make([]int, 8)

	libhdate.HdateGetUtcSunTimeFull(libhdate.HdateGetGday(hDate), libhdate.HdateGetGmonth(hDate), libhdate.HdateGetGyear(hDate), latitude, longitude, &times[0], &times[1], &times[2], &times[3], &times[4], &times[5], &times[6], &times[7])

	/*
		var timeZone = 3 * 60;
		// adjust time zone
		var timeStrings = times.slice(1).map(function (t) {
		t += timeZone;
		return "" + Math.floor(t / 60) + ":" + (t % 60);
		});
		console.log(timeStrings);
	*/

	timeZone := 3 * 60
	timeStrings := make([]string, len(times))
	for idx, timeVal := range times {
		timeVal += timeZone
		timeStrings[idx] = fmt.Sprintf("%d:%02d", int64(math.Floor(float64(timeVal)/60)), (timeVal % 60)) //fmt.Sprintf("%d:%d", math.Floor(float64(timeVal)/60), (timeVal % 60))
	}

	fmt.Printf("date: %v\n", hDate)
	fmt.Printf("omer: %v holyday: %v reading: %v\n", omer, holydayName, readingName)
	fmt.Printf("sun: %v first_light: %v talit: %v sunrise: %v midday: %v sunset: %v first_stars: %v three_stars: %v\n", timeStrings[0], timeStrings[1], timeStrings[2], timeStrings[3], timeStrings[4], timeStrings[5], timeStrings[6], timeStrings[7])
}