package timeutils

import "time"

const LayoutISO = "2006-01-02"
const LayoutISOWithTime = "2006-01-02T15:04:05.000000Z"
const LayoutUserFriendly = "02-Jan-2006 15:04:05"

func Now() time.Time {
	return time.Now().UTC().Truncate(time.Microsecond)
}
