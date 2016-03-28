package tests

import (
	"testing"

	"github.com/dakom/libhdate-go/libhdate"
)

func TestDate(t *testing.T) {
	hDate := &libhdate.HebDate{}

	// Set the Date
	libhdate.HdateSetGdate(hDate, 25, 3, 2016)

	want := "15 Adar II 5776, Shushan Purim"
	got := hDate.String()

	if want != got {
		t.Errorf("No match: [%v] != [%v]", got, want)
	}
}
