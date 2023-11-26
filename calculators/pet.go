package calculators

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/tonydawhale/Go-Skyhelper-Networth/constants"
	"github.com/tonydawhale/Go-Skyhelper-Networth/structs"
	"github.com/tonydawhale/Go-Skyhelper-Networth/utils"
)

type PetLevelPrices struct {
	Id     string  `json:"id"`
	Lvl1   float64 `json:"lvl1"`
	Lvl100 float64 `json:"lvl100"`
	Lvl200 float64 `json:"lvl200"`
}

func getPetLevelPrices(pet map[string]interface{}, prices map[string]float64) PetLevelPrices {
	var tier string
	if pet["heldItem"] == "PET_ITEM_TIER_BOOST" {
		tier = constants.Tiers[slices.Index(constants.Tiers, pet["tier"].(string))+1]
	} else {
		tier = pet["tier"].(string)
	}
	skin := strings.ToLower(pet["skin"].(string))
	tierName := fmt.Sprintf("%s_%s", tier, skin)
	var basePrices = PetLevelPrices{
		Id: fmt.Sprintf("%s%s", tier, func() string {
			if skin != "" {
				return fmt.Sprintf("_skinned_%s", skin)
			} else {
				return ""
			}
		}()),
		Lvl1:    prices["lvl_1_"+tierName],
		Lvl100:  prices["lvl_100_"+tierName],
		Lvl200:  prices["lvl_200_"+tierName],
	}

	if skin != "" {
		return PetLevelPrices{
			Id: basePrices.Id,
			Lvl1: math.Max(prices["lvl_1_"+basePrices.Id], basePrices.Lvl1),
			Lvl100: math.Max(prices["lvl_100_"+basePrices.Id], basePrices.Lvl100),
			Lvl200: math.Max(prices["lvl_200_"+basePrices.Id], basePrices.Lvl200),
		}
	} else {
		return basePrices
	}
}

func CalculatePet(pet map[string]interface{}, prices map[string]float64) (*structs.NetworthItem, error) {
	var PetPriceData = getPetLevelPrices(pet, prices)
	pet["name"] = fmt.Sprintf("[Lvl %d] %s%s", 
		pet["level"], 
		utils.TitleCast(fmt.Sprintf("%s %s", pet["tier"], pet["type"])), 
		func() string {
			if pet["skin"] != "" {
				return " ✦"
			} else {
				return ""
			}
		}(),
	)
	if PetPriceData.Lvl1 == 0 || PetPriceData.Lvl100 == 0 {
		return nil, nil
	}

	

	if pet["level"]
}