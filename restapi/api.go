package restapi

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	mv "github.com/dev-crusader404/go-test-project/internal"
	"github.com/google/uuid"
)

type SearchRequest struct {
	Title string `json:"title,omitempty"`
	Year  string `json:"year,omitempty"`
}

type SearchResponse struct {
	MovieTitle   string   `json:"movieTitle,omitempty"`
	Year         string   `json:"year,omitempty"`
	Description  string   `json:"description,omitempty"`
	Rating       float32  `json:"rating,omitempty"`
	Genre        []string `json:"genre,omitempty"`
	ReleasedDate string   `json:"releasedDate,omitempty"`
	GrossIncome  int64    `json:"grossIncome,omitempty"`
}

func Fetcher(m mv.MovieFetcher, hd fetcherHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := hd(m, w, r); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}

func GetMovieHandler(m mv.MovieFetcher, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	reqID := ctx.Value("RequestID").(string)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer r.Body.Close()

	req := SearchRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	log.Printf("Search Request received for movie: %s year: %s with requestID: %s", req.Title, req.Year, reqID)
	id, err := m.GetMovie(ctx, req.Title, req.Year)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return err
	}

	result, err := m.GetDetails(ctx, id)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return err
	}
	response := SearchResponse{
		MovieTitle:   result.MovieTitle,
		Year:         result.Year,
		Description:  result.Overview,
		Rating:       result.Rating,
		Genre:        result.Genre,
		ReleasedDate: result.ReleasedDate,
		GrossIncome:  result.GrossIncome,
	}
	b, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return nil
}

type fetcherHandlerFunc func(mv mv.MovieFetcher, w http.ResponseWriter, r *http.Request) error

func Logger(next http.HandlerFunc) http.HandlerFunc {
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
