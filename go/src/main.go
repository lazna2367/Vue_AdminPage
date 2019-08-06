package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Indexs float32
	Id     string
	Pw     string
	Token  string
	Time   string
}

type Login struct {
	Id string
	Pw string
}

// 세션 정보
type LoginCheck struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type ListData struct {
	Id     float32
	Wallet struct {
		Id             float32
		Device_address string
		Wallet_address string
		Created_at     string
		Updated_at     string
	}
	Base_unit      string
	Base_amount    float32
	Target_unit    string
	Target_amount  float32
	Rate_by_base   float32
	Payout_account struct {
		Transaction   string
		IncomAddress  string
		OutcomAddress string
	}
	Confirm     string
	Complete    string
	Description string
	Created_at  string
	Updated_at  string
}

type UpdateData struct {
	Id       int32
	Confirm  string
	Complete string
}

type MainListCountObj struct {
	Buys struct {
		Confirm        int
		Acknowledgment int
		Complete       int
	}
	Sells struct {
		Confirm        int
		Acknowledgment int
		Complete       int
	}
}

type MainWeeklyCount struct {
	Today struct {
		Buy  int
		Sell int
	}
	DayAgo struct {
		Buy  int
		Sell int
	}
	TwoDayAgo struct {
		Buy  int
		Sell int
	}
	ThreeDayAgo struct {
		Buy  int
		Sell int
	}
	FourDayAgo struct {
		Buy  int
		Sell int
	}
	FiveDayAgo struct {
		Buy  int
		Sell int
	}
	SixDayAgo struct {
		Buy  int
		Sell int
	}
}

type MainWeeklyAmount struct {
	Today struct {
		Buy  float32
		Sell float32
	}
	DayAgo struct {
		Buy  float32
		Sell float32
	}
	TwoDayAgo struct {
		Buy  float32
		Sell float32
	}
	ThreeDayAgo struct {
		Buy  float32
		Sell float32
	}
	FourDayAgo struct {
		Buy  float32
		Sell float32
	}
	FiveDayAgo struct {
		Buy  float32
		Sell float32
	}
	SixDayAgo struct {
		Buy  float32
		Sell float32
	}
}

func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	decoder := json.NewDecoder(r.Body)
	var loginData Login
	err := decoder.Decode(&loginData)
	if err != nil {
		fmt.Println(err)
	}

	// DB 대조
	var user User
	flag := 1 // 0일 경우 ID 오류
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from user where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(&loginData.Id).Scan(&user.Indexs, &user.Id, &user.Pw, &user.Token, &user.Time)
	if err != nil {
		// 없는 아이디
		flag = 0
		fmt.Fprint(w, "No ID")
		log.Println(err)
	}
	//fmt.Println(data)

	//fmt.Println(loginData.Id)
	//fmt.Println(loginData.Pw)

	if flag != 0 {
		if loginData.Pw == user.Pw {
			// 로그인 성공
			Init()
			// sessionid := RandStringRunes(32)

			type LoginSession struct {
				Id    string
				Token string
			}
			var loginSession LoginSession
			loginSession.Id = loginData.Id
			loginSession.Token = RandStringRunes(32)

			// JSON 인코딩
			jsonBytes, err := json.Marshal(loginSession)
			if err != nil {
				fmt.Println(err)
			}
			jsonString := string(jsonBytes)
			fmt.Fprintf(w, jsonString)

			// 토큰을 세션DB에 등록
			session_insert(loginData.Id, loginSession.Token)

			log.Println("-> Login Success : ", loginSession.Id, " / ", loginSession.Token)
		} else {
			// 비밀번호 불일치
			fmt.Fprint(w, "Miss PW")
		}
	}
}

// 난수 생성
func Init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	// fmt.Println("-> RandStringRunes Function Start")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	// fmt.Println("-> RandStringRunes Function End\n")

	return string(b)
}

