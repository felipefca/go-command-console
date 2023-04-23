package service

import (
	"console/src/models"
	"console/src/notification"
	"console/src/repository"
	"encoding/json"
)

func ProcessMessages() {

	fiats, _ := repository.GetFiats()
	stocks, _ := repository.GetStocks()

	done := make(chan bool)

	for _, item := range fiats {
		go func(item models.Fiat) {
			jsonBody, err := json.Marshal(item)
			if err != nil {
				panic(err)
			}

			err = notification.SendFiatMessage(jsonBody)
			if err != nil {
				panic(err)
			}

			done <- true
		}(item)
	}

	for _, item := range stocks {
		go func(item models.Stock) {
			jsonBody, err := json.Marshal(item)
			if err != nil {
				panic(err)
			}

			err = notification.SendStockMessage(jsonBody)
			if err != nil {
				panic(err)
			}

			done <- true
		}(item)
	}

	for i := 0; i < len(fiats)+len(stocks); i++ {
		<-done
	}
}
