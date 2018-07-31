package main

import (
	"net/http"
	"os"

	"github.com/siddontang/go/log"
	"./lib"
	"crypto/sha256"
	"encoding/base64"
	_ "github.com/mattn/go-sqlite3"
)

const DBNAME = "./shorten.db"
const TBLNAME = "shorten"
const DOMAIN = "bus.com"

var db *lib.DB

func main() {

	// start db
	db = &lib.DB{
		DBName:DBNAME,
		TableName:TBLNAME,
	}
	db.Connect()
	db.CreateDb()

	// start server and listen for get
	startServer()

	// process and send back response with html and new url

}

func startServer() {
	log.Info("Starting server")
	http.HandleFunc("/process", returnShortenURL)
	http.HandleFunc("/", showHomePage)

	// we can wrap this in a log since it should never
	// leave the loop unless theres an error or its about to stop
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func readFromFile(filename string) (file *os.File) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}
	return file
}

func showHomePage(w http.ResponseWriter, r *http.Request) {
	//form := readFromFile("form.html")
	//logrus.Info("Handling homepage")
	w.Write([]byte("this is a sentence"))
}

func returnShortenURL(w http.ResponseWriter, r *http.Request) {
	log.Info("Processing...")
	originalURL := r.URL.Query().Get("s")
	encoded := makeShort(originalURL)
	encoded = encoded[:5]
	fullURL := "http://" + encoded + "." + DOMAIN

	sql :=  "select shorten_url from " + TBLNAME + " where original_url = '" + originalURL + "'"
	cached := db.QueryDb(sql)
	if cached == "-1"{
		sql = "insert into " + TBLNAME + "(original_url, shorten_url) values(1, '"+originalURL+"'), (2, '"+fullURL+"')"
		db.Save(sql)
	}

	w.Write([]byte(fullURL))
}

func makeShort(originalURL string) string {
	hasher := sha256.New()
	hasher.Reset()
	hasher.Write([]byte(originalURL))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}


