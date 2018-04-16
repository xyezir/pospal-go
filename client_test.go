package pospal

import (
	"os"
	"errors"
)

func init() {
	var AppID = os.Getenv("PosPalAppID")
	var AppKey = os.Getenv("PosPalAppKey")

	if AppID == "" {
		panic(errors.New("环境变量中找不到AppID"))
	}

	if AppKey == "" {
		panic(errors.New("环境变量中找不到AppKey"))
	}
}
