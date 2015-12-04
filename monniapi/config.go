package monniapi

import (
    "encoding/json"
    "fmt"
    "os"
//    "encoding/xml"
//    "crypto/x509/pkix"
//    "sort"
    "github.com/ransoni/logger"
    "math/rand"
)

var (
    debug = false
)

type Config struct {
    Sensu   []SensuConfig
}

type SensuConfig struct {
    Name        string
    Host        string
    Port        int
    Ssl         bool
    Insecure    bool
    URL         string
    User        string
    Path        string
    Pass        string
    Timeout     int
}

var pubConfig *Config

func (c *Config) initSensu() {
    if debug {
        fmt.Println("Function: initSensu")
        fmt.Println("C.Sensu: %s", c.Sensu[0].Name)
        fmt.Println("\n C: %s", c)
    }
    for i, api := range c.Sensu {
        //if c.Sensu[i].Name == "Client01" {
        if debug {
//            fmt.Println("Inside if...")
        }
        prot := "http"
        if api.Name == "" {
            logger.Warningf("Sensu API %s has no name property. Generating random one...", api.URL)
            c.Sensu[i].Name = fmt.Sprintf("sensu-%v", rand.Intn(100))
        }
        if api.Host == "" {
            logger.Fatalf("Sensu API %q Host is missing", api.Name)
        }
        if api.Timeout == 0 {
            c.Sensu[i].Timeout = 10
        } else if api.Timeout >= 1000 { // backward compatibility with < 0.3.0 version
            c.Sensu[i].Timeout = api.Timeout/1000
        }
        if api.Port == 0 {
            c.Sensu[i].Port = 4567
        }
        if api.Ssl {
            prot += "s"
        }
        c.Sensu[i].URL = fmt.Sprintf("%s://%s:%d%s", prot, api.Host, c.Sensu[i].Port, api.Path)
        //}
    }
}

// LoadConfig function loads a specified configuration file and return a Config struct
func LoadConfig(path string) (*Config, error) {
    logger.Infof("Loading configuration file %s", path)
    c := new(Config)
    file, err := os.Open(path)
    if err != nil {
        if len(path) > 1 {
            return nil, fmt.Errorf("Error: could not read config file %s.", path)
        }
    }

    decoder := json.NewDecoder(file)
    err = decoder.Decode(c)
    if err != nil {
        return nil, fmt.Errorf("Error decoding file %s: %s", path, err)
    }

//    c.initGlobal()
//    c.initUser()
    c.initSensu()

    pubConfig = c

    return c, nil
}