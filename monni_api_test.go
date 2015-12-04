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

//type company struct {
//    id  int
//    companyInfo
//}
type companyInfo struct {
    name    string
    creator string
    email   string
    phone   string
}

//func deleteClientHandler(w http.ResponseWriter, r *http.Request) {
//    u, _ := url.Parse(r.URL.String())
//    i := u.Query().Get("id")
//    d := u.Query().Get("dc")
//    if i == "" || d == "" {
//        http.Error(w, fmt.Sprint("Parameters 'id' and 'dc' are required"), http.StatusInternalServerError)
//    }
//
//    err := DeleteClient(i, d)
//    if err != nil {
//        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
//    }
//}
//
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
//func getAggregateByIssuedHandler(w http.ResponseWriter, r *http.Request) {
//    u, _ := url.Parse(r.URL.String())
//    c := u.Query().Get("check")
//    i := u.Query().Get("issued")
//    d := u.Query().Get("dc")
//    if c == "" || i == "" || d == "" {
//        http.Error(w, fmt.Sprint("Parameters 'check', 'issued' and 'dc' are required"), 500)
//    }
//
//    a, err := GetAggregateByIssued(c, i, d)
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
//func getClientHandler(w http.ResponseWriter, r *http.Request) {
//    u, _ := url.Parse(r.URL.String())
//    i := u.Query().Get("id")
//    d := u.Query().Get("dc")
//    if i == "" || d == "" {
//        http.Error(w, fmt.Sprint("Parameters 'id' and 'dc' are required"), http.StatusInternalServerError)
//    }
//
//    c, err := GetClient(i, d)
//    if err != nil {
//        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
//    } else {
//        encoder := json.NewEncoder(w)
//        if err := encoder.Encode(c); err != nil {
//            http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
//        }
//    }
//}
//
//func getConfigHandler(w http.ResponseWriter, r *http.Request) {
//    encoder := json.NewEncoder(w)
//    fmt.Println("getConfigHandler")
//    //	fmt.Println("PublicConfig: %s", PublicConfig)
//    if err := encoder.Encode(PublicConfig); err != nil {
//        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
//    }
//}
//
//func getSensuHandler(w http.ResponseWriter, r *http.Request) {
//    encoder := json.NewEncoder(w)
//
//    // HAETAAN KEKSISTÄ DATAA
//    cookie, _ := r.Cookie("uchiwa_auth")
//    cookieValue := "payload="
//    cookieValue += cookie.Value
//
//    c, _ := url.ParseQuery(cookieValue)
//    fmt.Println("C:", c)
//
//    var data map[string]interface{}
//
//    json.Unmarshal([]byte(c["payload"][0]), &data)
//    // KEKSIN KÄSITTELY LOPPU
//
//    /* VOI KÄYTTÄÄ KEKSIN TESTAUKSIIN
//        for k, v := range data {
//            fmt.Println("Key:", k)
//            fmt.Println("Value:", v)
//        }
//
//        fmt.Println("DATA:", data)
//        fmt.Println("Name:", data["FullName"])
//        fmt.Println("Email:", data["Email"])
//    */
//
//    tempResults := Results.Get()
//    outResults := newResult() //HAETAAN STRUCTI OUTTILLE
//
//
//    clients := tempResults.Clients
//    dcs := tempResults.Dc
//    evnts := tempResults.Events
//    //	aggrs := tempResults.Aggregates
//    //	subs := tempResults.Subscriptions
//
//    // TO OUTRESULT
//    // FILTERÖINTI TEHTÄVÄ VIELÄ Eventeille, Aggregationille
//    // ja Subscriptioneille
//    outResults.Aggregates = tempResults.Aggregates
//    outResults.Subscriptions = tempResults.Subscriptions
//    //	outResults.Events = tempResults.Events
//    //	outResults.Dc = tempResults.Dc
//
//    //	fmt.Println("OUTRESULTS:", outResults)
//    fmt.Printf("CLIENTS: %v\nLENGTH: %v\n", clients, len(clients))
//
//    //	DC filtering for tenant
//    for i, v := range dcs {
//        dc := v
//        fmt.Println("RANGE DC\n", dc)
//
//        if dc["name"] == data["Org"] {
//            fmt.Println("DC IF\n", dc)
//            outResults.Dc = append(outResults.Dc, dcs[i])
//        }
//    }
//
//    //	Client filtering for tenant
//    for i, v := range clients {
//        clnt := v.(map[string]interface {})
//
//        if clnt["dc"] == data["Org"] {
//            //fmt.Println("IF DC:", dc["dc"])
//            outResults.Clients = append(outResults.Clients, clients[i])
//        }
//    }
//
//    //	Event filtering for tenant
//    for i, v := range evnts {
//        evnt := v.(map[string]interface {})
//
//        if evnt["dc"] == data["Org"] {
//            //fmt.Println("IF DC:", dc["dc"])
//            outResults.Events = append(outResults.Events, evnts[i])
//        }
//    }
//
//    fmt.Println("*** FINAL OUTRESULT ***")
//    fmt.Println(outResults)
//
//
//    if err := encoder.Encode(outResults); err != nil {
//        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
//    }
//
//    /*
//        if err := encoder.Encode(Results.Get()); err != nil {
//            http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
//        }
//    */
//}


