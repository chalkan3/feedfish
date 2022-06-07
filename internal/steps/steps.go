package steps

type Initial struct {
	Script string `yam:"script,omitempty"`
}

type Proceed struct {
	Script string `yam:"script,omitempty"`
}

type Processing struct {
	Script string `yam:"script,omitempty"`
}

type Scripts struct {
	Initial    *Initial    `yaml:"initial,omitempty"`
	Proceed    *Proceed    `yaml:"proceed,omitempty"`
	Processing *Processing `yaml:"processing,omitempty"`
}
type Step struct {
	Name    string   `yaml:"name,omitempty"`
	Order   int      `yaml:"order,omitempty"`
	Scripts *Scripts `yaml:"scripts,omitempty"`
}

type Pipeline struct {
	Steps []Step `yaml:"steps,omitempty"`
}
