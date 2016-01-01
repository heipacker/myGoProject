package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
import (
	log "github.com/cihub/seelog"
)

type Email struct {
	Title   string
	Content string
	Names   []string
}

func testHttp() {
	log.Info("do http invoke")
	var alertEmail = Email{
		Title:   "test",
		Content: "test",
		Names:   []string{"pan.fu"},
	}
	emailBytes, _ := json.Marshal(alertEmail)
	body := bytes.NewBuffer([]byte(emailBytes))
	log.Info(body)
	response, err := http.Post("http://l-schedule1.qss.dev.cn0.qunar.com:8080/alarm/sendAll", "application/x-www-form-urlencoded", body)
	if err != nil {
		log.Criticalf("post error, please check.%s", err)
		return
	}
	defer response.Body.Close()
	ret, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Criticalf("get ret error, %s", err)
		return
	}
	fmt.Println(string(ret))
}