//func getUserHandler(w http.ResponseWriter, r *http.Request) {
//    var data map[string]interface{}
//
//    encoder := json.NewEncoder(w)
//
//    fmt.Println("==== USER HANDLER ====\n\n")
//
//    // HAETAAN KEKSISTÄ DATAA
//    cookie, _ := r.Cookie("uchiwa_auth")
//    cookieValue := "payload="
//    cookieValue += cookie.Value
//
//    c, _ := url.ParseQuery(cookieValue)
//    //	fmt.Println("C:", c)
//
//    json.Unmarshal([]byte(c["payload"][0]), &data)
//    // KEKSIN KÄSITTELY LOPPU
//
//    fmt.Println("ORG:", data["Org"])
//    fmt.Println("User email:", data["Email"])
//
//    var userOut map[string]string
//    userOut = make(map[string]string)
//    //userOut = getUserInfo(data["Org"].(string), data["Email"].(string))
//    //	conf := PublicConfig
//    userOut, _ = getUserInfo(PublicConfig, data["Org"].(string), data["Email"].(string))
//
//    //	FOR TESTING
//    /*
//        userOut["name"] = "Testi Taina"
//        userOut["tel"] = "+358505057890"
//        userOut["city"] = "Helsinki"
//    */
//
//    fmt.Println("TenantOut:", userOut)
//
//    //fmt.Println(outResults)
//    if err := encoder.Encode(userOut); err != nil {
//        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
//    }
//}

//func postUserHandler(w http.ResponseWriter, r *http.Request) {
//    var data map[string]interface{}
//
//    encoder := json.NewEncoder(w)
//
//    fmt.Println("==== USER POST HANDLER ====\n\n")
//    fmt.Printf("REQUEST: %s", r)
//
//    // HAETAAN KEKSISTÄ DATAA
//    cookie, _ := r.Cookie("uchiwa_auth")
//    cookieValue := "payload="
//    cookieValue += cookie.Value
//
//    c, _ := url.ParseQuery(cookieValue)
//    //	fmt.Println("C:", c)
//
//    json.Unmarshal([]byte(c["payload"][0]), &data)
//    // KEKSIN KÄSITTELY LOPPU
//
//    fmt.Println("ORG:", data["Org"])
//    fmt.Println("User email:", data["Email"])
//
//    var userOut map[string]string
//    userOut = make(map[string]string)
//    //userOut = getUserInfo(data["Org"].(string), data["Email"].(string))
//    //	conf := PublicConfig
//    //	userOut, _ = getUserInfo(PublicConfig, data["Org"].(string), data["Email"].(string))
//
//    //	FOR TESTING
//    userOut["status"] = "User's information updated."
//    userOut["error"] = ""
//
//    //	fmt.Println("TenantOut:", userOut)
//
//    //fmt.Println(outResults)
//    if err := encoder.Encode(userOut); err != nil {
//        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
//    }
//}

