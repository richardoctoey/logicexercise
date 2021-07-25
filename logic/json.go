package logic

import "encoding/json"

type WallInput struct {
	Walls [][]int `json:"walls"`
}

func ParseJson(s string) WallInput {
	var res WallInput
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		return WallInput{}
	}
	return res
}