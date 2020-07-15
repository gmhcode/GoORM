package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

//User struct
type User struct {
	gorm.Model
	Name  string
	Email string
}

//InitialMigration for database
func InitialMigration() {
	// db, err = gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	panic("Failed to connect to database")
	// }
	// // defer db.Close()

	// db.AutoMigrate(&User{})
}

//AllUsers Returns all the users
func AllUsers(w http.ResponseWriter, r *http.Request) {
	// db, err = gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("Could not connect to the database")
	// }
	// // defer db.Close()
	// //Create an empty array of users
	// var users []User
	// //Finds all users
	// db.Find(&users)
	// json.NewEncoder(w).Encode(users)
	fmt.Fprintf(w, "All Users Endpoint Hit")
}

//NewUser Creates a new user
func NewUser(w http.ResponseWriter, r *http.Request) {
	// db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	// defer db.Close()
	w.Header().Set("Content-Type", "application/json")

	var user User
	//decodes the user from the body and turns it into data
	_ = json.NewDecoder(r.Body).Decode(&user)
	// user.ID = uint(rand.Uint32())

	//reads the body data
	body, err := ioutil.ReadAll(r.Body)
	//prints the body data
	fmt.Println(string(body))

	//converts user into json
	str, err := json.Marshal(&user)
	//prints the user json
	fmt.Println(string(str))

	name := user.Name
	email := user.Email
	db.Create(&User{Name: name, Email: email})

	if err != nil {
		panic("Could not connect to the database")
	}
	fmt.Fprintf(w, "New User Endpoint Hit")
}

// //OldNewUserFunc getting user info from parameters
// func OldNewUserFunc(w http.ResponseWriter, r *http.Request) {
// 	if err != nil {
// 		panic("Could not connect to the database")
// 	}
// 	vars := mux.Vars(r)
// 	name := vars["name"]
// 	email := vars["email"]
// 	db.Create(&User{Name: name, Email: email})
// }

//DeleteUser Deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	// defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User

	//The second parameter "name" fills in the "?" spot
	db.Where("name = ?", name).Find(&user)

	db.Delete(&user)
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

//UpdateUser Updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		panic("Could not connect to the database")
	}
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
