# libhdate-go

A pure go implementation of libhdate (http://libhdate.sourceforge.net)

## Install

`go get github.com/dakom/libhdate-go/libhdate`

## Notes

This was done mostly in the blind, didn't touch the original libhdate calculations (at least I tried not to!)

It's more just a pure C to Go language port than anything else, with a few small tweaks due to the language and idiomatic differences

While all the functions from libhdate are there and ported over, they are not exported unless really intended to be used outside the package

There are a few additions (in extra.go) to make usability clearer and easier

## Usage

In general, use HDateExtended{}, run Calculate() and then grab properties or generated strings from there. See extra.go and the bottom half of julien.go for the available properties/methods. Here's an example:

```
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
	fmt.Printf("omer: %v holyday: %v reading: %v\n", h.GetOmerString(), h.GetHolydayString(), h.GetParshaString())
	fmt.Printf("sun: %v first_light: %v talit: %v sunrise: %v midday: %v sunset: %v first_stars: %v three_stars: %v\n", h.GetSunString(), h.GetFirstlightString(), h.GetTalitString(), h.GetSunriseString(), h.GetMiddayString(), h.GetSunsetString(), h.GetFirstStarsString(), h.GetThreeStarsString())
}
```

This is also runnable via `go run ./example/example.go`

## Testing

`go test ./tests`

## Todo

1. Cleanup some comments and function description which was left as-is from the C library
2. Better documentation, list all available exported methods
