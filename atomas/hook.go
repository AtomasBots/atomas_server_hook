package atomas
import (
	"net/http"
	"fmt"
	"os/exec"
)

func CreateHookHandler(hooks *Hooks) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hooks.Hooks = append(hooks.Hooks, r.RemoteAddr)
		exec.Command("bash", "onhook").Start()
		fmt.Fprint(w, ToJsonString(hooks))
	}
}