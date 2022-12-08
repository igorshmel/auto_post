package config

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// configPathFlag - флаг запуска приложения, указывающий путь до файла с настройками приложения
	configPathFlag = "."
)

var (
	errConfigPathEmpty    = fmt.Errorf("the configuration path is empty")
	errFileExtensionEmpty = fmt.Errorf("the file extension is empty")
	errWrongFileName      = fmt.Errorf("wrong the file name")
)

// AppConfig - общие настройки приложения
type AppConfig struct {
	Port string `mapstructure:"port" validate:"required"`
}

// DbConfig - настройки базы данных
type DbConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Name     string `mapstructure:"name" validate:"required"`
}

// KafkaConfig настройки для подключения к кластеру kafka (https://docs.confluent.io/platform/current/installation/configuration/index.html)
type KafkaConfig struct {
	// BootstrapServers - адреса для подключения к кластеру kafka
	BootstrapServers []string `mapstructure:"bootstrap_servers" validate:"required"`

	// SSL - настройки для SSL подключения
	SSL struct {
		// Enabled - флаг необходимости использовать SSL для подключения
		Enabled bool `mapstructure:"enabled"`

		// CertLocation - расположение сертификатов для подключения к кластеру
		CertLocation string `mapstructure:"ca_location" validate:"required"`
	} `mapstructure:"ssl"`

	// SASL - настройки для SASL аутентификации
	SASL struct {
		// Enabled - флаг необходимости использовать SASL для подключения
		Enabled bool `mapstructure:"enabled"`

		// Mechanisms - механизм аутентификации при подключении к кластеру
		Mechanisms string `mapstructure:"mechanisms" validate:"required"`

		// Username - имя пользователя для механизма аутентификации
		Username string `mapstructure:"username" validate:"required"`

		// Password - пароль для механизма аутентификации
		Password string `mapstructure:"password" validate:"required"`
	} `mapstructure:"sasl"`

	// Consumer - настройки для потребителя сообщений (https://docs.confluent.io/platform/current/installation/configuration/consumer-configs.html)
	Consumer struct {
		// SecurityProtocol - протокол общения с брокерами
		SecurityProtocol string `mapstructure:"security_protocol" validate:"required"`

		// GroupID - уникальный идентификатор группы потребителей, к которой принадлежит текущий потребитель
		GroupID string `mapstructure:"group_id" validate:"required"`

		// SessionTimeout - таймаут для обнаружения сбоя клиента kafka
		SessionTimeout string `mapstructure:"session_timeout_ms"`

		// AutoOffsetReset - стратегия выбора оффсета для вычитывания сообщений
		AutoOffsetReset string `mapstructure:"auto_offset_reset"`
	} `mapstructure:"consumer"`

	// Producer - настройки для производителя сообщений (https://docs.confluent.io/platform/current/installation/configuration/producer-configs.html)
	Producer struct {
		// SecurityProtocol - протокол общения с брокерами
		SecurityProtocol string `mapstructure:"security_protocol" validate:"required"`

		// SessionTimeout - таймаут для обнаружения сбоя клиента kafka
		SessionTimeout string `mapstructure:"session_timeout_ms"`
	} `mapstructure:"producer"`
}

// SchemaRegistryConfig - настройки для подключения к Schema Registry
type SchemaRegistryConfig struct {
	URL string `mapstructure:"url" validate:"required"`
}

// Config хранит конфигурацию приложения
type Config struct {
	// Настройки приложения
	App AppConfig `mapstructure:"app"`

	// Настройки БД
	Database DbConfig `mapstructure:"database"`

	// Настройки Kafka
	Kafka KafkaConfig `mapstructure:"kafka"`

	// Настройки Schema Registry
	SchemaRegistry SchemaRegistryConfig `mapstructure:"schema_registry"`
}

// CreateDSN create a dsn string for postgres
func (c *Config) CreateDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s database=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
	)
}

// Load загружает конфигурационный файл
func Load() (config Config, err error) {
	pflag.String(configPathFlag, "config.yaml", "name of the config file")

	// Парсинг флагов запуска приложения
	pflag.Parse()
	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return Config{}, err
	}

	// Путь до конфигурационного файла
	cfgPath := viper.GetString(configPathFlag)

	return LoadFromPath(cfgPath)
}

// LoadFromPath reads configuration from file or environment variables.
func LoadFromPath(cfgPath string) (config Config, err error) {
	dir, filename, ext, err := prepareConfigPathToViper(cfgPath)
	if err != nil {
		return Config{}, err
	}

	// set config file path, name and extension (e.g,'/path-to-config/app.env')
	viper.AddConfigPath(dir)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)

	// enable loading of env vars values
	// it'll override the values in the config file
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("config file '%s' not found", cfgPath)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	v := validator.New()
	err = v.Struct(&config)
	return config, err
}

// prepareConfigPathToViper разбирает путь до конфигурационного файла, согласно требованиям Viper
func prepareConfigPathToViper(cfgPath string) (string, string, string, error) {
	if len(cfgPath) == 0 {
		return "", "", "", errConfigPathEmpty
	}

	dir := filepath.Dir(cfgPath)
	base := filepath.Base(cfgPath)
	rawExt := filepath.Ext(base)
	if len(rawExt) == 0 {
		return "", "", "", errFileExtensionEmpty
	}
	ext := strings.Replace(rawExt, ".", "", 1)
	if len(ext) == 0 {
		return "", "", "", errFileExtensionEmpty
	}

	splitBase := strings.Split(base, filepath.Ext(base))
	filename := ""
	if len(splitBase) == 0 {
		return "", "", "", errWrongFileName
	}
	filename = splitBase[0]

	return dir, filename, ext, nil
}
