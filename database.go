package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
)

var cfg *goconfig.ConfigFile = nil
var err error = nil
var db *sql.DB = nil

func dbQuery() {
	log.Println("application start.")
	applicationInit()

	db = openDb()
	insert(db)
	query(db)
	// httpTest()
	applicationDestory()
	log.Println("application end.")
}

type Email struct {
	Title   string
	Content string
	Names   string
}

func httpTest() {
	var alertEmail = Email{
		Title:   "test",
		Content: "test",
		Names:   "pan.fu",
	}
	str, _ := json.Marshal(alertEmail)
	log.Println(str)
	body := bytes.NewBuffer([]byte(`"title":"test","content":"test","names":"pan.fu"`))

	response, err := http.Post("http://l-schedule1.qss.dev.cn0.qunar.com:8080/alarm/sendAll", "application/x-www-form-urlencoded", body)
	if err != nil {
		log.Fatalln("post error, please check.%s", err)
	}
	defer response.Body.Close()
	ret, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("get ret error, %s", err)
	}
	fmt.Println(string(ret))
}

func query(ldb *sql.DB) {
	rows, err := ldb.Query("select id, username from user where id = ?", 1)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func insert(ldb *sql.DB) int {
	stmt, err := ldb.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return 0
	}
	stmt.Exec("guotie1", "guotie")
	stmt.Exec("testuse1r", "123123")
	return 2
}

func openDb() *sql.DB {
	host, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "host")
	if host == "" {
		log.Fatalln("host is empty, please check")
	}
	user, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "user")
	if user == "" {
		log.Fatalln("use is empty, please check.")
	}
	password, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "password")
	if password == "" {
		log.Fatalln("password is empty, please check.")
	}
	port, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "port")
	if port == "" {
		log.Fatalln("port is empyt, please  check.")
	}
	database, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "database")
	if database == "" {
		log.Fatalln("database is empty, please check.")

	}

	ldb, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8")
	if err != nil {
		log.Println("Open database error %s\n", err)
	}

	err = ldb.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return ldb
}

var config_file_name *string = flag.String("config", "config/config.ini", "config file name")

func applicationInit() bool {
	flag.Parse()
	if config_file_name == nil {
		tmp_config := "config/config.ini"
		config_file_name = &tmp_config
	}
	cfg, err = goconfig.LoadConfigFile(*config_file_name)
	if err != nil {
		log.Fatalln("读取配置文件失败[config.ini]")
		return false
	}
	return true
}

func applicationDestory() bool {
	defer db.Close()
	return true
}
