package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetTokenV2() string {
	url := "https://open.fxiaoke.com/cgi/corpAccessToken/get/V2"
	method := "POST"

	payload := strings.NewReader("{\n    \"appId\": \"FSAID_13195d5\",\n    \"appSecret\":\"f87c5ce5d48445aa8374b45704cec0c6\",\n    \"permanentCode\":\"EF4F4ED2F216DA6E26B70C7D37580237\"\n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	rs := []rune(string(body))
	token := string(rs[20:52])
	return (token)
}

func DataQuery(token string) string {
	url := "https://open.fxiaoke.com/cgi/crm/v2/data/query"
	method := "POST"

	payload := strings.NewReader("{" +
		"  \"corpAccessToken\":" + "\"" + token + "\"," +
		"  \"corpId\": \"FSCID_9F52763B5A8CDEBF11DA258EC782951F\"," +
		"   \"currentOpenUserId\": \"FSUID_F9B084FB9053B7432DC5389F7BFDB5D4\"," +
		"   \"data\": {" +
		"       \"dataObjectApiName\": \"SalesOrderObj\"," +
		"      \"search_query_info\":  {\"limit\":1," +
		"            \"offset\":0," +
		"           \"filters\":[" +
		"                {\"field_name\":\"life_status\",\"field_values\":[\"normal\"],\"operator\":\"EQ\"}," +
		"                {\"field_name\":\"logistics_status\",\"field_values\":[\"3\"],\"operator\":\"EQ\"}," +
		"                {\"field_name\":\"field_8vDVZ__c\",\"field_values\":[\"option1\"],\"operator\":\"EQ\"}]," +
		"                \"orders\":[{\"fieldName\":\"last_modified_time\",\"isAsc\":false}]}" +
		"}}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return (string(body))
}
