package webapp

import "github.com/goccy/go-yaml"

type Config struct {
	Revision string

	Debug bool `yaml:"debug"`

	ServerIPAddress string `yaml:"server_ip_address"`
	ServerPort      int    `yaml:"server_port"`

	ExporterIPAddress string `yaml:"exporter_ip_address"`
	ExporterPort      int    `yaml:"exporter_port"`

	Https          bool   `yaml:"https"`
	CetificatePath string `yaml:"certificate_path"`
	PrivateKeyPath string `yaml:"private_key_path"`

	SkipBasicAuth     bool   `yaml:"skip_basic_auth"`
	BasicAuthUsername string `yaml:"basic_auth_username"`
	BasicAuthPassword string `yaml:"basic_auth_password"`

	LogDir    string `yaml:"log_dir"`
	LogName   string `yaml:"log_name"`
	LogStdout bool   `yaml:"log_stdout"`
}

func InitConfig(data []byte, config interface{}) error {
	if err := yaml.Unmarshal(data, config); err != nil {
		// パースに失敗した場合 Fatal で終了
		return err
	}

	// TODO(v): 初期値

	return nil
}
