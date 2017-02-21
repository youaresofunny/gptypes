package types

type Plugin struct {
	Name    string
	Actions map[string][]HookFunc
}

type HookFunc func()

//New - add your plugin here
func New(name string) *Plugin {
	return &Plugin{Name: name, Actions: make(map[string][]HookFunc)}
}

//AddAction - add action to your plugin
func (p *Plugin) AddAction(action string, f HookFunc) {
	p.Actions[action] = append(p.Actions[action], f)
}

func (p *Plugin) DoAction(action string) {
	for _, fn := range p.Actions[action] {
		fn()
	}
}
