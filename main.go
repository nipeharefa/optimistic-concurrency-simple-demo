package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/go-chi/chi/v5"
)

type (
	User struct {
		Name    string `json:"name"`
		ID      int    `json:"id"`
		Version int    `json:"version"`
	}

	Users []*User
)

var users = make(Users, 0)

func main() {
	r := chi.NewRouter()
	r.Get("/users/{id}/update", func(rw http.ResponseWriter, r *http.Request) {

		ids := chi.URLParam(r, "id")

		id, _ := strconv.ParseInt(ids, 10, 64)

		for _, u := range users {

			if u.ID == int(id) {
				user1 := u
				user1.Version++
				json.NewEncoder(rw).Encode(user1)
				return
			}
		}

		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("User not found"))
	})
	r.Get("/users/new", func(rw http.ResponseWriter, r *http.Request) {
		u := &User{
			ID:      len(users) + 1,
			Name:    faker.Name(),
			Version: 1,
		}
		users = append(users, u)
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(u)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(users)
	})

	r.Get("/optimis/read", func(rw http.ResponseWriter, r *http.Request) {
		user1 := users[0]

		version := user1.Version
		time.Sleep(2 * time.Second)

		// check before update
		if version != user1.Version {
			msg := fmt.Sprintf("Error read version %d, actual version: %d", version, user1.Version)
			rw.Write([]byte(msg))
			return
		}

		rw.Write([]byte(fmt.Sprintf("%d", version)))

	})
	http.ListenAndServe(":3000", r)
}
