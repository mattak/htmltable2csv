package main

import (
	"bufio"
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
)

func parseTr(_ int, tr *goquery.Selection) []string {
	var lineData []string

	tr.Find("th").Each(func(col int, td *goquery.Selection) {
		str := td.Text()
		lineData = append(lineData, str)
	})

	tr.Find("td").Each(func(col int, td *goquery.Selection) {
		str := td.Text()
		lineData = append(lineData, str)
	})

	return lineData
}

func parseTable(_ int, table *goquery.Selection) [][]string {
	var tableData [][]string

	table.Find("tr").Each(func(row int, selection *goquery.Selection) {
		tableData = append(tableData, parseTr(row, selection))
	})

	return tableData
}

func parseHtml(node *html.Node) [][]string {
	doc := goquery.NewDocumentFromNode(node)
	var arr [][]string
	targetIndex := 0

	doc.Find("table").Each(func(ntable int, selection *goquery.Selection) {
		if ntable == targetIndex {
			arr = parseTable(ntable, selection)
		}
	})

	return arr
}

func parseReader(r io.Reader) [][]string {
	doc, err := html.Parse(r)

	if err != nil {
		log.Println("error2 parseReader")
		log.Fatal(err)
	}

	return parseHtml(doc)
}

func printCsv(table [][]string) {
	writer := csv.NewWriter(os.Stdout)

	for y := 0; y < len(table); y++ {
		writer.Write(table[y])
	}

	writer.Flush()
}

func main() {
	app := cli.NewApp()
	app.Name = "htmltable2csv"
	app.Usage = "html table structure to csv"
	app.Description = "argument filename is optional. default input is from stdin."
	app.Version = "0.0.1"
	app.UsageText = "htmltable2csv [global options] [argument_filename]"
	app.Action = func(c *cli.Context) error {
		var reader io.Reader

		if len(c.Args()) <= 0 {
			reader = os.Stdin
		} else {
			filename := c.Args()[0]
			file, err := os.Open(filename)

			if err != nil {
				log.Fatal(err)
			}

			reader = bufio.NewReader(file)
		}

		data := parseReader(reader)
		printCsv(data)

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
