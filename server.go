package main

import "github.com/gin-gonic/gin"
import "github.com/PuerkitoBio/goquery"
import "strings"
import "strconv"

type Camp struct {
	Ort   string `json:"ort"`
	Datum string `json:"datum"`
}

func main() {
	m := gin.Default()
	m.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "Welcome to the code.camp API! Check out the latest code camps with /camps/all or get a specific one at /camp/0",
		})
	})

	m.GET("/camps", func(c *gin.Context) {
		camps, err := fetchCamps()
		if err != nil {
			c.JSON(500, gin.H{
				"error":   true,
				"message": "Could not fetch data",
			})
			return
		}
		c.JSON(200, camps)
	})

	m.GET("/camp/:camp", func(c *gin.Context) {
		index, err1 := strconv.Atoi(c.Param("camp"))
		camp, err2 := fetchCamps()
		if err1 != nil || err2 != nil {
			c.JSON(500, gin.H{
				"error":   true,
				"message": "Could not fetch data",
			})
			return
		}
		if index <= len(camp) {
			c.JSON(200, camp[index])
		} else {
			c.JSON(400, gin.H{
				"error":   true,
				"message": "Invalid index!",
			})
		}
	})

	m.Run(":3001")
}

func fetchCamps() ([]Camp, error) {
	doc, err := goquery.NewDocument("https://code.design/camps/")
	if err != nil {
		return nil, err
	}
	camps := []Camp{}
	doc.Find("section.box ul li").Not(".pdf").Each(func(i int, s *goquery.Selection) {
		camps = append(camps, Camp{Datum: strings.SplitAfter(s.Text(), ",")[0], Ort: strings.SplitAfter(s.Text(), ",")[1]})
	})
	return camps, nil
}
