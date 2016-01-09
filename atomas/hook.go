package atomas
import (
	"net/http"
	"fmt"
	"net"
)

func CreateHookHandler(hooks *Hooks, doOnHook func()) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteIP := r.RemoteAddr
		if (isOnWhiteList(remoteIP)) {
			hooks.Hooks = append(hooks.Hooks, remoteIP)
			doOnHook()
			fmt.Fprint(w, ToJsonString(hooks))
		}else {
			http.Error(w, "Frobidden", http.StatusForbidden)
		}
	}
}

func isOnWhiteList(remoteIP string) bool {
	whiteList := []net.IPNet{
		parseNet("192.30.252.0/22"),
		parseNet("192.168.0.0/24"),
		parseNet("127.0.0.1/8"),
	}
	return contains(whiteList, ParseIP(remoteIP))
}

func ParseIP(s string) net.IP {
	ip, _, err := net.SplitHostPort(s)
	if err == nil {
		return net.ParseIP(ip)
	} else {
		return net.ParseIP(s)
	}
}

func contains(list []net.IPNet, ip net.IP) bool {
	for _, net := range list {
		if (net.Contains(ip)) {
			return true
		}
	}
	return false
}

func parseNet(s string) net.IPNet {
	_, n, err := net.ParseCIDR(s)
	if err != nil {
		panic(err)
	}
	return *n
}