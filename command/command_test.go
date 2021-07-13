package command

import "testing"

func TestCommand(t *testing.T){
	mb := &MotherBoard{}
	start := NewStartCommand(mb)
	reboot := NewRebootCommand(mb)

	box1 := NewBox(start,reboot)
	box1.PressButton1()
	box1.PressButton2()

	box2 := NewBox(reboot,start)
	box2.PressButton1()
	box2.PressButton2()
}
