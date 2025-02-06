package emailsuploader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mryeibis/indexer/internal/models"
)

var headerSetters = map[string]func(*models.Email, string){
	"Message-ID": func(email *models.Email, value string) { email.MessageID += value },
	"Date": func(email *models.Email, value string) {
		if email.Date != "" {
			return
		}

		email.Date += parseEmailDate(value)
	},
	"From":                      func(email *models.Email, value string) { email.From += value },
	"To":                        func(email *models.Email, value string) { email.To += value },
	"Subject":                   func(email *models.Email, value string) { email.Subject += value },
	"Mime-Version":              func(email *models.Email, value string) { email.MimeVersion += value },
	"Content-Type":              func(email *models.Email, value string) { email.ContentType += value },
	"Content-Transfer-Encoding": func(email *models.Email, value string) { email.ContentTransferEncoding += value },
	"X-From":                    func(email *models.Email, value string) { email.XFrom += value },
	"X-To":                      func(email *models.Email, value string) { email.XTo += value },
	"X-cc":                      func(email *models.Email, value string) { email.XCC += value },
	"X-bcc":                     func(email *models.Email, value string) { email.XBCC += value },
	"X-Folder":                  func(email *models.Email, value string) { email.XFolder += value },
	"X-Origin":                  func(email *models.Email, value string) { email.XOrigin += value },
	"X-FileName":                func(email *models.Email, value string) { email.XFileName += value },
}

func ParseEmail(path string) (*models.Email, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	email := models.Email{}
	parseEmailHeaders(scanner, &email)
	parseEmailBody(scanner, &email)

	return &email, nil
}

func parseEmailHeaders(scanner *bufio.Scanner, email *models.Email) {
	var currentHeaderKey string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) != 2 {
			setter := headerSetters[currentHeaderKey]
			if setter != nil {
				setter(email, " "+strings.TrimSpace(line))
			}

			continue
		}

		currentHeaderKey = headerParts[0]
		setter := headerSetters[currentHeaderKey]
		if setter != nil {
			setter(email, strings.TrimSpace(headerParts[1]))
		}
	}
}

func parseEmailBody(scanner *bufio.Scanner, email *models.Email) {
	var contentBuilder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		contentBuilder.WriteString(strings.TrimSpace(line) + "\n")
	}

	email.Content = contentBuilder.String()
}

func parseEmailDate(dateTime string) string {
	dateTime = strings.Split(dateTime, " (")[0]
	date, err := time.Parse("Mon, _2 Jan 2006 15:04:05 -0700", dateTime)
	if err != nil {
		fmt.Println(err.Error())
		return dateTime
	}

	return date.Format(time.RFC3339)
}
