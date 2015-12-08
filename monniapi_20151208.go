package main

import (
    "encoding/json"
    "fmt"
    "net/http"
//    "net/url"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "io"

    "github.com/palourde/logger"
//    "monniapi"
    "github.com/ransoni/monniapi/monniapi"
    "flag"
//    "github.com/bencaron/gosensu"
)

var (
    id              string
    company_name    string
    creator_name    string
    company_email   string
    company_phone   string
    pubConfig *monniapi.Config
    debug           bool = false
)

//type pubConfig struct {
//    Sensu   []sen
//}


//func deleteStashHandler(w http.ResponseWriter, r *http.Request) {
//    decoder := json.NewDecoder(r.Body)
//    var data interface{}
//    err := decoder.Decode(&data)
//    if err != nil {
//        http.Error(w, fmt.Sprint("Could not decode body"), http.StatusInternalServerError)
//    }
//
//    err = DeleteStash(data)
//    if err != nil {
//        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
//    }
//}
//
//func getAggregateHandler(w http.ResponseWriter, r *http.Request) {
//    u, _ := url.Parse(r.URL.String())
//    c := u.Query().Get("check")
//    d := u.Query().Get("dc")
//    if c == "" || d == "" {
//        http.Error(w, fmt.Sprint("Parameters 'check' and 'dc' are required"), 500)
//    }
//
//    a, err := GetAggregate(c, d)
//    if err != nil {
//        http.Error(w, fmt.Sprint(err), 500)
//    } else {
//        encoder := json.NewEncoder(w)
//        if err := encoder.Encode(a); err != nil {
//            http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), 500)
//        }
//    }
//}
//

func index(w http.ResponseWriter, r *http.Request) {
    page := "<!DOCTYPE html><html><body><h1>Monni API</h1><p><a href=\"/getCompany\">getCompany</a></p><p><a href=\"/getClients\">getClients</a></p></body></html>"
    io.WriteString(w, page)
    //    if err := io.WriteString(w, "Hello World!"); err != nil {
    //        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
    //    }
}

func getCompanyHandler(w http.ResponseWriter, r *http.Request) {
    encoder := json.NewEncoder(w)

    fmt.Println("Get:", r.URL.Query())

    fmt.Println("==== getCompany HANDLER ====\n\n")

    companyOut := make(map[string]string)

    company := make(map[string]map[string]string)
    companyInfo := make(map[string]string)


//    BEGIN OF SQL SEARCH
    var query string

    if r.URL.Query()["company"] != nil {
        query = "select id, company_name, creator_name, company_email, company_phone from subscriptions where company_name='" + r.URL.Query()["company"][0] + "'"
    } else {
        query = "select id, company_name, creator_name, company_email, company_phone from subscriptions"
        fmt.Println("QUERY:", query)
    }


    db, err := sql.Open("mysql", "monnidb:kantamonni@tcp(192.168.15.120:3306)/monni")
    if err != nil {
        fmt.Println("Database connection error.\n", err)
    }
    defer db.Close()

    rows, err := db.Query(query)
    //err = db.QueryRow("select id, ip, name from sensuclientvapp").Scan(&str)

    if err != nil && err != sql.ErrNoRows {
        fmt.Println("Error on query.\n", err) // add http.Error
    } else {
        fmt.Println("Query succeeded.")
    }


    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&id, &company_name, &creator_name, &company_email, &company_phone)
        if err != nil {
            fmt.Println("Fetching data from row went haywire.\n", err)
        }
//        fmt.Printf("COMPANY NAME: %v\nCREATOR NAME: %v\nCOMPANY_EMAIL: %v\nCOMPANY_PHONE: %v\n", company_name, creator_name, company_email, company_phone)

        companyOut["id"] = id
        companyOut["company_name"] = company_name
        companyOut["creator_name"] = creator_name
        companyOut["company_email"] = company_email
        companyOut["company_phone"] = company_phone
        companyInfo["company_name"] = company_name
        companyInfo["creator_name"] = creator_name
        companyInfo["company_email"] = company_email
        companyInfo["company_phone"] = company_phone

        company[id] = companyInfo


//        fmt.Println("Company:", companyOut)
    }

//    END OF SQL SEARCH

    //	fmt.Println("*** FINAL COMPANYOUT ***")
    fmt.Println("TenantOut:", company)

    //	fmt.Println(outResults)

    //fmt.Println(outResults)
    if err := encoder.Encode(company); err != nil {
        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
    }
}

func getClientsHandler(w http.ResponseWriter, r *http.Request) {
    encoder := json.NewEncoder(w)
    clients := monniapi.GetClients()

    logger.Infof("Got results for Tenants: %s", clients)

    if err := encoder.Encode(clients); err != nil {
        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
    }
}

