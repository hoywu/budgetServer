package utils

// import (
// 	"reflect"
// )

// func CopyStructField(from interface{}, to interface{}) {
// 	fromType := reflect.TypeOf(from)
// 	toType := reflect.TypeOf(to)

// 	if fromType.Kind() != reflect.Ptr || fromType.Elem().Kind() != reflect.Struct ||
// 		toType.Kind() != reflect.Ptr || toType.Elem().Kind() != reflect.Struct {
// 		panic("CopyStructField: must be a struct pointer")
// 	}

// 	fromMap := getFieldMap(from)
// 	toMap := getFieldMap(to)

// 	for k, v := range fromMap {
// 		if _, ok := toMap[k]; !ok {
// 			continue
// 		}
// 		if v.Kind() != toMap[k].Kind() {
// 			continue
// 		}
// 		if !toMap[k].CanSet() {
// 			continue
// 		}
// 		toMap[k].Set(v)
// 	}
// }

// func getFieldMap(o interface{}) map[string]reflect.Value {
// 	fieldMap := make(map[string]reflect.Value)
// 	getMap(reflect.ValueOf(o), &fieldMap)
// 	return fieldMap
// }

// func getMap(v reflect.Value, fieldMap *map[string]reflect.Value) {
// 	if v.Kind() == reflect.Ptr {
// 		v = v.Elem()
// 	}
// 	if v.Kind() != reflect.Struct {
// 		panic("getMap: must be a struct")
// 	}
// 	vType := v.Type()
// 	for i := 0; i < vType.NumField(); i++ {
// 		vField := v.Field(i)
// 		if vField.Kind() == reflect.Struct {
// 			getMap(vField, fieldMap)
// 		} else {
// 			(*fieldMap)[vType.Field(i).Name] = vField
// 		}
// 	}
// }
