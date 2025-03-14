package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are read by viper from a config  file or environment varialbes
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"` // the ENV name as they appear in .env file
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads configurations from a config file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // could be json, xml  etc.

	// reads configuration from environment variables
	// overrides configuration in config file with provided environment variable
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
