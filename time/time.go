package time

import (
	"strconv"
	"time"
)

func NowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func GetCurrentTimestampSec() int {
	ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	return ts
}
