package mysql

type Config struct {
	Provider string   `yaml:"provider"`
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	Database string   `yaml:"database"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	Options  []string `yaml:"options"`
}
