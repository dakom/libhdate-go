package libhdate

import (
	"time"
)

func HdateSetGoTime(h *HebDate, t time.Time) *HebDate {
	year, month, day := t.Date()

	return HdateSetGdate(h, day, int(month), year)
}
