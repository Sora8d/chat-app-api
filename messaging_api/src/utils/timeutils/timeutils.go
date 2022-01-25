package timeutils

import "time"

func GetNow() int64 {
	return time.Now().UTC().Unix()
}
