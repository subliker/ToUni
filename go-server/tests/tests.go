package tests

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

var log = logrus.New()

const CountTests = 1

func Start() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	var file *os.File
	file, err = os.Create(os.Getenv("TEST_FILENAME"))
	if err != nil {
		panic(err)
	}
	log.Out = file
	err = test1()
	if err != nil {
		log.Fatal("ОШИБКА!	", err)
		return
	}
}

func test1() error {
	log.Printf("Test #1 of #%v(add user, get user):\n", CountTests)
	log.Println("Creating random user data...")
	username := RandStringRunes(10)
	password := RandStringRunes(10)
	postData := url.Values{
		"username": {username},
		"password": {password},
	}
	res, err := http.PostForm(fmt.Sprintf("http://%s:%s/api/user", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")), postData)
	if err != nil {
		return err
	}
	fmt.Print(res.Body)
	return nil
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-.")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
