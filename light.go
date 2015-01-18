package gohue

type State struct {
	On        bool
	Bri       int
	Hue       int
	Sat       int
	Effect    string
	Xy        []int
	Ct        int
	Alert     string
	Colormode string
	Reachable bool
}

type Light struct {
	State       State
	Type        string
	Name        string
	Modelid     string
	Uniqueid    string
	Swversion   string
	Pointsymbol map[string]string
}
