package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type IntFace struct {
	IfIndex   int      `json:"ifindex"`
	IfName    string   `json:"ifname"`
	Glags     []string `json:"flags"`
	Mtu       int      `json:"mtu"`
	Gdisc     string   `json:"qdisc"`
	Txqlen    int      `json:"txqlen"`
	LinkType  string   `json:"link_type"`
	Address   string   `json:"address"`
	Broadcast string   `json:"broadcast"`
	AddrInfo  []struct {
		Family            string `json:"inet"`
		Local             string `json:"local"`
		Prefixlen         int    `json:"prefixlen"`
		Host              string `json:"scope"`
		Label             string `json:"label"`
		ValidLifeTime     int    `json:"valid_life_time"`
		PreferredLifeTime int    `json:"preferred_life_time"`
	} `json:"addr_info"`
}

func main() {

	out, err := exec.Command("ip", "-j", "-p", "addr").Output()
	if err != nil {
		fmt.Println(err)
	}

	ip := []IntFace{}
	jsonErr := json.Unmarshal(out, &ip)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	for i := 0; i < len(ip); i++ {
		for j := 0; j < len(ip[i].AddrInfo); j++ {
			if j == 0 {
				fmt.Printf("===========[%v]===========\n\n", ip[i].IfName)
			}
			fmt.Println(ip[i].AddrInfo[j].Local)
		}

		if i == len(ip)-1 {

			fmt.Printf("===========================\n")
		}

	}
}
