package framework

import "fmt"

type PluginFunc func(program string) string

type Machine struct {
	tape       *Tape
	plugins    []PluginFunc
	loop_jumps map[uint]uint
}

func NewMachine(tape *Tape) *Machine {
	return &Machine{
		tape:       tape,
		loop_jumps: make(map[uint]uint),
	}
}

func (m *Machine) Use(plugins ...PluginFunc) {
	m.plugins = append(m.plugins, plugins...)
}

func (m *Machine) Run(program string) {

	for _, pluginFunc := range m.plugins {
		program = pluginFunc(program)
	}

	var stack []uint = make([]uint, 2048)
	var top int = -1
	var pi uint = 0
	for pi < uint(len(program)) {
		switch program[pi] {
		case LOOP_BEGIN:
			top++
			stack[top] = pi
		case LOOP_END:
			ti := stack[top]
			top--
			m.loop_jumps[ti] = pi
			m.loop_jumps[pi] = ti
		}
		pi++
	}

	pi = 0
	for pi < uint(len(program)) {
		switch program[pi] {
		case WRITE:
			var c byte
			fmt.Scanf("%c", &c)
			m.tape.Set(c)
		case READ:
			c := m.tape.Get()
			fmt.Printf("%c", c)
		case INC:
			m.tape.Inc()
		case DEC:
			m.tape.Dec()
		case INC_PTR:
			m.tape.IncPtr()
		case DEC_PTR:
			m.tape.DecPtr()
		case LOOP_BEGIN:
			if m.tape.Get() == 0x0 {
				pi = m.loop_jumps[pi]
			}
		case LOOP_END:
			if m.tape.Get() != 0x0 {
				pi = m.loop_jumps[pi]
			}
		}
		pi++
	}
}
