// Package config contains application configuration structures.
package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config 전체 설정 구조체
type Config struct {
	App     *AppConfig     `yaml:"app"     mapstructure:"app"`
	DB      *DBConfig      `yaml:"db"      mapstructure:"db"`
	Server  *ServerConfig  `yaml:"server"  mapstructure:"server"`
	Storage *StorageConfig `yaml:"storage" mapstructure:"storage"`
	Auth    *AuthConfig    `yaml:"auth"    mapstructure:"auth"`
	SMTP    *SMTPConfig    `yaml:"smtp"    mapstructure:"smtp"`
}

// AppConfig : 애플리케이션 관련 설정
type AppConfig struct {
	Env string `yaml:"env" mapstructure:"env"`
}

// DBConfig : 데이터베이스 관련 설정
type DBConfig struct {
	Driver     string `yaml:"driver" mapstructure:"driver"`
	DBHost     string `yaml:"db_host" mapstructure:"db_host"`
	DBPort     string `yaml:"db_port" mapstructure:"db_port"`
	DBUser     string `yaml:"db_user" mapstructure:"db_user"`
	DBPassword string `yaml:"db_password" mapstructure:"db_password"`
	DBName     string `yaml:"db_name" mapstructure:"db_name"`
}

// ServerConfig : 서버 관련 설정
type ServerConfig struct {
	Port int    `yaml:"port" mapstructure:"port"`
	Host string `yaml:"host" mapstructure:"host"`
}

type StorageConfig struct {
	UploadPath string `yaml:"upload_path" mapstructure:"upload_path"`
	ImagePath string `yaml:"image_path" mapstructure:"image_path"`
}

type AuthConfig struct {
	// Access Token 서명 키
    AccessSecret string `yaml:"access_secret" mapstructure:"access_secret"`
    
    // Refresh Token 서명 키
    RefreshSecret string `yaml:"refresh_secret" mapstructure:"refresh_secret"`
    
    // Access Token 유효 기간 (예: 15m)
    AccessDuration time.Duration `yaml:"access_duration" mapstructure:"access_duration"`
    
    // Refresh Token 유효 기간 (예: 7d)
    RefreshDuration time.Duration `yaml:"refresh_duration" mapstructure:"refresh_duration"`

	// Email Verification 코드 유효 기간 (예: 5m)
	EmailVerificationExpiry time.Duration `yaml:"email_verification_expiry" mapstructure:"email_verification_expiry"`
}

// SMTP 서버 관련 설정
type SMTPConfig struct {
	HostServer  string `yaml:"host_server" mapstructure:"host_server"`
	Port        string `yaml:"port" mapstructure:"port"`
	From        string `yaml:"from" mapstructure:"from"`
	AppPassword string `yaml:"app_password" mapstructure:"app_password"`
}

// LoadConfig 설정 파일을 로드하고 환경변수를 치환하여 Config 구조체로 반환
func LoadConfig(path string) (*Config, error) {
	_ = godotenv.Load()

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	// 1. 기본 환경변수 매핑 (DB_SOURCE 등)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config file read error: %w", err)
	}

	// 2. YAML 내의 ${VAR} 형태를 환경변수로 치환
	content, _ := os.ReadFile(v.ConfigFileUsed())
	replacedContent := os.ExpandEnv(string(content))

	// 3. 치환된 내용을 다시 Viper에 로드
	if err := v.ReadConfig(strings.NewReader(replacedContent)); err != nil {
		return nil, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "mapstructure" // viper와의 호환성
	}); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// GetDSN 데이터베이스 연결 문자열(DSN)을 반환
func (d *DBConfig) GetDSN() string {
	switch d.Driver {
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			d.DBHost, d.DBPort, d.DBUser, d.DBPassword, d.DBName)
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
			d.DBUser, d.DBPassword, d.DBHost, d.DBPort, d.DBName)
	default:
		return ""
	}
}