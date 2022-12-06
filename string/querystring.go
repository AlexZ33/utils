package string

import (
	"net/url"
	"reflect"
	"strings"
)

//querystring parsing and stringifying

// Marshal nested struct or map to url query strings
func Marshal(obj interface{}) (string, error) {
	values := make(url.Values)
	err := encode(obj, "", values)
}

func encode(val interface{}, prefix string, values url.Values) error {
	v := reflect.ValueOf(val)

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct:
		err := encodeStruct(v, prefix, values)

	}

}

func parseFieldTag(tag string) (string, []string) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}

func encodeStruct(v reflect.Value, prefix string, values url.Values) error {
	typ := v.Type()

	for i := 0; i < v.NumMethod(); i++ {
		tf := typ.Field(i)
		sv := v.Field(i)

		tag, _ := tf.Tag.Lookup("url")
	}
}
