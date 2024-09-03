package util

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"strings"
	"time"
)

func GetUUIDStr() string {
	item, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(item.String(), "-", "")
}
func GetUUID() string {
	item, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return item.String()
}

func GetRandInt64() string {
	return time.Now().Format("20060102") + lo.RandomString(10, lo.NumbersCharset)
}
