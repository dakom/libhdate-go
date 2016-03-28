package libhdate

import (
	"math"
)

/**
@brief days from 1 january

@parm day this day of month
@parm month this month
@parm year this year
@return the days from 1 jan
*/
func HdateGetDayOfYear(day int, month int, year int) int {
	var jd int

	/* get todays julian day number */
	jd = (1461*(year+4800+(month-14)/12))/4 + (367*(month-2-12*((month-14)/12)))/12 - (3*((year+4900+(month-14)/12)/100))/4 + day

	/* substruct the julian day of 1/1/year and add one */
	jd = jd - ((1461*(year+4799))/4 + 367*11/12 - (3*((year+4899)/100))/4)

	return jd
}

/**
@brief utc sun times for altitude at a gregorian date

Returns the sunset and sunrise times in minutes from 00:00 (utc time)
if sun altitude in sunrise is deg degries.
This function only works for altitudes sun realy is.
If the sun never get to this altitude, the returned sunset and sunrise values
will be negative. This can happen in low altitude when latitude is
nearing the pols in winter times, the sun never goes very high in
the sky there.

@param day this day of month
@param month this month
@param year this year
@param longitude longitude to use in calculations
@param latitude latitude to use in calculations
@param deg degrees of sun's altitude (0 -  Zenith .. 90 - Horizon)
@param sunrise return the utc sunrise in minutes
@param sunset return the utc sunset in minutes
*/
func HdateGetUtcSunTimeDeg(day int, month int, year int, latitude float64, longitude float64, deg float64, sunrise *int, sunset *int) {
	var gama float64                                  /* location of sun in yearly cycle in radians */
	var eqtime float64                                /* diffference betwen sun noon and clock noon */
	var decl float64                                  /* sun declanation */
	var ha float64                                    /* solar hour engle */
	var sunrise_angle float64 = math.Pi * deg / 180.0 /* sun angle at sunrise/set */

	var day_of_year int

	/* get the day of year */
	day_of_year = HdateGetDayOfYear(day, month, year)

	/* get radians of sun orbit around erth =) */
	gama = 2.0 * math.Pi * ((float64)(day_of_year-1) / 365.0)

	/* get the diff betwen suns clock and wall clock in minutes */
	eqtime = 229.18 * (0.000075 + 0.001868*math.Cos(gama) - 0.032077*math.Sin(gama) - 0.014615*math.Cos(2.0*gama) - 0.040849*math.Sin(2.0*gama))

	/* calculate suns declanation at the equater in radians */
	decl = 0.006918 - 0.399912*math.Cos(gama) + 0.070257*math.Sin(gama) - 0.006758*math.Cos(2.0*gama) + 0.000907*math.Sin(2.0*gama) - 0.002697*math.Cos(3.0*gama) + 0.00148*math.Sin(3.0*gama)

	/* we use radians, ratio is 2pi/360 */
	latitude = math.Pi * latitude / 180.0

	/* the sun real time diff from noon at sunset/rise in radians */
	ha = math.Acos(math.Cos(sunrise_angle)/(math.Cos(latitude)*math.Cos(decl)) - math.Tan(latitude)*math.Tan(decl))

	/* check for too high altitudes and return negative values */
	/* currently disabled, let the user give sensible data? */
	/*
		if errno == EDOM {
			*sunrise = -720
			*sunset = -720

			return
		}
	*/

	/* we use minutes, ratio is 1440min/2pi */
	ha = 720.0 * ha / math.Pi

	/* get sunset/rise times in utc wall clock in minutes from 00:00 time */
	*sunrise = (int)(720.0 - 4.0*longitude - ha - eqtime)
	*sunset = (int)(720.0 - 4.0*longitude + ha - eqtime)

	return
}

/**
@brief utc sunrise/set time for a gregorian date

@parm day this day of month
@parm month this month
@parm year this year
@parm longitude longitude to use in calculations
@parm latitude latitude to use in calculations
@parm sunrise return the utc sunrise in minutes
@parm sunset return the utc sunset in minutes
*/
func HdateGetUtcSunTime(day int, month int, year int, latitude float64, longitude float64, sunrise *int, sunset *int) {
	HdateGetUtcSunTimeDeg(day, month, year, latitude, longitude, 90.833, sunrise, sunset)

	return
}

/**
@brief utc sunrise/set time for a gregorian date

@parm day this day of month
@parm month this month
@parm year this year
@parm longitude longitude to use in calculations
@parm latitude latitude to use in calculations
@parm sun_hour return the length of shaa zaminit in minutes
@parm first_light return the utc alut ha-shachar in minutes
@parm talit return the utc tphilin and talit in minutes
@parm sunrise return the utc sunrise in minutes
@parm midday return the utc midday in minutes
@parm sunset return the utc sunset in minutes
@parm first_stars return the utc tzeit hacochavim in minutes
@parm three_stars return the utc shlosha cochavim in minutes
*/
func HdateGetUtcSunTimeFull(day int, month int, year int, latitude float64, longitude float64, sun_hour *int, first_light *int, talit *int, sunrise *int, midday *int, sunset *int, first_stars *int, three_stars *int) {
	var place_holder int

	/* sunset and rise time */
	HdateGetUtcSunTimeDeg(day, month, year, latitude, longitude, 90.833, sunrise, sunset)

	/* shaa zmanit by gara, 1/12 of light time */
	*sun_hour = (*sunset - *sunrise) / 12
	*midday = (*sunset + *sunrise) / 2

	/* get times of the different sun angles */
	HdateGetUtcSunTimeDeg(day, month, year, latitude, longitude, 106.01, first_light, &place_holder)
	HdateGetUtcSunTimeDeg(day, month, year, latitude, longitude, 101.0, talit, &place_holder)
	HdateGetUtcSunTimeDeg(day, month, year, latitude, longitude, 96.0, &place_holder, first_stars)
	HdateGetUtcSunTimeDeg(day, month, year, latitude, longitude, 98.5, &place_holder, three_stars)

	return
}
