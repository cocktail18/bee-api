package util

import (
	"github.com/google/uuid"
	"strings"
)

func GetUUIDStr() string {
	item, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(item.String(), "-", "")
}
