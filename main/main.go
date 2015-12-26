package main
import (
	"net/http"
	"log"
	"os"
	"github.com/AtomasBots/atomas_server_hook/atomas"
)


func main() {
	hooks := atomas.Hooks{[]string{}}
	http.HandleFunc("/hook", atomas.CreateHookHandler(&hooks))
	http.HandleFunc("/hook_list", atomas.CreateHookListHandler(&hooks))
	port := os.Args[1]
	log.Fatal(http.ListenAndServe(port, nil))
}