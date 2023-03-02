package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Quote struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func getRandomCategory() string {
	categories := []string{
		"age",
		"alone",
		"amazing",
		"anger",
		"architecture",
		"art",
		"attitude",
		"beauty",
		"best",
		"birthday",
		"business",
		"car",
		"change",
		"communications",
		"computers",
		"cool",
		"courage",
		"dad",
		"dating",
		"death",
		"design",
		"dreams",
		"education",
		"environmental",
		"equality",
		"experience",
		"failure",
		"faith",
		"family",
		"famous",
		"fear",
		"fitness",
		"food",
		"forgiveness",
		"freedom",
		"friendship",
		"funny",
		"future",
		"god",
		"good",
		"government",
		"graduation",
		"great",
		"happiness",
		"health",
		"history",
		"home",
		"hope",
		"humor",
		"imagination",
		"inspirational",
		"intelligence",
		"jealousy",
		"knowledge",
		"leadership",
		"learning",
		"legal",
		"life",
		"love",
		"marriage",
		"medical",
		"men",
		"mom",
		"money",
		"morning",
		"movies",
		"success",
	}
	rand.Seed(time.Now().UnixNano())
	return categories[rand.Intn(len(categories))]
}

func getQuote(category string) (*Quote, error) {
	apiURL := fmt.Sprintf("https://api.api-ninjas.com/v1/quotes?category=%s", category)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Key", loadEnv("API_NINJA_KEY"))
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request returned status %d %s", res.StatusCode, res.Status)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var quotes []Quote
	err = json.Unmarshal(bodyBytes, &quotes)
	if err != nil {
		return nil, err
	}

	if len(quotes) == 0 {
		return nil, fmt.Errorf("no quotes found for category %s", category)
	}

	return &quotes[0], nil
}

func translateText(text, sourceLang, targetLang string) (string, error) {
	// 翻訳APIのURLとパラメータ
	apiURL := "https://api-free.deepl.com/v2/translate"
	params := url.Values{}
	params.Set("auth_key", loadEnv("DEEPL_API_KEY"))
	params.Set("source_lang", sourceLang)
	params.Set("target_lang", targetLang)
	params.Set("text", text)

	// 翻訳APIにリクエストを送信
	res, err := http.Get(apiURL + "?" + params.Encode())
	if err != nil {
		return "", fmt.Errorf("failed to send translation request: %v", err)
	}
	defer res.Body.Close()

	// レスポンスのJSONをパースして翻訳結果を取得
	var result struct {
		Translations []struct {
			Text string `json:"text"`
		} `json:"translations"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to parse translation response: %v", err)
	}
	if len(result.Translations) == 0 {
		return "", fmt.Errorf("no translations found")
	}
	return result.Translations[0].Text, nil
}

func main() {
	category := getRandomCategory()
	quote, err := getQuote(category)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 引用文と著者名を翻訳
	translatedQuote, err := translateText(quote.Quote, "en", "ja")
	if err != nil {
		fmt.Printf("Failed to translate quote: %v", err)
		return
	}
	translatedAuthor, err := translateText(quote.Author, "en", "ja")
	if err != nil {
		fmt.Printf("Failed to translate author: %v", err)
		return
	}

	fmt.Printf("%q - %s\n", quote.Quote, quote.Author)
	// 翻訳結果を出力
	fmt.Printf("%q - %s\n", translatedQuote, translatedAuthor)
}

func loadEnv(keyName string) string {
	err := godotenv.Load(".env")
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	// .envの SAMPLE_MESSAGEを取得して、messageに代入します。
	message := os.Getenv(keyName)

	return message
}
