package starter

type config struct {
	// 地址
	Addr string `json:"addr" xml:"addr" yaml:"addr" toml:"addr" validate:"hostname_port"`
	// 插值
	Iv string `json:"iv" xml:"iv" yaml:"iv" toml:"iv"`
}
