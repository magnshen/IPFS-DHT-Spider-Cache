package main

import (
	"./model"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)
type configuration struct {
	USERNAME string
	PASSWORD string
	NETWORK string
	SERVER  string
	PORT    int
	DATABASE string
}

func main() {
	var daysInfo bool
	flag.BoolVar(&daysInfo,"daysInfo",false,"Update DaysInfo,define Update NewsInfo. DaysInfo update once every day. NewInfo update every 2 seconds")
	flag.Parse()
	cpath,err:= filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("config file load error1")
		os.Exit(2)
	}
	file, _ := os.Open(cpath+"/config.cnf")
	decoder := json.NewDecoder(file)
	Setting := configuration{}
	err = decoder.Decode(&Setting)
	file.Close()
	if err != nil {
		fmt.Println("config file load error")
		os.Exit(2)
	}
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",Setting.USERNAME,Setting.PASSWORD,Setting.NETWORK,Setting.SERVER,Setting.PORT,Setting.DATABASE)
	db ,err:= models.NewDbWorker(dsn)
	if err != nil{
		fmt.Printf("Mysql connent error")
		os.Exit(2)
	}
	if daysInfo{
			fmt.Println("Update Days Info")
			db.UpdateDaysInfo()
	}else{
		for {
			fmt.Println("Update News Info")
			db.UpdateNewsInfo()
			time.Sleep(time.Second * 2)
		}
	}
}


