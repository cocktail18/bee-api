package util

import (
	"crypto/md5"
	"encoding/hex"
)

func ReGroup[T1 any, T2 comparable](list []T1, fn func(a T1) T2) map[T2][]T1 {
	ret := make(map[T2][]T1)
	for _, t1 := range list {
		key := fn(t1)
		if _, ok := ret[key]; !ok {
			ret[key] = make([]T1, 0, 2)
		}
		ret[key] = append(ret[key], t1)
	}
	return ret
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
