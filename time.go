package convert

import (
	"log"
	"time"
)

type Time struct {
	time.Time
}

func Timestamp(timestamp int64) Time {
	t := time.Unix(timestamp, 0)
	return Time{t}
}

func Date(date string) Time {
	t, err := time.ParseInLocation("2006-01-02", date, time.Local)
	if err != nil {
		log.Println("err date: ", err.Error())
	}

	return Time{t}
}

func (t *Time) ToString() string {
	return t.Format("2006-01-02 15:04:05")
}

func (t *Time) ToDate() string {
	return t.Format("2006-01-02")
}

func (t *Time) ToTimestamp() int64 {
	return t.UnixNano()
}