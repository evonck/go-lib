package common

import (
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

var (
	configPath string
	v          *viper.Viper
	handler    *Handlers
	// App the application object
	App            *cli.App
	appName        string
	appUsage       string
	appVersion     string
	appDescription string
	initFunc       string
)

// SetAppName set the app name
func SetAppName(name string) {
	appName = name
}

// SetAppUsage set the app usage
func SetAppUsage(usage string) {
	appUsage = usage
}

// SetAppVersion set the ap pversion
func SetAppVersion(version string) {
	appVersion = version
}

// SetAppDescription set the app description
func SetAppDescription(description string) {
	appDescription = description
}

// Main client installation
func Main(handlerFunc *Handlers) {
	handler = handlerFunc
	App = cli.NewApp()
	App.EnableBashCompletion = true
	App.Name = appName
	App.Usage = appUsage
	App.Version = appVersion

	// global flags
	//Log Level and Config Path
	App.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Value:  "./config.yml",
			EnvVar: "CONFIG",
			Usage:  "config Path",
		},
	}
	App.Commands = []cli.Command{{
		Name:   "start",
		Usage:  appDescription,
		Action: server,
		Flags:  App.Flags,
	}}
}

// Run the application
func Run() {
	App.Before = func(ctx *cli.Context) error {
		configPath = ctx.String("config")
		return nil
	}
	err := App.Run(os.Args)
	if err != nil {
		log.Warn("an eror occur while running the program")
	}
}

// Server launch the API
func server(ctx *cli.Context) {
	configInit()
	initLogLevel()
	if handler.Handle != nil {
		http.Handle("/", &MyServer{handler.Handle()})
		log.Fatal(http.ListenAndServe(Config().GetString("addr"), nil))
		return
	}
	if handler.HttpHandle != nil {
		http.Handle("/", &MyServer{handler.HttpHandle()})
		log.Fatal(http.ListenAndServe(Config().GetString("addr"), nil))
		return
	}
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Origin", origin)
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		rw.Header().Set("Access-Control-Allow-Methods", "GET,PATCH,PUT,POST,DELETE,OPTIONS,HEAD")
		rw.Header().Set("Access-Control-Max-Age", "1800")
		rw.Header().Set("Allow", "HEAD,POST,GET,OPTIONS")
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}

// configInit initialiaze the loading of the config
func configInit() {
	path, name := getConfigPathAndName()
	LoadConfig(path, name)
	Config().SetDefault("log-level", "debug")
	Config().BindEnv("log-level")
	Config().SetDefault("addr", "localhost:8081")
	Config().BindEnv("addr")
	Config().SetDefault("global_seed", "")
	Config().BindEnv("global_seed")
	if handler.Init != nil {
		handler.Init()
	}
}

func getConfigPathAndName() (string, string) {
	path := strings.Split(configPath, "/")
	configFilePath := "./"
	configFileName := "config"
	name := ""
	if len(path) > 0 {
		name = path[len(path)-1]
		names := strings.Split(name, ".")
		configFileName = names[0]
	}
	if len(path) > 1 {
		ConfigFileNameLength := len(name)
		configFilePath = configPath[:len(configPath)-ConfigFileNameLength]
	}
	return configFilePath, configFileName
}

// getLogLevel Retrieves log level
func getLogLevel() log.Level {
	lvl, err := log.ParseLevel(Config().GetString("log-level"))
	if err != nil {
		log.WithFields(log.Fields{
			"passed":  lvl,
			"default": "fatal",
		}).Warn("Log level is not valid, fallback to default level")
		return log.FatalLevel
	}
	return lvl
}

// LoadConfig load configuration file
func LoadConfig(pathFile, fileName string) {
	if v == nil {
		v = viper.New()
		v.SetConfigName(fileName)
		v.AddConfigPath(pathFile)
		err := v.ReadInConfig() // Find and read the config file
		if err != nil {         // Handle errors reading the config file
			log.Fatalf("Fatal error config file: %s \n", err)
		}
	}
}

// Config returns the current viper instance
func Config() *viper.Viper {
	if v == nil {
		path, name := getConfigPathAndName()
		LoadConfig(path, name)
	}
	return v
}

// initLogLevel Static method to init log level
func initLogLevel() {
	log.SetLevel(getLogLevel())
}
