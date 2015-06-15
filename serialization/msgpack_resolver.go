package serialization
import "reflect"

type MsgPackResolver map[string]reflect.Value

func (self MsgPackResolver) Resolve(name string, arguments []reflect.Value) (reflect.Value, error) {
	return self[name], nil
}
