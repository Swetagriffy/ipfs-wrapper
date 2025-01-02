package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPFSClient struct {
	apiURL string
}

func NewIPFSClient(apiURL string) *IPFSClient {
	return &IPFSClient{apiURL: apiURL}
}

func (client *IPFSClient) StoreJSON(data interface{}) (string, error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %v", err)
	}

	url := fmt.Sprintf("%s/api/v0/add", client.apiURL)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	cid, ok := response["Hash"].(string)
	if !ok {
		return "", fmt.Errorf("response does not contain a valid CID")
	}

	return cid, nil
}

func (client *IPFSClient) RetrieveJSON(cid string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/api/v0/cat?arg=%s", client.apiURL, cid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from IPFS: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the JSON
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return data, nil
}

func main() {

	client := NewIPFSClient("http://localhost:5001")

	data := map[string]interface{}{
		"name":  "Alice",
		"email": "alice@example.com",
		"age":   25,
	}

	cid, err := client.StoreJSON(data)
	if err != nil {
		fmt.Printf("Error storing JSON on IPFS: %v\n", err)
		return
	}
	fmt.Printf("Data stored on IPFS with CID: %s\n", cid)

	retrievedData, err := client.RetrieveJSON(cid)
	if err != nil {
		fmt.Printf("Error retrieving JSON from IPFS: %v\n", err)
		return
	}
	fmt.Printf("Data retrieved from IPFS: %v\n", retrievedData)
}
