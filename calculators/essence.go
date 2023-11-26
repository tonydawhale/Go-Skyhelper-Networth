package calculators

import (
	"errors"
	"strings"

	"github.com/tonydawhale/Go-Skyhelper-Networth/structs"
	"github.com/tonydawhale/Go-Skyhelper-Networth/utils"
)

func CalculateEssence(item map[string]interface{}, prices *map[string]float64) (*structs.NetworthItem, error) {
	price := 0.0
	if _, ok := (*prices)[strings.ToLower(item["id"].(string))]; ok {
		price = (*prices)[strings.ToLower(item["id"].(string))]
	}

	if price > 0 {
		return &structs.NetworthItem{
			Name:        utils.TitleCast(strings.Split(item["id"].(string), "_")[1]),
			Id:          item["id"].(string),
			Price:       float64(item["amount"].(int)) * price,
			Calculation: []interface{}{},
			Count:       item["amount"].(int),
			Soulbound:   false,
		}, nil
	} else {
		return nil, errors.New("price not found for item " + item["id"].(string))
	}
}