//func postPasswdHandler(w http.ResponseWriter, r *http.Request) {
//    var data map[string]interface{}
//
//    encoder := json.NewEncoder(w)
//
//    decoder := json.NewDecoder(r.Body)
//    var req map[string]string
//    err := decoder.Decode(&req)
//    if err != nil {
//        http.Error(w, fmt.Sprint("Could not decode body"), http.StatusInternalServerError)
//    }
//    fmt.Println("REQUEST:", req)
//    fmt.Println("OLDPASSWORD:", req)
//    for key, value := range req {
//        fmt.Printf("Key: %s\nValue: %s", key, value)
//    }
//
//    //	var resp map[string]string
//    //	resp = make(map[string]string)
//    _, err = changePasswd(PublicConfig, req)
//
//    //	fmt.Println("FORM DATA:", r.FormValue("oldPassword"))
//    fmt.Println("FORM DATA:", r.PostForm)
//    fmt.Println("METHOD:", r.Method)
//    fmt.Println("BODY:", r.Body)
//
//    fmt.Println("\n\n\n==== PASSWORD POST HANDLER ====\n")
//    fmt.Printf("REQUEST: %s", r)
//
//    // HAETAAN KEKSISTÄ DATAA
//    cookie, _ := r.Cookie("uchiwa_auth")
//    cookieValue := "payload="
//    cookieValue += cookie.Value
//
//    c, _ := url.ParseQuery(cookieValue)
//    //	fmt.Println("C:", c)
//
//    json.Unmarshal([]byte(c["payload"][0]), &data)
//    // KEKSIN KÄSITTELY LOPPU
//
//    fmt.Println("ORG:", data["Org"])
//    fmt.Println("User email:", data["Email"])
//
//    var userOut map[string]string
//    userOut = make(map[string]string)
//    //userOut = getUserInfo(data["Org"].(string), data["Email"].(string))
//    //	conf := PublicConfig
//    //	userOut, _ = getUserInfo(PublicConfig, data["Org"].(string), data["Email"].(string))
//
//    //	FOR TESTING
//    userOut["status"] = "Password changed."
//    userOut["error"] = err.Error()
//
//    //	fmt.Println("TenantOut:", userOut)
//
//    //fmt.Println(outResults)
//    if err := encoder.Encode(userOut); err != nil {
//        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
//    }
//}

//func healthHandler(w http.ResponseWriter, r *http.Request) {
//    encoder := json.NewEncoder(w)
//    var err error
//    if r.URL.Path[1:] == "health/sensu" {
//        //		fmt.Println("healthHandler")
//        //		fmt.Println("Health.Sensu: %s", Health.Sensu)
//
//        /*		for i := range Health.Sensu {
//                    fmt.Println("Health01:", Health.Sensu[i])
//                    if Health.Sensu[i] == "Client01" {
//                        err = encoder.Encode(Health.Sensu[i])
//                        break
//                    } else {
//                    // NOTHING TO SEE
//                    }
//                }
//                } else if r.URL.Path[1:] == "health/uchiwa" {
//                    err = encoder.Encode(Health.Uchiwa)
//                } else {
//                    fmt.Println("Health02: %s", Health)
//                    err = encoder.Encode(Health)
//                }
//        //	}
//        */
//
//        err = encoder.Encode(Health.Sensu)
//    } else if r.URL.Path[1:] == "health/uchiwa" {
//        err = encoder.Encode(Health.Uchiwa)
//    } else {
//        err = encoder.Encode(Health)
//    }
//
//
//    if err != nil {
//        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
//    }
//}

//func postEventHandler(w http.ResponseWriter, r *http.Request) {
//    decoder := json.NewDecoder(r.Body)
//    var data interface{}
//    err := decoder.Decode(&data)
//    if err != nil {
//        http.Error(w, fmt.Sprint("Could not decode body"), http.StatusInternalServerError)
//    }
//
//    err = ResolveEvent(data)
//
//    if err != nil {
//        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
//    }
//}

//func postStashHandler(w http.ResponseWriter, r *http.Request) {
//    decoder := json.NewDecoder(r.Body)
//    var data interface{}
//    err := decoder.Decode(&data)
//    if err != nil {
//        http.Error(w, fmt.Sprint("Could not decode body"), http.StatusInternalServerError)
//    }
//
//    err = CreateStash(data)
//
//    if err != nil {
//        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
//    }
//}

