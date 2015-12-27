package main
import (
	"net/http"
	"log"
	"os"
	"github.com/AtomasBots/atomas_server_hook/atomas"
)

var version string

func main() {
	hooks := atomas.Hooks{[]string{}}
	http.HandleFunc("/hook", atomas.CreateHookHandler(&hooks))
	http.HandleFunc("/hook_list", atomas.CreateHookListHandler(&hooks))
	http.HandleFunc("/health", atomas.CreateHealthService())
	http.HandleFunc("/version", atomas.CreateVersionHandler(version))
	port := os.Args[1]
	log.Fatal(http.ListenAndServe(port, nil))
}