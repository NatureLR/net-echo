//go:build !windows && !drawin

package netecho

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/vishvananda/netlink"
)

type Info struct {
	ClientAddr     string
	ClientReqPath  string
	ClientReqMeth  string
	ServerHostName string
	ServerAddr     string
}

func (i *Info) Output(w http.ResponseWriter) {
	var ret []string

	vr := reflect.ValueOf(*i)
	for i := 0; i < vr.NumField(); i++ {
		fieldInfo := vr.Type().Field(i)
		ret = append(ret, fieldInfo.Name+": "+vr.Field(i).String())
	}

	data := fmt.Sprintf("%s\n", strings.Join(ret, "\n"))

	fmt.Println(data)
	w.Write([]byte(data))
}

func getGwAddr() (addrs []string, err error) {
	links, err := netlink.LinkList()
	if err != nil {
		return
	}

	for _, link := range links {
		rs, e := netlink.RouteList(link, netlink.FAMILY_V4)
		if e != nil {
			err = e
			return
		}
		for _, r := range rs {
			if r.Gw != nil {
				as, e := netlink.AddrList(link, netlink.FAMILY_V4)
				if e != nil {
					err = e
					return
				}
				for _, a := range as {
					addrs = append(addrs, a.IP.String())
				}
				return
			}
		}
	}
	return
}

func AddrByName(name string) string {
	list, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return ""
	}

	for _, iface := range list {
		if iface.Name != name {
			continue
		}

		as, err := iface.Addrs()
		if err != nil {
			log.Println(err)
			continue
		}
		for _, a := range as {
			return a.String()
		}
	}
	return ""
}

func handle(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		log.Println(err)
	}
	saddr, err := getGwAddr()
	if err != nil {
		log.Fatalln(err)
		return
	}

	info := &Info{
		ClientAddr:     r.RemoteAddr,
		ClientReqPath:  r.RequestURI,
		ClientReqMeth:  r.Method,
		ServerHostName: host,
		ServerAddr:     saddr[0],
	}
	info.Output(w)
}

func Run() {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalln(err)
	}
}
