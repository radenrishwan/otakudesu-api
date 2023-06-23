package scrape_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/radenrishwan/otakudesu-api/scrape"
	"github.com/radenrishwan/otakudesu-api/utils"
)

func TestHomePageSuccess(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/home", nil)
	if err != nil {
		t.Error("Error creating request")
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(scrape.HomePage)

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Status code does not match")
	}

	// get body response
	var result utils.DefaultResponse[utils.HomeResponse]
	body := recorder.Body.String()

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		t.Error(err)
	}

	if len(result.Data.Complete) != 10 {
		t.Error("Complete anime list not match len : ", len(result.Data.Complete))
	}

	if len(result.Data.Ongoing) != 15 {
		t.Error("Ongoing anime list not match len : ", len(result.Data.Ongoing))
	}
}

func TestAnimeListSuccess(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/anime-list", nil)
	if err != nil {
		t.Error("Error creating request")
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(scrape.AnimeList)

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Status code does not match")
	}

	// get body response
	var result utils.DefaultResponse[[]utils.AnimeListResponse]
	body := recorder.Body.String()

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		t.Error(err)
	}

	if result.Code != 200 {
		t.Error("Code not match")
	}
}

func TestAnimeDetail(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/anime/yusinda-sub-indo", nil)
	if err != nil {
		t.Error("Error creating request")
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(scrape.AnimeDetail)

	req = mux.SetURLVars(req, map[string]string{
		"id": "yusinda-sub-indo",
	})

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Status code does not match")
	}

	// get body response
	var result utils.DefaultResponse[utils.AnimeDetailResponse]
	body := recorder.Body.String()

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		t.Error(err)
	}

	r := utils.AnimeDetailResponse{
		Id:        "yusinda-sub-indo",
		Title:     "Yuusha ga Shinda!",
		Thumbnail: "https://otakudesu.lol/wp-content/uploads/2023/04/Yuusha-ga-Shinda-Sub-Indo.jpg",
		Score:     "6.33",
		Producer:  "Pony Canyon, Tohokushinsha Film Corporation, Shogakukan, WOWMAX",
	}

	if result.Data.Id != r.Id {
		t.Error("Id not match")
	}

	if result.Data.Title != r.Title {
		t.Error("Title not match")
	}

	if result.Data.Thumbnail != r.Thumbnail {
		t.Error("Thumbnail not match")
	}

	if result.Data.Score != r.Score {
		t.Error("Score not match")
	}

	if result.Data.Producer != r.Producer {
		t.Error("Producer not match")
	}
}
