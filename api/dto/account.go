package dto

type AccountDetails struct {
    Balance int64 `json:"value"`
    Counter    int64      `json:"counter"`
    LastReward RewardInfo `json:"last_rewards"`
}
