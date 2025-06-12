package providers

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const TotalColumns = 55

var (
	colRegexp = regexp.MustCompile(`^[A-Z]+`)
)

type Cell struct {
	R string `xml:"r,attr"`
	T string `xml:"t,attr"`
	V string `xml:"v"`
}

type RowXML struct {
	R     int    `xml:"r,attr"`
	Cells []Cell `xml:"c"`
}

type Worksheet struct {
	SheetData struct {
		Rows []RowXML `xml:"row"`
	} `xml:"sheetData"`
}

type XLSXParserProvider struct{}

func NewXLSXParserProvider() *XLSXParserProvider {
	return &XLSXParserProvider{}
}

func (p *XLSXParserProvider) colToIndex(col string) int {
	idx := 0
	for _, ch := range col {
		idx = idx*26 + int(ch-'A'+1)
	}
	return idx - 1
}

func (p *XLSXParserProvider) ParseXLSXRows(file io.ReaderAt, size int64) ([][]string, error) {
	zr, err := zip.NewReader(file, size)
	if err != nil {
		return nil, err
	}

	sharedStrings, err := p.parseSharedStrings(zr)
	if err != nil {
		return nil, err
	}

	for _, f := range zr.File {
		if strings.HasPrefix(f.Name, "xl/worksheets/sheet") {
			rows, err := p.parseWorksheet(f, sharedStrings)
			if err != nil {
				return nil, err
			}
			return rows, nil
		}
	}

	return nil, fmt.Errorf("worksheet not found")
}

func (p *XLSXParserProvider) parseSharedStrings(zr *zip.Reader) ([]string, error) {
	for _, f := range zr.File {
		if f.Name == "xl/sharedStrings.xml" {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			var s struct {
				SI []struct {
					T string `xml:"t"`
				} `xml:"si"`
			}

			if err := xml.NewDecoder(rc).Decode(&s); err != nil {
				return nil, err
			}

			shared := make([]string, len(s.SI))
			for i, si := range s.SI {
				shared[i] = si.T
			}
			return shared, nil
		}
	}
	return nil, nil
}

func (p *XLSXParserProvider) parseWorksheet(f *zip.File, sharedStrings []string) ([][]string, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	var ws Worksheet
	buf := bytes.NewBuffer(make([]byte, 0, 1024*1024))
	if _, err := io.Copy(buf, rc); err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(buf.Bytes(), &ws); err != nil {
		return nil, err
	}

	rows := make([][]string, 0, len(ws.SheetData.Rows))
	emptyStringArr := make([]string, TotalColumns)
	row := make([]string, TotalColumns)

	for _, r := range ws.SheetData.Rows {
		copy(row, emptyStringArr)

		for _, c := range r.Cells {
			colRef := colRegexp.FindString(c.R)
			if colRef == "" {
				continue
			}

			idx := p.colToIndex(colRef)
			if idx < 0 || idx >= TotalColumns {
				continue
			}

			if c.T == "s" {
				if n, err := strconv.Atoi(c.V); err == nil && n < len(sharedStrings) {
					row[idx] = sharedStrings[n]
				}
			} else {
				row[idx] = c.V
			}
		}

		rowCopy := make([]string, TotalColumns)
		copy(rowCopy, row)
		rows = append(rows, rowCopy)
	}

	return rows, nil
}
