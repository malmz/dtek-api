package lunch

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/dtekcth/dtek-api/ent/schema"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

const freedomFormat = "1/2/2006 03:04:05 PM"

type FreedomTime struct {
	time.Time
}

func (t *FreedomTime) UnmarshalJSON(data []byte) error {
	// Copied from go stdlib time.Time.UnmarshalJSON
	if string(data) == "null" {
		return nil
	}

	var err error
	// Looking back at this i did not know why the quotes were there
	// It's because `data` has strings in it. It's trimming them off like pattern matching
	t.Time, err = time.Parse(`"`+freedomFormat+`"`, string(data))
	return err
}

type DishOccurrence struct {
	StartDate FreedomTime `json:"startDate" validate:"required"`
	EndDate   FreedomTime `json:"endDate" validate:"required"`
	Resturant struct {
		Name string `json:"mealProvidingUnitName" validate:"required"`
	} `json:"mealProvidingUnit" validate:"required"`
	DishBody []struct {
		Text     string `json:"dishDisplayName" validate:"required"`
		Category struct {
			Name string `json:"displayNameCategoryName" validate:"required"`
		} `json:"displayNameCategory" validate:"required"`
	} `json:"displayNames" validate:"required"`
	Title *struct {
		Swedish string `json:"dishTypeName" validate:"required"`
		English string `json:"dishTypeNameEnglish" validate:"required"`
	} `json:"dishType"`
	DishInfo struct {
		Emission float64 `json:"totalEmission" validate:"required"`
		Price    string  `json:"prices" validate:"required"`
		Recipes  []struct {
			Allergens []struct {
				Code string `json:"allergenCode" validate:"required"`
				Url  string `json:"allergenUrl" validate:"required"`
			} `json:"allergens"`
		} `json:"recipes" validate:"gt=0"`
	} `json:"dish" validate:"required"`
}

var validate = validator.New()
var defaultApiUrl = "http://carbonateapiprod.azurewebsites.net/api/v1/mealprovidingunits"

type KarenApiClient struct {
	baseUrl  string
	validate *validator.Validate
}

func NewKarenApiClient(baseUrl string) *KarenApiClient {
	if baseUrl == "" {
		baseUrl = defaultApiUrl
	}
	return &KarenApiClient{baseUrl: baseUrl, validate: validate}
}

func (k *KarenApiClient) FetchMenu(id string, startDate *time.Time, endDate *time.Time) ([]DishOccurrence, error) {
	dateFormat := "2006-01-02"

	query := url.Values{}
	if startDate != nil {
		query.Add("startDate", startDate.Format(dateFormat))
	}
	if endDate != nil {
		query.Add("endDate", endDate.Format(dateFormat))
	}
	req, err := url.JoinPath(k.baseUrl, id, "dishoccurrences")
	if err != nil {
		panic(err)
	}
	requ, err := url.Parse(req)
	if err != nil {
		panic(err)
	}
	requ.RawQuery = query.Encode()

	resp, err := http.Get(requ.String())
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("failed to fetch menu")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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
	validate.Struct(dishes)
	return dishes, nil
}

func (k *KarenApiClient) FetchToday(id string) ([]DishOccurrence, error) {
	return k.FetchMenu(id, nil, nil)
}

func (k *KarenApiClient) FetchDay(id string, date time.Time) ([]DishOccurrence, error) {
	return k.FetchMenu(id, &date, nil)
}

type KarenFetcher struct {
	id     string
	client *KarenApiClient
}

func NewKarenFetcher(id string) *KarenFetcher {
	return &KarenFetcher{id: id, client: NewKarenApiClient("")}
}

func (f *KarenFetcher) FetchByDate(date time.Time, lang string) (*LunchFetchResult, error) {
	dishes, err := f.client.FetchDay(f.id, date)
	if err != nil {
		return nil, err
	}

	if len(dishes) == 0 {
		return nil, errors.New("no dishes found")
	}

	var language string
	if lang == "en" {
		language = "English"
	} else {
		language = "Swedish"
	}

	name := dishes[0].Resturant.Name

	items := make([]schema.LunchMenuItem, len(dishes))

	for i, dish := range dishes {
		var body, title string
		for _, v := range dish.DishBody {
			if v.Category.Name == language {
				body = v.Text
				break
			}
		}
		if dish.Title != nil {
			if lang == "en" {
				title = dish.Title.English
			} else {
				title = dish.Title.Swedish
			}
		}

		var allergens []schema.Allergen

		if len(dish.DishInfo.Recipes) > 0 {
			allergens = make([]schema.Allergen, len(dish.DishInfo.Recipes[0].Allergens))
			for i, v := range dish.DishInfo.Recipes[0].Allergens {
				allergens[i] = schema.Allergen{
					Code:     v.Code,
					ImageUrl: v.Url,
				}
			}
		}

		items[i] = schema.LunchMenuItem{
			Title:        title,
			Body:         body,
			Preformatted: false,
			Allergens:    allergens,
			Emission:     dish.DishInfo.Emission,
			Price:        dish.DishInfo.Price,
		}
	}

	return &LunchFetchResult{
		Name:  name,
		Items: items,
	}, nil
}
