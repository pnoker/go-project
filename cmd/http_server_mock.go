/*
 * Copyright Pnoker. All Rights Reserved.
 */

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	http.HandleFunc("/", say)
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		log.Fatal("err", err)
	}
}

func say(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error.read request body error:", err.Error())
		return
	}
	log.Println("request:", string(request))

	if _, err := w.Write([]byte("ok")); err != nil {
		log.Println("error.response error:", err.Error())
	}
}
