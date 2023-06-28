package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Post struct {
	Contents string `json:"contents"`
}

type AllVulns interface{}
type FinalReport struct {
	Hash       string   `json:"hash"`
	Everything string   `json:"everything"`
	Vulns      []Report `json:"vulns"`
}

type Report struct {
	Type        string   `json:"type"`
	Description string   `json:"discription"`
	Position    int      `json:"position"`
	Source      AllVulns `json:"source"`
	Severity int `json:"severity"`
}
var myClient = &http.Client{Timeout: 20 * time.Second}


func getJson(url string, target interface{}, src []byte) error {
	r, err := myClient.Post(
		url,
		"application/json",
		bytes.NewBuffer(src),
	)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}