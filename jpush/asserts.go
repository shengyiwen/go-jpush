package jpush

import (
	"reflect"
)

// 当输入值不为空的时候，进行函数调用
func AssertNotNull(val interface{}, thenFunc func()) {
	vi := reflect.ValueOf(val)

	// 对string类型处理
	if vi.Kind() == reflect.String {
		if val != "" {
			thenFunc()
		}
		return
	}

	// 对slice和array类型处理
	if vi.Kind() == reflect.Slice || vi.Kind() == reflect.Array {
		if vi.Len() > 0 {
			thenFunc()
		}
		return
	}

	// 对指针类型处理
	if vi.Kind() == reflect.Ptr {
		if !vi.IsNil() {
			thenFunc()
		}
		return
	}

	// 对其他类型的处理
	if val != nil {
		thenFunc()
	}
}
