package typemapper

import (
	"fmt"
	"reflect"
)

func GetGenericTypeByT[T interface{}]() reflect.Type {
	res := reflect.TypeOf((*T)(nil)).Elem()
	return res
}

func getInstanceFromType(typ reflect.Type) interface{} {
	if typ.Kind() == reflect.Ptr {
		res := reflect.New(typ.Elem()).Interface()
		return res
	}

	return reflect.Zero(typ).Interface()
}

func GenericInstanceByT[T any]() T {
	typ := GetGenericTypeByT[T]()
	return getInstanceFromType(typ).(T)
}

func GetGenericTypeNameByT[T any]() string {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if t.Kind() != reflect.Ptr {
		return t.Name()
	}

	return fmt.Sprintf("*%s", t.Elem().Name())
}
