# Common

The common library provide a generic main class to easily start a new go program.
It automatically set up :
- Config(read from ./config.yml by default)
- Log-level
- Application command
- Application flag

## Dependency
-  [client](https://github.com/codegangsta/cli)
-  [viper](https://github.com/spf13/viper)
-  [logrus](https://github.com/Sirupsen/logrus)
-  [router](https://github.com/julienschmidt/httprouter)


## How To Use
To create a new go program simply use this main:
```go
 // main client installation
func main() {
	common.SetAppName("applicationName")
	common.SetAppUsage("applicationUsage")
	common.SetAppVersion("applicationVersion")
	common.SetAppDescription("applicationDescription")
	common.Main(&common.Handlers{Handle: Handlers})
	common.Run()
}

// Handlers Returns httprouter handlers
func Handlers() *httprouter.Router {
// Set up your router
	r := httprouter.New()
	r.POST("/path", handlerFunction)
	return r
}
```

## Help
this main will automatically generate help text :
```bash
NAME:
   applicationName - applicationUsage

USAGE:
   ./applicationName [global options] command [command options] [arguments...]

VERSION:
   applicationVersion

COMMANDS:
   start	applicationDescription
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config, -c "./"		absolute config path [$CONFIG]
   --help, -h			show help
   --generate-bash-completion
   --version, -v		print the version
  ```

## Config
  The generic main will also set up a configuration reader. By default the program will try to read from ./config.yml.
  A --config, -c flag is available so you can define another path to the configuration file.

## Test

  The library also provide a test package where you can find a simple test tool to test API endpoints.

### How To Use

```go
  // TestMain set up configuration for the test
func TestMain(m *testing.M) {
	common.LoadConfig("./", "configTest")
	test.Init(&common.Handlers{Handle: Handlers})
	os.Exit(m.Run())
}


func TestingEndpoint(t *testing.T, directory, file string) {
	w := test.Events(t, "POST", "route", directory, file)
}
```

The test.Events function read a file and send its informations to the endpoint of your choice. In order to make sure the test are working you need to define 2 type of file for every test :
data.json :
```bash
{
  "commentEvent": null,
  "eventType": "create",
  "pullRequestEvent": null,
  "pushEvent": null,
  "repository": {
    "branch": "test2",
    "fullName": "evonck/webhook-tes",
    "name": "webhook-tes"
  },
  "repositoryType": "github"
}
```

and a data.header :
```bash
X-Nexway-type: repository
```
