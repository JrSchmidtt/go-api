package controllers

import "net/http"

//Create User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

//Find all users
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find all users"))
}

//Find one user
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find user"))
}

//Update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

//Delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
