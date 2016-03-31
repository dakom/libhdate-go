package libhdate

import (
	"fmt"
)

/**
@brief Return a string, with the hebrew date.

@return NULL pointer upon failure or, upon success, a pointer to a
string containing the short ( e.g. "1 Tishrey" ) or long (e.g. "Tuesday
18 Tishrey 5763 Hol hamoed Sukot" ) formated date. You must free() the
pointer after use.

@param short_format A short flag (true - returns a short string, false returns a long string).

*/

func (h *HebDate) GetFormatDate(short_format bool) string {
	var hebrew_format bool = HDATE_STRING_LOCAL
	var omer_day, holiday int
	var bet_h, hebrew_buffer1, hebrew_buffer2, hday_int_str, hyear_int_str, omer_str string
	var hebrew_buffer1_len, hebrew_buffer2_len int = -1, -1

	if h.IsHebrewLocale {
		bet_h = "ב"
		hebrew_format = HDATE_STRING_HEBREW
	}

	hday_int_str = GetString(HDATE_STRING_INT, h.hd_day, HDATE_STRING_LONG, hebrew_format)
	if hday_int_str == "" {
		return ""
	}

	hyear_int_str = GetString(HDATE_STRING_INT, h.hd_year, HDATE_STRING_LONG, hebrew_format)
	if hyear_int_str == "" {
		return ""
	}

	/************************************************************
	* short format
	************************************************************/
	if short_format {
		hebrew_buffer1 = fmt.Sprintf("%s %s %s\n", hday_int_str, GetString(HDATE_STRING_HMONTH, h.hd_mon, HDATE_STRING_LONG, hebrew_format), hyear_int_str)
		hebrew_buffer1_len = len(hebrew_buffer1)
	} else {
		hebrew_buffer1 = fmt.Sprintf("%s %s%s %s", hday_int_str, bet_h, GetString(HDATE_STRING_HMONTH, h.hd_mon, HDATE_STRING_LONG, hebrew_format), hyear_int_str)
		hebrew_buffer1_len = len(hebrew_buffer1)

		/* if a day in the omer print it */
		if hebrew_buffer1_len != -1 {
			omer_day = h.GetOmerDay()
		}
		if omer_day != 0 {
			omer_str = GetString(HDATE_STRING_OMER, omer_day, HDATE_STRING_LONG, hebrew_format)

			hebrew_buffer2 = fmt.Sprintf("%s, %s", hebrew_buffer1, omer_str)
			hebrew_buffer2_len = len(hebrew_buffer2)

			if hebrew_buffer2_len != -1 {
				hebrew_buffer1 = hebrew_buffer2
			}
			hebrew_buffer1_len = hebrew_buffer2_len
		}

		/* if holiday print it */
		if hebrew_buffer1_len != -1 {
			holiday = h.GetHolyday()
		}
		if holiday != 0 {
			hebrew_buffer2 = fmt.Sprintf("%s, %s", hebrew_buffer1, GetString(HDATE_STRING_HOLIDAY, holiday, HDATE_STRING_LONG, hebrew_format))
			hebrew_buffer2_len = len(hebrew_buffer2)

			if hebrew_buffer2_len != -1 {
				hebrew_buffer1 = hebrew_buffer2
			}
			hebrew_buffer1_len = hebrew_buffer2_len
		}

	}

	if hebrew_buffer1_len != -1 {
		return hebrew_buffer1
	}
	return ""
}

/**
@brief Return a static string, with the package name and version

@return a a static string, with the package name and version
*/
func GetVersionString() string {
	return ("")
}

/**
@brief Return a static string, with the name of translator

@return a a static string, with the name of translator
*/
func GetTranslatorString() string {
	return ""
}

