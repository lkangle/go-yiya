package types

const (
	QuaNormal = 1
	QuaBest   = 2
)

type SystemOptions struct {
	WindowWidth   int  `json:"windowWidth" yaml:"window_width"`
	WindowHeight  int  `json:"windowHeight" yaml:"window_height"`
	WindowPosX    int  `json:"windowPosX" yaml:"window_pos_x"`
	WindowPosY    int  `json:"windowPosY" yaml:"window_pos_y"`
	IsAlwaysOnTop bool `json:"isAlwaysOnTop" yaml:"is_always_on_top"`
}

type CompressOptions struct {
	Quality   int    `json:"quality" yaml:"quality"`
	Override  bool   `json:"override" yaml:"override"`
	NewSuffix string `json:"newSuffix" yaml:"new_suffix"`
}

type AppOptions struct {
	System   SystemOptions   `json:"system" yaml:"system"`
	Compress CompressOptions `json:"compress" yaml:"compress"`
}
