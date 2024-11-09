package types

type SystemOptions struct {
	WindowWidth   int  `json:"windowWidth" yaml:"window_width"`
	WindowHeight  int  `json:"windowHeight" yaml:"window_height"`
	WindowPosX    int  `json:"windowPosX" yaml:"window_pos_x"`
	WindowPosY    int  `json:"windowPosY" yaml:"window_pos_y"`
	IsAllowsOnTop bool `json:"isAllowsOnTop" yaml:"isAllowsOnTop"`
}
