package utils

import "reflect"

// StructCopy 拷贝结构体,不指定具体类型，解耦
// from 和 to 必须为指针类型
// TODO: BUG 对于嵌套结构体，无法拷贝
func StructCopy(from, to any) {
	fValue := reflect.ValueOf(from)
	tValue := reflect.ValueOf(to)

	// 判断是否为指针类型
	if fValue.Kind() != reflect.Ptr || tValue.Kind() != reflect.Ptr {
		return
	}
	// 如果为空
	if fValue.IsNil() || tValue.IsNil() {
		return
	}

	fElem := fValue.Elem()
	tElem := tValue.Elem()

	for i := 0; i < tElem.NumField(); i++ {
		tField := tElem.Type().Field(i)
		fField, ok := fElem.Type().FieldByName(tField.Name)

		// 字段名称相同，类型相同
		if ok && fField.Type == tField.Type {
			tElem.Field(i).Set(fElem.FieldByName(tField.Name))
		}
	}
}
