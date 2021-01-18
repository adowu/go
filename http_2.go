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

const (
	hydraURL = "online"
	goURL    = "offline"
	logURL   = "/Users/shaojun7/adowu/go/compare.txt"
)

var (
	// UIDS ...
	UIDS = []string{
		"436457567",
	}
)

func httpMain() {
	hydraUIDResults, goUIDResults := getCompareData()
	results := make(map[string]map[string][][]string)
	for _, uid := range UIDS {
		hydraResults := hydraUIDResults[uid]
		goResults := goUIDResults[uid]
		hydraMID2Index := mid2Index(hydraResults)
		goMID2Index := mid2Index(goResults)
		// 过滤 mid 一致 得分一致
		hydraDiffes, goDiffes := compareMID(hydraMID2Index, goMID2Index, hydraResults, goResults)
		hydraSrc2Index := src2Index(hydraDiffes)
		goSrc2Index := src2Index(goDiffes)
		// 过滤 src 一致 得分一致
		hydraDiffes, goDiffes = compareSrc(hydraSrc2Index, goSrc2Index, hydraDiffes, goDiffes)

		results[uid] = map[string][][]string{
			"HY_UNIQUE": hydraDiffes,
			"GO_UNIQUE": goDiffes,
		}

	}
	printResult(results)
}

func printResult(result map[string]map[string][][]string) {
	fout, err := os.OpenFile(logURL, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	now := time.Now().Local().Format("2006-01-02-15:04:05")
	io.WriteString(fout, "\n============================"+now+"==========================\n")
	if err != nil {
		fmt.Println(err)
	}
	str, err := json.Marshal(result)
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
func compareMID(hydra2Index, go2Index map[string]int, hydraResults, goResults [][]string) ([][]string, [][]string) {
	hydraReturnDiffes := [][]string{}
	goReturnDiffes := [][]string{}
	cacheGoIndex := map[int]struct{}{}
	for compare, hydraIndex := range hydra2Index {
		hydraResult := hydraResults[hydraIndex]
		if goIndex, ok := go2Index[compare]; ok {
			goResult := goResults[goIndex]
			if hydraResult[1] != goResult[1] {
				hydraReturnDiffes = append(hydraReturnDiffes, hydraResult)
				goReturnDiffes = append(goReturnDiffes, goResult)
			}
			cacheGoIndex[goIndex] = struct{}{}
		} else {
			hydraReturnDiffes = append(hydraReturnDiffes, hydraResult)
		}
	}

	for _, goIndex := range go2Index {
		if _, ok := cacheGoIndex[goIndex]; !ok {
			goReturnDiffes = append(goReturnDiffes, goResults[goIndex])
		}
	}
	return hydraReturnDiffes, goReturnDiffes
}

func compareSrc(hydra2Index, go2Index map[string][]int, hydraResults, goResults [][]string) ([][]string, [][]string) {
	hydraReturnDiffes := [][]string{}
	goReturnDiffes := [][]string{}
	cacheGoIndex := map[int]struct{}{}
	for compare, hydraIndexes := range hydra2Index {
		for _, eachHydraIndex := range hydraIndexes {
			hydraResult := hydraResults[eachHydraIndex]
			if goIndexes, ok := go2Index[compare]; ok {
				flag := true
				for _, eachGoIndex := range goIndexes {
					goResult := goResults[eachGoIndex]
					if hydraResult[1] == goResult[1] {
						flag = false
					}
					if _, ok := cacheGoIndex[eachGoIndex]; !ok {
						goReturnDiffes = append(goReturnDiffes, goResult)
						cacheGoIndex[eachGoIndex] = struct{}{}
					}
					break
				}
				if flag {
					hydraReturnDiffes = append(hydraReturnDiffes, hydraResult)
				}
			} else {
				hydraReturnDiffes = append(hydraReturnDiffes, hydraResult)
			}
		}
	}

	for _, goIndexes := range go2Index {
		for _, goIndex := range goIndexes {
			if _, ok := cacheGoIndex[goIndex]; !ok {
				goReturnDiffes = append(goReturnDiffes, goResults[goIndex])
			}
		}

	}
	return hydraReturnDiffes, goReturnDiffes
}

func mid2Index(result [][]string) map[string]int {
	mid2index := make(map[string]int)
	for index, values := range result {
		mid2index[values[0]] = index
	}
	return mid2index
}
func src2Index(result [][]string) map[string][]int {
	src2index := make(map[string][]int)
	for index, values := range result {
		src2index[values[2]] = append(src2index[values[2]], index)
	}
	return src2index
}

func getCompareData() (map[string][][]string, map[string][][]string) {
	hydraResults := make(map[string][][]string)
	getBody(hydraResults, hydraURL)
	goResults := make(map[string][][]string)
	getBody(goResults, goURL)
	for _, uid := range UIDS {
		_, ok1 := hydraResults[uid]
		_, ok2 := goResults[uid]
		if !ok1 || !ok2 {
			delete(hydraResults, uid)
			delete(goResults, uid)
		}
	}
	return hydraResults, goResults
}
func getBody(result map[string][][]string, localhost string) {
	for index, uid := range UIDS {
		url := localhost + uid
		fmt.Println(index+1, url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
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
				eachUIDResult := [][]string{}
				for _, itemInter := range items {
					item := itemInter.(map[string]interface{})
					mid := item["mid"].(string)
					score := validScore(item)
					if strings.HasPrefix(score, "16") {
						skipCount++
						continue
					}
					src := validSrc(item)
					eachUIDResult = append(eachUIDResult, []string{mid, score, src})
				}
				if skipCount == len(items) {
					fmt.Printf("skip %s\n", uid)
				} else {
					result[uid] = eachUIDResult
				}
			}
		}
	}
}

func validScore(item map[string]interface{}) string {
	score := item["score"].(float64)
	return fmt.Sprintf("%.6f", score)
}

func validSrc(item map[string]interface{}) string {
	src := item["rc_srcs"].(string)
	if flag := strings.Contains(src, "redis-"); flag {
		src = strings.ReplaceAll(src, "redis-", "")
	}
	return src
}
