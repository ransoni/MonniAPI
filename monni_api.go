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
)

var (
    id              string
    company_name    string
    creator_name    string
    company_email   string
    company_phone   string
)


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

func index(w http.ResponseWriter, r *http.Request) {
    page := "<!DOCTYPE html><html><body><h1>Monni API</h1><p><a href=\"/getCompany\">getCompany</a></p></body></html>"
    io.WriteString(w, page)
//    if err := io.WriteString(w, "Hello World!"); err != nil {
//        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
//    }
}


// WebServer starts the web server and serves GET & POST requests
func main() {


    http.Handle("/", http.HandlerFunc(index))
    http.Handle("/getCompany", http.HandlerFunc(getCompanyHandler))
//    http.Handle("/addCompany/", http.HandlerFunc(addCompany))

    listen := fmt.Sprintf("%s:%d", "0.0.0.0", 8088)
    logger.Infof("Uchiwa is now listening on %s", listen)
    err := http.ListenAndServe(listen, nil)
    if err != nil {
        fmt.Println(err)
    }

}

