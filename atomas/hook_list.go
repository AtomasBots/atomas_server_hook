package atomas
import (
	"net/http"
	"fmt"
	"encoding/json"
)

func CreateHookListHandler(hooks *Hooks) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, ToJsonString(hooks))
	}
}

func ToJsonString(any interface{}) string {
	json, _ := json.Marshal(any)
	return string(json)
}