package tests

import (
	"testing"
	"time"

	"github.com/dakom/libhdate-go/libhdate"
)

func TestDate(t *testing.T) {
	hDate := &libhdate.HebDate{}

	// Set the Date
	hDate.SetGdate(25, 3, 2016)

	want := "15 Adar II 5776, Shushan Purim"
	got := hDate.String()

	checkStringMatch(t, "date", want, got)
}

func TestNow(t *testing.T) {

	hDate := &libhdate.HebDate{}
	hDate.SetGdate(0, 0, 0)

	now := time.Now()
	year, month, day := now.Date()

	checkIntMatch(t, "day", day, hDate.GetGday())
	checkIntMatch(t, "month", int(month), hDate.GetGmonth())
	checkIntMatch(t, "year", year, hDate.GetGyear())

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
