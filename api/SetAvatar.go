package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type assetsColors struct {
	Head     string
	Torso    string
	LeftArm  string `json:"left_arm"`
	LeftLeg  string `json:"left_leg"`
	RightArm string `json:"right_arm"`
	RightLeg string `json:"right_leg"`
}

type assetsItems struct {
	Face   uint32
	Hats   []uint32
	Head   uint32
	Tool   uint32
	Pants  uint32
	Shirt  uint32
	Figure uint32
	TShirt uint32
}

type assetsJSON struct {
	UserID uint32       `json:"user_id"`
	Items  assetsItems  `json:"items"`
	Colors assetsColors `json:"colors"`
}

const (
	avatarAPI = "https://api.brick-hill.com/v1/games/retrieveAvatar?id="
)

func SetAvatar(id uint32) {
	res, err := http.Get(avatarAPI + fmt.Sprint(id))
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res.Body)
	a := assetsJSON{}
	json.NewDecoder(res.Body).Decode(&a)
	fmt.Println(a)
}
