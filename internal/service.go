package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/dev-crusader404/go-test-project/client"
	props "github.com/dev-crusader404/go-test-project/startup"
)

type MovieFetcher interface {
	GetMovie(ctx context.Context, title, year string) (int, error)
	GetDetails(ctx context.Context, id int) (MovieResult, error)
	GetMovieNowScreening(ctx context.Context, pageSize int32, result chan<- MovieResult) error
}

type Movie struct {
	Client client.RestClient
}

type MovieResult struct {
	ID           int      `json:"id"`
	MovieTitle   string   `json:"title,omitempty"`
	Overview     string   `json:"overview"`
	Year         string   `json:"year,omitempty"`
	Rating       float32  `json:"vote_average,omitempty"`
	Genre        []string `json:"genre,omitempty"`
	Origin       string   `json:"origin_country"`
	ReleasedDate string   `json:"release_date"`
	GrossIncome  int64    `json:"revenue,omitempty"`
}

func (sm *Movie) GetMovie(ctx context.Context, title, year string) (int, error) {
	URL := props.GetAll().GetString("MOVIE-URL", "")
	if URL == "" {
		log.Panic("no url found")
	}
	params := url.Values{}
	params.Add("query", title)
	params.Add("include_adult", "false")
	params.Add("language", "en-US")
	params.Add("page", "1")
	params.Add("year", year)
	u, err := url.Parse(fmt.Sprintf("%s/search/movie?%s", URL, params.Encode()))
	if err != nil {
		log.Panic("error parsing url: " + URL)
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    u,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Content-Type":  {"application/json"},
			"Authorization": {GetAuth()},
		},
	}

	resp, err := sm.Client.Do(req)
	if err != nil {
		log.Printf("error during call: %s", err.Error())
		return -1, err
	}

	if resp == nil || resp.Body == nil {
		err := fmt.Errorf("nil respose/body received")
		log.Println(err)
		return -1, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("unable to read response body")
		log.Println(err)
		return -1, err
	}

	if resp.StatusCode != 200 {
		err := fmt.Errorf("\nunexpected status code: %d", resp.StatusCode)
		log.Println(err)
		return -1, err
	}

	resMsg := struct {
		Results []struct {
			ID int `json:"id"`
		} `json:"results"`
	}{}
	err = json.Unmarshal(b, &resMsg)
	if err != nil {
		err := fmt.Errorf("error while unmarshalling")
		log.Println(err)
		return -1, err
	}
	log.Printf("Received Response: %+v", resMsg)

	if len(resMsg.Results) == 0 {
		return -1, fmt.Errorf("no result found")
	}

	return resMsg.Results[0].ID, nil
}

func (sm *Movie) GetDetails(ctx context.Context, id int) (MovieResult, error) {
	var result MovieResult
	URL := props.GetAll().GetString("MOVIE-URL", "")
	if URL == "" {
		log.Panic("no url found")
	}
	params := url.Values{}
	params.Add("language", "en-US")
	u, err := url.Parse(fmt.Sprintf("%s/movie/%d?%s", URL, id, params.Encode()))
	if err != nil {
		log.Panic("error parsing url: " + URL)
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    u,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Content-Type":  {"application/json"},
			"Authorization": {GetAuth()},
		},
	}

	resp, err := sm.Client.Do(req)
	if err != nil {
		log.Printf("error during call: %s", err.Error())
		return result, err
	}

	if resp == nil || resp.Body == nil {
		err := fmt.Errorf("nil respose/body received")
		log.Println(err)
		return result, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("unable to read response body")
		log.Println(err)
		return result, err
	}

	if resp.StatusCode != 200 {
		err := fmt.Errorf("\nunexpected status code: %d", resp.StatusCode)
		log.Println(err)
		return result, err
	}

	resMsg := struct {
		Genre []struct {
			Name string `json:"name"`
		} `json:"genres"`
		MovieTitle   string   `json:"title,omitempty"`
		Overview     string   `json:"overview"`
		Year         string   `json:"year,omitempty"`
		Rating       float32  `json:"vote_average,omitempty"`
		Origin       []string `json:"origin_country"`
		ReleasedDate string   `json:"release_date"`
		GrossIncome  int64    `json:"revenue,omitempty"`
	}{}
	err = json.Unmarshal(b, &resMsg)
	if err != nil {
		err := fmt.Errorf("error while unmarshalling")
		log.Println(err)
		return result, err
	}
	resMsg.Year = strings.Split(resMsg.ReleasedDate, "-")[0]
	log.Printf("Received Response: %+v", resMsg)
	for _, v := range resMsg.Genre {
		result.Genre = append(result.Genre, v.Name)
	}
	result.MovieTitle = resMsg.MovieTitle
	result.Overview = resMsg.Overview
	result.Year = resMsg.Year
	result.Rating = resMsg.Rating
	result.Origin = resMsg.Origin[0]
	result.ReleasedDate = resMsg.ReleasedDate
	result.GrossIncome = resMsg.GrossIncome
	return result, nil
}

func (sm *Movie) GetMovieNowScreening(ctx context.Context, pageSize int32, result chan<- MovieResult) error {
	defer close(result)

	URL := props.GetAll().GetString("MOVIE-URL", "")
	if URL == "" {
		log.Panic("no url found")
	}

	params := url.Values{}
	params.Add("language", "en-US")
	params.Add("page", fmt.Sprintf("%d", pageSize))
	u, err := url.Parse(fmt.Sprintf("%s/movie/now_playing?%s", URL, params.Encode()))
	if err != nil {
		log.Panic("error parsing url: " + URL)
		return err
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    u,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Content-Type":  {"application/json"},
			"Authorization": {GetAuth()},
		},
	}

	resp, err := sm.Client.Do(req)
	if err != nil {
		log.Printf("error during call: %s", err.Error())
		return err
	}

	if resp == nil || resp.Body == nil {
		err := fmt.Errorf("nil respose/body received")
		log.Println(err)
		return err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("unable to read response body")
		log.Println(err)
		return err
	}

	if resp.StatusCode != 200 {
		err := fmt.Errorf("\nunexpected status code: %d", resp.StatusCode)
		log.Println(err)
		return err
	}

	res := struct {
		Result []MovieResult `json:"results"`
	}{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		err := fmt.Errorf("error while unmarshalling")
		log.Println(err)
		return err
	}
	if len(res.Result) == 0 {
		return errors.New("empty response")
	}

	for _, description := range res.Result {
		result <- description
	}
	return nil
}

func GetAuth() string {
	token := props.GetAll().MustGetString("TOKEN")

	return "Bearer " + token
}

func closeChannel(result chan<- MovieResult, err chan<- error) {
	close(result)
}
