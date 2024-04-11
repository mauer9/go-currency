package http

import "encoding/xml"

type rate struct {
	XMLName     xml.Name `xml:"item"`
	FullName    string   `xml:"fullname"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
}

type rates struct {
	XMLName xml.Name `xml:"rates"`
	Date    string   `xml:"date"`
	Items   []rate   `xml:"item"`
}
