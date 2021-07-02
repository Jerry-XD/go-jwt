package config

type (
	Config struct {
		App App
	}

	App struct {
		Port string
	}
)

func New() *Config {
	return &Config{
		App{
			Port: "8080",
		},
	}
}
