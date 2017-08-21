package GOWikia_Wubsy

import (
	"encoding/json"
	"bytes"
	"github.com/valyala/fasthttp"
	"time"
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

var client = fasthttp.Client{ReadTimeout: time.Second * 10, WriteTimeout: time.Second * 10}

type WikiaApi struct {
	url string
}

type ContentResult struct {
	Sections []Section `json:"sections"`
}

type Section struct {
	Title string `json:"title"`
	Level int `json:"level"`
	Content []SectionContent `json:"SectionContent"`
	Images []SectionImages `json:"images"`
}

type SectionContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Elements []ListElement `json:"elements"`
}

type ListElement struct {
	Text string `json:"text"`
	Elements []ListElement `json:"elements"`
}

type SectionImages struct {
	Src string `json:"src"`
	Caption string `json:"caption"`
}

type ExpandedArticleResultSet struct {
	Items []ExpandedArticle `json:"items"`
	Basepath string `json:"basepath"`
}

type ExpandedArticle struct {
	Original_dimensions *OriginalDimension `json:"original_dimensions"`
	Url string `json:"url"`
	Ns int `json:"ns"`
	Abstract string `json:"abstract"`
	Thumbnail string `json:"thumbnail"`
	Revision *Revision `json:"revision"`
	ID int `json:"id"`
	Title string `json:"title"`
	Type string `json:"type"`
	Comments int `json:"comments"`
}

type OriginalDimension struct {
	Width int
	Height int
}

type Revision struct {
	ID int
	User string
	User_ID int
	Timestamp int
}

type ExpandedArticleParams struct {
	IDs string
	Titles string
	Abstract int
	Width int
	Height int
}

type UnexpandedListArticleResultSet struct {
	Items []UnexpandedArticle
	Offset string
	Basepath string
}

type UnexpandedArticle struct {
	ID int
	Title string
	URL string
	Ns int
}

type CrossWikiSearchResultSet struct {
	Items []CrossWikiSearchResult `json:"items"`
}

type CrossWikiSearchResult struct {
	Id int `json:"id"`
	Language string `json:"language"`
}

type SearchListItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Ns int `json:"ns"`
	Quality int `json:"quality"`
	}

type SearchResult struct {
	Total int `json:"total"`
	Batches int `json:"batches"`
	CurrentBatch int `json:"currentBatch"`
	Next int `json:"next"`
	Items []SearchListItem `json:"items"`
	Id int `json:"id"`
}

type GetAsSimpleJsonParams struct {
	IDs int
}

type QueryParams struct {
	Query string //Search query
	Hub string //Comma-separated list of verticals (e.g. Gaming, Entertainment, Lifestyle)
	Lang string //Comma separated language codes (e.g. en,de,fr)
	Rank string //The ranking to use in fetching the list of results, one of default, newest, oldest, recently-modified, stable, most-viewed, freshest, stalest
	Limit int //Limit the number of results
	Batch int //The batch(page) of results to fetch
}
func (w *WikiaApi) SearchList(params QueryParams) (*SearchResult, error) {

	//err := GetJson(w.url + "api/v1/Search/List/?query=" + params.Query, &params)
	urn, err := fetchUrl(w.url + "api/v1/Search/List/?query=" + params.Query)
	var searchResult *SearchResult
	if err != nil {
		return nil, err
	}

	return searchResult, json.Unmarshal(urn, &searchResult)
}

func (w *WikiaApi) GetArticleInfoByID(params GetAsSimpleJsonParams) (*ContentResult, error) {
	if params.IDs == 0{
		return nil, errors.New("ID is required")
	}

	urn, err := fetchUrl(w.url + "api/v1/Articles/AsSimpleJson?id=" + strconv.Itoa(params.IDs))

	var contentResult *ContentResult
	if err != nil {
		return nil, err
	}

	return contentResult, json.Unmarshal(urn, &contentResult)
}

//Credit to github.com/Time6628
func GetJson(url string, target interface{}) error {
	fmt.Println(url)
	stat, body, err := client.Get(nil, url)
	if err != nil || stat != 200 {
		return errors.New("Could not obtain json response")
	}
	return json.NewDecoder(bytes.NewReader(body)).Decode(target)


}

func fetchUrl(u string) ([]byte, error) {
	response, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}