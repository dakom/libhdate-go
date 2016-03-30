# libhdate-go
A pure go implementation of libhdate (http://libhdate.sourceforge.net)

With a few extra additions to make useability easier

## Testing

`go test ./tests`

## Usage

Runnable via `go run ./example/example.go`

```
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
```
## TODO

1. Cleanup some comments and function description which was left as-is from the C library
