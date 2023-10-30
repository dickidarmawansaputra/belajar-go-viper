package test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestViperJson(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("../")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar-go-viper", config.GetString("app.name"))
	assert.Equal(t, "Dicki Darmawan Saputra", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
}

func TestViperYaml(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("../config.yaml")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar-go-viper", config.GetString("app.name"))
	assert.Equal(t, "Dicki Darmawan Saputra", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))

}

func TestViperEnvFile(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("../config.env")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar-go-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Dicki Darmawan Saputra", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DB_HOST"))
}

func TestViperEnvVariable(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("../config.env")
	// untuk baca env variable dr OS
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar-go-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Dicki Darmawan Saputra", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DB_HOST"))

	// ENV VARIABLE
	// RUN COMMAND DI TERMINAL
	// export FROM_ENV=Hello
	// run go test
	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
	assert.Equal(t, "Hello Dicki", config.GetString("ENV_VAR_NAME"))
}
