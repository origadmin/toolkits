package bootstrap

type Bootstrap struct {
	WorkDir    string
	ConfigPath string
	Env        string
	Daemon     bool
}

func DefaultBootstrap() Bootstrap {
	return Bootstrap{
		WorkDir:    ".",
		ConfigPath: "configs/config.toml",
		Env:        "dev",
		Daemon:     false,
	}
}
