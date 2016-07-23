package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var version = "master"

type Settings struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	Prices           []int  `json:"prices"`
	SundayToThursday bool     `json:"sunday_to_thursday_working_days"` //if false then MondayToFriday will be true

}

func isWeekDay(date time.Time, sundayToThursday bool) bool {
	switch date.Weekday() {
	case time.Sunday:
		if (sundayToThursday == false) {
			return false;
		}
	case time.Friday:
		if (sundayToThursday == true ) {
			return false;
		}
	case time.Saturday:
		return false;
	default:
		return true;
	}
	return true;
}

func isWeekEnd(date time.Time, sundayToThursday bool) bool {
	return isWeekDay(date, sundayToThursday) == false
}

func daysLeft(date time.Time, sundayToThursday bool) int {

	workingDays := 0
	year, m, d := date.Date()
	daysInMonth := time.Date(year, m + 1, 0, 0, 0, 0, 0, time.UTC).Day()

	totalDays := daysInMonth - d

	if (totalDays < 3 &&  isWeekEnd(date, sundayToThursday)) {
		return workingDays
	} else {
		if (isWeekDay(date, sundayToThursday) == true) {
			workingDays++ //include current day
		}

		for ; totalDays > 0; totalDays-- {
			date = date.AddDate(0, 0, 1)

			switch date.Weekday() {
			case time.Sunday:
				if (sundayToThursday == true) {
					workingDays++
				}

			case time.Friday:
				if (sundayToThursday == false) {
					workingDays++
				}
			case time.Saturday:
			//do nothing
			//all other days
			default:workingDays++

			}
		}
		return workingDays
	}

}

func buildPredictions(budget float64, settings *Settings) string {
	if budget <= 0 {
		return ""
	}
	var buffer bytes.Buffer

	// food days left.
	// compute number of days left in this month, reduce the
	// no-food days per week
	daysLeftThisMonth := daysLeft(time.Now(), settings.SundayToThursday)

	mealdays := daysLeftThisMonth

	buffer.WriteString(fmt.Sprintf("🍔 you have to eat for %v more days.\n---\n", mealdays))
	for _, price := range settings.Prices {
		// how much will be off our current budget for this price?
		prediction := budget - float64(mealdays * price)
		gooddays := int(budget / float64(price))

		if mealdays <= gooddays {
			buffer.WriteString(fmt.Sprintf("₪%v: %v (₪+%v left)", price, gooddays, prediction))
		} else {
			buffer.WriteString(fmt.Sprintf("₪%v: %v + %v for ₪%v", price, gooddays, mealdays - gooddays, prediction * -1))
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

	budget, err := strconv.ParseFloat(strings.Replace(prettyAmount, "₪", "", -1), 64)
	if err != nil {
		log.Fatal(err)
	}
	submenu := buildPredictions(budget, settings)
	fmt.Printf("%s\n---\n%s---\n%s", prettyAmount, submenu, version)
}
