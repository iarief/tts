package tts

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

const (
	// ErrorInvalidWidth is string representation of the error returned
	// when the function is called with an empty string
	ErrorInvalidWidth = "invalid width"
	// ErrorEmptyString is string representatino of the error returned
	// when there are one or more invalid txt_width tag in the struct
	ErrorEmptyString = "string is empty"
	// ErrorEmptyStruct is string representation of the error returned
	// when marshal is called with an empty struct
	ErrorEmptyStruct = "struct is empty"

	widthTag  = "txt_width"
	padDirTag = "pad_dir"
	padStrTag = "pad_str"
	defPadDir = "left"
	defPadStr = " "
)

// Unmarshal will Unmarshal a fixed width string to a struct, the struct need to have the tag "txt_width" that contains an int, the field would be parsed by it's definition order
// Use `txt_width:"-"` if you don't want the field to be parsed
func Unmarshal(str string, obj interface{}) error {
	if len(str) < 1 {
		return errors.New(ErrorEmptyString)
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
				return errors.New(ErrorInvalidWidth)
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
				case reflect.Float64:
					valFloat, _ := strconv.ParseFloat(val, 64)
					elemVal.SetFloat(valFloat)
				case reflect.String:
					elemVal.SetString(val)
				default: // default we're handling it as string
					elemVal.SetString(val)
				}
			} else {
				break
			}
		}
	}
	return nil
}

// Marshal is a function to marshal struct into string of fixed width determined by the struct tag
func Marshal(obj interface{}) (str string, err error) {

	result := ""
	if obj == nil {
		return "", errors.New(ErrorEmptyStruct)
	}
	elemsType := reflect.TypeOf(obj)
	elemsVal := reflect.ValueOf(obj)
	if elemsType.Kind() == reflect.Ptr {
		elemsType = elemsType.Elem()
		elemsVal = elemsVal.Elem()
	}
	for i := 0; i < elemsType.NumField(); i++ {
		elemVal := elemsVal.Field(i)
		elemType := elemsType.Field(i)
		tag := elemType.Tag
		widthTxt := tag.Get(widthTag)
		padStr := tag.Get(padStrTag)
		padDir := tag.Get(padDirTag)

		if widthTxt == "-" {
			continue
		}
		if padStr == "" {
			padStr = defPadStr
		}
		if padDir == "" {
			padDir = defPadDir
		}
		width, err := strconv.Atoi(widthTxt)
		if err != nil {
			return "", errors.New(ErrorInvalidWidth)
		}
		if !elemVal.CanInterface() {
			continue
		}
		strVal := ""
		kind := elemVal.Kind()
		switch kind {
		case reflect.Int:
			strVal = strconv.FormatInt(elemVal.Int(), 10)
		case reflect.Float64:
			strVal = strconv.FormatFloat(elemVal.Float(), 'f', -1, 64)
		case reflect.String:
			strVal = elemVal.String()
		default: // default we're handling it as string
			strVal = elemVal.String()
		}

		if len(strVal) > width {
			strVal = strVal[:width]
		} else if len(strVal) < width {
			strPad := strings.Repeat(padStr, width-len(strVal))
			if padDir == "left" {
				strVal = strPad + strVal
			} else if padDir == "right" {
				strVal = strVal + strPad
			}
		}
		result += strVal

	}
	return result, nil
}

/* TODO
1. Add more type to handle
*/
