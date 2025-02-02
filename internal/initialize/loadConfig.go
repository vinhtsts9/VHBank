package initialize

import (
	"Golang-Masterclass/simplebank/global"
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Tự động lấy biến môi trường nếu có
	viper.AutomaticEnv()

	// Đọc file cấu hình
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}

	// Ghi đè giá trị từ file cấu hình vào global.Config
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode configuration: %v", err)
	}
}
