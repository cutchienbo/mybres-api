package helper

import "time"

func GetCurrentTimeVN() time.Time {
	location, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	currentTime := time.Now().In(location)

	// currentDate := currentTime.Format("0000-00-00 00:00:00")

	return currentTime
}