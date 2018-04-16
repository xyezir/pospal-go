package pospal

import (
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"time"
	"encoding/json"
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
)

const (
	// OpenAPI url
	ApiPrefix = "https://area12-win.pospal.cn:443"

	// Pospal 接口调用参数
	//Account = "233333aa"
	//AppID = "13371E7086543A4D283262CA356C94A6"
	//AppKey = "1135156339676827980"
)

type Params map[string]interface{}
type Result map[string]interface{}

type PPClient struct {
	AppID      string
	AppKey     string
	HttpClient *http.Client
}

// RawResponse the envelope response
type RawResponse map[string]interface{}

func ParseRawResponse(retBytes []byte) (RawResponse, error) {
	var jsonObject RawResponse
	err := json.Unmarshal(retBytes, &jsonObject)
	return jsonObject, err
}

func getSign(AppKey string, params Params) string {
	jsonData, _ := json.Marshal(params)
	jsonStr := string(jsonData)
	return getMd5String(AppKey + jsonStr)
}

func getMd5String(src string) string {
	h := md5.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}

func getTimestamp() int64 {
	cur := time.Now()
	return cur.UnixNano() / 1000000 //UnitNano
}

func (c *PPClient) post(rawURL string, params Params) (*http.Response, error) {
	httpClient := &http.Client{}

	var req *http.Request
	var err error
	jsonData, _ := json.Marshal(params)
	dataReader := bytes.NewReader(jsonData)
	req, err = http.NewRequest("POST", rawURL, dataReader)
	req.Header.Add("User-Agent", "openApi")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("accept-encoding", "gzip,deflate")
	req.Header.Set("time-stamp", string(getTimestamp()))
	req.Header.Set("data-signature", strings.ToUpper(getSign(c.AppKey, params)))
	//println(jsonData)

	if err != nil {
		panic(err)
	}
	return httpClient.Do(req)
}

// Invoke PosPal API
func (c *PPClient) Invoke(apiUrl string, params Params) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(ApiPrefix)
	buf.WriteString(apiUrl)
	httpURL := buf.String()

	resp, err := c.post(httpURL, params)
	defer resp.Body.Close()

	if err == nil {
		if resp.StatusCode != http.StatusOK {
			err = errors.New("http error code: " + string(resp.StatusCode) + " reason: " + resp.Status)
		}
	}

	var result []byte
	if err == nil {
		result, err = ioutil.ReadAll(resp.Body)
	}

	//println(string(result))
	return result, err
}
