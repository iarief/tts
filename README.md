# tts
Simple package to marshall and unmarshall fixed width text to struct


How to use:

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/iarief/tts"
)

type User struct {
	FirstName    string `txt_width:"10" pad_dir:"right" pad_str:" "`
	LastName     string `txt_width:"10" pad_dir:"right" pad_str:" "`
	IgnoredValue string `txt_width:"-" json:"-"`
	Age          int    `txt_width:"3" pad_dir:"left" pad_str:"0"`
}

func main() {
	user := User{"john", "doe", "ignore this", 20}
	str, _ := tts.Marshal(&user)
	fmt.Println(str) // "john      doe       020"

	user2 := User{}
	_ = tts.Unmarshal(str, &user2)
	byteJSON, _ := json.Marshal(user2)
	fmt.Println(string(byteJSON)) // "{\"FirstName\":\"john\",\:LastName\":\"doe\",\"Age\":20}"
}
```

Please note while marshaling the string is forced to be the length that is defined in txt_width tag, either by substring or padding (which could be configured in pad tag)
