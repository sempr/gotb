package gotb

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

type TopParams map[string]interface{}

type TopClient struct {
	httpClient http.Client
	baseUrl    string
	appKey     string
	appSecret  string
}

type TBAPIError struct {
	Code      int
	Msg       string
	SubCode   string
	SubMsg    string
	RequestId string
}

func (e TBAPIError) Error() string {
	return fmt.Sprintf("%s|%s|%s|%s|%s", e.Code, e.Msg, e.SubCode, e.SubMsg, e.RequestId)
}

func (cli *TopClient) Call(method string, token string, params *TopParams) (res interface{}, e error) {
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
		if k != "" && value != "" {
			pb.Add(k, value)
		}
	}
	mk := make([]string, len(pb))
	i := 0
	for k, _ := range pb {
		mk[i] = k
		i++
	}
	sort.Strings(mk)
	var buffer bytes.Buffer

	for k := range mk {
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
	resp, err := cli.httpClient.PostForm(requestUri, pb)
	if err != nil {
		e = err
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		e = err
		return
	}

	js, _ := simplejson.NewJson(bytes)

	response_key := strings.Replace(fmt.Sprintf("%s_response", strings.Replace(method, ".", "_", -1)), "taobao_", "", -1)

	res = js.Get(response_key).Interface()
	e_res := js.Get("error_response")
	if e_res.Interface() != nil {
		e = TBAPIError{
			e_res.Get("code").MustInt(),
			e_res.Get("msg").MustString(""),
			e_res.Get("sub_code").MustString(""),
			e_res.Get("sub_msg").MustString(""),
			e_res.Get("request_id").MustString(""),
		}
	}
	return
}

func (cli *TopClient) Init(baseUrl string, appKey string, appSecret string) {
	cli.baseUrl = baseUrl
	cli.appKey = appKey
	cli.appSecret = appSecret
	cli.httpClient = http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
}

func (cli *TopClient) Show() {
	fmt.Println("baseUrl:", cli.baseUrl, "\nappKey:", cli.appKey, "\nappSecret:", cli.appSecret)
}
