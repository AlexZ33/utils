package files

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadCSV(filename string) ([]string, error) {
	ff, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer ff.Close()
	r := csv.NewReader(ff)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, rec := range records {
		if len(rec) < 5 {
			continue
		}

		if strings.TrimSpace(rec[4]) == "正常" {
			fmt.Printf("%s => %s, \n", strconv.Quote(strings.TrimSpace(rec[1])), strings.TrimSpace(rec[0]))
		}
	}
	return nil, nil
}
