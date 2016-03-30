package main

import (
	"fmt"
	"time"

	"github.com/dakom/libhdate-go/libhdate"
)

func main() {
	latitude := 31.8903
	longitude := 35.0104
	timezone := 3

	h := &libhdate.HDateExtended{}

	h.SetTime(time.Now())
	//or a specific date
	//h.SetGdate(25, 3, 2016) //Shushan Purim
	//h.SetGdate(26, 3, 2016) //Shabbos (parsha reading tzav)
	//h.SetGdate(7, 5, 2016) //14th day of Omer parsha kedoshim

	h.Calculate(latitude, longitude, timezone)

	fmt.Printf("date: %v\n", h)
	fmt.Printf("Gregorian dd-mm-yyyy: %02d-%02d-%04d\n", h.GetGday(), h.GetGmonth(), h.GetGyear())
	fmt.Printf("Hebrew dd-mm-yyyy: %02d-%02d-%04d\n", h.GetHday(), h.GetHmonth(), h.GetHyear())
	fmt.Printf("omer: %v holyday: %v reading: %v\n", h.GetOmerString(), h.GetHolydayString(), h.GetParshaString())
	fmt.Printf("sun: %v first_light: %v talit: %v sunrise: %v midday: %v sunset: %v first_stars: %v three_stars: %v\n", h.GetSunString(), h.GetFirstlightString(), h.GetTalitString(), h.GetSunriseString(), h.GetMiddayString(), h.GetSunsetString(), h.GetFirstStarsString(), h.GetThreeStarsString())
}
