package libhdate

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type HDateExtended struct {
	HebDate

	HolydayIndex      int
	OmerIndex         int
	ParshaIndex       int
	SunMinutes        int
	FirstLightMinutes int
	TalitMinutes      int
	SunriseMinutes    int
	MiddayMinutes     int
	SunsetMinutes     int
	FirstStarsMinutes int
	ThreeStarsMinutes int
}

func (h *HDateExtended) SetTime(t time.Time) {
	year, month, day := t.Date()

	h.SetGdate(day, int(month), year)
}

func (h *HDateExtended) GetTime() time.Time {
	//uses mid-day just to keep things simple...
	return time.Date(h.GetGyear(), time.Month(h.GetGmonth()), h.GetGday(), 12, 0, 0, 0, time.UTC)
}

func (h *HDateExtended) String() string {
	return h.GetFormatDate(false)
}

func (h *HDateExtended) Calculate(latitude float64, longitude float64, timezone int) {
	h.HolydayIndex = h.GetHolyday()
	h.OmerIndex = h.GetOmerDay()
	h.ParshaIndex = h.GetParasha()

	sun, firstLight, talit, sunrise, midday, sunset, firstStars, threeStars := getUtcSunTimeFull(h.GetGday(), h.GetGmonth(), h.GetGyear(), latitude, longitude)

	timezone *= 60
	h.SunMinutes = sun + timezone
	h.FirstLightMinutes = firstLight + timezone
	h.TalitMinutes = talit + timezone
	h.SunriseMinutes = sunrise + timezone
	h.MiddayMinutes = midday + timezone
	h.SunsetMinutes = sunset + timezone
	h.FirstStarsMinutes = firstStars + timezone
	h.ThreeStarsMinutes = threeStars + timezone

}

func (h *HDateExtended) GetHolydayString() string {
	return getStringByType(h, zmanimTypeHolyday)
}
func (h *HDateExtended) GetOmerString() string {
	return getStringByType(h, zmanimTypeOmer)
}
func (h *HDateExtended) GetParshaString() string {
	return getStringByType(h, zmanimTypeParsha)
}
func (h *HDateExtended) GetSunString() string {
	return getStringByType(h, zmanimTypeSun)
}
func (h *HDateExtended) GetFirstlightString() string {
	return getStringByType(h, zmanimTypeFirstlight)
}
func (h *HDateExtended) GetTalitString() string {
	return getStringByType(h, zmanimTypeTalit)
}
func (h *HDateExtended) GetSunriseString() string {
	return getStringByType(h, zmanimTypeSunrise)
}
func (h *HDateExtended) GetMiddayString() string {
	return getStringByType(h, zmanimTypeMidday)
}
func (h *HDateExtended) GetSunsetString() string {
	return getStringByType(h, zmanimTypeSunset)
}
func (h *HDateExtended) GetFirstStarsString() string {
	return getStringByType(h, zmanimTypeFirstStars)
}
func (h *HDateExtended) GetThreeStarsString() string {
	return getStringByType(h, zmanimTypeThreeStars)
}

const (
	_ = iota
	zmanimTypeHolyday
	zmanimTypeOmer
	zmanimTypeParsha
	zmanimTypeSun
	zmanimTypeFirstlight
	zmanimTypeTalit
	zmanimTypeSunrise
	zmanimTypeMidday
	zmanimTypeSunset
	zmanimTypeFirstStars
	zmanimTypeThreeStars
)

func getStringByType(h *HDateExtended, typeIndex int) string {
	switch typeIndex {
	case zmanimTypeHolyday:
		return GetString(HDATE_STRING_HOLIDAY, h.HolydayIndex, false, false)
	case zmanimTypeOmer:
		return intToOrdinalString(h.OmerIndex)
	case zmanimTypeParsha:
		return GetString(HDATE_STRING_PARASHA, h.ParshaIndex, false, false)
	case zmanimTypeSun:
		return minutesToString(h.SunMinutes)
	case zmanimTypeFirstlight:
		return minutesToString(h.FirstLightMinutes)
	case zmanimTypeTalit:
		return minutesToString(h.TalitMinutes)
	case zmanimTypeSunrise:
		return minutesToString(h.SunriseMinutes)
	case zmanimTypeMidday:
		return minutesToString(h.MiddayMinutes)
	case zmanimTypeSunset:
		return minutesToString(h.SunsetMinutes)
	case zmanimTypeFirstStars:
		return minutesToString(h.FirstStarsMinutes)
	case zmanimTypeThreeStars:
		return minutesToString(h.ThreeStarsMinutes)
	}

	return ""
}

func minutesToString(m int) string {
	return fmt.Sprintf("%d:%02d", int64(math.Floor(float64(m)/60)), (m % 60))
}

func intToOrdinalString(i int) string {
	j := i % 10
	k := i % 100
	istr := strconv.FormatInt(int64(i), 10)

	if j == 1 && k != 11 {
		return istr + "st"
	}
	if j == 2 && k != 12 {
		return istr + "nd"
	}
	if j == 3 && k != 13 {
		return istr + "rd"
	}
	return istr + "th"
}
