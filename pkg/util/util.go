package util

import "log"

func CheckFatalError(err error) {
	if err != nil {
		log.Fatalf("alert: %v", err)
	}
}
