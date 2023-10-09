// Package appctx
package appctx

import (
	"fmt"
	"sync"
	"time"

	"technical_test_go/technical_test_go/internal/consts"
	"technical_test_go/technical_test_go/pkg/file"
	"technical_test_go/technical_test_go/pkg/logger"
)

var (
	once sync.Once
	_cfg *Config
)

// NewConfig initialize config object
func NewConfig() *Config {
	fpath := []string{consts.ConfigPath}
	once.Do(func() {
		c, err := readCfg("app.yaml", fpath...)
		if err != nil {
			logger.Fatal(err)
		}

		_cfg = c
	})

	return _cfg
}

// Config object contract
//
//go:generate easytags $GOFILE yaml,json
type Config struct {
	App        *Common    `yaml:"app" json:"app"`
	Logger     Logging    `yaml:"logger" json:"logger"`
	WriteDB    *Database  `yaml:"db_write" json:"db_write"`
	ReadDB     *Database  `yaml:"db_read" json:"read_db"`
	APM        APM        `yaml:"apm" json:"apm"`
	Cloudinary Cloudinary `yaml:"cloudinary" json:"cloudinary"`
}

// Common general config object contract
type Common struct {
	AppName      string        `yaml:"name" json:"name"`
	ApiKey       string        `yaml:"key" json:"api_key"`
	Debug        bool          `yaml:"debug" json:"debug"`
	Maintenance  bool          `yaml:"maintenance" json:"maintenance"`
	Timezone     string        `yaml:"timezone" json:"timezone"`
	Env          string        `yaml:"env" json:"env"`
	Port         int           `yaml:"port" json:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout_second" json:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" json:"write_timeout"`
	DefaultLang  string        `yaml:"default_lang" json:"default_lang"`
}

// Database configuration structure
type Database struct {
	Name         string        `yaml:"name" json:"name"`
	User         string        `yaml:"user" json:"user"`
	Pass         string        `yaml:"pass" json:"pass"`
	Host         string        `yaml:"host" json:"host"`
	Port         int           `yaml:"port" json:"port"`
	MaxOpen      int           `yaml:"max_open" json:"max_open"`
	MaxIdle      int           `yaml:"max_idle" json:"max_idle"`
	DialTimeout  time.Duration `yaml:"dial_timeout" json:"dial_timeout"`
	MaxLifeTime  time.Duration `yaml:"life_time" json:"max_life_time"`
	ReadTimeout  time.Duration `yaml:"read_timeout" json:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" json:"write_timeout"`
	Charset      string        `yaml:"charset" json:"charset"`
	Driver       string        `yaml:"driver" json:"driver"`
	Timezone     string        `yaml:"timezone" json:"timezone"`
}

// SQS config contract for aws sqs
type SQS struct {
	QueueName      string `yaml:"queue_name" json:"queue_name"`
	QueueURL       string `yaml:"queue_url" json:"queue_url"`
	MaxMessage     int    `yaml:"max_message" json:"max_message"`
	WaitTimeSecond int    `yaml:"wait_time_second" json:"wait_time_second"`
}

// readCfg reads the configuration from file
// args:
//
//	fname: filename
//	ps: full path of possible configuration files
//
// returns:
//
//	*config.Configuration: configuration ptr object
//	error: error operation
func readCfg(fname string, ps ...string) (*Config, error) {
	var cfg *Config
	var errs []error

	for _, p := range ps {
		f := fmt.Sprint(p, fname)

		err := file.ReadFromYAML(f, &cfg)
		if err != nil {
			errs = append(errs, fmt.Errorf("file %s error %s", f, err.Error()))
			continue
		}
		break
	}

	if cfg == nil {
		return nil, fmt.Errorf("file config parse error %v", errs)
	}

	return cfg, nil
}

// Client is a config contract for third party  service provider
type Client struct {
	URL       string        `yaml:"url" json:"url"`
	ApiKey    string        `yaml:"api_key" json:"api_key"`
	ApiSecret string        `yaml:"api_secret" json:"api_secret"`
	Version   string        `yaml:"version" json:"version"`
	Timeout   time.Duration `yaml:"timeout" json:"timeout"`
	VendorID  int           `yaml:"vendor_id" json:"vendor_id"`
}

// Logging config
type Logging struct {
	Name  string `yaml:"name" json:"name"`
	Level string `yaml:"level" json:"level"`
}

// SALS secure connection config
type SASL struct {
	// Whether or not to use SASL authentication when connecting to the broker
	// (defaults to false).
	Enable bool `yaml:"enable" json:"enable"`
	// SASLMechanism is the name of the enabled SASL mechanism.
	// Possible values: OAUTHBEARER, PLAIN (defaults to PLAIN).
	Mechanism string `yaml:"mechanism" json:"mechanism"`
	// Version is the SASL Protocol Version to use
	// Kafka > 1.x should use V1, except on Azure EventHub which use V0
	Version int16 `yaml:"version" json:"version"`
	// Whether or not to send the Kafka SASL handshake first if enabled
	// (defaults to true). You should only set this to false if you're using
	// a non-Kafka SASL proxy.
	Handshake bool `yaml:"handshake" json:"handshake"`
	// User is the authentication identity (authcid) to present for
	// SASL/PLAIN or SASL/SCRAM authentication
	User string `yaml:"user" json:"user"`
	// Password for SASL/PLAIN authentication
	Password string `yaml:"password" json:"password"`
}

// TLS config
type TLS struct {
	Enable     bool   `yaml:"enable" json:"enable"`
	CaFile     string `yaml:"ca_file" json:"ca_file"`
	KeyFile    string `yaml:"key_file" json:"key_file"`
	CertFile   string `yaml:"cert_file" json:"cert_file"`
	SkipVerify bool   `yaml:"skip_verify" json:"skip_verify"`
}

// APM config
type APM struct {
	Address string `yaml:"address" json:"address"`
	Enable  bool   `yaml:"enable" json:"enable"`
	Name    string `yaml:"name" json:"name"`
}

type Cloudinary struct {
	CloudName string `yaml:"cloud_name" json:"cloud_name"`
	ApiKey    string `yaml:"api_key" json:"api_key"`
	ApiSecret string `yaml:"api_secret" json:"api_secret"`
}
