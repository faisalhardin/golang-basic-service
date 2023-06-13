package repo

import (
	"fmt"
	"reflect"
	"strings"
)

func GetFieldNames(s interface{}) (fieldNames []string, err error) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		err = fmt.Errorf("error type not struct")
		return
	}

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get("json"), ",")[0]
		if v == "" {
			continue
		}
		fieldNames = append(fieldNames, v)
	}

	return
}

// func ConvertFromHGetAllToStruct ()
