package libhdate

/** @struct hdate_struct
  @brief libhdate Hebrew date struct
*/
type HebDate struct {
	/** The number of day in the hebrew month (1..31). */
	hd_day int
	/** The number of the hebrew month 1..14 (1 - tishre, 13 - adar 1, 14 - adar 2). */
	hd_mon int
	/** The number of the hebrew year. */
	hd_year int
	/** The number of the day in the month. (1..31) */
	gd_day int
	/** The number of the month 1..12 (1 - jan). */
	gd_mon int
	/** The number of the year. */
	gd_year int
	/** The day of the week 1..7 (1 - sunday). */
	hd_dw int
	/** The length of the year in days. */
	hd_size_of_year int
	/** The week day of Hebrew new year. */
	hd_new_year_dw int
	/** The number type of year. */
	hd_year_type int
	/** The Julian day number */
	hd_jd int
	/** The number of days passed since 1 tishrey */
	hd_days int
	/** The number of weeks passed since 1 tishrey */
	hd_weeks int
}

func (h *HebDate) String() string {
	return HdateGetFormatDate(h, false, false)
}
