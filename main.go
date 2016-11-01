package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var version = "master"

// Settings Holds the settings for bitbar
type Settings struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Prices            []int  `json:"prices"`
	NoFoodDaysPerWeek int    `json:"no_food_days_per_week"`
}

func daysLeft(date time.Time) int {
	year, m, d := date.Date()
	daysInMonth := time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()

	return daysInMonth - d
}

func buildPredictions(budget float64, settings *Settings) string {
	if budget <= 0 {
		return ""
	}
	var buffer bytes.Buffer

	// food days left.
	// compute number of days left in this month, reduce the
	// no-food days per week
	daysLeftThisMonth := daysLeft(time.Now())
	mealdays := daysLeftThisMonth - (daysLeftThisMonth/7)*settings.NoFoodDaysPerWeek

	buffer.WriteString(fmt.Sprintf("🍔 you have to eat for %v more days.\n---\n", mealdays))
	buffer.WriteString(fmt.Sprintf("It's %.2f a day.\n---\n", budget/float64(mealdays+1)))
	for _, price := range settings.Prices {
		// how much will be off our current budget for this price?
		prediction := budget - float64(mealdays*price)
		gooddays := int(budget / float64(price))

		if mealdays <= gooddays {
			buffer.WriteString(fmt.Sprintf("₪%v: %v (₪+%v left)", price, gooddays, prediction))
		} else {
			buffer.WriteString(fmt.Sprintf("₪%v: %v + %v for ₪%v", price, gooddays, mealdays-gooddays, prediction*-1))
		}
		if prediction < 0 {
			buffer.WriteString("|color=red")
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dot10bis := filepath.Join(u.HomeDir, ".10bis.json")
	data, err := ioutil.ReadFile(dot10bis)
	if err != nil {
		log.Fatal(err)
	}

	settings := &Settings{}
	json.Unmarshal(data, settings)

	options := &cookiejar.Options{}
	jar, err := cookiejar.New(options)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}
	resp, err := client.Post(
		"https://www.10bis.co.il/Account/LogonAjax",
		"application/json",
		strings.NewReader(
			fmt.Sprintf("{\"timestamp\":%d,\"model\":{\"UserName\":\"%s\",\"Password\":\"%s\",\"SocialLoginUID\":\"\",\"FacebookUserId\":\"undefined\"},\"returnUrl\":\"\"}",
				time.Now().Unix(),
				settings.Username,
				settings.Password),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = client.Get("https://www.10bis.co.il/Account/UserReport")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := doc.Find(".userReportDataTbl th.currency").First()
	prettyAmount := strings.TrimSpace(res.Text())
	prettyAmount = strings.Replace(prettyAmount, ",", "", -1)

	budget, err := strconv.ParseFloat(strings.Replace(prettyAmount, "₪", "", -1), 64)
	if err != nil {
		log.Fatal(err)
	}
	submenu := buildPredictions(budget, settings)
	fmt.Printf("%s\n---\n%s---\n%s", prettyAmount, submenu, version)
}
