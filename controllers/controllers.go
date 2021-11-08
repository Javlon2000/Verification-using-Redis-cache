package controllers

import (
	"log"
	"time"
	"strings"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"crypto/rand"
	"encoding/json"

	"github.com/Javlon2000/Verification-using-Redis-database/models"
	"github.com/Javlon2000/Verification-using-Redis-database/utils"

	"github.com/go-redis/redis"
)

type SignUPInput struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string 
}

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "",
	DB: 0,
})

func SignUP(w http.ResponseWriter, r *http.Request) {

	var input SignUPInput

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("Cannot raed the request body: %v", err)
	}
	
	json.Unmarshal(body, &input)

	randomNumber, _ := rand.Prime(rand.Reader, 18)

	message := []byte(randomNumber.String())

	Sender := "goguruh01@gmail.com"
	Password := "Qwertyu!op"

	receivers := []string {
		input.Email,
	}

	host := "smtp.gmail.com"
	port := "587"

	auth := smtp.PlainAuth("", Sender, Password, host)

	err = smtp.SendMail(host + ":" + port, auth, Sender, receivers, message)

	if err != nil {

		log.Printf("Cannot send to the email: %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Email was not found!"))
	
	} else{

		w.Write([]byte("Sended! Check your email!"))

	}

	input.Password = randomNumber.String()

	user, err := json.Marshal(SignUPInput {Username: strings.ToLower(input.Username), Email: input.Email, Password: input.Password})

	if err != nil {
		log.Fatalf("Cannot marshalling: %v", err)
	}

	Redis(strings.ToLower(input.Username), user)
}

func Redis(key string, value []uint8) {
	
	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatalf("Cannot connecting to the redis: %v", err)
	}

	log.Println(pong)

	err = client.Set(strings.ToLower(key), value, 1 * time.Hour).Err()

	if err != nil {
		log.Fatalf("Cannot inserting into to the Redis: %v", err)
	}
}

func Verify(w http.ResponseWriter, r *http.Request) {

	db, err := utils.DB()

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Cannot connecting to the database"))
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalf("Cannot read the request body: %v", err)
	}
	
	var input models.Check

	json.Unmarshal(body, &input)

	key := strings.ToLower(input.Username)

	values, err := client.Get(key).Result()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Username was not found!"))
	}

	var check SignUPInput

	var isFound bool

	err = json.Unmarshal([]byte(values), &check)

	if err != nil {
		log.Printf("Cannot get the data from the Redis: %v", err)
	}

	// getUser := models.Check{}



	if input.Password == check.Password{

		user := models.InsertDatabase { Username: check.Username, Email: check.Email, Password: check.Password }

		db.Table("users").Create(&user)
		w.Write([]byte("Verified!"))

		isFound = true
		
	}
	// row := db.Table("users").Select("username").Where("username = ?", check.Username).Find(&getUser)
 
	encoder := json.NewEncoder(w)

	var getUser models.User

	w.Header().Set("Content-Type", "application/json")

	if isFound {
		_ = db.Where("username = ?", check.Username).First(&getUser)

		encoder.Encode(getUser)

	}else{
		w.Write([]byte("Username or Password wrong!"))
	}

	// } else {

	// 	w.WriteHeader(http.StatusForbidden)
	// 	w.Write([]byte("Username or email already exits!"))

	// }

	// } else {

	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte("Password was wrong!"))
	// }

}