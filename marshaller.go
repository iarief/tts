package tts

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

const (
	widthTag  = "txt_width"
	padDirTag = "pad_dir"
	padStrTag = "pad_str"
)

// Unmarshall will unmarshall a fixed width string to a struct, the struct need to have the tag "txt_width" that contains an int, the field would be parsed by it's definition order
// Use `txt_width:"-"` if you don't want the field to be parsed
func Unmarshall(str string, obj interface{}) error {
	if len(str) < 1 {
		return errors.New("string is empty")
	}
	elemsType := reflect.TypeOf(obj).Elem()
	elemsVal := reflect.ValueOf(obj).Elem()

	for i := 0; i < elemsType.NumField(); i++ {
		elemVal := elemsVal.Field(i)
		elemType := elemsType.Field(i)
		if elemVal.CanSet() {
			tag := elemType.Tag
			widthTxt := tag.Get(widthTag)
			if widthTxt == "-" {
				continue
			}
			width, err := strconv.Atoi(widthTxt)
			if err != nil {
				return errors.New("invalid width")
			}
			if len(str) >= width {
				val := strings.TrimSpace(str[:width])
				str = str[width:]
				kind := elemVal.Kind()
				switch kind {
				case reflect.Int:
					// try to convert to int, but continue anyway if it failed
					valInt, _ := strconv.ParseInt(val, 10, 64)
					elemVal.SetInt(valInt)
				case reflect.String:
					elemVal.SetString(val)
				}
			} else {
				break
			}
		}
	}
	return nil
}

/* TODO
1. Create marshaller function
2. Add more type to handle in unmarshall
3. Go for a vacation somewhere far away
*/
