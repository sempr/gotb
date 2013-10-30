package gotb

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"strings"
	"time"
)

type TopParams map[string]interface{}

type TopClient struct {
	httpClient http.Client
	baseUrl	string
	appKey string
	appSecret string
}

func (cli *TopClient) Call(method string, token string, params *TopParams) error {
	pb := url.Values{} 
	
	pb.Add("app_key", cli.appKey)
	pb.Add("sign_method", "hmac")
	pb.Add("format", "json")
	pb.Add("v", "2.0")
	pb.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	pb.Add("method", method)

	if token != "" {
		pb.Add("session", token)
	}
	for k, v := range *params {
		value := fmt.Sprint(v)
		if k!= "" && value != "" {
			pb.Add(k, value)
		}
	}
	mk := make([]string, len(pb))
	i := 0
	for k, _ := range pb{
		mk[i] = k
		i++
	}
	sort.Strings(mk)
	var buffer bytes.Buffer

	for k := range mk{
		key := mk[k]
		buffer.WriteString(key)
		buffer.WriteString(pb.Get(key))
	}
	mac := hmac.New(md5.New, []byte(cli.appSecret))
	mac.Write(buffer.Bytes())
	expectedMAC := mac.Sum(nil)
	sign := hex.EncodeToString(expectedMAC)
	pb.Add("sign", strings.ToUpper(sign))
	requestUri := fmt.Sprintf("http://%s/router/rest", cli.baseUrl)
	resp, _ := cli.httpClient.PostForm(requestUri, pb)
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	return nil
}


func (cli *TopClient) Init(baseUrl string, appKey string, appSecret string) {
	cli.baseUrl = baseUrl
	cli.appKey = appKey
	cli.appSecret = appSecret
}

func (cli *TopClient) Show() {
	fmt.Println("baseUrl:", cli.baseUrl,"\nappKey:", cli.appKey,"\nappSecret:", cli.appSecret)
}
