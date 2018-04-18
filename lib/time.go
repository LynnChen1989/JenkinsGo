package lib

import "time"

func Now() (now string) {
	now = time.Now().Format("2006-01-02 15:04:05")
	return
}
