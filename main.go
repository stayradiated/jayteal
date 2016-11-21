package main

import (
	"encoding/csv"
	"encoding/xml"
	"flag"
	"log"
	"os"

	"gopkg.in/cheggaaa/pb.v1"
)

type HttpSample struct {
	ResponseCode string `xml:"rc,attr"`
	Label        string `xml:"lb,attr"`
	Timestamp    string `xml:"ts,attr"`
	URI          string `xml:"tn,attr"`

	Lt string `xml:"lt,attr"`
	S  string `xml:"s,attr"`
	Ng string `xml:"ng,attr"`
	Dt string `xml:"dt,attr"`
	Na string `xml:"na,attr"`
	T  string `xml:"t,attr"`
	Rm string `xml:"rm,attr"`
	By string `xml:"by,attr"`
	It string `xml:"it,attr"`

	// Name    string `xml:"assertionResult>name"`
	// Failure string `xml:"assertionResult>failure"`
	// Error   string `xml:"assertionResult>error"`
	// ResponseHeader string `xml:"responseHeader"`
	// RequestHeader  string `xml:"requestHeader"`
}

func (h HttpSample) Export() []string {
	return []string{
		h.Label,
		h.ResponseCode,
		// h.ResponseHeader,
	}
}

func main() {
	var srcPath string
	var dstPath string
	var count int

	flag.StringVar(&srcPath, "src", "", "Path to source file")
	flag.StringVar(&dstPath, "dst", "", "Path to destination file")
	flag.IntVar(&count, "n", 1000000, "Number of items")
	flag.Parse()

	if srcPath == "" {
		log.Fatal("Must specify --src argument")
	}

	if dstPath == "" {
		log.Fatal("Must specify --dst argument")
	}

	src, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	dst, err := os.Create(dstPath)

	decoder := xml.NewDecoder(src)
	writer := csv.NewWriter(dst)

	bar := pb.StartNew(count)

	for {
		// read tokens from the XML document in a stream
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		// inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// if we just read a StartElement token
			// ...and it's name is "page"
			if se.Name.Local == "httpSample" {
				bar.Increment()

				var sample HttpSample
				decoder.DecodeElement(&sample, &se)

				writer.Write(sample.Export())
			}
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}

	bar.FinishPrint("The End!")
}
