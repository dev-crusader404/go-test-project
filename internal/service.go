package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	c "github.com/dev-crusader404/go-test-project/client"
	props "github.com/dev-crusader404/go-test-project/startup"
)

type MovieFetcher interface {
	GetMovie(title string) string
	GetDetails(id string) MovieResult
}

type Movie struct {
	Title string
	Year  string
	Id    string
}

type MovieResult struct {
	ID           int      `json:"id"`
	MovieTitle   string   `json:"title,omitempty"`
	Overview     string   `json:"overview"`
	Year         string   `json:"year,omitempty"`
	Rating       int32    `json:"vote_average,omitempty"`
	Genre        []string `json:"genre,omitempty"`
	Origin       string   `json:"origin_country"`
	ReleasedDate string   `json:"release_date"`
	GrossIncome  int64    `json:"revenue,omitempty"`
}

func (sm *Movie) GetMovie(client c.RestClient) (int, error) {
	URL := props.GetAll().GetString("MOVIE-URL", "")
	if URL == "" {
		log.Panic("no url found")
	}
	params := url.Values{}
	params.Add("query", sm.Title)
	params.Add("include_adult", "false")
	params.Add("language", "en-US")
	params.Add("page", "1")
	params.Add("year", sm.Year)
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

	resp, err := client.Do(req)
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

func GetAuth() string {
	token := props.GetAll().MustGetString("TOKEN")

	return "Bearer " + token
}
