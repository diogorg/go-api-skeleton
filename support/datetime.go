package support

import "time"

func CurrentDateTime() time.Time {
	return time.Now()
}

func ParseDbToBr(dt time.Time) string {
	return dt.Format("02/01/2006 15:04")
}
