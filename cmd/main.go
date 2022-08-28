package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type configs struct {
	ServiceInfo serviceInfo
}

type serviceInfo struct {
	ServiceGroupId    string
	ServiceInstanceID string
}

func main() {
	configs := loadConfigs()
	fmt.Println(configs)
}

func loadConfigs() *configs {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("An error encountered while reading config.yaml file. Did you removed .example suffix?")
	}

	configs := &configs{}

	configs.ServiceInfo = serviceInfo{
		ServiceGroupId:    viper.GetString("service-group-id"),
		ServiceInstanceID: viper.GetString("service-instance-id"),
	}

	return configs
}