func addCompany(w http.ResponseWriter, r *http.Request) {
    fmt.Println("==== addCompany HANDLER ====\n\n")
    // Test valid input
    if r.URL.Query()["companyname"] == nil || r.URL.Query()["name"] == nil || r.URL.Query()["email"] == nil || r.URL.Query()["phone"] == nil{
          http.Error(w, fmt.Sprintf("Input missing"), http.StatusBadRequest)
          return
    }

    // Read companyName, creator name, email and phone
    encoder := json.NewEncoder(w)

    fmt.Println("Get:", r.URL.Query())
    fmt.Println("Company name:", r.URL.Query()["companyname"][0])
    fmt.Println("Name:", r.URL.Query()["name"][0])
    fmt.Println("Email:", r.URL.Query()["email"][0])
    fmt.Println("Phone:", r.URL.Query()["phone"][0])

    subscription := make(map[string]string)
    subscription["status"] = "false"
    subscription["message"] = "Unknown error"
    subscription["companyname"] = r.URL.Query()["companyname"][0]
    subscription["name"] = r.URL.Query()["name"][0]
    subscription["email"] = r.URL.Query()["email"][0]
    subscription["phone"] = r.URL.Query()["phone"][0]


    //Query company from db
    // TODO: Replace with ping and and move connection to main?
    db, err := sql.Open("mysql", "monnidb:kantamonni@tcp(192.168.15.120:3306)/monni")
    if err != nil {
        fmt.Println("Database connection error.\n", err)
    }
    rows, err := db.Query("SELECT id, company_name FROM subscriptions WHERE company_name=?", r.URL.Query()["companyname"][0])
    if err != nil {
        fmt.Println("Database query error.\n", err)
    }
    var sizeOfResult int
    sizeOfResult = 0
    for rows.Next() {
      sizeOfResult ++
    }

    // Check if the company is already there
    if sizeOfResult != 0 {
      // Company is there
      fmt.Println("Company already there.\n", err)
      subscription["message"] = "Duplicate company found"
      if err := encoder.Encode(subscription); err != nil {
           http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
      }
      return
    }

    // Add the company
    _, err = db.Exec("INSERT INTO subscriptions SET order_time=NOW(), company_name=?, creator_name=?, company_email=?, company_phone=?", r.URL.Query()["companyname"][0], r.URL.Query()["name"][0], r.URL.Query()["email"][0], r.URL.Query()["phone"][0])

    if err != nil {
        fmt.Println("Database insert error.\n", err)
        subscription["message"] = "Company insert failed"
        if err := encoder.Encode(subscription); err != nil {
             http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
        }
        return
    }

    subscription["status"] = "true"
    subscription["message"] = "Company successfully added"

    defer db.Close()

    if err := encoder.Encode(subscription); err != nil {
         http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
    }

}

func deleteCompany(w http.ResponseWriter, r *http.Request) {
    fmt.Println("==== deleteCompany HANDLER ====\n\n")
    // Test valid input
    if r.URL.Query()["companyname"] == nil {
          http.Error(w, fmt.Sprintf("Input missing"), http.StatusBadRequest)
          return
    }

    // Read companyName
    encoder := json.NewEncoder(w)

    fmt.Println("Get:", r.URL.Query())
    fmt.Println("Company name:", r.URL.Query()["companyname"][0])

    subscription := make(map[string]string)
    subscription["status"] = "false"
    subscription["message"] = "Unknown error"
    subscription["companyname"] = r.URL.Query()["companyname"][0]

    //Query company from db
    // TODO: Replace with ping and and move connection to main?
    db, err := sql.Open("mysql", "monnidb:kantamonni@tcp(192.168.15.120:3306)/monni")
    if err != nil {
        fmt.Println("Database connection error.\n", err)
    }
    rows, err := db.Query("SELECT id, company_name FROM subscriptions WHERE company_name=?", r.URL.Query()["companyname"][0])
    if err != nil {
        fmt.Println("Database query error.\n", err)
    }
    var sizeOfResult int
    sizeOfResult = 0
    for rows.Next() {
      sizeOfResult ++
    }

    // Check if the company is already there
    if sizeOfResult != 1 {
      // Company is there
      fmt.Println("Company not found.\n", err)
      subscription["message"] = "Requested company not found"
      if err := encoder.Encode(subscription); err != nil {
           http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
      }
      return
    }

    // Delete the company
    _, err = db.Exec("DELETE FROM subscriptions WHERE company_name=? LIMIT 1", r.URL.Query()["companyname"][0])

    if err != nil {
        fmt.Println("Database delete error.\n", err)
        subscription["message"] = "Company delete failed"
        if err := encoder.Encode(subscription); err != nil {
             http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
        }
        return
    }

    subscription["status"] = "true"
    subscription["message"] = "Company successfully deleted"

    defer db.Close()

    if err := encoder.Encode(subscription); err != nil {
         http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
    }

}



// WebServer starts the web server and serves GET & POST requests
func main() {
    var confErr error
    confFile := flag.String("c", "./config.json", "relative or full path to configuration file")
    pubConfig, confErr = monniapi.LoadConfig(*confFile)
    if confErr != nil {
        logger.Fatal(confErr)
    }

    if debug {
        for i := 0; i < len(pubConfig.Sensu); i++ {
            fmt.Printf("Name: %s\nURL: %s\nUser: %s\nPass: %s\n", pubConfig.Sensu[i].Name, pubConfig.Sensu[i].URL, pubConfig.Sensu[i].User, pubConfig.Sensu[i].Pass)
            //        pubConfig.Sensu[i]
        }
    }

//    sensu.Sensu.GetClients()


//    monniapi.New(config)


    // THIS PART BELOW SHOULD BE MOVED TO WEBSERVER FUNCTION, OR SO...
    http.Handle("/", http.HandlerFunc(index))
    http.Handle("/getCompany", http.HandlerFunc(getCompanyHandler))
    http.Handle("/getClients", http.HandlerFunc(getClientsHandler))
    http.Handle("/addCompany/", http.HandlerFunc(addCompany))
    http.Handle("/deleteCompany/", http.HandlerFunc(deleteCompany))

    listen := fmt.Sprintf("%s:%d", "0.0.0.0", 8088)
    logger.Infof("Monni API is now listening on %s", listen)
    err := http.ListenAndServe(listen, nil)
    if err != nil {
        fmt.Println(err)
    }

}
