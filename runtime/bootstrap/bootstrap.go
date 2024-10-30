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

func NewBootstrap(dir, path string) Bootstrap {
	return Bootstrap{
		WorkDir:    dir,
		ConfigPath: path,
		Env:        "dev",
		Daemon:     false,
	}
}
