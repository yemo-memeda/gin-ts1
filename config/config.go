package config

type Configuration struct {
	App      App      `json:"app"`
	Log      Log      `json:"log"`
	Database Database `json:"db"`
}
