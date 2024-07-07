package util

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
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

func IsEmptyArrayOrSlice(i interface{}) bool {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		// 如果不是切片或数组类型，直接返回false，因为不可能是空数组或切片
		return false
	}

	// 使用Len方法获取切片或数组的长度
	return v.Len() == 0
}
