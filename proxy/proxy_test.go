package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T){
	var sub Subject
	sub = &Proxy{}
	res := sub.Do()
	fmt.Println(res)

	var pro Proxy
	res1 := pro.Do()
	fmt.Println(res1)
}
