package searches

import (
	"fmt"

	"net/http"

	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Goods struct {
	Name  string
	Image string
	Price float64
}

func SearchOnEbay(search string, prices ...int) []Goods {

	search = strings.Replace(search, " ", "+", -1)

	url := "https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=" + search + "&_sacat=0"

	if len(prices) == 1 {
		url += "&_udlo=" + strconv.Itoa(prices[0])
	}
	if len(prices) == 2 && prices[0] != 0 {
		url += "&_udlo=" + strconv.Itoa(prices[0])
		url += "&_udhi=" + strconv.Itoa(prices[1])

	} else if len(prices) == 2 && prices[0] == 0 {
		url += "&_udlo=" + "0"
		url += "&_udhi=" + strconv.Itoa(prices[1])
	}

	res, errorEbay := http.Get(url)

	if errorEbay != nil {
		fmt.Println("Err", errorEbay)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	goodsList := []Goods{}

	doc.Find("li.s-item").Each(func(index int, s *goquery.Selection) {

		s.Find("span[role='heading']").Each(func(index int, q *goquery.Selection) {
			goodsList = append(goodsList, Goods{Name: q.Text()})
		})
		s.Find("span.s-item__price").Each(func(index int, w *goquery.Selection) {
			str := w.Text()
			delimiter := " to "

			ind := strings.Index(str, delimiter)
			if ind != -1 {
				trimmedStr := strings.TrimSpace(str[:ind])
				num, err := strconv.ParseFloat(trimmedStr[1:], 64)
				if err == nil {
					goodsList[len(goodsList)-1].Price = num
				} else {
					fmt.Println("Ошибка", err)
				}
			} else {
				str := w.Text()[1:]
				num, err := strconv.ParseFloat(str, 64)
				if err == nil {
					goodsList[len(goodsList)-1].Price = num
				} else {
					fmt.Println("Ошибка", err)
				}
			}
		})

		s.Find("img[alt='" + goodsList[len(goodsList)-1].Name + "']").Each(func(index int, i *goquery.Selection) {
			src, _ := i.Attr("src")

			goodsList[len(goodsList)-1].Image = src

		})

	})
	sort.Slice(goodsList, func(i, j int) bool {
		return goodsList[i].Price < goodsList[j].Price
	})
	return goodsList

}
