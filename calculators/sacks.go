package calculators

import (
	"errors"
	"slices"
	"strings"

	"github.com/tonydawhale/Go-Skyhelper-Networth/constants"
	"github.com/tonydawhale/Go-Skyhelper-Networth/structs"
	"github.com/tonydawhale/Go-Skyhelper-Networth/utils"
)

func CalculateSackItem(item map[string]interface{}, prices *map[string]float64) (*structs.NetworthItem, error) {
	price := 0.0
	if _, ok := (*prices)[strings.ToLower(item["id"].(string))]; ok {
		price = (*prices)[strings.ToLower(item["id"].(string))]
	}
	if strings.HasPrefix(item["id"].(string), "RUNE_") && !slices.Contains(constants.ValidRunes, item["id"].(string)) {
		return nil, nil
	}
	if price > 0 {
		return &structs.NetworthItem{
			Name:        getDisplayName(item),
			Id:          item["id"].(string),
			Price:       price * float64(item["amount"].(int)),
			Calculation: []interface{}{},
			Count:       item["amount"].(int),
			Soulbound:   false,
		}, nil
	} else {
		return nil, errors.New("price not found for item " + item["id"].(string))
	}
}

func getDisplayName(item map[string]interface{}) string {
	if item["display_name"] != nil {
		return utils.TitleCast(item["display_name"].(string))
	}
	return utils.TitleCast(constants.ItemsMap[item["id"].(string)].Name)
}
