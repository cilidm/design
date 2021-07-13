package command

import "fmt"

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{mb: mb}
}

type RebootCommand struct {
	mb *MotherBoard
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{mb: mb}
}

type MotherBoard struct{}

func (c *MotherBoard) Start() {
	fmt.Print("system starting\n")
}

func (c *MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

type Box struct {
	button1 Command
	button2 Command
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}

func NewBox(button1, button2 Command) *Box {
	return &Box{button1, button2}
}
