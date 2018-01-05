package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	accountSid := "AC3d7847bd6c1b3be4ba46156053b1e37a"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", "19173279483")
	msgData.Set("From", "+12013081917")
	//msgData.Set("Body", "\u2B50\u2764\u2B50\u2764\u2B50\u2764\u2B50\u2764\n\n\u2709" +" Save the date! "+ "\u2709\n\nLauren Pang and Thomas Curtis are delighted to invite you to our wedding.\n\nSaturday 3rd September 2016. \n\nColville Hall,\nChelmsford Road,\nWhite Roding,\nCM6 1RQ.\n\nThe Ceremony begins at 2pm.\n\nMore details will follow shortly!\n\nPlease text YES if you are saving the date and can join us or text NO if sadly, you won't be able to be with us.\n\n" + "\u2B50" + "\u2764" + "\u2B50" + "\u2764" + "\u2B50" + "\u2764" + "\u2B50" + "\u2764")
	msgData.Set("Body", "No you are :)")

	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth("SK47e5bd39198ca1710284af1d9f87e452", "YCjFoaqBgKU3lpMuuGB2HsO76dGiFeQS")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if (err == nil) {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status);
	}
}
