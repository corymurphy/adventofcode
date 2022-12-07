package shared

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(err)
	}
	return
}
