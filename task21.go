package main

import "fmt"

type Sender interface {
	xmlSend()
}

type xmlWorker struct {
}

func (xml *xmlWorker) xmlSend() {
	fmt.Println("Send xml file")
}

type jsonWorker struct {
}

func (js *jsonWorker) jsonSend() {
	fmt.Println("Send json file")
}

type jsonAdupter struct {
	jsWorker *jsonWorker
}

func (ja *jsonAdupter) xmlSend() {
	fmt.Println("Convert json to xml")
	ja.jsWorker.jsonSend()
}

type client struct {
}

func (cl *client) sendFileInXMLFormat(sender Sender) {
	sender.xmlSend()
}

// Реализовать паттерн «адаптер» на любом примере
func main() {
	client := &client{}
	xmlWorker := &xmlWorker{}
	client.sendFileInXMLFormat(xmlWorker)
	jsWorker := &jsonWorker{}
	jsAdupter := &jsonAdupter{jsWorker}
	client.sendFileInXMLFormat(jsAdupter)
}
