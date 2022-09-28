package scrape

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"github.com/radenrishwan/otakudesu-api/utils"
)

var (
	ENDPOINT = os.Getenv("ENDPOINT")
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(ENDPOINT)
	utils.PanicIfError(err, w, r)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		utils.PanicIfError(err, w, r)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	utils.PanicIfError(err, w, r)

	var result []utils.HomePageResponse
	document.Find(".detpost").Each(func(i int, selection *goquery.Selection) {
		var anime utils.HomePageResponse

		anime.Episode = selection.Find(".epz").Text()
		anime.Url = selection.Find(".thumb a").AttrOr("href", "")
		anime.Thumbnail = selection.Find(".thumb a .thumbz img").AttrOr("src", "")
		anime.Title = selection.Find(".thumb a .thumbz h2").Text()

		// get anime id
		d := strings.Split(anime.Url, "/")
		anime.Id = d[4]

		result = append(result, anime)
	})

	bytes, err := json.Marshal(utils.DefaultResponse[[]utils.HomePageResponse]{
		Code: 200,
		Data: result,
	})
	utils.PanicIfError(err, w, r)

	utils.NewSuccessResponse(string(bytes), w, r)
}

func AnimeList(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(ENDPOINT + "anime-list-2")
	utils.PanicIfError(err, w, r)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		utils.PanicIfError(err, w, r)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	utils.PanicIfError(err, w, r)

	var result []utils.AnimeListResponse
	document.Find("ul li .hodebgst").Each(func(i int, selection *goquery.Selection) {
		var anime utils.AnimeListResponse

		anime.Title = selection.Text()
		anime.Url = selection.AttrOr("href", "")

		// get anime id
		d := strings.Split(anime.Url, "/")
		anime.Id = d[4]

		result = append(result, anime)
	})

	bytes, err := json.Marshal(utils.DefaultResponse[[]utils.AnimeListResponse]{
		Code: 200,
		Data: result,
	})
	utils.PanicIfError(err, w, r)

	utils.NewSuccessResponse(string(bytes), w, r)
}

func AnimeDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	resp, err := http.Get(ENDPOINT + "anime/" + params["id"])
	utils.PanicIfError(err, w, r)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		utils.PanicIfError(err, w, r)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	utils.PanicIfError(err, w, r)

	var result utils.AnimeDetailResponse
	root := document.Find(".fotoanime")

	// TODO: check if id is correct

	result.Thumbnail = root.Find("img").AttrOr("src", "")
	result.Id = params["id"]
	root.Find(".infozin .infozingle p span").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			result.Title = utils.GetDetailInfo(s.Text())
		case 1:
			result.JapaneseTitle = utils.GetDetailInfo(s.Text())
		case 2:
			result.Score = utils.GetDetailInfo(s.Text())
		case 3:
			result.Producer = utils.GetDetailInfo(s.Text())
		case 4:
			result.Type = utils.GetDetailInfo(s.Text())
		case 5:
			result.Status = utils.GetDetailInfo(s.Text())
		case 7:
			result.Duration = utils.GetDetailInfo(s.Text())
		case 8:
			result.RealeaseDate = utils.GetDetailInfo(s.Text())
		case 9:
			result.Studio = utils.GetDetailInfo(s.Text())
		case 10:
			result.Genre = utils.GetDetailInfo(s.Text())
		}
	})

	var synopsis []string
	root.Find(".sinopc p").Each(func(i int, s *goquery.Selection) {
		synopsis = append(synopsis, s.Text())
	})
	result.Synopsis = strings.Join(synopsis, "")

	var episodes []utils.AnimeDetailEpisode
	document.Find(".venser .episodelist ul li").Each(func(i int, s *goquery.Selection) {
		var episode utils.AnimeDetailEpisode

		episode.Episode = s.Find("span a").Text()
		episode.Url = s.Find("span a").AttrOr("href", "")
		episode.UploadDate = s.Find(".zeebr").Text()

		// get episode id
		d := strings.Split(episode.Url, "/")
		episode.Id = d[4]

		episodes = append(episodes, episode)
	})
	result.Episode = episodes

	bytes, err := json.Marshal(utils.DefaultResponse[utils.AnimeDetailResponse]{
		Code: 200,
		Data: result,
	})
	utils.PanicIfError(err, w, r)

	utils.NewSuccessResponse(string(bytes), w, r)
}

func EpisodeDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	resp, err := http.Get(ENDPOINT + "episode/" + params["id"])
	utils.PanicIfError(err, w, r)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		utils.PanicIfError(err, w, r)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	utils.PanicIfError(err, w, r)

	var result utils.EpisodeDetail
	root := document.Find(".venser .venutama")
	result.Id = params["id"]
	result.Title = root.Find("h1").Text()

	var urls []utils.UrlInfo
	root.Find(".download ul li").Each(func(i int, s *goquery.Selection) {
		var url utils.UrlInfo
		url.Resolution = s.Find("strong").Text()

		s.Find("a").Each(func(i int, se *goquery.Selection) {
			url.Url = se.AttrOr("href", "")
			url.Host = se.Text()

			url.Size = s.Find("i").Text()

			urls = append(urls, url)
		})

	})
	result.Urls = urls

	bytes, err := json.Marshal(utils.DefaultResponse[utils.EpisodeDetail]{
		Code: 200,
		Data: result,
	})
	utils.PanicIfError(err, w, r)

	utils.NewSuccessResponse(string(bytes), w, r)
}
