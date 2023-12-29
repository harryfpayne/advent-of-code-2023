package main

type Pulse bool

const (
	LowPulse  Pulse = false
	HighPulse Pulse = true
)

type IBroadcaster interface {
	Process(string, Pulse) (Pulse, bool)
	GetLabel() string
	GetOutputs() []string
	AddInput(string)
}

type Broadcaster struct {
	inputs  []string
	outputs []string
	label   string
}

func (b Broadcaster) Process(_ string, p Pulse) (Pulse, bool) {
	return p, true
}

func NewBroadcaster(label string, inputs []string, outputs []string) *Broadcaster {
	return &Broadcaster{
		inputs:  inputs,
		outputs: outputs,
		label:   label,
	}
}

type FlipFlop struct {
	Broadcaster
	state bool
}

func (f *FlipFlop) Process(_ string, p Pulse) (Pulse, bool) {
	if p == HighPulse {
		return LowPulse, false
	}
	if f.state == false {
		f.state = true
		return HighPulse, true
	}
	f.state = false
	return LowPulse, true
}

func NewFlipFlop(label string, inputs []string, outputs []string) *FlipFlop {
	return &FlipFlop{
		Broadcaster: *NewBroadcaster(label, inputs, outputs),
		state:       false,
	}
}

type Conjunction struct {
	Broadcaster
	inputMemory map[string]Pulse
}

func (c *Conjunction) Process(i string, p Pulse) (Pulse, bool) {
	c.inputMemory[i] = p

	for _, pulse := range c.inputMemory {
		if pulse == LowPulse {
			return HighPulse, true
		}
	}
	return LowPulse, true
}

func NewConjunction(label string, inputs []string, outputs []string) *Conjunction {
	mem := make(map[string]Pulse, len(inputs))
	for _, input := range inputs {
		mem[input] = LowPulse
	}

	return &Conjunction{
		Broadcaster: *NewBroadcaster(label, inputs, outputs),
		inputMemory: mem,
	}
}

func (b Broadcaster) GetOutputs() []string {
	return b.outputs
}
func (f *FlipFlop) GetOutputs() []string {
	return f.Broadcaster.GetOutputs()
}
func (c *Conjunction) GetOutputs() []string {
	return c.Broadcaster.GetOutputs()
}

func (b *Broadcaster) AddInput(input string) {
	b.inputs = append(b.inputs, input)
}
func (f *FlipFlop) AddInput(input string) {
	f.Broadcaster.AddInput(input)
}
func (c *Conjunction) AddInput(input string) {
	c.Broadcaster.AddInput(input)
	c.inputMemory[input] = LowPulse
}

func (b Broadcaster) GetLabel() string {
	return b.label
}
func (f *FlipFlop) GetLabel() string {
	return f.Broadcaster.GetLabel()
}
func (c *Conjunction) GetLabel() string {
	return c.Broadcaster.GetLabel()
}
