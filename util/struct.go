package util

import (
	"reflect"
)

// CopyStruct ,浅复制！
// dst 目标结构体，src 源结构体
// 必须传入指针，且不能为nil
// 它会把src与dst的相同字段名的值，复制到dst中
func CopyStruct(dst, src interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		srcName := srcValue.Type().Field(i).Name
		dstFieldByName := dstValue.FieldByName(srcName)
		if dstFieldByName.IsValid() {
			if srcField.Type() != dstFieldByName.Type() {
				continue
			}
			switch dstFieldByName.Kind() {
			case reflect.Ptr:
				switch srcField.Kind() {
				case reflect.Ptr:
					if srcField.IsNil() {
						//dstFieldByName.Set(reflect.New(dstFieldByName.Type().Elem()))
					} else {
						dstFieldByName.Set(srcField)
					}
				default:
					dstFieldByName.Set(srcField.Addr())
				}
			default:
				switch srcField.Kind() {
				case reflect.Ptr:
					if srcField.IsNil() {
						dstFieldByName.Set(reflect.Zero(dstFieldByName.Type()))
					} else {
						dstFieldByName.Set(srcField.Elem())
					}
				default:
					dstFieldByName.Set(srcField)
				}
			}
		}
	}
}

func ToAnyArray[T any](list []T) []any {
	items := make([]any, len(list))
	for i, a := range list {
		items[i] = a
	}
	return items
}
