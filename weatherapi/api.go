package weatherapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ruanpelissoli/weatherapi-golang/model"
)

func Call(city string) (model.Weather, error) {
	url := os.Getenv("wheatherapiurl")

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url+city, nil)

	if err != nil {
		log.Fatal(err)
		return model.Weather{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return model.Weather{}, err
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return model.Weather{}, err
	}

	res := model.Weather{}
	json.Unmarshal(responseBody, &res)

	defer resp.Body.Close()

	return res, nil
}
