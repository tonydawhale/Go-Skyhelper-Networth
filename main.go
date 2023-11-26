package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tonydawhale/Go-Skyhelper-Networth/structs"
	"github.com/tonydawhale/Go-Skyhelper-Networth/utils"
)

func main() {
	file, err := os.Open("profile.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	val, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	// Your code here
	var data map[string]interface{}
	json.Unmarshal(val, &data)

	members, ok := data["members"].(map[string]interface{})
	if !ok {
		panic("Members not found in profile")
	}

	memberData, ok := members["27a9fd8dcd8b4beca3753c2e318f44f1"].(map[string]interface{})
	if !ok {
		panic("Member data not found")
	}

	var options = structs.CalculatorOptions{
		V2Endpoint:     false,
		Cache:          false,
		OnlyNetworth:   false,
		Prices:         nil,
		ReturnItemData: false,
		MuseumData:     nil,
	}

	networth, err := GetNetworth(memberData, 0, &options)
	if err != nil {
		panic(err)
	}
	fmt.Println(networth)
}

func GetNetworth(profile map[string]interface{}, bankBalance int, options *structs.CalculatorOptions) (map[string]interface{}, error) {
	if profile == nil {
		return nil, errors.New("profile is nil")
	}
	var purse int
	if options.V2Endpoint {
		purse = profile["currencies"].(map[string]interface{})["coin_purse"].(int)
	} else {
		purse = profile["coin_purse"].(int)
	}
	prices, err := ParsePrices(options.Prices, options.Cache)
	if err != nil {
		return nil, err
	}
	items, err := utils.ParseItems(profile, options.MuseumData, options.V2Endpoint)
	if err != nil {
		return nil, err
	}
	return utils.CalculateNetworth(items, purse, bankBalance, prices, options.OnlyNetworth, options.ReturnItemData)
}

func ParsePrices(prices map[string]int, cache bool) (map[string]interface{}, error) {
	return nil, nil
}

type PriceCache struct {
	LastCache int64
	Prices    map[string]float64
}

var CachedPrices PriceCache
var isLoadingPrices = false

func GetPrices(cache bool) (map[string]float64, error){
	if CachedPrices.LastCache > (time.Now().UnixNano() / int64(time.Millisecond) - 1000 * 60 * 5) && !cache {
		return CachedPrices.Prices, nil
	}
	if isLoadingPrices {
		for isLoadingPrices {
			time.Sleep(1000)
		}
		return CachedPrices.Prices, nil
	}

	isLoadingPrices = true
	resp, err := http.Get("https://raw.githubusercontent.com/SkyHelperBot/Prices/main/prices.json")
	if err != nil {
		return nil, err
	}
	var pricesResp map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&pricesResp); err != nil {
		return nil, err
	}
	var prices = make(map[string]float64)
	for key, value := range pricesResp {
		prices[strings.ToLower(key)] = value
	}
	CachedPrices = PriceCache{
		LastCache: time.Now().UnixNano() / int64(time.Millisecond),
		Prices:    prices,
	}
	isLoadingPrices = false
	return prices, nil
}
