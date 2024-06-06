package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigManager interface {
	GetAppConfig() (AppConfig, error)
	GetHttpServerConfig() (config HttpServerConfig, err error)
}

type AppConfig struct {
	Mode      string `mapstructure:"MODE"`
	LogPath   string `mapstructure:"LOG_PATH"`
	LogMaxAge string `mapstructure:"LOG_MAX_AGE"`
}

type HttpServerConfig struct {
	Port string `mapstructure:"SERVER_PORT"`
	Host string `mapstructure:"SERVER_HOST"`
}

type DBConfig struct {
	Host         string `mapstructure:"DB_HOST"`
	Port         string `mapstructure:"DB_PORT"`
	Username     string `mapstructure:"DB_USERNAME"`
	Password     string `mapstructure:"DB_PASSWORD"`
	Database     string `mapstructure:"DB_NAME"`
	Dialect      string `mapstructure:"DB_DIALECT"`
	MigrationDir string `mapstructure:"MIGRATION_DIR"`
}

type Redis struct {
	Host     string `mapstructure:"REDIS_HOST" env-default:"127.0.0.1"`
	Port     string `mapstructure:"REDIS_PORT" env-default:"6379"`
	Password string `mapstructure:"REDIS_PASS" validate:"required"`
	// Username string `mapstructure:"REDIS_USER" validate:"" `
	// DB       int    `mapstructure:"REDIS_DATABASE" validate:""`
}

type ServiceAuthConfig struct {
	WhiteList   string `mapstructure:"SERVICE_WHITE_LIST_IP"`
	SecretToken string `mapstructure:"SERVICE_SECRET_TOKEN"`
}

type PwsCallbackAuthConfig struct {
	WhiteList string `mapstructure:"PWS_CALLBACK_WHITE_LIST_IP"`
}

type PwsConfig struct {
	ServiceId  string `mapstructure:"PWS_SERVICE_ID"`
	Url        string `mapstructure:"PWS_URL"`
	Username   string `mapstructure:"PWS_USERNAME"`
	Password   string `mapstructure:"PWS_PASSWORD"`
	PrivateKey string `mapstructure:"PWS_PRIVATE_KEY"`
	PublicKey  string `mapstructure:"PWS_PUBLIC_KEY"`
}

type PwsActionConfig struct {
	WaitForActionInfoMin    string `mapstructure:"PWS_WAIT_FOR_ACTION_INFO_MIN"`
	GetInfoCycleDurationSec string `mapstructure:"PWS_GET_INFO_CYCLE_DURATION_SEC"`
	ActionCallback          string `mapstructure:"PWS_ACTION_CALLBACK"`
}

type IranicardConfig struct {
	Url string `mapstructure:"IRC_URL"`
}

type CacheConfig struct {
	DefaultCacheExpiration    string `mapstructure:"DEFAULT_CACHE_EXPIRATION"`
	UserAuthCacheExpiration   string `mapstructure:"USER_AUTH_CACHE_EXPIRATION"`
	PwsBalanceCacheExpiration string `mapstructure:"PWS_BALANCE_CACHE_EXPIRATION"`
}

func LoadEnv() error {
	return godotenv.Load("config/.env")
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetAppConfig() AppConfig {
	return AppConfig{
		Mode:      GetEnv("MODE", "development"),
		LogPath:   GetEnv("LOG_PATH", "./custodial_wallet.log"),
		LogMaxAge: GetEnv("LOG_MAX_AGE", "30"),
	}
}

func GetHttpServerConfig() HttpServerConfig {
	return HttpServerConfig{
		Port: GetEnv("SERVER_PORT", "8080"),
		Host: GetEnv("SERVER_HOST", "127.0.0.1"),
	}
}

func GetDBConfig() DBConfig {
	return DBConfig{
		Host:         GetEnv("DB_HOST", "127.0.0.1"),
		Port:         GetEnv("DB_PORT", "3306"),
		Username:     GetEnv("DB_USERNAME", "root"),
		Password:     GetEnv("DB_PASSWORD", "1234"),
		Database:     GetEnv("DB_NAME", "custodial_wallet"),
		Dialect:      GetEnv("DB_DIALECT", "mysql"),
		MigrationDir: GetEnv("MIGRATION_DIR", "./migrations/"),
	}
}

func GetRedisConfig() Redis {
	return Redis{
		Host:     GetEnv("REDIS_HOST", "127.0.0.1"),
		Port:     GetEnv("REDIS_PORT", "6379"),
		Password: GetEnv("REDIS_PASS", "redis"),
	}
}

func GetServiceAuthConfig() ServiceAuthConfig {
	return ServiceAuthConfig{
		WhiteList:   GetEnv("SERVICE_WHITE_LIST_IP", "*"),
		SecretToken: GetEnv("SERVICE_SECRET_TOKEN", "SECRET"),
	}
}

func GetPwsCallbackAuthConfig() PwsCallbackAuthConfig {
	return PwsCallbackAuthConfig{
		WhiteList: GetEnv("PWS_CALLBACK_WHITE_LIST_IP", "*"),
	}
}

func GetPwsConfig() PwsConfig {
	return PwsConfig{
		ServiceId:  GetEnv("PWS_SERVICE_ID", "IRC-TEST"),
		Url:        GetEnv("PWS_URL", "https://main.pws.plzdev.ir"),
		Username:   GetEnv("PWS_USERNAME", "wallet-developers"),
		Password:   GetEnv("PWS_PASSWORD", "F]qZvoY}tlNNV}B&QQVoopymfcT,ojFJ"),
		PrivateKey: GetEnv("PWS_PRIVATE_KEY", "private.pem"),
		PublicKey:  GetEnv("PWS_PUBLIC_KEY", "public.pem"),
	}
}

func GetPwsActionConfig() PwsActionConfig {
	return PwsActionConfig{
		WaitForActionInfoMin:    GetEnv("PWS_WAIT_FOR_ACTION_INFO_MIN", "5"),
		GetInfoCycleDurationSec: GetEnv("PWS_GET_INFO_CYCLE_DURATION_SEC", "10"),
		ActionCallback:          GetEnv("PWS_ACTION_CALLBACK", "http://127.0.0.1:8080"),
	}
}

func GetIranicardConfig() IranicardConfig {
	return IranicardConfig{
		Url: GetEnv("IRC_URL", "https://api.iranicard.ir"),
	}
}

func GetCacheConfig() CacheConfig {
	return CacheConfig{
		DefaultCacheExpiration:    GetEnv("DEFAULT_CACHE_EXPIRATION", "600"),
		UserAuthCacheExpiration:   GetEnv("USER_AUTH_CACHE_EXPIRATION", "600"),
		PwsBalanceCacheExpiration: GetEnv("PWS_BALANCE_CACHE_EXPIRATION", "600"),
	}
}
