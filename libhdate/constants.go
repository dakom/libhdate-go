package libhdate

/** @def HDATE_DIASPORA_FLAG
  @brief use diaspora dates and holydays flag
*/
const HDATE_DIASPORA_FLAG int = -1

/** @def HDATE_ISRAEL_FLAG
  @brief use israel dates and holydays flag
*/
const HDATE_ISRAEL_FLAG int = 0

/** @def HDATE_SHORT_FLAG
  @brief use short strings flag
*/
const HDATE_SHORT_FLAG int = -1

/** @def HDATE_LONG_FLAG
  @brief use long strings flag
*/
const HDATE_LONG_FLAG int = 0

/** @def HEBREW_NUMBER_BUFFER_SIZE
  @brief for HdateGetint_string_ and HdateGetint_wstring
  @note
  How large should the buffer be? Hebrew year 10,999 would
  be י'תתקצ"ט, eight characters, each two bytes, plus an
  end-of-string delimiter, equals 17. This could effectively
  yield a range extending to Hebrew year 11,899, י"א תתצ"ט,
  due to the extra ק needed for the '900' century. However,
  for readability, I would want a an extra space at that
  point between the millenium and the century...
*/
const HEBREW_NUMBER_BUFFER_SIZE int = 17
const HEBREW_WNUMBER_BUFFER_SIZE int = 9

/** @def HDATE_STRING_INT
  @brief for function HdateString: identifies string type: integer
*/
const HDATE_STRING_INT int = 0

/** @def HDATE_STRING_DOW
  @brief for function HdateString: identifies string type: day of week
*/
const HDATE_STRING_DOW int = 1

/** @def HDATE_STRING_PARASHA
  @brief for function HdateString: identifies string type: parasha
*/
const HDATE_STRING_PARASHA int = 2

/** @def HDATE_STRING_HMONTH
  @brief for function HdateString: identifies string type: hebrew_month
*/
const HDATE_STRING_HMONTH int = 3

/** @def HDATE_STRING_GMONTH
  @brief for function HdateString: identifies string type: gregorian_month
*/
const HDATE_STRING_GMONTH int = 4

/** @def HDATE_STRING_HOLIDAY
  @brief for function HdateString: identifies string type: holiday
*/
const HDATE_STRING_HOLIDAY int = 5

/** @def HDATE_STRING_HOLIDAY
  @brief for function HdateString: identifies string type: holiday
*/
const HDATE_STRING_OMER int = 6

/** @def HDATE_STRING_SHORT
  @brief for function HdateString: use short form, if one exists
*/
const HDATE_STRING_SHORT int = 1

/** @def HDATE_STRING_LONG
  @brief for function HdateString: use long form
*/
const HDATE_STRING_LONG bool = false

/** @def HDATE_STRING_HEBREW
  @brief for function HdateString: use embedded hebrew string
*/
const HDATE_STRING_HEBREW bool = true

/** @def HDATE_STRING_LOCAL
  @brief for function HdateString: use local locale string
*/
const HDATE_STRING_LOCAL bool = false
