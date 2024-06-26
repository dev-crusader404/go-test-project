package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type DB interface {
	Store(s string) error
}

type Store struct{}

func NewDB() DB {
	return &Store{}
}

func (s *Store) Store(a string) error {
	fmt.Printf("\nStoring the value: %s", a)
	return nil
}

func MakeHTTPFunc(db DB, hd myHandlerFunc) http.HandlerFunc {
	fmt.Println("creating the makeHTTPFunc")
	return func(w http.ResponseWriter, r *http.Request) {
		if err := hd(db, w, r); err != nil {
			http.Error(w, err.Error(), 500)
		}
		db.Store("Key to success")
	}
}

func Handler(db DB, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	reqID := ctx.Value("RequestID").(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer r.Body.Close()

	request := struct {
		Message string `json:"body"`
	}{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	fmt.Println("I am at handler function")
	db.Store("Way through the key: " + reqID)
	time.Sleep(2 * time.Second)
	resp, _ := json.Marshal(map[string]any{"status": fmt.Sprintf("Successfully processed requestID: %s", reqID),
		"message": request.Message})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return nil
}

type myHandlerFunc func(db DB, w http.ResponseWriter, r *http.Request) error

func Logger(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("logger init")
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "RequestID", requestID)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
