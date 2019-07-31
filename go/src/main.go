package main

import (
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

			fmt.Println("-> Login Success : ", loginData.Id)
		} else {
			// 비밀번호 불일치
			fmt.Fprint(w, "Miss PW")
		}
	}

	fmt.Println(loginData)
}

// 난수 생성
func Init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	fmt.Println("-> RandStringRunes Function Start")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	fmt.Println("-> RandStringRunes Function End\n")

	return string(b)
}

// 토큰을 세션DB에 등록
func session_insert(loginData string, sessionid string) {

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(currentTime)
	fmt.Println("-> Session_insert Function Start")

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

	fmt.Println("-> Session_insert Function End\n")
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
		if (List[i].Confirm != "confirmed") && (List[i].Complete != "confirmed") {
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
		if (List[i].Confirm == "confirmed") && (List[i].Complete != "confirmed") {
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
}

func withdrawList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func withdrawUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func withdrawCheckList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func withdrawCheckUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func withdrawConfirmList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func main() {
	router := httprouter.New()
	router.POST("/login", login)
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

	log.Println("\n-> Server Running.....\n")
	log.Fatal(http.ListenAndServe(":8091", router))
}
