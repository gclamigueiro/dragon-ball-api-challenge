package db

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func NewConfig(host, port, user, password, name string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
	}
}

func (c *Config) IsValid() bool {
	return c.Host != "" && c.Port != "" && c.User != "" && c.Password != "" && c.Name != ""
}
