package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
)

// CurrencyModel represents currency data.
type CurrencyModel struct {
	CurName         string
	CurOfficialRate *big.Float
}

type CoinMarketCapResponse struct {
	Data map[string]struct {
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

func GetCurrencyRate(cryptoName string) (string, error) {
	apiKey := os.Getenv("API_KEY")

	apiUrl := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=%s&convert=USD", cryptoName)

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		var response CoinMarketCapResponse
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			return "", err
		}

		// Extract the price of the cryptocurrency
		price := response.Data[cryptoName].Quote["USD"].Price
		priceString := fmt.Sprintf("%.15f", price)

		return priceString, nil
	}

	// Handle other HTTP status codes and errors here, if needed.
	return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
}