/**
 @brief   Return string values for hdate information
 @return  a pointer to a string containing the information. In the cases
          integers and omer, the strings will NOT be static, and the
          caller must free() them after use.
 @param type_of_string 	0 = integer, 1 = day of week, 2 = parshaot,
						3 = hmonth, 4 = gmonth, 5 = holiday, 6 = omer
 @param index			integer		( 0 < n < 11000)
						day of week ( 0 < n <  8 )
						parshaot	( 0 , n < 62 )
						hmonth		( 0 < n < 15 )
						gmonth		( 0 < n < 13 )
						holiday		( 0 < n < 37 )
						omer		( 0 < n < 50 )
 @param short_form   0 = short format
 @param hebrew_form  0 = not hebrew (native/embedded)
*/

// TODO - Number days of chol hamoed, and maybe have an entry for shabbat chol hamoed
// DONE - (I hope) change short to be = 1 long = 0, and switch order of data structures
//        this way user app opt.short = 0/FALSE will work as a parameter to pass here

// These definitions are in hdate.h
//
// HDATE_STRING_INT     0
// HDATE_STRING_DOW     1
// HDATE_STRING_PARASHA 2
// HDATE_STRING_HMONTH  3
// HDATE_STRING_GMONTH  4
// HDATE_STRING_HOLIDAY 5
// HDATE_STRING_OMER    6
// HDATE STRING_SHORT   1
// HDATE_STRING_LONG    0
// HDATE_STRING_HEBREW  1
// HDATE_STRING_LOCAL   0
func GetString(type_of_string int, index int, input_short_form bool, input_hebrew_form bool) string {
	var short_form, hebrew_form int
	var return_string_len int = -1
	var return_string, h_int_string string

	if input_short_form {
		short_form = 1
	}
	if input_hebrew_form {
		hebrew_form = 1
	}

	digits := [3][10]string{
		{" ", "א", "ב", "ג", "ד", "ה", "ו", "ז", "ח", "ט"},
		{"ט", "י", "כ", "ל", "מ", "נ", "ס", "ע", "פ", "צ"},
		{" ", "ק", "ר", "ש", "ת"},
	}

	days := [2][2][7]string{
		{
			{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		},
		{
			{"ראשון", "שני", "שלישי", "רביעי", "חמישי", "שישי", "שבת"},
			{"א", "ב", "ג", "ד", "ה", "ו", "ש"},
		},
	}

	parashaot := [2][2][62]string{
		{ // begin english
			{ // begin english long
				"none", "Bereshit", "Noach",
				"Lech-Lecha", "Vayera", "Chayei Sara",
				"Toldot", "Vayetzei", "Vayishlach",
				"Vayeshev", "Miketz", "Vayigash", /* 11 */
				"Vayechi", "Shemot", "Vaera",
				"Bo", "Beshalach", "Yitro",
				"Mishpatim", "Terumah", "Tetzaveh", /* 20 */
				"Ki Tisa", "Vayakhel", "Pekudei",
				"Vayikra", "Tzav", "Shmini",
				"Tazria", "Metzora", "Achrei Mot",
				"Kedoshim", "Emor", "Behar", /* 32 */
				"Bechukotai", "Bamidbar", "Nasso",
				"Beha'alotcha", "Sh'lach", "Korach",
				"Chukat", "Balak", "Pinchas", /* 41 */
				"Matot", "Masei", "Devarim",
				"Vaetchanan", "Eikev", "Re'eh",
				"Shoftim", "Ki Teitzei", "Ki Tavo", /* 50 */
				"Nitzavim", "Vayeilech", "Ha'Azinu",
				"Vezot Habracha", /* 54 */
				"Vayakhel-Pekudei", "Tazria-Metzora", "Achrei Mot-Kedoshim",
				"Behar-Bechukotai", "Chukat-Balak", "Matot-Masei",
				"Nitzavim-Vayeilech"},
			{ // begin english short
				"none", "Bereshit", "Noach",
				"Lech-Lecha", "Vayera", "Chayei Sara",
				"Toldot", "Vayetzei", "Vayishlach",
				"Vayeshev", "Miketz", "Vayigash", /* 11 */
				"Vayechi", "Shemot", "Vaera",
				"Bo", "Beshalach", "Yitro",
				"Mishpatim", "Terumah", "Tetzaveh", /* 20 */
				"Ki Tisa", "Vayakhel", "Pekudei",
				"Vayikra", "Tzav", "Shmini",
				"Tazria", "Metzora", "Achrei Mot",
				"Kedoshim", "Emor", "Behar", /* 32 */
				"Bechukotai", "Bamidbar", "Nasso",
				"Beha'alotcha", "Sh'lach", "Korach",
				"Chukat", "Balak", "Pinchas", /* 41 */
				"Matot", "Masei", "Devarim",
				"Vaetchanan", "Eikev", "Re'eh",
				"Shoftim", "Ki Teitzei", "Ki Tavo", /* 50 */
				"Nitzavim", "Vayeilech", "Ha'Azinu",
				"Vezot Habracha", /* 54 */
				"Vayakhel-Pekudei", "Tazria-Metzora", "Achrei Mot-Kedoshim",
				"Behar-Bechukotai", "Chukat-Balak", "Matot-Masei",
				"Nitzavim-Vayeilech"},
		},
		{ // begin hebrew
			{ // begin hebrew long
				"none", "בראשית", "נח",
				"לך לך", "וירא", "חיי שרה",
				"תולדות", "ויצא", "וישלח",
				"וישב", "מקץ", "ויגש", /* 11 */
				"ויחי", "שמות", "וארא",
				"בא", "בשלח", "יתרו",
				"משפטים", "תרומה", "תצוה", /* 20 */
				"כי תשא", "ויקהל", "פקודי",
				"ויקרא", "צו", "שמיני",
				"תזריע", "מצורע", "אחרי מות",
				"קדושים", "אמור", "בהר", /* 32 */
				"בחוקתי", "במדבר", "נשא",
				"בהעלתך", "שלח", "קרח",
				"חקת", "בלק", "פנחס", /* 41 */
				"מטות", "מסעי", "דברים",
				"ואתחנן", "עקב", "ראה",
				"שופטים", "כי תצא", "כי תבוא", /* 50 */
				"נצבים", "וילך", "האזינו",
				"וזאת הברכה", /* 54 */
				"ויקהל-פקודי", "תזריע-מצורע", "אחרי מות-קדושים",
				"בהר-בחוקתי", "חוקת-בלק", "מטות מסעי",
				"נצבים-וילך"},
			{ // begin hebrew short
				"none", "בראשית", "נח",
				"לך לך", "וירא", "חיי שרה",
				"תולדות", "ויצא", "וישלח",
				"וישב", "מקץ", "ויגש", /* 11 */
				"ויחי", "שמות", "וארא",
				"בא", "בשלח", "יתרו",
				"משפטים", "תרומה", "תצוה", /* 20 */
				"כי תשא", "ויקהל", "פקודי",
				"ויקרא", "צו", "שמיני",
				"תזריע", "מצורע", "אחרי מות",
				"קדושים", "אמור", "בהר", /* 32 */
				"בחוקתי", "במדבר", "נשא",
				"בהעלתך", "שלח", "קרח",
				"חקת", "בלק", "פנחס", /* 41 */
				"מטות", "מסעי", "דברים",
				"ואתחנן", "עקב", "ראה",
				"שופטים", "כי תצא", "כי תבוא", /* 50 */
				"נצבים", "וילך", "האזינו",
				"וזאת הברכה", /* 54 */
				"ויקהל-פקודי", "תזריע-מצורע", "אחרי מות-קדושים",
				"בהר-בחוקתי", "חוקת-בלק", "מטות מסעי",
				"נצבים-וילך"},
		},
	}

	hebrew_months := [2][2][14]string{
		{ // begin english
			{ // begin english long
				"Tishrei", "Cheshvan", "Kislev", "Tevet",
				"Sh'vat", "Adar", "Nisan", "Iyyar",
				"Sivan", "Tammuz", "Av", "Elul", "Adar I",
				"Adar II"},
			{ // begin english short
				"Tishrei", "Cheshvan", "Kislev", "Tevet",
				"Sh'vat", "Adar", "Nisan", "Iyyar",
				"Sivan", "Tammuz", "Av", "Elul", "Adar I",
				"Adar II"},
		},
		{ // begin hebrew
			{ // begin hebrew long
				"תשרי", "חשון", "כסלו", "טבת", "שבט", "אדר", "ניסן", "אייר",
				"סיון", "תמוז", "אב", "אלול", "אדר א", "אדר ב"},
			{ // begin hebrew short
				"תשרי", "חשון", "כסלו", "טבת", "שבט", "אדר", "ניסן", "אייר",
				"סיון", "תמוז", "אב", "אלול", "אדר א", "אדר ב"},
		},
	}

	gregorian_months := [2][12]string{
		{"January", "February", "March",
			"April", "May", "June",
			"July", "August", "September",
			"October", "November", "December"},
		{"Jan", "Feb", "Mar", "Apr", "May",
			"Jun", "Jul", "Aug", "Sep", "Oct",
			"Nov", "Dec"},
	}

	holidays := [2][2][37]string{
		{ // begin english
			{ // begin english long
				"Rosh Hashana I", "Rosh Hashana II",
				"Tzom Gedaliah", "Yom Kippur",
				"Sukkot", "Hol hamoed Sukkot",
				"Hoshana raba", "Simchat Torah",
				"Chanukah", "Asara B'Tevet",
				"Tu B'Shvat", "Ta'anit Esther",
				"Purim", "Shushan Purim",
				"Pesach", "Hol hamoed Pesach",
				"Yom HaAtzma'ut", "Lag B'Omer",
				"Erev Shavuot", "Shavuot",
				"Tzom Tammuz", "Tish'a B'Av",
				"Tu B'Av", "Yom HaShoah",
				"Yom HaZikaron", "Yom Yerushalayim",
				"Shmini Atzeret", "Pesach VII",
				"Pesach VIII", "Shavuot II",
				"Sukkot II", "Pesach II",
				"Family Day", "Memorial day for fallen whose place of burial is unknown",
				"Yitzhak Rabin memorial day", "Zeev Zhabotinsky day",
				"Erev Yom Kippur"},
			{ // begin english short
				"Rosh Hashana I", "Rosh Hashana II",
				"Tzom Gedaliah", "Yom Kippur",
				"Sukkot", "Hol hamoed Sukkot",
				"Hoshana raba", "Simchat Torah",
				"Chanukah", "Asara B'Tevet", /* 10 */
				"Tu B'Shvat", "Ta'anit Esther",
				"Purim", "Shushan Purim",
				"Pesach", "Hol hamoed Pesach",
				"Yom HaAtzma'ut", "Lag B'Omer",
				"Erev Shavuot", "Shavuot", /* 20 */
				"Tzom Tammuz", "Tish'a B'Av",
				"Tu B'Av", "Yom HaShoah",
				"Yom HaZikaron", "Yom Yerushalayim",
				"Shmini Atzeret", "Pesach VII",
				"Pesach VIII", "Shavuot II", /* 30 */
				"Sukkot II", "Pesach II",
				"Family Day", "Memorial day for fallen whose place of burial is unknown",
				"Rabin memorial day", "Zhabotinsky day",
				"Erev Yom Kippur"},
		},
		{ // begin hebrew
			{ // begin hebrew long
				"א ר\"ה", "ב' ר\"ה",
				"צום גדליה", "יוה\"כ",
				"סוכות", "חוה\"מ סוכות",
				"הוש\"ר", "שמח\"ת",
				"חנוכה", "י' בטבת", /* 10 */
				"ט\"ו בשבט", "תענית אסתר",
				"פורים", "שושן פורים",
				"פסח", "חוה\"מ פסח",
				"יום העצמאות", "ל\"ג בעומר",
				"ערב שבועות", "שבועות", /* 20 */
				"צום תמוז", "ט' באב",
				"ט\"ו באב", "יום השואה",
				"יום הזכרון", "יום י-ם",
				"שמיני עצרת", "ז' פסח",
				"אחרון של פסח", "ב' שבועות", /* 30 */
				"ב' סוכות", "ב' פסח",
				"יום המשפחה", "יום זכרון...",
				"יום הזכרון ליצחק רבין", "יום ז'בוטינסקי",
				"עיוה\"כ"},
			{ // begin hebrew short
				"א' ראש השנה", "ב' ראש השנה",
				"צום גדליה", "יום הכפורים",
				"סוכות", "חול המועד סוכות",
				"הושענא רבה", "שמחת תורה",
				"חנוכה", "צום עשרה בטבת", /* 10 */
				"ט\"ו בשבט", "תענית אסתר",
				"פורים", "שושן פורים",
				"פסח", "חול המועד פסח",
				"יום העצמאות", "ל\"ג בעומר",
				"ערב שבועות", "שבועות", /* 20 */
				"צום שבעה עשר בתמוז", "תשעה באב",
				"ט\"ו באב", "יום השואה",
				"יום הזכרון", "יום ירושלים",
				"שמיני עצרת", "שביעי פסח",
				"אחרון של פסח", "שני של שבועות", /* 30 */
				"שני של סוכות", "שני של פסח",
				"יום המשפחה", "יום זכרון...",
				"יום הזכרון ליצחק רבין", "יום ז'בוטינסקי",
				"עיוה\"כ"},
		},
	}

	switch type_of_string {
	case HDATE_STRING_DOW:
		if index >= 1 && index <= 7 {
			return (days[hebrew_form][short_form][index-1])
		}
	case HDATE_STRING_PARASHA:
		if index >= 1 && index <= 61 {
			return (parashaot[hebrew_form][short_form][index])
		}
	case HDATE_STRING_HMONTH:
		if index >= 1 && index <= 14 {
			return (hebrew_months[hebrew_form][short_form][index-1])
		}
	case HDATE_STRING_GMONTH:
		if index >= 1 && index <= 12 {
			return (gregorian_months[short_form][index-1])
		}
	case HDATE_STRING_HOLIDAY:
		if index >= 1 && index <= 37 {
			return (holidays[hebrew_form][short_form][index-1])
		}
	case HDATE_STRING_OMER:
		if index > 0 && index < 50 {
			h_int_string = GetString(HDATE_STRING_INT, index, HDATE_STRING_LONG, input_hebrew_form)
			if h_int_string == "" {
				return ""
			}

			return_string = fmt.Sprintf("%s %s", h_int_string, "in the Omer")
			return_string_len = len(return_string)

			if return_string_len != -1 {
				return return_string
			}
		}
		return ""

	case HDATE_STRING_INT:
		if (index > 0) && (index < 11000) {
			// not hebrew form - return the number in decimal form
			if !input_hebrew_form {
				return_string = fmt.Sprintf("%d", index)
				return_string_len = len(return_string)

				if return_string_len == -1 {
					return ""
				}
				return return_string
			}

			var n int = index

			if n >= 1000 {
				return_string += digits[0][n/1000]
				n %= 1000
			}
			for n >= 400 {
				return_string += digits[2][4]
				n -= 400
			}
			if n >= 100 {
				return_string += digits[2][n/100]
				n %= 100
			}
			if n >= 10 {
				if n == 15 || n == 16 {
					n -= 9
				}
				return_string += digits[1][n/10]
				n %= 10
			}
			if n > 0 {
				return_string += digits[0][n]
			}
			// possibly add the ' and " to hebrew numbers
			if !input_short_form {
				return_string_len = len(return_string)
				if return_string_len <= 1 {
					return_string += "'"
				} else {
					return_string = replaceAtIndex(return_string, return_string[return_string_len], return_string_len+1)

					return_string = replaceAtIndex(return_string, return_string[return_string_len-1], return_string_len)

					return_string = replaceAtIndex(return_string, return_string[return_string_len-2], return_string_len-1)

					return_string = replaceAtIndex(return_string, '"', return_string_len-2)

				}
			}
			return return_string
		}
		return ""

	} // end of switch(type_of_string)

	return ""
}

func replaceAtIndex(str string, replacement byte, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}