// 토큰을 세션DB에 등록
func session_insert(loginData string, sessionid string) {

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	// fmt.Println(currentTime)
	// fmt.Println("-> Session_insert Function Start")

	// 현재 시간
	// now := time.Now()
	// secs := now.Unix()
	// currentTime := time.Unix(secs, 0)

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// db 트랜젝션 시작
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	// db ptmt SQL문 삽입
	stmt, err := tx.Prepare("update user set token = ?, time = ? where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionid, currentTime, loginData)
	if err != nil {
		fmt.Println(err)
	}

	tx.Commit() // 커밋

	// fmt.Println("-> Session_insert Function End\n")
}

func loginCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// currentTime := time.Now().Format("2006-01-02 15:04:05")
	// fmt.Println(currentTime)
	// fmt.Println("-> IDSessionCheck Function Start")

	var userData LoginCheck
	var dbData LoginCheck

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userData)
	if err != nil {
		fmt.Println(err)
	}

	// currentTime := time.Now().Format("2006-01-02 15:04:05")
	// fmt.Println(currentTime)
	// log.Println("-> ", userData.Id, " (ID Session Check Start)")

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// 사용자 정보와 DB 정보 대조
	ptmt, err := db.Prepare("select id, token from user where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer ptmt.Close()

	rows, err := ptmt.Query(userData.Id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&dbData.Id, &dbData.Token)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(userData)
	// fmt.Println(dbData)

	if userData.Token == dbData.Token {
		// 계정 승인
		fmt.Fprintf(w, "SUCCESS")
		log.Println("-> ", userData.Id, "Login Check is Administrator User")
	} else { // 사용자와 DB의 토큰 불일치, 해킹당함
		fmt.Fprintf(w, "WARING")
		log.Println("-> ", userData.Id, "Login Check is Fail")
	}

	// fmt.Println("-> IDSessionCheck Function End\n")
}

func depositList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	res, err := http.Get("http://api.dev.cint-corp.com:1337/buys")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var List []ListData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&List)
	if err != nil {
		log.Println(err)
	}

	var List_new []ListData
	for i := range List {
		if (List[i].Confirm != "confirmed") && (List[i].Complete != "completed") {
			List_new = append(List_new, List[i])
		} else {
		}
	}

	jsonBytes, err := json.Marshal(List_new)
	if err != nil {
		log.Println(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func depositUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var updateData UpdateData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateData)
	if err != nil {
		log.Println(err)
	}

	jsonBytes, err := json.Marshal(updateData)
	if err != nil {
		log.Println(err)
	}
	jsonBuffer := bytes.NewBuffer(jsonBytes)

	baseURI := "http://api.dev.cint-corp.com:1337/buys/"
	fixURL := baseURI + fmt.Sprint(updateData.Id)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fixURL, jsonBuffer)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode == 200 {
		fmt.Fprintf(w, "ok")
	} else {
		fmt.Fprintf(w, "fail")
	}

	log.Println(fixURL)
	log.Println(updateData)
	log.Println(res.StatusCode)
}

