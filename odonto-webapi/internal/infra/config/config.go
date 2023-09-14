package config

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

//=======================================================
// Settings Errors
//=======================================================

var (
	//ErrDBConfigNotExists databases configuration not exists
	ErrDBConfigNotExists = errors.New("databases configuration not exists")
	//ErrDBConfigNotFound database configuration not found
	ErrDBConfigNotFound = errors.New("database configuration not found")
	//ErrYamlFileNotFound settings YAML file not found or empty
	ErrYamlFileNotFound = errors.New("settings YAML file not found or empty")
)

//=======================================================
// Setting structure
//=======================================================

type AWSS3 struct {
Bucket string `yaml:"bucket"`
Region string `yaml:"region"` 
ApiKey string `yaml:"apikey"`
Secret string `yaml:"secret"`
}

//Databases a map of database configurations
type Databases map[string]DBConfig

//DBConfig the structure of configuration of a database
type DBConfig struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pswd string `yaml:"pswd"`
	DBNm string `yaml:"dbnm"`
	Port string `yaml:"port"`
}

type Connections struct {
	MaxOpenConns    int  `yaml:"maxOpenConns"`
	MaxIdleConns    int  `yaml:"maxIdleConns"`
	ConnMaxIdleTime int  `yaml:"connMaxIdleTime"`
	ConnMaxLifetime int  `yaml:"connMaxLifetime"`
}

type Ecosystem struct {
	Host   string `yaml:"host"`   
	GID    string `yaml:"gid"`    
	Tenant string `yaml:"tenant"` 
	Name   string `yaml:"name"`   
	Token  string `yaml:"token"`  
}

type PixConfig struct {
	AppKey         string `yaml:"appKey"`
	AuthUrl        string `yaml:"authUrl"`
	QrCodeUrl      string `yaml:"qrCodeUrl"`
	CheckChargeUrl string `yaml:"chargeUrl"`
	Expiration     int    `yaml:"expiration"`
}

type CieloConfig struct {
	URL         string `yaml:"url"`
	MerchantId  string `yaml:"mId"`
	MerchantKey string `yaml:"mKey"`
}

//
type SlackConfig struct {
	Token   string `yaml:"token"`
	Channel string `yaml:"channel"`
}

//
type FirebaseConfig struct {
	ID   string `yaml:"projectID"`
	Path string `yaml:"configFile"`
}

//
type Integrations struct {
	Slack    SlackConfig    `yaml:"slack"`
	Firebase FirebaseConfig `yaml:"firebase"`
}

//
type MailConfig struct {
	From     string `yaml:"from"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"user"`
	Password string `yaml:"pass"`
}

//Settings settings of applications running
type Settings struct {
	Mode         string         `yaml:"mode"`
	JWTSecret    string         `yaml:"jwtSecret"`
	AWSS3        AWSS3          `yaml:"awss3"`
	Ecosystem    Ecosystem      `yaml:"ecosystem"`
	Pix          PixConfig      `yaml:"pix"`
	Integrations Integrations   `yaml:"integrations"`
	Databases    []Databases    `yaml:"databases"`
	Mail         MailConfig     `yaml:"mail"`
	Connections  Connections  `yaml:"connections"`
}

//Write marshall settings to file yaml
func (s *Settings) Write(filename string) error {
	bs, err := yaml.Marshal(s)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bs, os.FileMode(0755))
}

//GetMode get enviroment mode (dev or prod)
func (s *Settings) GetMode() string {
	return s.Mode
}

//GetDatabases get enviroment mode (dev or prod)
func (s *Settings) GetDatabases(name string) (DBConfig, error) {
	if s.Databases == nil {
		return DBConfig{}, ErrDBConfigNotExists
	}
	for _, dbc := range s.Databases {
		if dbc[name].Host != "" {
			return dbc[name], nil
		}
	}
	return DBConfig{}, ErrDBConfigNotFound
}

//=======================================================
// Statics
//=======================================================

func checkAndAdaptTypeError(err error) error {
	terr, ok := err.(*yaml.TypeError)
	if !ok {
		return err
	}
	build := strings.Builder{}
	for i, estr := range terr.Errors {
		if i > 0 {
			build.WriteString(", ")
		}
		build.WriteString(estr)
	}
	return errors.New(build.String())
}

//Write marshall settings to file yaml. An alias for settings.Write(...)
func Write(s *Settings, filename string) error {
	return s.Write(filename)
}

//Load unmarshall settings from file yaml
func Load(filename string) (*Settings, error) {
	// read file
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, ErrYamlFileNotFound
	}
	// unmarshall
	inst := &Settings{
		Mode: "dev",
	}
	err = yaml.Unmarshal(bs, inst)
	if err != nil {
		return nil, checkAndAdaptTypeError(err)
	}
	//success
	return inst, nil
}

//=======================================================
// Settings as Singleton
//=======================================================

var settingsInst *Settings = nil
var settingsInstOnce = sync.Once{}

// Init initializes settings as singleton instance from yaml file
func Init(filename string) {
	settingsInstOnce.Do(func() {
		s, e := Load(filename)
		if e != nil {
			panic(e)
		}
		settingsInst = s
	})
}

// GetSettings get settings as singleton instance
func GetSettings() *Settings {
	return settingsInst
}

