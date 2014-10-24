package public

import (
	"fmt"
	"reflect"
	"strconv"
	//"strings"
)

func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return reflect.ValueOf(value).String()
	case []byte:
		return string(reflect.ValueOf(value).Bytes())
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%ud", reflect.ValueOf(value).Uint())
	case float32:
		return strconv.FormatFloat(reflect.ValueOf(value).Float(), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(reflect.ValueOf(value).Float(), 'f', -1, 64)
	default:
		return ``
	}
}

func ToFloat32(value interface{}) float32 {
	switch value.(type) {
	case string:
		var f float64
		f = 0
		s := reflect.ValueOf(value).String()
		f, _ = strconv.ParseFloat(s, 32)
		return float32(f)
	case float32, float64:
		return float32(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return float32(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return float32(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToFloat64(value interface{}) float64 {
	switch value.(type) {
	case string:
		var f float64
		f = 0
		s := reflect.ValueOf(value).String()
		f, _ = strconv.ParseFloat(s, 64)
		return f
	case float32, float64:
		return float64(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToInt(value interface{}) int {
	switch value.(type) {
	case string:
		var n int
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.Atoi(s); err != nil {
			return 0
		}
		return n
	case float32, float64:
		return int(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToInt8(value interface{}) int8 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 8); err != nil {
			return 0
		}
		return int8(n)
	case float32, float64:
		return int8(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int8(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int8(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToInt16(value interface{}) int16 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 16); err != nil {
			return 0
		}
		return int16(n)
	case float32, float64:
		return int16(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int16(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int16(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToInt32(value interface{}) int32 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			return 0
		}
		return int32(n)
	case float32, float64:
		return int32(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int32(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int32(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToInt64(value interface{}) int64 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 64); err != nil {
			return 0
		}
		return n
	case float32, float64:
		return int64(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int64(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToUInt(value interface{}) uint {
	switch value.(type) {
	case string:
		var n int
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.Atoi(s); err != nil || n < 0 {
			return 0
		}
		return uint(n)
	case float32, float64:
		return uint(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToUInt8(value interface{}) uint8 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 8); err != nil || n < 0 {
			return 0
		}
		return uint8(n)
	case float32, float64:
		return uint8(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint8(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint8(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToUInt16(value interface{}) uint16 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 16); err != nil || n < 0 {
			return 0
		}
		return uint16(n)
	case float32, float64:
		return uint16(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint16(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint16(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToUInt32(value interface{}) uint32 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 32); err != nil || n < 0 {
			return 0
		}
		return uint32(n)
	case float32, float64:
		return uint32(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint32(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint32(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToUInt64(value interface{}) uint64 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 64); err != nil || n < 0 {
			return 0
		}
		return uint64(n)
	case float32, float64:
		return uint64(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint64(reflect.ValueOf(value).Uint())
	default:
		return 0
	}
}

func ToIntDef(value interface{}, r int) int {
	switch value.(type) {
	case string:
		var n int
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.Atoi(s); err != nil {
			return r
		}
		return n
	case float32, float64:
		return int(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToInt8Def(value interface{}, r int8) int8 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 8); err != nil {
			return r
		}
		return int8(n)
	case float32, float64:
		return int8(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int8(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int8(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToInt16Def(value interface{}, r int16) int16 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 16); err != nil {
			return r
		}
		return int16(n)
	case float32, float64:
		return int16(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int16(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int16(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToInt32Def(value interface{}, r int32) int32 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 32); err != nil {
			return r
		}
		return int32(n)
	case float32, float64:
		return int32(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int32(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int32(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToInt64Def(value interface{}, r int64) int64 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 64); err != nil {
			return r
		}
		return n
	case float32, float64:
		return int64(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return int64(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToUIntDef(value interface{}, r uint) uint {
	switch value.(type) {
	case string:
		var n int
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.Atoi(s); err != nil || n < 0 {
			return r
		}
		return uint(n)
	case float32, float64:
		return uint(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToUInt8Def(value interface{}, r uint8) uint8 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 8); err != nil || n < 0 {
			return r
		}
		return uint8(n)
	case float32, float64:
		return uint8(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint8(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint8(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToUInt16Def(value interface{}, r uint16) uint16 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 16); err != nil || n < 0 {
			return r
		}
		return uint16(n)
	case float32, float64:
		return uint16(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint16(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint16(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToUInt32Def(value interface{}, r uint32) uint32 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 32); err != nil || n < 0 {
			return r
		}
		return uint32(n)
	case float32, float64:
		return uint32(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint32(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint32(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToUInt64Def(value interface{}, r uint64) uint64 {
	switch value.(type) {
	case string:
		var n int64
		var err error
		s := reflect.ValueOf(value).String()
		if n, err = strconv.ParseInt(s, 10, 64); err != nil || n < 0 {
			return r
		}
		return uint64(n)
	case float32, float64:
		return uint64(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return uint64(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToFloat32Def(value interface{}, r float32) float32 {
	switch value.(type) {
	case string:
		var f float64
		var err error
		s := reflect.ValueOf(value).String()
		if f, err = strconv.ParseFloat(s, 32); err != nil {
			return r
		}
		return float32(f)
	case float32, float64:
		return float32(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return float32(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return float32(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func ToFloat64Def(value interface{}, r float64) float64 {
	switch value.(type) {
	case string:
		var f float64
		var err error
		s := reflect.ValueOf(value).String()
		if f, err = strconv.ParseFloat(s, 64); err != nil {
			return r
		}
		return f
	case float32, float64:
		return float64(reflect.ValueOf(value).Float())
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(value).Int())
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(value).Uint())
	default:
		return r
	}
}

func IntToStr(value int64) string {
	//return fmt.Sprintf("%d", value)
	return strconv.FormatInt(value, 10)
}

func UIntToStr(value uint64) string {
	return fmt.Sprintf("%ud", value)
}

func FloatToStr(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func StrToInt64(value string) int64 {
	var n int64
	var err error
	if n, err = strconv.ParseInt(value, 10, 64); err != nil {
		panic(err)
	}
	return n
}

func StrToUInt64(value string) uint64 {
	var n int64
	var err error
	if n, err = strconv.ParseInt(value, 10, 64); err != nil || n < 0 {
		panic(`string to uint convert error`)
	}
	return uint64(n)
}
