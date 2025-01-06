package date

import (
	"testing"
	"time"
)

func TestMoments(t *testing.T) {
	now := time.Now()
	t.Logf("BeginningOfTheDay: %s", BeginningOfTheDay(now).Format(time.Stamp))
	t.Logf("EndOfTheDay: %s", EndOfTheDay(now).Format(time.Stamp))
	t.Logf("BeginningOfWeek: %s", BeginningOfWeek(now).Format(time.Stamp))
	t.Logf("EndOfWeek: %s", EndOfWeek(now).Format(time.Stamp))
	t.Logf("BeginningOfMonth: %s", BeginningOfMonth(now).Format(time.Stamp))
	t.Logf("EndOfMonth: %s", EndOfMonth(now).Format(time.Stamp))

}
