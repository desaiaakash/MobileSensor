package station

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	//"log"
	"net/http"
)
	
	func init() {
    db, err = sql.Open("mysql", "root:root@/test")
    if err != nil {
        fmt.Println(err)
    }

    err = db.Ping()
	if err == nil {
		fmt.Println("DB Successfully connected")
	}

	}

func createUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u Users
	var uid int
	var result ResultSet
	
	

	json.NewDecoder(r.Body).Decode(&u)

	username := u.Username
	password := u.Password

	row := db.QueryRow("SELECT username FROM users WHERE Username = ? ", username)

	usr := new(Users)
	err = row.Scan(&usr.Username)

	if err == nil {
		fmt.Println("User already exists. Kindly login")
		return 
	} else if err == sql.ErrNoRows {
		row := db.QueryRow("SELECT max(userid) FROM users")
		err = row.Scan(&usr.Uid)
		if err == nil {
			uid = (usr.Uid + 1)
		} else if err == sql.ErrNoRows {
			uid = 1
		}
	}

	if username == "" || password == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	
	_, err = db.Exec("INSERT INTO users (userid,username,password) VALUES (?, ?, ?)", uid, username, password)

	if err == nil {
		fmt.Println("User login successfull")
		result.Result = true
		res,_ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", res)
	} else if err == sql.ErrNoRows {
		//fmt.Println("User Not Authenticated")
		//http.NotFound(w, r)
		result.Result = true
		res,_ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", res)
	} else if err != nil {
		//http.Error(w, http.StatusText(500), 500)
		result.Result = false
		res,_ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", res)
	}

}

func validateUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//mux.GET("/check/:username/:pwd", validateUsers)

	username := p.ByName("username")
	password := p.ByName("pwd")
	
	var result ResultSet
	
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if username == "" || password == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	fmt.Println("Testing before validation query", username, password)

	row := db.QueryRow("SELECT username FROM users WHERE username = ? and password = ?", username, password)

	usr := new(Users)
	err = row.Scan(&usr.Username)
	if err == nil {
		fmt.Println("User login successfull")
		result.Result = true;
		res,_ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", res)
	} else if err == sql.ErrNoRows {
		//fmt.Println("User Not Authenticated")
		//http.NotFound(w, r)
		result.Result = false;
		res,_ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", res)
	} else if err != nil {
		//http.Error(w, http.StatusText(500), 500)
		result.Result = false;
		res,_ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", res)
	}

	
}

func Authorization() {

	mux := httprouter.New()

	mux.POST("/create", createUsers)
	mux.GET("/check/:username/:pwd", validateUsers)

	server := http.Server{
		Addr:    "0.0.0.0:3501",
		Handler: mux,
	}
	server.ListenAndServe()

}


