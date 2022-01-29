package time

import "time"

func NowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}
