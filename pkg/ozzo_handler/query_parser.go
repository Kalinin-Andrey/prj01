package ozzo_handler

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

func ParseQueryParams(ctx *routing.Context, out interface{}) error {
	v := make(map[string]string)

	for key, vals := range ctx.Request.URL.Query() {
		if len(vals) > 0 {
			v[key] = vals[0]
		}
	}

	return strings2struct(v, out)
}

func ParseUintParam(ctx *routing.Context, paramName string) (uint, error) {
	str := ctx.Param(paramName)
	if str == "" {
		return 0, errors.New("empty")
	}

	paramVal, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(paramVal), nil
}

func ParseUintQueryParam(ctx *routing.Context, paramName string) (uint, error) {
	str := ctx.Query(paramName)
	if str == "" {
		return 0, errors.New("empty")
	}

	paramVal, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(paramVal), nil
}

func strings2struct(data interface{}, out interface{}) error {
	outType := reflect.TypeOf(out)

	if outType.Kind() != reflect.Ptr {
		return fmt.Errorf("Parameter out must be a Ptr")
	}

	outVal := reflect.ValueOf(out)
	outValElem := outVal.Elem()

	if !outValElem.CanSet() {
		return fmt.Errorf("!outValElem.CanSet()")
	}
	outPtrType := reflect.Indirect(outVal).Kind()
	dataVal := reflect.ValueOf(data)

	switch outPtrType {
	case reflect.Bool:
		str, ok := data.(string)
		if !ok {
			return fmt.Errorf("Data mast be a string")
		}
		paramVal, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}
		outValElem.SetBool(paramVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str, ok := data.(string)
		if !ok {
			return fmt.Errorf("Data mast be a string")
		}

		paramVal, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		outValElem.SetUint(paramVal)
	case reflect.Int:
		str, ok := data.(string)
		if !ok {
			return fmt.Errorf("Data mast be a string")
		}

		paramVal, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		outValElem.SetInt(paramVal)
	case reflect.Float32, reflect.Float64:
		str, ok := data.(string)
		if !ok {
			return fmt.Errorf("Data mast be a string")
		}

		paramVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		outValElem.SetFloat(paramVal)
	case reflect.String:
		res, ok := data.(string)
		if !ok {
			return fmt.Errorf("Data mast be a string")
		}
		outValElem.SetString(res)
	case reflect.Slice:
		dataKind := dataVal.Kind()
		if dataKind != reflect.Slice {
			return fmt.Errorf("Wrong type of data: %v", dataKind)
		}
		len := dataVal.Len()
		elemType := outType.Elem().Elem()
		slice := reflect.MakeSlice(outValElem.Type(), 0, len)

		for i := 0; i < len; i++ {
			dataElem := dataVal.Index(i)
			elem := reflect.New(elemType)

			err := strings2struct(dataElem.Interface(), elem.Interface())
			if err != nil {
				return err
			}
			e := reflect.Indirect(elem)
			slice = reflect.Append(slice, e)
		}
		outValElem.Set(slice)
	case reflect.Map:
		return fmt.Errorf("Kind Map\n")
	case reflect.Struct:
		dataKind := dataVal.Kind()
		if dataKind != reflect.Map { //	структура после анмаршалинга распознаётся как мапа
			return fmt.Errorf("Wrong type of data: %v", dataKind)
		}
		iter := dataVal.MapRange()

		indexesByNames := structFieldIndexesByJsonName(outValElem.Type())

		for iter.Next() {
			k := iter.Key()
			v := iter.Value()

			i, ok := indexesByNames[k.String()]
			if !ok {
				continue
			}
			field := outValElem.Field(i)

			if !field.CanAddr() {
				return fmt.Errorf("Cannot get address!")
			}
			err := strings2struct(v.Interface(), field.Addr().Interface())
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("Kind unknown! Kind() = %v ; data = %#v\n", dataVal.Kind(), data)
	}

	return nil
}

func structFieldIndexesByJsonName(s reflect.Type) map[string]int {
	numField := s.NumField()
	res := make(map[string]int, numField)

	for i := 0; i < numField; i++ {
		field := s.Field(i)
		name := field.Tag.Get("json")
		if name == "" {
			name = field.Name
		}
		res[name] = i
	}
	return res
}
