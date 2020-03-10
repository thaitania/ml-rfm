package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	rfm "github.com/thaitania/ml-rfm"
)

func main() {
	rfmStruct := rfm.RFM{ClusterSize: 4}
	rd := rfm.InitRawData()

	//Load Data from CSV file
	csvfile, err := os.Open("./dataset/OnlineRetail.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(csvfile)

	rawData := make(map[string]rfm.DataRaw)

	//Date time of day to visualization
	t2, err := time.Parse("1/2/2006 15:04", "12/9/2011 12:50")
	if err != nil {
		panic(err)
	}
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		//536388, 22469, HEART OF WICKER SMALL, 12,12/1/2010 9:59, 1.65, 16250, United Kingdom
		if record[7] == "United Kingdom" {
			//Group Customer ID to unique
			t1, err := time.Parse("1/2/2006 15:04", record[4])
			if err != nil {
				panic(err)
			}
			rcc := int(t2.Sub(t1).Hours() / 24)
			if rawData[record[6]].UserID != "" {
				if rcc < rawData[record[6]].Recency {
					tmp := rawData[record[6]]
					tmp.Recency = rcc
					rawData[record[6]] = tmp
				}
			} else {
				quantity, _ := strconv.Atoi(record[3])
				unitprice, _ := strconv.ParseFloat(record[5], 64)
				tmp := rfm.DataRaw{
					UserID:              record[6],
					Recency:             rcc,
					PurchaseTransaction: 1,
					Price:               (float64(quantity) * unitprice),
				}
				rawData[record[6]] = tmp
			}
		}
	}

	println(len(rawData))
	for i, e := range rawData {
		/*if i == "17850" {
			println(fmt.Sprintf("%v", e.UserID), fmt.Sprintf("%v", e.Recency))
		}
		if i == "13047" {
			println(fmt.Sprintf("%v", e.UserID), fmt.Sprintf("%v", e.Recency))
		}
		if i == "15291" {
			println(fmt.Sprintf("%v", e.UserID), fmt.Sprintf("%v", e.Recency))
		}*/
		rd.AddRFMRawData(i, e.Recency, e.PurchaseTransaction, e.Price)
	}

	rfmStruct.GenerateRecency(rd)
	for _, e := range rfmStruct.Recency.Cluster {
		println("Cluster: ", e.ClusterNum, ", Size: ", len(e.Points))
		println("Count: ", fmt.Sprintf("%v", e.Stat.Count), ",Min: ", fmt.Sprintf("%.2f", e.Stat.Min), ",Max: ", fmt.Sprintf("%.2f", e.Stat.Max),
			",STD: ", fmt.Sprintf("%.2f", e.Stat.STD),
			",Mean: ", fmt.Sprintf("%.2f", e.Stat.Mean),
			",P25: ", fmt.Sprintf("%.2f", e.Stat.P25),
			",P50: ", fmt.Sprintf("%.2f", e.Stat.P50),
			",P75: ", fmt.Sprintf("%.2f", e.Stat.P75))
	}
}
