package typemapper

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
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

// GetNonePointerTypeName returns the name of the type without its package name and its pointer.
func GetNonePointerTypeName(input interface{}) string {
	if input == nil {
		return ""
	}

	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Ptr {
		return t.Name()
	}

	return t.Elem().Name()
}

func GetSnakeTypeName(input interface{}) string {
	if input == nil {
		return ""
	}

	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Ptr {
		return t.Name()
	}

	return strcase.ToSnake(t.Elem().Name())
}

func GetPackageName(value interface{}) string {
	inputType := reflect.TypeOf(value)
	if inputType.Kind() == reflect.Ptr {
		inputType = inputType.Elem()
	}

	packagePath := inputType.PkgPath()

	parts := strings.Split(packagePath, "./")

	return parts[len(parts)-1]
}

// GetTypeName returns the name of type without its package name
func GetTypeName(input interface{}) string {
	if input == nil {
		return ""
	}

	t := reflect.TypeOf(input)
	if t.Kind() != reflect.Ptr {
		return t.Name()
	}

	return fmt.Sprintf("*%s", t.Elem().Name())
}

func GetGenericFullTypeNameByT[T any]() string {
	t := reflect.TypeOf((*T)(nil)).Elem()

	return t.String()
}

// GetFullTypeName returns the full name of the type by its package name
func GetFullTypeName(input interface{}) string {
	if input == nil {
		return ""
	}

	t := reflect.TypeOf(input)
	return t.String()
}
