package main

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	arg_num := len(os.Args)
	if arg_num < 6 {
		fmt.Println("No Enough Args")
		return
	}

	url := os.Args[1]
	token := os.Args[2]
	from := os.Args[3]
	to := os.Args[4]
	text := os.Args[5]

	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)
	nonce := randStringRunes(8)

	sign := signature(timestampStr, nonce, token)

	url = fmt.Sprintf("%s?signature=%s&timestamp=%s&nonce=%s", url, sign, timestampStr, nonce)

	var message Text
	message.FromUserName = strToCDATA(from)
	message.ToUserName = strToCDATA(to)
	message.MsgType = strToCDATA("text")
	message.MsgId = rand.Int63()
	message.Content = strToCDATA(text)
	message.CreateTime = timestamp

	xml, err := xml.Marshal(message)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := send(url, string(xml))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("URL:", url)
	fmt.Println("--------------------")
	fmt.Println("Send Message:")

	x, err := formatXML(xml)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(x))
	}

	fmt.Println("--------------------")
	fmt.Println("Response:")

	if resp == nil {
		fmt.Println("--------------------")
		return
	}

	x, err = formatXML(resp)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(x))
	}

	fmt.Println("--------------------")
}
