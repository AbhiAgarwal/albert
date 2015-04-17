package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SSR_GET_CLASSES_REQ struct {
	CLASS_SEARCH_REQUEST struct {
		TypeRequest      string `xml:"type"`
		LanguageCd       string `xml:"languageCd"`
		STRM             string `xml:"STRM"`
		SUBJECT          string `xml:"SUBJECT"`
		INSTITUTION      string `xml:"INSTITUTION"`
		CAMPUS           string `xml:"CAMPUS"`
		ACAD_CAREER      string `xml:"ACAD_CAREER"`
		SSR_EXACT_MATCH1 string `xml:"SSR_EXACT_MATCH1"`
		CATALOG_NBR      string `xml:"CATALOG_NBR"`
	} `xml:"CLASS_SEARCH_REQUEST"`
}

func GetRequestHTTPS() {
	timeout := time.Duration(1000 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	AlbertXML := new(SSR_GET_CLASSES_REQ)
	AlbertXML.CLASS_SEARCH_REQUEST.TypeRequest = "complete"
	AlbertXML.CLASS_SEARCH_REQUEST.LanguageCd = "eng"
	AlbertXML.CLASS_SEARCH_REQUEST.STRM = "1154"
	AlbertXML.CLASS_SEARCH_REQUEST.SUBJECT = "MA-UY"
	AlbertXML.CLASS_SEARCH_REQUEST.INSTITUTION = "NYUNV"
	AlbertXML.CLASS_SEARCH_REQUEST.CAMPUS = "WS"
	AlbertXML.CLASS_SEARCH_REQUEST.ACAD_CAREER = "UGRD"
	AlbertXML.CLASS_SEARCH_REQUEST.SSR_EXACT_MATCH1 = "E"
	AlbertXML.CLASS_SEARCH_REQUEST.CATALOG_NBR = "1024"

	output, err := xml.MarshalIndent(AlbertXML, "", " ")
	output = []byte(xml.Header + string(output))

	currentURL := "https://sis-int.nyu.edu/PSIGW/RESTListeningConnector/SSR_GET_CLASSES_R.v1/classes/get?type=complete&languageCd=eng"
	req, err := http.NewRequest("POST", currentURL, bytes.NewBuffer([]byte(output)))
	req.Header.Set("Content-Type", "application/xml")
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}

func main() {
	GetRequestHTTPS()
}
