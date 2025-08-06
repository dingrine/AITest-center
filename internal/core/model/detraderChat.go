package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

func (m *Model) RunDetrader(q string, examId, modelId int64) (string, error) {
	m.Logger.Printf("RunDetrader begin")
	REQUEST_URL := "https://testapi.traderwtf.top/trader/core/v4/test/trader"

	code := fmt.Sprintf("%d%s%d%s", modelId, q, examId, m.cfg.DetraderHashKey)
	hash := sha256.Sum256([]byte(code))

	var params = DetraderParams{
		ContentHash: hex.EncodeToString(hash[:]),
		Question:    q,
		TestId:      examId,
		TraderId:    modelId,
	}

	// u := uuid.New()
	client := resty.New()
	// 设置超时时间
	client.SetTimeout(20 * time.Second)
	client.SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetTransport(&http.Transport{
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		})

	res, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		// SetHeader("User-Agent", u.String()).
		SetBody(params).
		Post(REQUEST_URL)
	// m.Logger.Printf("RunDetrader request")
	if err != nil {
		log.Fatalf("这条问题：%s，发送 Post 请求报错：%s", q, err)
		return "", err
	}

	type response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Answer string `json:"answer"`
		} `json:"data"`
	}
	var result response
	if len(res.Body()) > 0 {
		err = json.Unmarshal(res.Body(), &result)
		if err != nil {
			log.Fatalf("这个问题：%s的回答:%s的json格式转换报错:%s", q, res.Body(), err)
		}
		// return result.Answer
		return result.Data.Answer, nil
	}
	// m.Logger.Printf("%s", result)
	return "", nil
}
