package todos

import "time"

func getTodaysTodoList() string {
	return time.Now().Local().Format("20060102")
}
