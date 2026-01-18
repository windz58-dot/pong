package rewards_shadow

import "strings"

func ShadowReward(rewardType string, amount float64, display string) (newType string, newAmount float64, newDisplay string, shadow bool) {
	rt := strings.ToUpper(rewardType)
	switch rt {
	case "CRYPTO_PI":
		return "CREDITS", 25, "+25 Credits", true
	case "COSMETIC":
		return "COSMETIC", 1, "Cosmetic: Common Token (Shadow)", true
	default:
		return rewardType, amount, display, false
	}
}
