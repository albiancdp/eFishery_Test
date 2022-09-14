package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"fetch_service/service"

	"github.com/gin-gonic/gin"
)

type ResourceController struct{}

type ResponseBodyArray []struct {
	Uuid         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	PriceUSD     string `json:"price_usd"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

type ResponseBody struct {
	Uuid         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	PriceUSD     string `json:"price_usd"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

var ExchangeService = new(service.ExchangeMoney)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func UnmarshalUsers(r *http.Response) ResponseBodyArray {
	body, err := ioutil.ReadAll(r.Body)
	checkError(err)
	var responseBodyArray ResponseBodyArray
	err = json.Unmarshal(body, &responseBodyArray)
	checkError(err)

	return responseBodyArray
}

func getRequestEfishery() ResponseBodyArray {
	url := "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Set("apikey", "pAaRwfpM6l3I94U4pjhPAl74ig7aomI1")
	checkError(err)

	res, err := client.Do(req)
	checkError(err)
	if res.Body != nil {
		defer res.Body.Close()
	}

	responseArray := UnmarshalUsers(res)
	return responseArray
}

func (ctrl ResourceController) FetchResource(c *gin.Context) {
	var dataResponse []ResponseBody
	price := ExchangeService.GetCurrentPrice()
	if price == 0 {
		price = ExchangeService.RequestCurrentPrice()
	}
	responseArray := getRequestEfishery()
	for i := 1; i < len(responseArray); i++ {
		element := responseArray[i]
		if element.Uuid != "" {
			priceIDR, err := strconv.ParseFloat(element.Price, 64)
			checkError(err)
			priceUSD := float64(int((priceIDR/price)*100)) / 100
			responseArray[i].PriceUSD = "$" + strconv.FormatFloat(priceUSD, 'f', 2, 64)
			dataResponse = append(dataResponse, responseArray[i])
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    dataResponse,
		"message": "response found",
	})
}

func (ctrl ResourceController) AggregateResource(c *gin.Context) {
	var dataResponse []ResponseBody
	price := ExchangeService.GetCurrentPrice()
	if price == 0 {
		price = ExchangeService.RequestCurrentPrice()
	}
	responseArray := getRequestEfishery()
	for i := 1; i < len(responseArray); i++ {
		element := responseArray[i]
		if element.Uuid != "" {
			priceIDR, err := strconv.ParseFloat(element.Price, 64)
			checkError(err)
			priceUSD := float64(int((priceIDR/price)*100)) / 100
			responseArray[i].PriceUSD = "$" + strconv.FormatFloat(priceUSD, 'f', 2, 64)
			dataResponse = append(dataResponse, responseArray[i])
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    dataResponse,
		"message": "response found",
	})
}
