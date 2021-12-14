package app

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type UccClient struct {
	Client      *http.Client
	Environment *Config
}

func NewHTTPClient(e *Config) (c UccClient) {
	c.Client = &http.Client{}
	c.Environment = e
	return c
}

func (c UccClient) NewHTTPRequest(method, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Authorization", "Bearer "+c.Environment.UccToken)
	return req
}

func (c UccClient) getSitelist() (sitelist Sitelist) {
	url := c.Environment.UccBaseURL + c.Environment.UccSitelistEndpoint
	req := c.NewHTTPRequest(http.MethodGet, url, nil)
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Println("Can't get response from UCC middleware.", err)
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&sitelist)
	return sitelist
}

func (c UccClient) sendStartJobSignal() (jobId string) {
	url := c.Environment.UccBaseURL + c.Environment.UccStartJobEndpoint
	req := c.NewHTTPRequest(http.MethodPut, url, nil)
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Println("Can't get response from UCC middleware while sending STARTJOB.", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (c UccClient) sendEndJobSignal(checkId string) {
	url := c.Environment.UccBaseURL + c.Environment.UccEndJobEndpoint + "?startStamp=" + checkId
	req := c.NewHTTPRequest(http.MethodPut, url, nil)
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Println("Can't get response from UCC middleware while sending ENDJOB.", err)
	}
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(resp.StatusCode, string(body))
	}
}

func (c UccClient) sendResult(checkId, payload string) {
	url := c.Environment.UccBaseURL + c.Environment.UccReportEndpoint + "?startStamp=" + checkId
	req := c.NewHTTPRequest(http.MethodPut, url, bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Println("Can't send results to UCC middleware.", err)
	}
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(resp.StatusCode, string(body))
	}
}
