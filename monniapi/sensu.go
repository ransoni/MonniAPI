package monniapi

import (
    "fmt"
    "net/http"
//    "crypto/x509/pkix"
//    "crypto/cipher"
    "crypto/tls"
    "time"
    "io/ioutil"
    "encoding/json"
)


type Tenants struct {
    Tenant []TenantInfo `json:"tenants"`
}

type TenantInfo  struct {
    Name        string  `json:"name"`
    ClientCount int     `json:"clientcount"`
//    Clients     []Client  // For future statistics, debig needs, so on...

}
//type Client struct { // For future features on reporting, to gather more info about installed clients on tenant
//    Name        string
//    Version     string
//    IP          string
//}
type Sensu struct {
    Name     string
    Path     string
    URL      string
    Port     int
//    Timeout  int
    User     string
    Pass     string
    Client   http.Client
}

func GetClients() (*Tenants) {

    tenants := new(Tenants)

    s := new(Sensu)

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    s.Client = http.Client{Timeout: time.Duration(5) * time.Second, Transport: tr}

    for i, _ := range pubConfig.Sensu {
        s.URL = fmt.Sprintf("%s/%s", pubConfig.Sensu[i].URL, "clients")
        s.Name = pubConfig.Sensu[i].Name
        s.User = pubConfig.Sensu[i].User
        s.Pass = pubConfig.Sensu[i].Pass
        s.Port = pubConfig.Sensu[i].Port

        req, err :=  http.NewRequest("GET", s.URL, nil)
        if err != nil {
            fmt.Errorf("URL parsing gone south: %q returned: %v", s.URL, err)
        }

        res, err := s.doHTTP(req)
        if err != nil {
            fmt.Errorf("API call to %q returned: %v", s.URL, err)
        }
        result, err := s.doJSONArray(res)
        if err != nil {
            fmt.Errorf("Could not parse response to JSON: %v", err)
        }

        tenants.Tenant = append(tenants.Tenant, TenantInfo{s.Name, len(result)})

    }
    return tenants
}

func (s *Sensu) doHTTP(req *http.Request) ([]byte, error) {

    if s.User != "" && s.Pass != "" {
        req.SetBasicAuth(s.User, s.Pass)
    }

    res, err := s.Client.Do(req)

    if err != nil {
        return nil, fmt.Errorf("%v", err)
    }

    defer res.Body.Close()

    if res.StatusCode >= 400 {
        return nil, fmt.Errorf("%v", res.Status)
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        return nil, fmt.Errorf("Parsing response body returned: %v", err)
    }
    return body, nil
}

// doJsonArray Unmarshall JSON expecting an array
func (s *Sensu) doJSONArray(body []byte) ([]interface{}, error) {
    var results []interface{}
    if err := json.Unmarshal(body, &results); err != nil {
        return nil, fmt.Errorf("Parsing JSON-encoded response body: %v", err)
    }
    return results, nil
}
