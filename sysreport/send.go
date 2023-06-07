package sysreport

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

// SendReport produces JSON, gzips it, and sends it to the specified server.
func (r ScanResult) SendReport(server, assetID string) (err error) {
	gzr := r.gzipReport()
	postReport(server, assetID, gzr)
	return nil
}

func (r ScanResult) gzipReport() []byte {
	rjson := r.GetJSON()
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(rjson); err != nil {
		log.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}
	return b.Bytes()
}

// GetJSON converts the report into JSON and returns it.
func (r ScanResult) GetJSON() []byte {
	report, _ := json.MarshalIndent(r, "", "  ")
	return report
}

func postReport(server, assetID string, report []byte) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	b, w := createMultipartFormData(report, assetID)

	req, err := http.NewRequest("POST", server, &b)
	if err != nil {
		log.Fatal("Couldn't create HTTP request.")
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Couldn't read response body.")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if string(body) != "OK\r\n" {
		log.Fatal("Server did not return expected 'OK' string.")
	}
}

func createMultipartFormData(report []byte, assetID string) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	var fw io.Writer
	w := multipart.NewWriter(&b)

	// AgentKey
	if fw, err = w.CreateFormField("AgentKey"); err != nil {
		log.Fatalf("Error creating writer: %v", err)
	}
	if _, err = fw.Write([]byte(assetID)); err != nil {
		log.Fatalf("Error writing mime field: %v", err)
	}

	// OperatingSystem
	if fw, err = w.CreateFormField("OperatingSystem"); err != nil {
		log.Fatalf("Error creating writer: %v", err)
	}
	if _, err = fw.Write([]byte("Linux")); err != nil {
		log.Fatalf("Error writing mime field: %v", err)
	}

	// AssetId
	if fw, err = w.CreateFormField("AssetId"); err != nil {
		log.Fatalf("Error creating writer: %v", err)
	}
	if _, err = fw.Write([]byte(assetID)); err != nil {
		log.Fatalf("Error writing mime field: %v", err)
	}

	// Action
	if fw, err = w.CreateFormField("Action"); err != nil {
		log.Fatalf("Error creating writer: %v", err)
	}
	if _, err = fw.Write([]byte("ScanData")); err != nil {
		log.Fatalf("Error writing mime field: %v", err)
	}

	// Scan (JSON)
	if fw, err = w.CreateFormFile("Scan", "Scan"); err != nil {
		log.Fatalf("Error creating mime field writer: %v", err)
	}

	if _, err = fw.Write(report); err != nil {
		log.Fatalf("Error writing mime field: %v", err)
	}

	w.Close()
	return b, w
}
