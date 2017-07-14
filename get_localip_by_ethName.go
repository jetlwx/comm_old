package comm

import (
	"log"
	"net"
	"strings"
)

func LocalIP(ethName string) string {
	i, err := net.InterfaceByName(ethName)
	if err != nil {
		log.Println("[E] an error at method LocalIP", err)
		return ""
	}

	add, err := i.Addrs()
	if err != nil {
		log.Println("[E] an error at method LocalIP i.addr()", err)
		return ""
	}

	for _, v := range add {
		str := strings.Split(v.String(), "/")
		if len(str) == 2 {
			if len(strings.Split(str[0], ".")) == net.IPv4len {
				return str[0]

			}
		}
	}

	return ""
}
