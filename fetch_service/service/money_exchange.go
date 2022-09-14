package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type ExchangeMoney struct{}

type ResponseBody struct {
	Date string `json:"date"`
	Info struct {
		Rate      float64 `json:"rate"`
		Timestamp int     `json:"timestamp"`
	} `json:"info"`
	Query struct {
		Amount int    `json:"amount"`
		From   string `json:"from"`
		To     string `json:"to"`
	} `json:"query"`
	Result  float64 `json:"result"`
	Success bool    `json:"success"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func UnmarshalUsers(r *http.Response) ResponseBody {
	body, err := ioutil.ReadAll(r.Body)
	checkError(err)
	var responseBody ResponseBody
	err = json.Unmarshal(body, &responseBody)
	checkError(err)

	return responseBody
}

var c = cache.New(24*time.Hour, 24*time.Hour)

func (ctrl ExchangeMoney) RequestCurrentPrice() (price float64) {
	//cache

	url := "https://api.apilayer.com/exchangerates_data/convert?to=IDR&from=USD&amount=1"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", "pAaRwfpM6l3I94U4pjhPAl74ig7aomI1")
	checkError(err)

	res, err := client.Do(req)
	checkError(err)
	if res.Body != nil {
		defer res.Body.Close()
	}

	responseObj := UnmarshalUsers(res)
	price = responseObj.Result
	c.Set("price", price, cache.DefaultExpiration)

	return
}

func (ctrl ExchangeMoney) GetCurrentPrice() (price float64) {
	res, found := c.Get("foo")
	if found {
		price = res.(float64)
	} else {
		price = 0
	}

	return
}
