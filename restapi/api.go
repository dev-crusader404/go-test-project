package restapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	mv "github.com/dev-crusader404/go-test-project/internal"
	ds "github.com/dev-crusader404/go-test-project/models"
	md "github.com/dev-crusader404/go-test-project/startup/middleware"
)

type fetcherHandlerFunc func(mv mv.MovieFetcher, w http.ResponseWriter, r *http.Request) error

func Fetcher(m mv.MovieFetcher, hd fetcherHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := hd(m, w, r); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}

func GetMovieHandler(m mv.MovieFetcher, w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	reqID := ctx.Value(md.RequestIDKey).(string)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer r.Body.Close()

	req := ds.SearchRequest{}
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
	response := ds.SearchResponse{
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
