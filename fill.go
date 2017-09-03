package funk

import (
	"reflect"
	"errors"
	"fmt"
)

func Fill(in interface{}, fillValue interface{}) (interface{}, error) {
	inValue := reflect.ValueOf(in)
	inKind := inValue.Type().Kind()
	if inKind != reflect.Slice && inKind != reflect.Array {
		return nil, errors.New("Can only fill slices and arrays")
	}

	inType := reflect.TypeOf(in).Elem()
	value := reflect.ValueOf(fillValue)
	if inType != value.Type() {
		return nil, errors.New(fmt.Sprintf(
			"Cannot fill '%s' with '%s'", reflect.TypeOf(in), value.Type(),
		))
	}

	for i := 0; i < inValue.Len(); i++ {
		inValue.Index(i).Set(value)
	}
	return in, nil
}
