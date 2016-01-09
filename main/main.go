package main
import (
	"net/http"
	"log"
	"os"
	"github.com/AtomasBots/atomas_server_hook/atomas"
	"os/exec"
	"time"
	"io/ioutil"
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
	go doOnHook()
	go rebootWhenNoInternet()
	log.Fatal(http.ListenAndServe(port, nil))
}

func rebootWhenNoInternet() {
	time.Sleep(300)
	for {
		if (isWifiOk()) {
			time.Sleep(time.Second)
		} else {
			time.Sleep(time.Second * 10)
			if (!isWifiOk()) {
				ioutil.WriteFile(time.Now().String(), []byte(time.Now().String()), 0644)
				exec.Command("sudo", "reboot").Start()
			}
		}
	}
}

func isWifiOk() bool {
	command := exec.Command("ping", "-c", "1", "192.168.0.1")
	command.Start()
	err := command.Wait()
	return err == nil
}