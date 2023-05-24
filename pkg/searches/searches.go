package searches

import (
	"fmt"

	"net/http"
	//"os"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Goods struct {
	Name  string
	Price float64
	Image string
}

func SearchOnEbay(search string) []Goods {

	res, errorEbay := http.Get("https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=" + search + "&_sacat=0")

	if errorEbay != nil {
		fmt.Println("Err", errorEbay)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	goodsList := []Goods{}

	doc.Find("div.s-item__wrapper.clearfix").Each(func(index int, s *goquery.Selection) {

		s.Find("span[role='heading']").Each(func(index int, q *goquery.Selection) {
			goodsList = append(goodsList, Goods{Name: q.Text()})
		})
		s.Find("span.s-item__price").Each(func(index int, w *goquery.Selection) {
			str := w.Text()
			delimiter := " to "

			ind := strings.Index(str, delimiter)
			if ind != -1 {
				trimmedStr := strings.TrimSpace(str[ind+len(delimiter):])
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

		doc.Find("img").Each(func(index int, i *goquery.Selection) {
			src, _ := i.Attr("src")

			goodsList[len(goodsList)-1].Image = src

		})

	})
	return goodsList

}

func SortByPrice(priceFrom float64, priceTo float64, goods []Goods) []Goods {
	var result []Goods
	for _, item := range goods {
		if item.Price > priceFrom && item.Price < priceTo {
			result = append(result, item)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Price < result[j].Price
	})
	return result
}
