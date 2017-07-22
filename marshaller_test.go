package tts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var strUser = "irfani    arief     26"

type (
	User struct {
		FirstName  string `txt_width:"10" pad_dir:"left" pad_str:" "`
		MiddleName string `txt_width:"-" json:"-"`
		LastName   string `txt_width:"10" pad_dir:"left" pad_str:" "`
		Age        int    `txt_width:"2"`
	}

	UserWrongWidth struct {
		FirstName string `txt_width:"abc"`
	}
	UserTooWide struct {
		FirstName string `txt_width:"999"`
	}
)

func TestUnmarshallCorrect(t *testing.T) {
	user := User{}
	err := Unmarshall(strUser, &user)
	assert.NoError(t, err)
	byteJSON, err := json.Marshal(user)
	assert.NoError(t, err)
	// assert.Equal(t, "{\"FirstName\":\"irfani\",\"LastName\":\"arief\"}", string(byteJSON))
	assert.Equal(t, "{\"FirstName\":\"irfani\",\"LastName\":\"arief\",\"Age\":26}", string(byteJSON))
}
func TestUnmarshallEmptyString(t *testing.T) {

	user := User{}
	err := Unmarshall("", &user)
	assert.Error(t, err, "string is empty")
}

func TestUnmarshallWrongWidth(t *testing.T) {
	user := UserWrongWidth{}
	err := Unmarshall(strUser, &user)
	assert.Error(t, err, "invalid width")
}

func TestUnmarshallTooWide(t *testing.T) {
	user := UserTooWide{}
	err := Unmarshall(strUser, &user)
	assert.NoError(t, err)
	assert.Equal(t, UserTooWide{}, user)
}
