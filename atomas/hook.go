package atomas
import (
	"net/http"
	"fmt"
)

func CreateHookHandler(hooks *Hooks) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hooks.Hooks = append(hooks.Hooks, r.RemoteAddr)
		fmt.Fprint(w, ToJsonString(hooks))
	}
}