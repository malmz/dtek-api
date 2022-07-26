package lunch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dtekcth/dtek-api/model"
	"github.com/rs/zerolog/log"
)

const freedomFormat = "2/01/2006 03:04:05 PM"

type FreedomTime struct {
	time.Time
}

func (t *FreedomTime) UnmarshalJSON(data []byte) error {
	// Copied from go stdlib time.Time.UnmarshalJSON
	if string(data) == "null" {
		return nil
	}

	var err error
	t.Time, err = time.Parse(`"`+freedomFormat+`"`, string(data))
	return err
}

type DishOccurrence struct {
	StartDate FreedomTime `json:"startDate"`
	EndDate   FreedomTime `json:"endDate"`
	Resturant struct {
		Name string `json:"mealProvidingUnitName"`
	} `json:"mealProvidingUnit"`
	DishBody []struct {
		Text     string `json:"dishDisplayName"`
		Category struct {
			Name string `json:"displayNameCategoryName"`
		} `json:"displayNameCategory"`
	} `json:"displayNames"`
	Title struct {
		Swedish string `json:"dishTypeName"`
		English string `json:"dishTypeNameEnglish"`
	} `json:"dishType"`
	DishInfo struct {
		Emission float64 `json:"totalEmission"`
		Price    string  `json:"prices"`
		Recipes  []struct {
			Allergens []struct {
				Code string `json:"allergenCode"`
				Url  string `json:"allergenUrl"`
			} `json:"allergens"`
		} `json:"recipes"`
	}
}

const baseUrl = "http://carbonateapiprod.azurewebsites.net/api/v1/mealprovidingunits"

func fetchKarenApi(id string, startDate time.Time, endDate time.Time) ([]DishOccurrence, error) {
	url := fmt.Sprintf("%s/%s/dishoccurrences?startDate=%s&endDate=%s", baseUrl, id, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("failed to fetch menu")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("failed to read menu body")
		return nil, err
	}

	var dishes []DishOccurrence

	err = json.Unmarshal(body, &dishes)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("failed to unmarshal menu")
		return nil, err
	}
	return dishes, nil
}

type KarenFetcher struct {
	id string
}

func NewKarenFetcher(id string) *KarenFetcher {
	return &KarenFetcher{id: id}
}

func (f *KarenFetcher) Fetch(date time.Time, lang string) (*model.LunchMenu, error) {
	startDate := date
	endDate := startDate.AddDate(0, 0, 1)
	dishes, err := fetchKarenApi(f.id, startDate, endDate)
	if err != nil {
		return nil, err
	}

	var language string
	if lang == "en" {
		language = "English"
	} else {
		language = "Swedish"
	}

	items := make([]model.LunchMenuItem, len(dishes))

	for i, dish := range dishes {
		var body, title string
		for _, v := range dish.DishBody {
			if v.Category.Name == language {
				body = v.Text
				break
			}
		}
		if lang == "en" {
			title = dish.Title.English
		} else {
			title = dish.Title.Swedish
		}

		allergens := make([]model.Allergen, len(dish.DishInfo.Recipes[0].Allergens))
		for i, v := range dish.DishInfo.Recipes[0].Allergens {
			allergens[i] = model.Allergen{
				Code:     v.Code,
				ImageUrl: v.Url,
			}
		}

		items[i] = model.LunchMenuItem{
			Title:        title,
			Body:         body,
			Preformatted: false,
			Allergens:    allergens,
			Emission:     dish.DishInfo.Emission,
			Price:        dish.DishInfo.Price,
		}
	}

	return &model.LunchMenu{
		Items: items,
	}, nil
}
