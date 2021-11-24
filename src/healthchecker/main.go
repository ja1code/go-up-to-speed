package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

type configItem struct {
	Url      string `mapstructure:"url"`
	Priority int    `mapstructure:"priority"`
}

func main() {
	configs := readConfig()

	sortedConfigs := sortConfigs(configs)

	for _, url := range sortedConfigs {
		testUrl(url)
	}
}

func sortConfigs(configs []configItem) []string {
	var sortedConfigs []string

	sort.SliceStable(configs, func(i, j int) bool {
		return configs[i].Priority <= configs[j].Priority
	})

	for _, config := range configs {
		sortedConfigs = append(sortedConfigs, config.Url)
	}

	return sortedConfigs
}

func readConfig() []configItem {
	var configs []configItem

	config, err := os.Open("./src/healthchecker/config.json")

	if err != nil {
		fmt.Println("[ERROR] Config file not found")
		os.Exit(0)
	}

	byteValue, _ := ioutil.ReadAll(config)

	json.Unmarshal(byteValue, &configs)

	config.Close()

	return configs
}

func testUrl(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("[ERROR] Error while testing", url)
	}

	if response.StatusCode == 200 {
		fmt.Println("[SUCCESS] Successfully tested", url)
	} else {
		fmt.Println("[UNK] Different status for", url, "| STATUS:", response.StatusCode)
	}
}
