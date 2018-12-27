package main

import(
	"fmt"
	"../GOWikia_Wubsy"
	"strconv"
)

var (
	articleName string
	articleUrl string
	articleId int
	totalItems int
)

func main() {
	w, err := GOWikia_Wubsy.NewClient("http://slimerancher.wikia.com/")

	if err != nil {
		fmt.Println(err)
		return
	}

	searchParams := GOWikia_Wubsy.QueryParams{
		Query: "Plorts",
		Lang:  "en",
		Limit: 5,
	}

	results, err := w.SearchList(searchParams)
	if err != nil {
		fmt.Println("Fatal error: " + err.Error())
		return
	}

	if results.Total <= 25 && searchParams.Limit <= 25{
		totalItems = searchParams.Limit
	} else if results.Total > 25{
		totalItems = 25
	}
	fmt.Println(totalItems)
	for i := 0; i < totalItems; i++ {
		if results != nil {
			fmt.Println("~~~~~~~")
			articleName = results.Items[i].Title
			articleUrl = results.Items[i].Url
			articleId = results.Items[i].Id

			fmt.Println("Name: "+articleName)
			fmt.Println("URL: "+articleUrl)
			fmt.Println("ID: "+strconv.Itoa(articleId))

			articleParams := GOWikia_Wubsy.GetAsSimpleJsonParams{
				IDs: results.Items[i].Id,
			}



			article, err := w.GetArticleInfoByID(articleParams)
			if err != nil{
				fmt.Println("Fatal error: "+err.Error())
				return
			}

			if len(article.Sections) < 1 || len(article.Sections[0].Images) < 1{
				continue
				return
			}

			var articleThumbnail = article.Sections[0].Images[0].Src

			fmt.Println(articleThumbnail)
		}
	}
}
