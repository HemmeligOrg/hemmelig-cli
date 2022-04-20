package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type secret struct {
	Id  string
	Key string
}

func CreateSecret(text string, password string, ttl string, url string) (string, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("text", text)
	bodyWriter.WriteField("password", password)
	bodyWriter.WriteField("ttl", ttl)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url+"api/secret", contentType, bodyBuf)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var secrets secret
	jsonerr := json.Unmarshal(respData, &secrets)
	if jsonerr != nil {
		fmt.Println("error:", jsonerr)
	}

	return fmt.Sprintf("%ssecret/%s/%s", url, secrets.Key, secrets.Id), nil
}