func depositCheckList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	res, err := http.Get("http://api.dev.cint-corp.com:1337/buys")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var List []ListData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&List)
	if err != nil {
		log.Println(err)
	}

	var List_new []ListData
	for i := range List {
		if (List[i].Confirm == "confirmed") && (List[i].Complete != "completed") {
			List_new = append(List_new, List[i])
		} else {
		}
	}

	jsonBytes, err := json.Marshal(List_new)
	if err != nil {
		log.Println(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func depositCheckUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func depositConfirmList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	res, err := http.Get("http://api.dev.cint-corp.com:1337/buys")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var List []ListData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&List)
	if err != nil {
		log.Println(err)
	}

	var List_new []ListData
	for i := range List {
		if (List[i].Confirm == "confirmed") && (List[i].Complete == "completed") {
			List_new = append(List_new, List[i])
		} else {
		}
	}

	jsonBytes, err := json.Marshal(List_new)
	if err != nil {
		log.Println(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func withdrawList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	res, err := http.Get("http://api.dev.cint-corp.com:1337/sells")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var List []ListData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&List)
	if err != nil {
		log.Println(err)
	}

	var List_new []ListData
	for i := range List {
		if (List[i].Confirm != "confirmed") && (List[i].Complete != "completed") {
			List_new = append(List_new, List[i])
		} else {
		}
	}

	jsonBytes, err := json.Marshal(List_new)
	if err != nil {
		log.Println(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func withdrawUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var updateData UpdateData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateData)
	if err != nil {
		log.Println(err)
	}

	jsonBytes, err := json.Marshal(updateData)
	if err != nil {
		log.Println(err)
	}
	jsonBuffer := bytes.NewBuffer(jsonBytes)

	baseURI := "http://api.dev.cint-corp.com:1337/sells/"
	fixURL := baseURI + fmt.Sprint(updateData.Id)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fixURL, jsonBuffer)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode == 200 {
		fmt.Fprintf(w, "ok")
	} else {
		fmt.Fprintf(w, "fail")
	}

	log.Println(fixURL)
	log.Println(updateData)
	log.Println(res.StatusCode)
}

func withdrawCheckList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	res, err := http.Get("http://api.dev.cint-corp.com:1337/sells")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var List []ListData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&List)
	if err != nil {
		log.Println(err)
	}

	var List_new []ListData
	for i := range List {
		if (List[i].Confirm == "confirmed") && (List[i].Complete != "completed") {
			List_new = append(List_new, List[i])
		} else {
		}
	}

	jsonBytes, err := json.Marshal(List_new)
	if err != nil {
		log.Println(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func withdrawCheckUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func withdrawConfirmList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	res, err := http.Get("http://api.dev.cint-corp.com:1337/sells")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var List []ListData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&List)
	if err != nil {
		log.Println(err)
	}

	var List_new []ListData
	for i := range List {
		if (List[i].Confirm == "confirmed") && (List[i].Complete == "completed") {
			List_new = append(List_new, List[i])
		} else {
		}
	}

	jsonBytes, err := json.Marshal(List_new)
	if err != nil {
		log.Println(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func mainListCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// buys 시작
	res, err := http.Get("http://api.dev.cint-corp.com:1337/buys")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var buysList []ListData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&buysList)
	if err != nil {
		log.Println(err)
	}

	var buysList_Confirm []ListData
	for i := range buysList {
		if (buysList[i].Confirm != "confirmed") && (buysList[i].Complete != "completed") {
			buysList_Confirm = append(buysList_Confirm, buysList[i])
		}
	}
	var buysList_Acknowledgment []ListData
	for i := range buysList {
		if (buysList[i].Confirm == "confirmed") && (buysList[i].Complete != "completed") {
			buysList_Acknowledgment = append(buysList_Acknowledgment, buysList[i])
		}
	}
	var buysList_Complete []ListData
	for i := range buysList {
		if (buysList[i].Confirm == "confirmed") && (buysList[i].Complete == "completed") {
			buysList_Complete = append(buysList_Complete, buysList[i])
		}
	}
	// buys 끝

	// sells 시작
	res2, err := http.Get("http://api.dev.cint-corp.com:1337/sells")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var sellsList []ListData
	decoder2 := json.NewDecoder(res2.Body)
	err = decoder2.Decode(&sellsList)
	if err != nil {
		log.Println(err)
	}

	var sellsList_Confirm []ListData
	for i := range sellsList {
		if (sellsList[i].Confirm != "confirmed") && (sellsList[i].Complete != "completed") {
			sellsList_Confirm = append(sellsList_Confirm, sellsList[i])
		}
	}
	var sellsList_Acknowledgment []ListData
	for i := range sellsList {
		if (sellsList[i].Confirm == "confirmed") && (sellsList[i].Complete != "completed") {
			sellsList_Acknowledgment = append(sellsList_Acknowledgment, sellsList[i])
		}
	}
	var sellsList_Complete []ListData
	for i := range sellsList {
		if (sellsList[i].Confirm == "confirmed") && (sellsList[i].Complete == "completed") {
			sellsList_Complete = append(sellsList_Complete, sellsList[i])
		}
	}
	// sells 끝

	var mainListCountObj MainListCountObj
	mainListCountObj.Buys.Confirm = len(buysList_Confirm)
	mainListCountObj.Buys.Acknowledgment = len(buysList_Acknowledgment)
	mainListCountObj.Buys.Complete = len(buysList_Complete)
	mainListCountObj.Sells.Confirm = len(sellsList_Confirm)
	mainListCountObj.Sells.Acknowledgment = len(sellsList_Acknowledgment)
	mainListCountObj.Sells.Complete = len(sellsList_Complete)

	jsonBytes, err := json.Marshal(mainListCountObj)
	if err != nil {
		log.Println(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func mainWeeklyCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	timeNow := time.Now()
	var data MainWeeklyCount

	// 입금
	res, err := http.Get("http://api.dev.cint-corp.com:1337/buys")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var buysList []ListData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&buysList)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(buysList); i++ {
		timeParse, e := time.Parse(time.RFC3339, buysList[i].Created_at)
		if err != nil {
			log.Println(e)
		}

		// 오늘
		if timeParse.Year() == timeNow.AddDate(0, 0, 0).Year() && timeParse.Month() == timeNow.AddDate(0, 0, 0).Month() && timeParse.Day() == timeNow.AddDate(0, 0, 0).Day() {
			data.Today.Buy = data.Today.Buy + 1
		}
		// 오늘 끝
		// 1일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -1).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -1).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -1).Day() {
			data.DayAgo.Buy = data.DayAgo.Buy + 1
		}
		// 1일 전 끝
		// 2일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -2).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -2).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -2).Day() {
			data.TwoDayAgo.Buy = data.TwoDayAgo.Buy + 1
		}
		// 2일 전 끝
		// 3일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -3).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -3).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -3).Day() {
			data.ThreeDayAgo.Buy = data.ThreeDayAgo.Buy + 1
		}
		// 3일 전 끝
		// 4일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -4).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -4).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -4).Day() {
			data.FourDayAgo.Buy = data.FourDayAgo.Buy + 1
		}
		// 4일 전 끝
		// 5일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -5).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -5).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -5).Day() {
			data.FiveDayAgo.Buy = data.FiveDayAgo.Buy + 1
		}
		// 5일 전 끝
		// 6일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -6).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -6).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -6).Day() {
			data.SixDayAgo.Buy = data.SixDayAgo.Buy + 1
		}
		// 6일 전 끝
	}
	// 입금 끝

	// 출금
	res2, err := http.Get("http://api.dev.cint-corp.com:1337/sells")
	if err != nil {
		log.Println(err)
	}
	defer res2.Body.Close()

	var sellsList []ListData
	decoder2 := json.NewDecoder(res2.Body)
	err = decoder2.Decode(&sellsList)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(sellsList); i++ {
		timeParse, e := time.Parse(time.RFC3339, sellsList[i].Created_at)
		if err != nil {
			log.Println(e)
		}

		// 오늘
		if timeParse.Year() == timeNow.AddDate(0, 0, 0).Year() && timeParse.Month() == timeNow.AddDate(0, 0, 0).Month() && timeParse.Day() == timeNow.AddDate(0, 0, 0).Day() {
			data.Today.Sell = data.Today.Sell + 1
		}
		// 오늘 끝
		// 1일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -1).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -1).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -1).Day() {
			data.DayAgo.Sell = data.DayAgo.Sell + 1
		}
		// 1일 전 끝
		// 2일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -2).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -2).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -2).Day() {
			data.TwoDayAgo.Sell = data.TwoDayAgo.Sell + 1
		}
		// 2일 전 끝
		// 3일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -3).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -3).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -3).Day() {
			data.ThreeDayAgo.Sell = data.ThreeDayAgo.Sell + 1
		}
		// 3일 전 끝
		// 4일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -4).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -4).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -4).Day() {
			data.FourDayAgo.Sell = data.FourDayAgo.Sell + 1
		}
		// 4일 전 끝
		// 5일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -5).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -5).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -5).Day() {
			data.FiveDayAgo.Sell = data.FiveDayAgo.Sell + 1
		}
		// 5일 전 끝
		// 6일 전
		if timeParse.Year() == timeNow.AddDate(0, 0, -6).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -6).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -6).Day() {
			data.SixDayAgo.Sell = data.SixDayAgo.Sell + 1
		}
		// 6일 전 끝
	}
	// 출금 끝

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func mainWeeklyAmount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	timeNow := time.Now()
	var data MainWeeklyCount

	// 입금
	res, err := http.Get("http://api.dev.cint-corp.com:1337/buys")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	var buysList []ListData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&buysList)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(buysList); i++ {
		timeParse, e := time.Parse(time.RFC3339, buysList[i].Created_at)
		if err != nil {
			log.Println(e)
		}

		if buysList[i].Target_unit == "MO" {
			// 오늘
			if timeParse.Year() == timeNow.AddDate(0, 0, 0).Year() && timeParse.Month() == timeNow.AddDate(0, 0, 0).Month() && timeParse.Day() == timeNow.AddDate(0, 0, 0).Day() {
				data.Today.Buy = data.Today.Buy + int(buysList[i].Target_amount)
			}
			// 오늘 끝
			// 1일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -1).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -1).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -1).Day() {
				data.DayAgo.Buy = data.DayAgo.Buy + int(buysList[i].Target_amount)
			}
			// 1일 전 끝
			// 2일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -2).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -2).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -2).Day() {
				data.TwoDayAgo.Buy = data.TwoDayAgo.Buy + int(buysList[i].Target_amount)
			}
			// 2일 전 끝
			// 3일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -3).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -3).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -3).Day() {
				data.ThreeDayAgo.Buy = data.ThreeDayAgo.Buy + int(buysList[i].Target_amount)
			}
			// 3일 전 끝
			// 4일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -4).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -4).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -4).Day() {
				data.FourDayAgo.Buy = data.FourDayAgo.Buy + int(buysList[i].Target_amount)
			}
			// 4일 전 끝
			// 5일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -5).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -5).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -5).Day() {
				data.FiveDayAgo.Buy = data.FiveDayAgo.Buy + int(buysList[i].Target_amount)
			}
			// 5일 전 끝
			// 6일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -6).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -6).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -6).Day() {
				data.SixDayAgo.Buy = data.SixDayAgo.Buy + int(buysList[i].Target_amount)
			}
			// 6일 전 끝
		}

	}
	// 입금 끝

	// 출금
	res2, err := http.Get("http://api.dev.cint-corp.com:1337/sells")
	if err != nil {
		log.Println(err)
	}
	defer res2.Body.Close()

	var sellsList []ListData
	decoder2 := json.NewDecoder(res2.Body)
	err = decoder2.Decode(&sellsList)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(sellsList); i++ {
		timeParse, e := time.Parse(time.RFC3339, sellsList[i].Created_at)
		if err != nil {
			log.Println(e)
		}

		if sellsList[i].Base_unit == "MO" {
			// 오늘
			if timeParse.Year() == timeNow.AddDate(0, 0, 0).Year() && timeParse.Month() == timeNow.AddDate(0, 0, 0).Month() && timeParse.Day() == timeNow.AddDate(0, 0, 0).Day() {
				data.Today.Sell = data.Today.Sell + int(sellsList[i].Base_amount)
			}
			// 오늘 끝
			// 1일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -1).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -1).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -1).Day() {
				data.DayAgo.Sell = data.DayAgo.Sell + int(sellsList[i].Base_amount)
			}
			// 1일 전 끝
			// 2일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -2).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -2).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -2).Day() {
				data.TwoDayAgo.Sell = data.TwoDayAgo.Sell + int(sellsList[i].Base_amount)
			}
			// 2일 전 끝
			// 3일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -3).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -3).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -3).Day() {
				data.ThreeDayAgo.Sell = data.ThreeDayAgo.Sell + int(sellsList[i].Base_amount)
			}
			// 3일 전 끝
			// 4일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -4).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -4).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -4).Day() {
				data.FourDayAgo.Sell = data.FourDayAgo.Sell + int(sellsList[i].Base_amount)
			}
			// 4일 전 끝
			// 5일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -5).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -5).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -5).Day() {
				data.FiveDayAgo.Sell = data.FiveDayAgo.Sell + int(sellsList[i].Base_amount)
			}
			// 5일 전 끝
			// 6일 전
			if timeParse.Year() == timeNow.AddDate(0, 0, -6).Year() && timeParse.Month() == timeNow.AddDate(0, 0, -6).Month() && timeParse.Day() == timeNow.AddDate(0, 0, -6).Day() {
				data.SixDayAgo.Sell = data.SixDayAgo.Sell + int(sellsList[i].Base_amount)
			}
			// 6일 전 끝
		}
	}
	// 출금 끝

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	jsonString := string(jsonBytes)
	fmt.Fprintf(w, jsonString)
}

func main() {
	router := httprouter.New()
	router.POST("/login", login)
	router.POST("/loginCheck", loginCheck)
	router.POST("/depositList", depositList)
	router.POST("/depositUpdate", depositUpdate)
	router.POST("/depositCheckList", depositCheckList)
	router.POST("/depositCheckUpdate", depositCheckUpdate)
	router.POST("/depositConfirmList", depositConfirmList)
	router.POST("/withdrawList", withdrawList)
	router.POST("/withdrawUpdate", withdrawUpdate)
	router.POST("/withdrawCheckList", withdrawCheckList)
	router.POST("/withdrawCheckUpdate", withdrawCheckUpdate)
	router.POST("/withdrawConfirmList", withdrawConfirmList)
	router.POST("/mainListCount", mainListCount)
	router.POST("/mainWeeklyCount", mainWeeklyCount)
	router.POST("/mainWeeklyAmount", mainWeeklyAmount)

	port := ":8091"
	fmt.Println("\n\n")
	log.Println("-> Server Running.....")
	log.Println("-> Try connecting / http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
