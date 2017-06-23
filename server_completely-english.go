package main

import "github.com/gin-gonic/gin"
import "github.com/PuerkitoBio/goquery"
import "strings"
import "strconv"

type Event struct {
	Ort   string `json:"place"`
	Datum string `json:"timespan"`
}

func main() {
	m := gin.Default()
	m.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "Welcome to the code.camp API! Check out the docs at https://github.com/fronbasal/codecamps",
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

	m.GET("/camp/:index", func(c *gin.Context) {
		index, err1 := strconv.Atoi(c.Param("index"))
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

	m.GET("/workshops", func(c *gin.Context) {
		workshops, err := fetchWorkshops()
		if err != nil {
			c.JSON(500, gin.H{
				"error":   true,
				"message": "Could not fetch data",
			})
			return
		}
		c.JSON(200, workshops)
	})

	m.GET("/workshop/:index", func(c *gin.Context) {
		index, err1 := strconv.Atoi(c.Param("index"))
		workshop, err2 := fetchCamps()
		if err1 != nil || err2 != nil {
			c.JSON(500, gin.H{
				"error":   true,
				"message": "Could not fetch data",
			})
			return
		}
		if index <= len(workshop) {
			c.JSON(200, workshop[index])
		} else {
			c.JSON(400, gin.H{
				"error":   true,
				"message": "Invalid index!",
			})
		}
	})

	m.Run(":5000")
}

func fetchCamps() ([]Event, error) {
	doc, err := goquery.NewDocument("https://code.design/camps/")
	if err != nil {
		return nil, err
	}
	camps := []Event{}
	doc.Find("section.box ul li").Not(".pdf").Each(func(i int, s *goquery.Selection) {
		camps = append(camps, Event{Datum: strings.Replace(strings.SplitAfter(s.Text(), ",")[0], ",", "", -1), place: strings.Replace(strings.SplitAfter(s.Text(), ",")[1], ",", "", -1)})
	})
	return camps, nil
}

func fetchWorkshops() ([]Event, error) {
	doc, err := goquery.NewDocument("https://code.design/workshops/")
	if err != nil {
		return nil, err
	}
	workshops := []Event{}
	doc.Find("section.box ul li").Each(func(i int, s *goquery.Selection) {
		workshops = append(workshops, Event{Datum: strings.Replace(strings.SplitAfter(s.Text(), ",")[0], ",", "", -1), place: strings.Replace(strings.SplitAfter(s.Text(), ",")[1], ",", "", -1)})
	})
	return workshops, nil
}
