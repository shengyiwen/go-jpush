package jpush

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	AppKey string

	MasterSecret string
}

func NewClient(appKey, masterSecret string) *Client {
	return &Client{
		AppKey:       appKey,
		MasterSecret: masterSecret,
	}
}

func (client *Client) GenerateAuthorization() string {
	return base64.StdEncoding.EncodeToString([]byte(client.AppKey + ":" + client.MasterSecret))
}

var httpClient = &http.Client{
	Timeout: time.Second * 5,
}

func (client *Client) Send(request *Request) (result *PushResult, err error) {
	body := &bytes.Buffer{}

	requestByte, err := json.Marshal(request)
	if err != nil {
		return
	}

	_, err = body.Write(requestByte)
	if err != nil && err != io.EOF {
		return
	}

	httpRequest, err := http.NewRequest("POST", PushUrl, body)
	if err != nil {
		return
	}

	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("Authorization", "Basic "+client.GenerateAuthorization())
	httpRequest.Header.Add("User-Agent", "JPush-API-Go-Client")
	httpRequest.Header.Add("Connection", "Keep-Alive")
	httpRequest.Header.Add("Accept-Charset", "UTF-8")
	httpRequest.Header.Add("Charset", "UTF-8")

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return handleResponse(resp)
}

func handleResponse(resp *http.Response) (result *PushResult, err error) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	content := string(bodyBytes)

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode >= 300 && resp.StatusCode < 400 {
			logrus.WithFields(logrus.Fields{
				"responseCode":    resp.StatusCode,
				"responseContent": content,
			}).Warn("Normal response but unexpected")
		} else {
			logrus.WithFields(logrus.Fields{
				"responseCode":    resp.StatusCode,
				"responseContent": content,
			}).Warn("Got error response.")

			switch resp.StatusCode {
			case 400:
				logrus.Error("your request param is invalid. please check them according to error msg.")
			case 401:
				logrus.Error("authentication failed. please check authentication params according to docs.")
			case 403:
				logrus.Error("request is forbidden. maybe you app key list in black list or your param is invalid.")
			case 404:
				logrus.Error("request page is not found. maybe your param is invalid.")
			case 410:
				logrus.Error("request resources is no longer in service. please according to notice on official website.")
			case 429:
				logrus.Error("too many request. please review your app key's request quota.")
			case 500:
				fallthrough
			case 502:
				fallthrough
			case 503:
				fallthrough
			case 504:
				logrus.Error("seems encountered server error. maybe jpush is in maintenance? please retry later.")
			default:
				logrus.Error("unexpected response.")
			}
		}
	}
	return FromResponse(resp.StatusCode, content)
}
