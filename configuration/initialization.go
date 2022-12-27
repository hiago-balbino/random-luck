package configuration

import "github.com/spf13/viper"

// InitConfigurations starts the viper settings and read all app configurations.
func InitConfigurations() {
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
}
