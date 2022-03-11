package helpers

import (
	"log"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ChangeFormatDate(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
