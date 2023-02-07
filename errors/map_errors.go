package errors

// MapError Map格式的错误
type MapError map[string]interface{}

//func (e *MapError) Error() string {
//	errb, _ := json.Marshal(map[string]interface{}(*e))
//	return string(errb)
//}

// NewMapError 新建MapError
//func NewMapError(args ...map[string]interface{}) *MapError {
//	var mape MapError
//	if len(args)> 0 {
//		mape = MapError(args[0])
//	} else {
//		mape = MapError(map[string] interface{})
//	}
//	return &mape
//}

func (e *MapError) Add(field string, msg interface{}) *MapError {
	ee := map[string]interface{}(*e)
	ee[field] = msg
	*e = MapError(ee)
	return e
}
