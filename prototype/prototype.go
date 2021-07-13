package prototype

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (m *PrototypeManager) Get(name string) Cloneable {
	return m.prototypes[name]
}

func (m *PrototypeManager) Set(name string, value Cloneable) {
	m.prototypes[name] = value
}
