package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	onLine    = "online"
	offLine   = "offline"
	whitelist = []string{
		"12324235436",
	}
)

func httpmain() {
	onResults := make(map[string]float64)
	onCache := loop(onResults, onLine)
	offResults := make(map[string]float64)
	ofCache := loop(offResults, offLine)
	cache := make(map[string]struct{})
	for _, uid := range whitelist {
		_, ok1 := onCache[uid]
		_, ok2 := ofCache[uid]
		if ok1 && ok2 {
			cache[uid] = struct{}{}
		}
	}
	uidFinalResult := map[string]interface{}{}
	uidFinalResult["LENGTH_TOTAL"] = map[string]int{"HYDRA_TOTAL": len(onResults), "GO_TOTAL": len(offResults)}
	for midUID, offScore := range offResults {
		finalResult := map[string][]interface{}{}
		uid := strings.Split(midUID, "_")[1]

		if _, ok := cache[uid]; !ok {
			continue
		}
		if _, ok := uidFinalResult[uid]; !ok {
			uidFinalResult[uid] = finalResult
		} else {
			finalResult = uidFinalResult[uid].(map[string][]interface{})
		}

		offScore := fmt.Sprintf("%.6f", offScore)

		if onScore, ok := onResults[midUID]; ok {
			onScore := fmt.Sprintf("%.6f", onScore)
			delete(offResults, midUID)
			delete(onResults, midUID)
			if onScore == offScore {
				continue
			} else {
				finalResult["SCORE_DIFF"] = append(finalResult["SCORE_DIFF"], []string{midUID, onScore, offScore})
			}

		} else {
			finalResult["GO_UNIQUE"] = append(finalResult["GO_UNIQUE"], []string{midUID, offScore})
		}
	}
	if len(onResults) > 0 {
		for midUID, onScore := range onResults {
			uid := strings.Split(midUID, "_")[1]
			if _, ok := cache[uid]; !ok {
				continue
			}
			finalResult := map[string][]interface{}{}
			if _, ok := uidFinalResult[uid]; !ok {
				uidFinalResult[uid] = finalResult
			} else {
				finalResult = uidFinalResult[uid].(map[string][]interface{})
			}
			onScore := fmt.Sprintf("%.6f", onScore)
			finalResult["HYDRA_UNIQUE"] = append(finalResult["HYDRA_UNIQUE"], []string{midUID, onScore})
		}
	}
	fout, err := os.OpenFile("/Users/shaojun7/adowu/go/compare.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	now := time.Now().Local().Format("2006-01-02")
	io.WriteString(fout, "\n============================"+now+"==========================\n")
	if err != nil {
		fmt.Println(err)
	}
	str, err := json.Marshal(uidFinalResult)
	if err == nil {
		_, err := io.WriteString(fout, string(str))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(str))
	}
	io.WriteString(fout, "\n============================"+now+"==========================\n")
	fmt.Println("FINISHED..." + now)
}

func loop(result map[string]float64, localhost string) map[string]struct{} {
	cache := make(map[string]struct{})
	for index, uid := range whitelist {
		url := localhost + uid
		fmt.Println(index+1, url)
		resp, err := http.Get(url)
		// time.Sleep(time.Duration(2000) * time.Millisecond)
		if err != nil {
			fmt.Println(err)
			return cache
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		data := make(map[string]interface{})
		err = json.Unmarshal(body, &data)
		if err == nil {
			state := data["state"].(string)
			if state == "succ" {
				items := data["items"].([]interface{})
				skipCount := 0
				for _, itemInter := range items {
					item := itemInter.(map[string]interface{})
					mid := item["mid"].(string)
					score := item["score"].(float64)
					scoreStr := fmt.Sprintf("%f", score)
					if strings.HasPrefix(scoreStr, "16") {
						skipCount++
						continue
					}
					src := item["rc_srcs"].(string)
					if flag := strings.Contains(src, "redis-"); flag {
						src = strings.ReplaceAll(src, "redis-", "")
					}
					result[mid+"_"+uid+"_"+src] = score
				}
				if skipCount == len(items) {
					fmt.Printf("skip %s\n", uid)
				} else {
					cache[uid] = struct{}{}
				}
			}
		}
	}
	return cache
}
