package utils

import "log"

func IsError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
