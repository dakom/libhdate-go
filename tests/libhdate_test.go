package tests

import (
	"testing"
	"time"

	"github.com/dakom/libhdate-go/libhdate"
)

func TestDateString(t *testing.T) {
	hDate := &libhdate.HebDate{}

	// Set the Date
	hDate.SetGdate(25, 3, 2016)

	want := "15 Adar II 5776, Shushan Purim"
	got := hDate.String()

	checkStringMatch(t, "date", want, got)
}

func TestNowDefaults(t *testing.T) {

	hDate := &libhdate.HebDate{}
	hDate.SetGdate(0, 0, 0)

	now := time.Now()
	year, month, day := now.Date()

	checkIntMatch(t, "day", day, hDate.GetGday())
	checkIntMatch(t, "month", int(month), hDate.GetGmonth())
	checkIntMatch(t, "year", year, hDate.GetGyear())

}

func TestConversion(t *testing.T) {

	hDate := &libhdate.HebDate{}
	hDate.SetGdate(14, 5, 1948)

	checkIntMatch(t, "day", 5, hDate.GetHday())
	checkIntMatch(t, "month", 8, hDate.GetHmonth())
	checkIntMatch(t, "year", 5708, hDate.GetHyear())

	hDate = &libhdate.HebDate{}
	hDate.SetHdate(5, 8, 5708)

	checkIntMatch(t, "day", 14, hDate.GetGday())
	checkIntMatch(t, "month", 5, hDate.GetGmonth())
	checkIntMatch(t, "year", 1948, hDate.GetGyear())

	hDate = &libhdate.HebDate{}
	now := time.Now()
	hDate.SetTime(now)
	tm := hDate.GetTime()

	checkIntMatch(t, "day", tm.Day(), now.Day())
	checkIntMatch(t, "month", int(tm.Month()), int(now.Month()))
	checkIntMatch(t, "year", tm.Year(), now.Year())
	checkIntMatch(t, "yearday", tm.YearDay(), now.YearDay())
}

func TestExtendedCalculation(t *testing.T) {
	latitude := 31.8903
	longitude := 35.0104
	timezone := 3

	shushanPurim := &libhdate.HDateExtended{}
	shushanPurim.SetGdate(25, 3, 2016) //Shushan Purim
	shushanPurim.Calculate(latitude, longitude, timezone)

	parshaTzav := &libhdate.HDateExtended{}
	parshaTzav.SetGdate(26, 3, 2016) //Shabbos (parsha reading tzav)
	parshaTzav.Calculate(latitude, longitude, timezone)

	omerCountKedoshim := &libhdate.HDateExtended{}
	omerCountKedoshim.SetGdate(7, 5, 2016) //14th day of Omer parsha kedoshim
	omerCountKedoshim.Calculate(latitude, longitude, timezone)

	//checkIntMatch(t, "shushan purim", 14, shushanPurim.HolydayIndex)
	checkIntMatch(t, "shushan purim", 15, shushanPurim.HolydayIndex)
	checkIntMatch(t, "parsha tzav", 25, parshaTzav.ParshaIndex)
	checkIntMatch(t, "parsha kedoshim", 30, omerCountKedoshim.ParshaIndex)
	checkIntMatch(t, "day 14 of omer", 14, omerCountKedoshim.OmerIndex)

}

func checkStringMatch(t *testing.T, label string, want string, got string) {
	if want != got {
		t.Errorf("[%s] No match: [%s] != [%s]", label, got, want)
	}
}

func checkIntMatch(t *testing.T, label string, want int, got int) {
	if want != got {
		t.Errorf("[%s] No match: [%d] != [%d]", label, got, want)
	}
}
