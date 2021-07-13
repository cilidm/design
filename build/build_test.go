package build

import (
	"fmt"
	"testing"
)

func TestBuild1(t *testing.T){
	build := &Build1{}
	director := NewDirector(build)
	director.Construct()
	res := build.GetResult()
	fmt.Println(res)
	if res != "123"{
		t.Fatalf("Builder1 fail expect 123 acture %s", res)
	}
}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	fmt.Println(res)
	if res != 6 {
		t.Fatalf("Builder2 fail expect 6 acture %d", res)
	}
}