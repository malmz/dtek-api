package lunch

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dtekcth/dtek-api/model"
	"github.com/ledongthuc/pdf"
)

const pdfUrl = "http://www.cafelinsen.se/menyer/cafe-linsen-lunch-meny.pdf"

var weekReg = regexp.MustCompile(`^Cafè Linsen Vecka (\d+)$`)
var dayReg = regexp.MustCompile(`^(Måndag|Tisdag|Onsdag|Torsdag|Fredag|Lördag|Söndag):\s+(.+)`)

type LinsenFetcher struct {
	karenFetcher *KarenFetcher
}

func NewLinsenFetcher() *LinsenFetcher {
	return &LinsenFetcher{
		karenFetcher: NewKarenFetcher("b672efaf-032a-4bb8-d2a5-08d558129279"),
	}
}

func parseLinsenPdf(data []byte) ([]model.LunchMenu, error) {
	// Parse pdf
	p, err := pdf.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	// Extract plain text
	d, err := p.GetPlainText()
	if err != nil {
		return nil, err
	}

	// Iterator over lines
	scanner := bufio.NewScanner(d)
	scanner.Scan()
	line := scanner.Text()

	// Check that the week of the menu matches this week
	weekMatch := weekReg.FindStringSubmatch(line)
	if weekMatch == nil {
		return nil, errors.New("no week in pdf")
	}
	menuWeek, err := strconv.Atoi(weekMatch[1])
	if err != nil {
		return nil, err
	}

	_, currentWeek := time.Now().ISOWeek()

	if currentWeek != menuWeek {
		return nil, errors.New("menu is not for current week")
	}

	builder := strings.Builder{}
	menu := make([]model.LunchMenu, 0, 5)

	for scanner.Scan() {
		line := scanner.Text()
		dayMatch := dayReg.FindStringSubmatch(line)
		if dayMatch != nil {

			menu = append(menu, model.LunchMenu{
				Items: []model.LunchMenuItem{
					{
						Body:         builder.String(),
						Preformatted: true,
					},
				},
			})
			builder.Reset()
			builder.WriteString(dayMatch[2])
		} else {
			builder.WriteString(line)
		}
	}

	return menu, nil
}

func fetchLinsenFromPdf(date time.Time) (*model.LunchMenu, error) {
	_, currentWeek := time.Now().ISOWeek()
	_, dateWeek := date.ISOWeek()
	if currentWeek != dateWeek {
		return nil, errors.New("pdf is only available for current week")
	}

	resp, err := http.Get(pdfUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pdfData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	menu, err := parseLinsenPdf(pdfData)
	if err != nil {
		return nil, err
	}
	weekday := date.Weekday()
	return &menu[weekday-1], nil
}

func (f *LinsenFetcher) Fetch(date time.Time, lang string) (*model.LunchMenu, error) {
	menu, err := f.karenFetcher.Fetch(date, lang)
	if err != nil {
		menu, err := fetchLinsenFromPdf(date)
		return menu, err
	}
	for _, item := range menu.Items {
		item.Title = ""
	}
	return menu, nil
}
