package main

import (
	"database/sql"
	"flag"

	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
)
import (
	log "github.com/cihub/seelog"
)

var cfg *goconfig.ConfigFile = nil
var err error = nil
var db *sql.DB = nil

func dbQuery() {
	log.Info("application start.")
	applicationInit()

	db = openDb()
	insert(db)
	query(db)
	applicationDestory()
	log.Info("application end.")
}

func query(ldb *sql.DB) {
	rows, err := ldb.Query("select id, username from user where id = ?", 1)
	if err != nil {
		log.Criticalf("query error")
	}

	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Criticalf("scan error")
		}
		log.Info(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Criticalf("err")
	}
}

func insert(ldb *sql.DB) int {
	stmt, err := ldb.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Criticalf("prepare err")
		return 0
	}
	stmt.Exec("guotie1", "guotie")
	stmt.Exec("testuse1r", "123123")
	return 2
}

func openDb() *sql.DB {
	host, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "host")
	if host == "" {
		log.Criticalf("host is empty, please check")
	}
	user, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "user")
	if user == "" {
		log.Criticalf("use is empty, please check.")
	}
	password, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "password")
	if password == "" {
		log.Criticalf("password is empty, please check.")
	}
	port, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "port")
	if port == "" {
		log.Criticalf("port is empyt, please  check.")
	}
	database, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, "database")
	if database == "" {
		log.Criticalf("database is empty, please check.")

	}

	ldb, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8")
	if err != nil {
		log.Criticalf("Open database error %s\n", err)
	}

	err = ldb.Ping()
	if err != nil {
		log.Criticalf("open err")
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
		log.Criticalf("读取配置文件失败[config.ini]")
		return false
	}
	return true
}

func applicationDestory() bool {
	defer db.Close()
	return true
}
