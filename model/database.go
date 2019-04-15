package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
type Hash struct {
	Hash string
	Heat string
	CreateTime string
}

type NewsInfo struct {
	Total_hashs string `json:"total_hashs"`
	Yesterday_hashs string `json:"yesterday_hashs"`
	Today_hashs string `json:"today_hashs"`
	Working_spiders int `json:"spiders"`
	New_hash []Hash `json:"new_hash"`
}
type DaysLinePoint struct {
	Date_time string `json:"date_time"`
	Hashs string  `json:"hashs"`
}
type DaysInfo struct {
	Days_line []DaysLinePoint `json:"days_line"`
	History_heats []Hash  `json:"history_heats"`
	Lastweek_heats []Hash  `json:"lastweek_heats"`
}
type DbWorker struct {
	DB       *sql.DB
}
func (self *DbWorker)getHistoryHeats()(res []Hash){
	rows ,err := self.DB.Query("SELECT * FROM `Hash_List` ORDER BY Hits DESC LIMIT 10")
	defer rows.Close()
	if err != nil {
		fmt.Printf("GetHistoryHeats error: %v\n", err)
		return
	}
	for rows.Next() {
		var id ,objget  int
		var consult bool
		var hash Hash
		rows.Scan(&id, &hash.Hash, &consult,&objget,&hash.Heat,&hash.CreateTime)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		//fmt.Printf("get data, id: %d  Hash: %s  consult:%t Objectcont:%d  Heats:%d CreateTime: %s \n",id,hash.Hash,consult,objget,hash.Heat,hash.CreateTime)
		res = append(res,hash)
	}
	return
}
func (self *DbWorker)getLastWeekHeats()(res []Hash){
	now := time.Now().Unix()
	datetime := time.Unix(now - 3600*24*7, 0).Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("SELECT * FROM `Hash_List` WHERE CreateTime >= '%s' ORDER BY Hits DESC LIMIT 10",datetime)
	rows ,err := self.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		fmt.Printf("GetLastWeekHeats error: %v\n", err)
		return
	}
	for rows.Next() {
		var id ,objget  int
		var consult bool
		var hash Hash
		rows.Scan(&id, &hash.Hash, &consult,&objget,&hash.Heat,&hash.CreateTime)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		//fmt.Printf("get data, id: %d  Hash: %s  consult:%t Objectcont:%d  Heats:%d CreateTime: %s \n",id,hash.Hash,consult,objget,hash.Heat,hash.CreateTime)
		res = append(res,hash)
	}
	return
}

func (self *DbWorker)get14Days()(res []DaysLinePoint){
	t := time.Now()
	for i := 14;i > 0;i--{
		thatDay := time.Date(t.Year(), t.Month(), t.Day()-i, 0, 0, 0, 0, t.Location())
		thatDayEnd := time.Date(t.Year(), t.Month(), t.Day()-i+1, 0, 0, 0, 0, t.Location())
		thatDayStr := thatDay.Format("2006-01-02 15:04:05")
		thatDayEndStr := thatDayEnd.Format("2006-01-02 15:04:05")
		sql := fmt.Sprintf("SELECT COUNT(*) FROM `Hash_List` where CreateTime >= '%s' AND CreateTime < '%s'",thatDayStr,thatDayEndStr)
		//fmt.Println(sql)
		row := self.DB.QueryRow(sql)
		var count string
		row.Scan(&count);
		date_time := thatDay.Format("02/01")
		res = append(res, DaysLinePoint{date_time,count})
	}
	return
}

func (self *DbWorker)getTotalHashCount()(count string){
	sql := fmt.Sprintf("SELECT COUNT(*) FROM `Hash_List`")
	row := self.DB.QueryRow(sql)
	row.Scan(&count);
	return
}

func (self *DbWorker)getYesterdayHashCount()(count string){
	t := time.Now()
	thatDay := time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, t.Location())
	thatDayEnd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	thatDayStr := thatDay.Format("2006-01-02 15:04:05")
	thatDayEndStr := thatDayEnd.Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("SELECT COUNT(*) FROM `Hash_List` where CreateTime >= '%s' AND CreateTime < '%s'",thatDayStr,thatDayEndStr)
	row := self.DB.QueryRow(sql)
	row.Scan(&count);
	return
}

func (self *DbWorker)getTodayHashCount()(count string){
	t := time.Now()
	thatDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	thatDayStr := thatDay.Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("SELECT COUNT(*) FROM `Hash_List` where CreateTime >= '%s'",thatDayStr)
	row := self.DB.QueryRow(sql)
	row.Scan(&count);
	return
}

func (self *DbWorker)getSpiderCount()(count int){
	//count Spider ,with data submit in 5 minute
	now := time.Now().Unix()
	thatDayStr := time.Unix(now - 60*5, 0).Format("2006-01-02 15:04:05")
	sql := fmt.Sprintf("SELECT COUNT(*) FROM `Spider_List` where UpdateTime >= '%s'",thatDayStr)
	row := self.DB.QueryRow(sql)
	row.Scan(&count)
	return
}

func (self *DbWorker)getNewHash()(res []Hash){
	rows ,err := self.DB.Query("SELECT * FROM `Hash_List` ORDER BY ID DESC LIMIT 8")
	defer rows.Close()
	if err != nil {
		fmt.Printf("GetNewHash error: %v\n", err)
		return
	}
	for rows.Next() {
		var id ,objget  int
		var consult bool
		var hash Hash
		var timestr string
		rows.Scan(&id, &hash.Hash, &consult,&objget,&hash.Heat,&timestr)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		timeer ,_:= time.Parse("2006-01-02 15:04:05",timestr)
		hash.CreateTime = timeer.Format("15:04:05")
		//fmt.Printf("get data, id: %d  Hash: %s  consult:%t Objectcont:%d  Heats:%d CreateTime: %s \n",id,hash.Hash,consult,objget,hash.Heat,hash.CreateTime)
		res = append(res,hash)
	}
	return
}


func (self *DbWorker)UpdateNewsInfo(){
	var news = NewsInfo{}
	news.Total_hashs = self.getTotalHashCount()
	news.New_hash = self.getNewHash()
	news.Working_spiders = self.getSpiderCount()
	news.Today_hashs = self.getTodayHashCount()
	news.Yesterday_hashs = self.getYesterdayHashCount()
	data, _ := json.Marshal(news)
	//fmt.Println(string(data))
	//fmt.Println(new)
	sql := fmt.Sprintf("UPDATE Web_Data SET NewsInfo='%s' WHERE ID =1",string(data))
	rows ,_ :=self.DB.Query(sql)
	defer rows.Close()
}

func (self *DbWorker)UpdateDaysInfo(){
	var days = DaysInfo{}
	days.Lastweek_heats = self.getLastWeekHeats()
	days.History_heats = self.getHistoryHeats()
	days.Days_line = self.get14Days()
	data, _ := json.Marshal(days)
	//fmt.Println(string(data))
	sql := fmt.Sprintf("UPDATE Web_Data SET DaysInfo='%s' WHERE ID =1",string(data))
	rows ,_ :=self.DB.Query(sql)
	defer rows.Close()
}
func NewDbWorker(dsn string)(*DbWorker,error){
	dbwk := &DbWorker{}
	DB,err := sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return dbwk ,err
	}
	dbwk.DB=DB
	return dbwk ,err
}