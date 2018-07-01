package emotions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Emo API client
type Emo struct {
	Key    string
	Client *http.Client
}

// NewClient - Create new API client
func NewClient(key string) *Emo {
	emo := new(Emo)
	emo.Key = key
	emo.Client = &http.Client{}

	return emo
}

// GetEmotions - return array of faces
func (emo *Emo) GetEmotions(photoURL string) ([]Face, error) {
	log.Printf("Get emotions: %s", photoURL)

	// create request
	body := []byte(`{"url": "` + photoURL + `"}`)
	req, _ := http.NewRequest("POST", "https://westcentralus.api.cognitive.microsoft.com/face/v1.0/detect", bytes.NewBuffer(body))
	req.Header.Add("Ocp-Apim-Subscription-Key", emo.Key)
	req.Header.Add("Content-Type", "application/json")

	// exec request
	resp, errorHTTP := emo.Client.Do(req)

	if errorHTTP != nil {
		log.Panic(errorHTTP)
		return nil, errors.New("API service unavaiable")
	}

	defer resp.Body.Close()

	// read response stream to bytes
	response, errorIO := ioutil.ReadAll(resp.Body)
	if errorIO != nil {
		log.Panic(errorIO)
		return nil, errors.New("Cannot read response")
	}

	// api service error
	if resp.StatusCode != 200 {
		var errorAPI ErrorAPI
		if jsonErrorAPI := json.Unmarshal(response, &errorAPI); jsonErrorAPI != nil {
			log.Panic(jsonErrorAPI)
			return nil, jsonErrorAPI
		}

		log.Printf("API error: %s", errorAPI.Message)
		return nil, errors.New(errorAPI.Message)
	}

	// try to parse result
	var results []Face
	if jsonErrorFaces := json.Unmarshal(response, &results); jsonErrorFaces != nil {
		log.Panic(jsonErrorFaces)
		return nil, jsonErrorFaces
	}

	log.Printf("Found %d faces", len(results))
	return results, nil
}
