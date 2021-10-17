package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRecord struct {
	Id            string    `json:"id" bson:"_id"`
	UserName      string    `json:"uname" bson:"uname"`
	DisplayName   string    `json:"dname" bson:"dname"`
	DeveloperType string    `json:"developer_type" bson:"developer_type"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
	DetectedAt    time.Time `json:"detected_at" bson:"detected_at"`
}
type AvatarRecord struct {
	Id            string    `json:"id" bson:"_id"`
	Name          string    `json:"name" bson:"name"`
	AssetURL      string    `json:"asset_url" bson:"asset_url"`
	ImageURL      string    `json:"image_url" bson:"image_url"`
	ThumbnailURL  string    `json:"thumb_url" bson:"thumb_url"`
	AuthorId      string    `json:"author_id" bson:"author_id"`
	ReleaseStatus string    `json:"release" bson:"release"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
	DetectedAt    time.Time `json:"detected_at" bson:"detected_at"`
}
type WorldRecord struct {
	Id           string    `json:"id" bson:"_id"`
	Name         string    `json:"name" bson:"name"`
	AssetURL     string    `json:"asset_url" bson:"asset_url"`
	ImageURL     string    `json:"image_url" bson:"image_url"`
	ThumbnailURL string    `json:"thumb_url" bson:"thumb_url"`
	AuthorId     string    `json:"author_id" bson:"author_id"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
	DetectedAt   time.Time `json:"detected_at" bson:"detected_at"`
}

func httpUserPostHandler(w http.ResponseWriter, r *http.Request) {
	var userReq UserRecord
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\":\"Invalid json\"}"))
		return
	}

	_, err = dbVrcUsers.InsertOne(dbCtx, userReq)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"inserted\":false}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"inserted\":true}"))
}
func httpUserGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"inserted\":true}"))
}

func httpAvatarPostHandler(w http.ResponseWriter, r *http.Request) {
	var avatarReq AvatarRecord
	err := json.NewDecoder(r.Body).Decode(&avatarReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\":\"Invalid json\"}"))
		return
	}

	_, err = dbVrcAvatars.InsertOne(dbCtx, avatarReq)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"inserted\":false}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"inserted\":true}"))
}
func httpAvatarGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"inserted\":true}"))
}

func httpWorldPostHandler(w http.ResponseWriter, r *http.Request) {
	var worldReq WorldRecord
	err := json.NewDecoder(r.Body).Decode(&worldReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\":\"Invalid json\"}"))
		return
	}

	_, err = dbVrcWorlds.InsertOne(dbCtx, worldReq)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"inserted\":false}"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"inserted\":true}"))
}
func httpWorldGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"inserted\":true}"))
}

func httpSearchUsersGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "PUT /auth/login\n")
}
func httpSearchAvatarsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "PUT /auth/login\n")
}
func httpSearchWorldsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "PUT /auth/login\n")
}

func httpLoggerHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("<< %s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func httpAuthenticatorHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, headerFound := r.Header["Uploader-Key"]
		if !headerFound {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Please provide a uploader key"))
			return
		}

		uploaderKey, err := primitive.ObjectIDFromHex(header[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed uploader key"))
		}

		err = dbVrcUploaders.FindOne(dbCtx, bson.M{"_id": uploaderKey}).Err()
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Invalid uploader key"))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func httpRecoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
