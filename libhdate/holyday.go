package libhdate

/**
@brief Return number of hebrew holyday.

@param h The HebDate of the date to use.
@param diaspora if True give diaspora holydays
@return the number of holyday.
*/
func (h *HebDate) GetHolyday() int {
	var holyday int

	/* holydays table */
	holydays_table := [14][30]int{
		{ /* Tishrey */
			1, 2, 3, 3, 0, 0, 0, 0, 37, 4,
			0, 0, 0, 0, 5, 31, 6, 6, 6, 6,
			7, 27, 8, 0, 0, 0, 0, 0, 0, 0},
		{ /* Heshvan */
			0, 0, 0, 0, 0, 0, 0, 0, 0, 35,
			35, 35, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{ /* Kislev */
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 9, 9, 9, 9, 9, 9},
		{ /* Tevet */
			9, 9, 9, 0, 0, 0, 0, 0, 0, 10,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{ /* Shvat */
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 11, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 33},
		{ /* Adar */
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			12, 0, 12, 13, 14, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{ /* Nisan */
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 15, 32, 16, 16, 16, 16,
			28, 29, 0, 0, 0, 24, 24, 24, 0, 0},
		{ /* Iyar */
			0, 17, 17, 17, 17, 17, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 18, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 26, 0, 0},
		{ /* Sivan */
			0, 0, 0, 0, 19, 20, 30, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{ /* Tamuz */
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 21, 21, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 36, 36},
		{ /* Av */
			0, 0, 0, 0, 0, 0, 0, 0, 22, 22,
			0, 0, 0, 0, 23, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{ /* Elul */
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{ /* Adar 1 */
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{ /* Adar 2 */
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			12, 0, 12, 13, 14, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	/* sanity check */
	if h.hd_mon < 1 || h.hd_mon > 14 || h.hd_day < 1 || h.hd_day > 30 {
		return 0
	}

	holyday = holydays_table[h.hd_mon-1][h.hd_day-1]

	/* if tzom on sat delay one day */
	/* tzom gdalyaho on sat */
	if (holyday == 3) && (h.hd_dw == 7 || (h.hd_day == 4 && h.hd_dw != 1)) {
		holyday = 0
	}
	/* 17 of Tamuz on sat */
	if (holyday == 21) && ((h.hd_dw == 7) || (h.hd_day == 18 && h.hd_dw != 1)) {
		holyday = 0
	}
	/* 9 of Av on sat */
	if (holyday == 22) && ((h.hd_dw == 7) || (h.hd_day == 10 && h.hd_dw != 1)) {
		holyday = 0
	}

	/* Hanukah in a long year */
	if (holyday == 9) && (h.hd_size_of_year%10 != 3) && (h.hd_day == 3) {
		holyday = 0
	}

	/* if tanit ester on sat mov to Thu */
	if (holyday == 12) && ((h.hd_dw == 7) || (h.hd_day == 11 && h.hd_dw != 5)) {
		holyday = 0
	}

	/* yom yerushalym after 68 */
	if holyday == 26 {
		if h.gd_year < 1968 {
			holyday = 0
		}
	}

	/* yom ha azmaot and yom ha zicaron */
	if holyday == 17 {
		if h.gd_year < 1948 {
			holyday = 0
		} else if h.gd_year < 2004 {
			if (h.hd_day == 3) && (h.hd_dw == 5) {
				holyday = 17
			} else if (h.hd_day == 4) && (h.hd_dw == 5) {
				holyday = 17
			} else if (h.hd_day == 5) && (h.hd_dw != 6 && h.hd_dw != 7) {
				holyday = 17
			} else if (h.hd_day == 2) && (h.hd_dw == 4) {
				holyday = 25
			} else if (h.hd_day == 3) && (h.hd_dw == 4) {
				holyday = 25
			} else if (h.hd_day == 4) && (h.hd_dw != 5 && h.hd_dw != 6) {
				holyday = 25
			} else {
				holyday = 0
			}
		} else {
			if (h.hd_day == 3) && (h.hd_dw == 5) {
				holyday = 17
			} else if (h.hd_day == 4) && (h.hd_dw == 5) {
				holyday = 17
			} else if (h.hd_day == 6) && (h.hd_dw == 3) {
				holyday = 17
			} else if (h.hd_day == 5) && (h.hd_dw != 6 && h.hd_dw != 7 && h.hd_dw != 2) {
				holyday = 17
			} else if (h.hd_day == 2) && (h.hd_dw == 4) {
				holyday = 25
			} else if (h.hd_day == 3) && (h.hd_dw == 4) {
				holyday = 25
			} else if (h.hd_day == 5) && (h.hd_dw == 2) {
				holyday = 25
			} else if (h.hd_day == 4) && (h.hd_dw != 5 && h.hd_dw != 6 && h.hd_dw != 1) {
				holyday = 25
			} else {
				holyday = 0
			}
		}
	}

	/* yom ha shoaa, on years after 1958 */
	if holyday == 24 {
		if h.gd_year < 1958 {
			holyday = 0
		} else {
			if (h.hd_day == 26) && (h.hd_dw != 5) {
				holyday = 0
			}
			if (h.hd_day == 28) && (h.hd_dw != 2) {
				holyday = 0
			}
			if (h.hd_day == 27) && (h.hd_dw == 6 || h.hd_dw == 1) {
				holyday = 0
			}
		}
	}

	/* Rabin day, on years after 1997 */
	if holyday == 35 {
		if h.gd_year < 1997 {
			holyday = 0
		} else {
			if (h.hd_day == 10 || h.hd_day == 11) && (h.hd_dw != 5) {
				holyday = 0
			}
			if (h.hd_day == 12) && (h.hd_dw == 6 || h.hd_dw == 7) {
				holyday = 0
			}
		}
	}

	/* Zhabotinsky day, on years after 2005 */
	if holyday == 36 {
		if h.gd_year < 2005 {
			holyday = 0
		} else {
			if (h.hd_day == 30) && (h.hd_dw != 1) {
				holyday = 0
			}
			if (h.hd_day == 29) && (h.hd_dw == 7) {
				holyday = 0
			}
		}
	}

	/* diaspora holidays */

	/* simchat tora only in diaspora in israel just one day shmini+simchat tora */
	if holyday == 8 && !h.Diaspora {
		holyday = 0
	}
	/* sukkot II holiday only in diaspora */
	if holyday == 31 && !h.Diaspora {
		holyday = 6
	}

	/* pesach II holiday only in diaspora */
	if holyday == 32 && !h.Diaspora {
		holyday = 16
	}

	/* shavot II holiday only in diaspora */
	if holyday == 30 && !h.Diaspora {
		holyday = 0
	}

	/* pesach VIII holiday only in diaspora */
	if holyday == 29 && !h.Diaspora {
		holyday = 0
	}

	return holyday
}

/**
@brief Return the day in the omer of the given date

@param h The HebDate of the date to use.
@return The day in the omer, starting from 1 (or 0 if not in sfirat ha omer)
*/
func (h *HebDate) GetOmerDay() int {
	var omer_day int
	sixteen_nissan := &HebDate{}

	sixteen_nissan.SetHdate(16, 7, h.hd_year)
	omer_day = h.hd_jd - sixteen_nissan.hd_jd + 1

	if (omer_day > 49) || (omer_day < 0) {
		omer_day = 0
	}

	return omer_day
}

/**
@brief Return number of hebrew holyday type.

 Holiday types:
   0 - Regular day
   1 - Yom tov (plus yom kippor)
   2 - Erev yom kippur
   3 - Hol hamoed
   4 - Hanuka and purim
   5 - Tzomot
   6 - Independance day and Yom yerushalaim
   7 - Lag baomer ,Tu beav, Tu beshvat
   8 - Tzahal and Holocaust memorial days
   9 - National days

@param holyday the holyday number
@return the number of holyday type.
*/

func getHolydayType(holyday int) int {
	var holyday_type int

	switch holyday {
	case 0: /* regular day */
		holyday_type = 0

	case 1, 2, 4, 5, 8, 15, 20, 27, 28, 29, 30, 31, 32: /* Yom tov, To find erev yom tov, check if tomorrow returns 1 */
		holyday_type = 1

	case 37: /* Erev yom kippur */
		holyday_type = 2

	case 6, 7, 16: /* Hol hamoed */
		holyday_type = 3

	case 9, 13, 14: /* Hanuka and purim */
		holyday_type = 4

	case 3, 10, 12, 21, 22: /* tzom */
		holyday_type = 5

	case 17, 26: /* Independance day and Yom yerushalaim */
		holyday_type = 6

	case 18, 23, 11: /* Lag baomer ,Tu beav, Tu beshvat */
		holyday_type = 7

	case 24, 25: /* Tzahal and Holocaust memorial days */
		holyday_type = 8

	default: /* National days */
		holyday_type = 9

	}

	return holyday_type
}
