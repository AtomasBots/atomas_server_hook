package atomas
import (
	"net/http"
	"os/exec"
	"fmt"
	"io/ioutil"
	"strings"
)

func CreateHealthService() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("bash", "system_health")
		cmd.Start()
		cmd.Wait()
		dat, _ := ioutil.ReadFile("system_health.out")
		for _, element := range strings.Split(string(dat),"\n"){
			fmt.Fprintln(w, ToJsonString(element))
		}
	}
}