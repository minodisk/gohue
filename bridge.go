package gohue

import "fmt"

type Bridge struct {
	host string
	user string
}

func NewBridge(host, user string) *Bridge {
	return &Bridge{host: host, user: user}
}

func (b Bridge) Lights() {
	lights := make(map[string]Light)
	if err := Fetch("GET", "http://192.168.1.8/api/minodisk00/lights", &lights); err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%#", lights)
	fmt.Printf("%s", lights["1"].State.On)
}
