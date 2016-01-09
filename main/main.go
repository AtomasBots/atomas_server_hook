package main
import (
	"net/http"
	"log"
	"os"
	"github.com/AtomasBots/atomas_server_hook/atomas"
	"os/exec"
)

var version string

func doOnHook() {
	exec.Command("bash", "onhook").Start()
}

func main() {
	hooks := atomas.Hooks{[]string{}}
	http.HandleFunc("/hook", atomas.CreateHookHandler(&hooks, doOnHook))
	http.HandleFunc("/hook_list", atomas.CreateHookListHandler(&hooks))
	http.HandleFunc("/health", atomas.CreateHealthService())
	http.HandleFunc("/version", atomas.CreateVersionHandler(version))
	port := os.Args[1]
	doOnHook()
	log.Fatal(http.ListenAndServe(port, nil))
}