package cfg

import (
	"github.com/spf13/viper"
)

func New() *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix("BARBAND")
	return v
}
