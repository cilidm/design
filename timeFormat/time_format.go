package timeFormat

import (
	"fmt"
	"time"
)

func GetTimestampByShortString(str string, isEnd bool) int64 {
	timeTemplate1 := "2006-01-02"
	retTime, err := time.ParseInLocation(timeTemplate1, str, time.Local)
	if err != nil {
		return 0
	}
	if isEnd {
		fmt.Println( retTime.Add(time.Second * 86399).Unix())
		return retTime.Add(time.Second * 86399).Unix()
	}

	return retTime.Unix()
}
