package util

func BoolPointer(val bool) *bool {
	return &val
}

func BoolValue(val *bool) bool {
	if val != nil {
		return *val
	}
	return false
}

func Int32Pointer(val int32) *int32 {
	return &val
}

func Int32Value(val *int32) int32 {
	if val != nil {
		return *val
	}
	return 0
}

func Int64Pointer(val int64) *int64 {
	return &val
}

func Int64Value(val *int64) int64 {
	if val != nil {
		return *val
	}
	return 0
}

func Float64Pointer(val float64) *float64 {
	return &val
}

func Float64Value(val *float64) float64 {
	if val != nil {
		return *val
	}
	return 0
}
