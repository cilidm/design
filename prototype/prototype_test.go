package prototype

import (
	"fmt"
	"testing"
)

type Proto1 struct {
	name string
}

func (p *Proto1) Clone() Cloneable{
	tc := *p
	return &tc
}

var manager *PrototypeManager

func init(){
	manager = NewPrototypeManager()
	t1 := &Proto1{name: "type1"}
	manager.Set("t1",t1)
}

func TestPrototype(t *testing.T){
	t1 := manager.Get("t1")
	t2 := t1.Clone()
	t3 := t1
	fmt.Println(t1,t2,t3)
	if t1 == t2{
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManager(t *testing.T){
	c := manager.Get("t1").Clone()
	t1 := c.(*Proto1)
	fmt.Println(t1,t1.name)
	if t1.name != "type1" {
		t.Fatal("error")
	}
}