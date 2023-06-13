package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

const apiUrl = "https://api.sgok.dev/doviz/api.json"

type Currency struct {
	Name   string `json:"name"`
	Rate   string `json:"rate"`
	Symbol string `json:"symbol"`
}

type Data struct {
	Time       map[string]string `json:"time"`
	Disclaimer string            `json:"disclaimer"`
	Currencies map[string]Currency `json:"currencies"`
}

func main() {
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var data Data
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	usdValStr := strings.Replace(strings.TrimSpace(data.Currencies["USD/TRY"].Rate), ",", ".", -1)
	usdVal, _ := strconv.ParseFloat(usdValStr, 64)
	eurValStr := strings.Replace(strings.TrimSpace(data.Currencies["EUR/TRY"].Rate), ",", ".", -1)
	eurVal, _ := strconv.ParseFloat(eurValStr, 64)
	goldValStr := strings.Replace(strings.TrimSpace(data.Currencies["GAU/TRY"].Rate), ".", "", -1)
	goldVal, _ := strconv.ParseFloat(strings.Replace(goldValStr, ",", ".", -1), 64)
	bitcoinValStr := strings.Replace(strings.TrimSpace(data.Currencies["BTC/USD"].Rate), ",", "", -1)
	bitcoinVal, _ := strconv.ParseFloat(bitcoinValStr, 64)

	for range time.Tick(time.Second) {
		resp, err := http.Get(apiUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}

		var data Data
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}

		newUsdValStr := strings.Replace(strings.TrimSpace(data.Currencies["USD/TRY"].Rate), ",", ".", -1)
		newUsdVal, _ := strconv.ParseFloat(newUsdValStr, 64)
		newEurValStr := strings.Replace(strings.TrimSpace(data.Currencies["EUR/TRY"].Rate), ",", ".", -1)
		newEurVal, _ := strconv.ParseFloat(newEurValStr, 64)
		newGoldValStr := strings.Replace(strings.TrimSpace(data.Currencies["GAU/TRY"].Rate), ".", "", -1)
		newGoldVal, _ := strconv.ParseFloat(strings.Replace(newGoldValStr, ",", ".", -1), 64)
		newBitcoinValStr := strings.Replace(strings.TrimSpace(data.Currencies["BTC/USD"].Rate), ",", "", -1)
		newBitcoinVal, _ := strconv.ParseFloat(newBitcoinValStr, 64)

		if newUsdVal == usdVal && newEurVal == eurVal && newGoldVal == goldVal && newBitcoinVal == bitcoinVal {
			continue
		}
	
		usdVal = newUsdVal
		eurVal = newEurVal
		goldVal = newGoldVal
		bitcoinVal = newBitcoinVal
	
		t, err := time.Parse("Jan 02, 2006 15:04:05 UTC-07:00", data.Time["updated"])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}
		now := t.Format("02 Jan 2006 15:04:05")
	
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Döviz Cinsi", "Döviz Değeri", "Tarih-Saat"})
		table.SetBorder(true) // True değeri ile çizgileri ekliyoruz
		table.SetRowLine(true)
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.FgHiRedColor},
		)
		table.SetColumnColor(
			tablewriter.Colors{},
			tablewriter.Colors{},
			tablewriter.Colors{tablewriter.FgHiGreenColor},
		)
	
		table.Append([]string{"DOLAR", fmt.Sprintf("%.4f%s", usdVal, data.Currencies["USD/TRY"].Symbol), now})
		table.Append([]string{"EURO", fmt.Sprintf("%.4f%s", eurVal, data.Currencies["EUR/TRY"].Symbol), now})
		table.Append([]string{"GOLD", fmt.Sprintf("%.2f%s", goldVal, data.Currencies["GAU/TRY"].Symbol), now})
		table.Append([]string{"BITCOIN", fmt.Sprintf("%.3f%s", bitcoinVal, data.Currencies["BTC/USD"].Symbol), now})
		fmt.Println("\033c")
		table.Render()
	}
	}