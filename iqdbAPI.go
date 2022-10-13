package iqdbAPI

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

func API(input Options) string {
	if input.FilePath != "" {
		return API2(input.FilePath)
	} else {
		if input.Url != "" {
			return API1(input.Url)
		}
	}
	return "Invalid Input!"
}

func API1(target string) string {
	// Request the HTML page.

	res, err := http.PostForm("https://iqdb.org/",
		url.Values{"url": {target}})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	result := HtmlProcessor(res)
	return result
}

func API2(filename string) string {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		log.Fatal(err)
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		log.Fatal(err)
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		log.Fatal(err)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post("https://iqdb.org/", contentType, bodyBuf)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	result := HtmlProcessor(resp)
	return result

}
