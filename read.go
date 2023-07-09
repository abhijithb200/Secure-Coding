package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"net/http"
	"time"
)



type AllVulns interface{}
type FinalReport struct {
	Hash string 	`json:"hash"`
	Everything string `json:"everything"`
	Vulns []Report	`json:"vulns"`
}

type Report struct {
	Type   string   `json:"type"`
	Description   string   `json:"discription"`
	Position int `json:"position"`
	Source AllVulns `json:"source"`
}



var myClient = &http.Client{Timeout: 20 * time.Second}

func getJson(url string, target interface{},src []byte) error {
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

type Post struct{
	Contents string `json:"contents"`
}

func main() {

	file := flag.String("f", "index.php", "Specify the file")
	hash := flag.String("hash", "", "Specify the file")
	flag.Parse()
	
	src, err := ioutil.ReadFile(*file)
	if err != nil{
		panic("[!!]File Not Found")
	}

	s := string(src)
	p := Post{
		Contents: s,
	}
	v,_ := json.Marshal(p)

	foo2 := FinalReport{}
	
	getJson("http://localhost:3000/analyze", &foo2,v)
	foo2.Hash = *hash

	time.Sleep(time.Second*5)
	f, err := os.Create("Codeguardian.json")
	encoder := json.NewEncoder(f)
	_ = encoder.Encode(foo2)
 
	fmt.Println(foo2.Everything)
}