func getCompanyHandler(w http.ResponseWriter, r *http.Request) {
    encoder := json.NewEncoder(w)

    fmt.Println("Get:", r.URL.Query())

    fmt.Println("==== getCompany HANDLER ====\n\n")

    companyOut := make(map[string]map[string]companyInfo)

//    company := make(map[string]map[string]string)
//    companyInfo := make(map[string]string)
//    tenantOut = getTenantInfo(data["Org"].(string))

//    BEGIN OF SQL SEARCH
    var query string

    if r.URL.Query()["company"] != nil {
        query = "select id, company_name, creator_name, company_email, company_phone from subscriptions where company_name='" + r.URL.Query()["company"][0] + "'"
    } else {
        query = "select id, company_name, creator_name, company_email, company_phone from subscriptions"
        fmt.Println("QUERY:", query)
    }

//    tenant = make(map[string]string)

    db, err := sql.Open("mysql", "monnidb:kantamonni@tcp(192.168.15.120:3306)/monni")
    if err != nil {
        fmt.Println("Database connection error.\n", err)
    }
    defer db.Close()

    rows, err := db.Query(query)
    //err = db.QueryRow("select id, ip, name from sensuclientvapp").Scan(&str)

    if err != nil && err != sql.ErrNoRows {
        fmt.Println("Error on query.\n", err)
    } else {
        fmt.Println("Query succeeded.")
    }

    //fmt.Println("STR:", str)

    defer rows.Close()

    companyOut["id"] = make(map[string]companyInfo)

    for rows.Next() {
        err := rows.Scan(&id, &company_name, &creator_name, &company_email, &company_phone)
        if err != nil {
            fmt.Println("Fetching data from row went haywire.\n", err)
        }
//        fmt.Printf("COMPANY NAME: %v\nCREATOR NAME: %v\nCOMPANY_EMAIL: %v\nCOMPANY_PHONE: %v\n", company_name, creator_name, company_email, company_phone)

        companyOut["id"][id] = companyInfo{name: company_name, creator: creator_name, email: company_email, phone: company_phone}
//        companyOut["company_name"] = company_name
//        companyOut["creator_name"] = creator_name
//        companyOut["company_email"] = company_email
//        companyOut["company_phone"] = company_phone
//        companyInfo["company_name"] = company_name
//        companyInfo["creator_name"] = creator_name
//        companyInfo["company_email"] = company_email
//        companyInfo["company_phone"] = company_phone
//
//        company[id] = companyInfo

//        fmt.Println("Company:", json.Indent(companyOut, "", "   "))


//        fmt.Println("Company:", companyOut)
    }

//    END OF SQL SEARCH

    fmt.Println("*** FINAL COMPANYOUT ***")

    fmt.Println("TenantOut:", string(json.Marshal(companyOut)))

    //	fmt.Println(outResults)

    //fmt.Println(outResults)
    if err := encoder.Encode(companyOut["id"]); err != nil {
        http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
    }

    /*
        if err := encoder.Encode(Results.Get()); err != nil {
            http.Error(w, fmt.Sprintf("Cannot encode response data: %v", err), http.StatusInternalServerError)
        }
    */
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

//    http.Handle("/delete_client",http.HandlerFunc(deleteClientHandler))
//    http.Handle("/delete_stash", auth.Authenticate(http.HandlerFunc(deleteStashHandler)))
//    http.Handle("/get_aggregate", auth.Authenticate(http.HandlerFunc(getAggregateHandler)))
//    http.Handle("/get_aggregate_by_issued", auth.Authenticate(http.HandlerFunc(getAggregateByIssuedHandler)))
//    http.Handle("/get_client", auth.Authenticate(http.HandlerFunc(getClientHandler)))
//    http.Handle("/get_config", auth.Authenticate(http.HandlerFunc(getConfigHandler)))
//    http.Handle("/get_sensu", auth.Authenticate(http.HandlerFunc(getSensuHandler)))
//    http.Handle("/get_tenant", auth.Authenticate(http.HandlerFunc(getTenantHandler)))
//    http.Handle("/get_user", auth.Authenticate(http.HandlerFunc(getUserHandler)))
//    http.Handle("/post_user", auth.Authenticate(http.HandlerFunc(postUserHandler)))
//    http.Handle("/post_passwd", auth.Authenticate(http.HandlerFunc(postPasswdHandler)))
//    http.Handle("/post_event", auth.Authenticate(http.HandlerFunc(postEventHandler)))
//    http.Handle("/post_stash", auth.Authenticate(http.HandlerFunc(postStashHandler)))
    http.Handle("/", http.HandlerFunc(index))
    http.Handle("/getCompany", http.HandlerFunc(getCompanyHandler))
//    http.Handle("/health/", http.HandlerFunc(healthHandler))

    listen := fmt.Sprintf("%s:%d", "0.0.0.0", 8088)
    logger.Infof("Uchiwa is now listening on %s", listen)
    err := http.ListenAndServe(listen, nil)
    if err != nil {
        fmt.Println(err)
    }

}

