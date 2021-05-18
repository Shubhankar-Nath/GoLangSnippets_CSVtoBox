package main

import (
	"encoding/csv"
	"fmt"
	"github.com/objectbox/objectbox-go/objectbox"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var count int
var waitGroup sync.WaitGroup 

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen
type CsvRecord struct {
	Id        uint64 `objectbox:"id"`
	serial    int
	firstname string
	lastname  string
	email     string
	phone     string
	message   string
}

func initObjectBox() *objectbox.ObjectBox {
	objectBox, err0 := objectbox.NewBuilder().Model(ObjectBoxModel()).Build()
	if err0 != nil {
		fmt.Println("Error happened", err0)
	}
	return objectBox
}

func processFileByLine(filename string) {
	ob := initObjectBox()
	defer ob.Close()
	box := BoxForCsvRecord(ob)

	recordFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while opening file ::", err)
		return
	}
	defer recordFile.Close()
	reader := csv.NewReader(recordFile)
	_, err2 := reader.Read() // Skipping the header
	if err2 != nil {
		fmt.Println("An error encountered while reading header::", err)
		return
	}
	waitGroup.Add(1000)
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		go writeRecord(&waitGroup, box, record)
	}
	waitGroup.Wait()
}

func writeRecord(wg *sync.WaitGroup, box *CsvRecordBox, record []string) {
	_, err := box.Put(&CsvRecord{
		serial:    0,
		firstname: record[1],
		lastname:  record[2],
		email:     record[3],
		phone:     record[4],
		message:   record[5],
	})
	if err != nil {
		count--
	}
	count++
	//fmt.Println("Record written for id:", id)
	wg.Done()
}

func main() {
	starttime := time.Now()
	fmt.Println("Processing started :", starttime)
	processFileByLine("DataFile.csv")
	fmt.Println("Total processed by now = ", count)
	fmt.Println("Processing Ended", time.Since(starttime))
}
