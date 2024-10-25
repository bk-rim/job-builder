package adapter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type BridgeStatusAPIAdapter struct {
}

type Passage struct {
	Bateau                    string `json:"bateau"`
	DatePassage               string `json:"date_passage"`
	FermetureALaCirculation   string `json:"fermeture_a_la_circulation"`
	ReOuvertureALaCirculation string `json:"re_ouverture_a_la_circulation"`
	TypeDeFermeture           string `json:"type_de_fermeture"`
	FermetureTotale           string `json:"fermeture_totale"`
	SensFermeture             string `json:"sens_fermeture"`
}

type Response struct {
	TotalCount int       `json:"total_count"`
	Results    []Passage `json:"results"`
}

func NewBridgeStatusAPIAdapter() *BridgeStatusAPIAdapter {
	return &BridgeStatusAPIAdapter{}
}

func (a *BridgeStatusAPIAdapter) GetBridgeData() (string, error) {

	var wg sync.WaitGroup
	wg.Add(1)

	bridgeStatus := make(chan string)
	errChan := make(chan error)

	go func() {
		defer wg.Done()

		resp, err := http.Get("https://opendata.bordeaux-metropole.fr/api/explore/v2.1/catalog/datasets/previsions_pont_chaban/records?limit=20")
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		bridgeStatus <- string(body)
	}()

	go func() {
		wg.Wait()
		close(bridgeStatus)
		close(errChan)
	}()

	select {
	case status := <-bridgeStatus:
		return status, nil
	case err := <-errChan:
		return "", err
	}
}

func (a *BridgeStatusAPIAdapter) ProcessBridgeStatus(data string) (string, error) {
	var response Response
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		return "", err
	}

	filteredResults := make([]Passage, 0)
	today := time.Now().Truncate(24 * time.Hour)

	for _, passage := range response.Results {
		passageDate, err := time.Parse("2006-01-02", passage.DatePassage)
		if err != nil {
			return "", err
		}

		if passageDate.After(today) && passageDate.Before(time.Now().AddDate(0, 0, 5)) {
			filteredResults = append(filteredResults, passage)
		}
	}

	if len(filteredResults) == 0 {
		return "le pont Chaban Delmas restera ouvert dans les 5 prochains jours", nil
	}

	for _, passage := range filteredResults {
		if passage.FermetureTotale == "oui" {
			return "le pont Chaban Delmas sera fermé le " + passage.DatePassage + " de " + passage.FermetureALaCirculation + " à " + passage.ReOuvertureALaCirculation, nil
		}
	}

	return "", nil
}
