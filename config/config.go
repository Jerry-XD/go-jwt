package config

type (
	Config struct {
		App App
	}

	App struct {
		Port                   string
		SigningKeyAdminAccess  string
		SigningKeyAdminRefresh string
		AccessTokenExp         int64
		RefreshTokenExp        int64
	}
)

func New() *Config {
	return &Config{
		App{
			Port:                   "8080",
			SigningKeyAdminAccess:  "some-salt-access-key",
			SigningKeyAdminRefresh: "some-salt-refresh-key",
			AccessTokenExp:         2400,
			RefreshTokenExp:        360,
		},
	}
}
