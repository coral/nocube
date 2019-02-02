package main

import (
	"github.com/coral/nocube/pkg/utils"
	kcp "github.com/xtaci/kcp-go"
)

func main() {

	kcpconn, err := kcp.DialWithOptions("192.168.0.1:10000", nil, 10, 3)

	var bytes = []byte{}
	for i := 0; i < 32; i++ {
		bytes = append(bytes, []byte{
			utils.Clamp255(127),
			utils.Clamp255(127),
			utils.Clamp255(127),
		}...)
	}

}
