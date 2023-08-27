package calculators

import (
	"strings"

	"github.com/tonydawhale/Go-Skyhelper-Networth/structs"
	"github.com/tonydawhale/Go-Skyhelper-Networth/utils"
)

func CalculateEssence(item structs.Item, prices map[string]float64) *structs.NetworthItem {
	price := 0.0
	if _, ok := prices[strings.ToLower(item.Id)]; ok {
		price = prices[strings.ToLower(item.Id)]
	}

	if price > 0 {
		return &structs.NetworthItem{
			Name:        utils.TitleCast(strings.Split(item.Id, "_")[1]),
			Id:          item.Id,
			Price:       float64(item.Amount) * price,
			Calculation: []interface{}{},
			Count:       item.Amount,
			Soulbound:   false,
		}
	} else {
		return nil
	}
}
