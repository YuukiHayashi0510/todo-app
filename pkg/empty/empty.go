package empty

import "reflect"

// Is checks if a value is empty. Returns true for:
// - Array/String: length 0
// - Bool: false
// - Numbers: 0
// - Interface/Pointer: nil
// - Map/Slice: nil or length 0
// - Other: nil
func Is(value interface{}) bool {
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Map, reflect.Slice:
		if v.IsNil() {
			return true
		}
		return v.Len() == 0
	default:
		return value == nil
	}
}

// Any returns true if any of the given values is empty.
// Empty values are:
// - zero values (0, "", false)
// - nil values (nil slice, nil map, nil pointer, nil interface)
// - empty containers (empty slice, empty map)
// If no values are provided, returns false.
func Any(values ...interface{}) bool {
	for _, v := range values {
		if Is(v) {
			return true
		}
	}

	return false
}
