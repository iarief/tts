package tts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var strUser = "john      doe       026  182.88"

type (
	User struct {
		FirstName   string `txt_width:"10" pad_dir:"right" pad_str:" "`
		MiddleName  string `txt_width:"-" json:"-"`
		LastName    string `txt_width:"10" pad_dir:"right" pad_str:" "`
		ignoredProp string `txt_width:"10"`
		Age         int    `txt_width:"3" pad_dir:"left" pad_str:"0"`
		Height      float64 `txt_width:"8" pad_dir:"left" pad_str:" "`
	}

	UserWrongWidth struct {
		FirstName string `txt_width:"abc"`
	}
	UserTooWide struct {
		FirstName string `txt_width:"999"`
	}
)

func TestUnmarshalDefault(t *testing.T) {
	user := User{}
	err := Unmarshal(strUser, &user)
	assert.NoError(t, err)
	byteJSON, err := json.Marshal(user)
	assert.NoError(t, err)
	// assert.Equal(t, "{\"FirstName\":\"irfani\",\"LastName\":\"arief\"}", string(byteJSON))
	assert.Equal(t, "{\"FirstName\":\"john\",\"LastName\":\"doe\",\"Age\":26,\"Height\":182.88}", string(byteJSON))
}
func TestUnmarshalEmptyString(t *testing.T) {

	user := User{}
	err := Unmarshal("", &user)
	assert.EqualError(t, err, ErrorEmptyString)
}

func TestUnmarshalWrongWidth(t *testing.T) {
	user := UserWrongWidth{}
	err := Unmarshal(strUser, &user)
	assert.EqualError(t, err, ErrorInvalidWidth)
}

func TestUnmarshalTooWide(t *testing.T) {
	user := UserTooWide{}
	err := Unmarshal(strUser, &user)
	assert.NoError(t, err)
	assert.Equal(t, UserTooWide{}, user)
}

func TestMarshalDefault(t *testing.T) {
	user := User{"john", "the", "doe", "smart", 26, 182.88}
	res, err := Marshal(user)
	assert.NoError(t, err)
	assert.Equal(t, "john      doe       026  182.88", res)
}

func TestMarshalEmptyStruct(t *testing.T) {
	res, err := Marshal(nil)
	assert.EqualError(t, err, ErrorEmptyStruct)
	assert.Empty(t, res)
}

func TestMarshalOverflowString(t *testing.T) {
	user := User{"johnjohnjohn", "the", "doe", "smart", 26, 182.88}
	res, err := Marshal(user)
	assert.NoError(t, err)
	assert.Equal(t, "johnjohnjodoe       026  182.88", res)
}

func TestMarshalWrongWidth(t *testing.T) {
	user := UserWrongWidth{"john"}
	res, err := Marshal(user)
	assert.EqualError(t, err, ErrorInvalidWidth)
	assert.Empty(t, res)
}

func TestMarshalPointer(t *testing.T) {
	user := User{"john", "the", "doe", "smart", 26, 182.88}
	res, err := Marshal(&user)
	assert.NoError(t, err)
	assert.Equal(t, "john      doe       026  182.88", res)
}
