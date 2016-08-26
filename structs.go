package main

import "encoding/xml"

type CDATAText struct {
	Text string `xml:",innerxml"`
}

type Base struct {
	FromUserName CDATAText
	ToUserName   CDATAText
	MsgType      CDATAText
	CreateTime   int64
	MsgId        int64 `xml:",omitempty"`
}

type Text struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Content CDATAText
}

type Image struct {
	XMLName xml.Name `xml:"xml"`
	Base
	PicUrl  CDATAText
	MediaId CDATAText
}

type Voice struct {
	XMLName xml.Name `xml:"xml"`
	Base
	MediaId     CDATAText
	Format      CDATAText
	Recognition CDATAText `xml:",omitempty"`
}

type Video struct {
	XMLName xml.Name `xml:"xml"`
	Base
	MediaId      CDATAText
	ThumbMediaId CDATAText
}

type ShortVideo struct {
	XMLName xml.Name `xml:"xml"`
	Base
	MediaId      CDATAText
	ThumbMediaId CDATAText
}

type Location struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Location_X float64
	Location_Y float64
	Scale      int
	Label      CDATAText
}
