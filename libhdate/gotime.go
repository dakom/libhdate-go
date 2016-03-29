package libhdate

import (
	"time"
)

func (h *HebDate) SetTime(t time.Time) {
	year, month, day := t.Date()

	h.SetGdate(day, int(month), year)
}
