package rfm

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// TestAddDataRawSet is function for test add recency raw data
func TestAddDataRawSet(t *testing.T) {
	rd := InitRawData()
	loop := 100000
	for i := 0; i < loop; i++ {
		pts := rand.Intn(6)
		rd.AddRFMRawData("test"+strconv.Itoa(i), rand.Intn(500), pts, float64(pts*rand.Intn(2000)))
	}
	println("= TestAddDataRawSet Process: 1/5 : must be not error : PASS")

	_, err := rd.GetRFMRawDataByRange(5, 0)
	if err != nil {
		t.Errorf("= TestAddDataRawSet Process: 2/5 : must be not error : ERROR :" + " " + err.Error())
	} else {
		println("= TestAddDataRawSet Process: 2/5 : must be not error : PASS")
	}

	_, err = rd.GetRFMRawDataByRange(35, 100)
	if err != nil {
		t.Errorf("= TestAddDataRawSet Process: 3/5 : must be not error : ERROR" + " " + err.Error())
	} else {
		println("= TestAddDataRawSet Process: 3/5 : must be not error : PASS")
	}

	_, err = rd.GetRFMRawDataByRange(5, loop)
	if err == nil {
		t.Errorf("= TestAddDataRawSet Process: 4/5 : must be error : ERROR")
	} else {
		println("= TestAddDataRawSet Process: 4/5 : must be error : PASS")
	}

	_, err = rd.GetRFMRawData()
	if err != nil {
		t.Errorf("= TestAddDataRawSet Process: 5/5 : must be not error : ERROR")
	} else {
		println("= TestAddDataRawSet Process: 5/5 : must be not error : PASS")
	}
}

func BenchmarkAddDataRawSet(b *testing.B) {
	rd := &DataRawSet{}
	for i := 0; i < b.N; i++ {
		pts := rand.Intn(6)
		rd.AddRFMRawData("test"+strconv.Itoa(i), rand.Intn(500), pts, float64(pts*rand.Intn(2000)))
	}
}

func TestGenerateRecency(t *testing.T) {
	rfm := &RFM{ClusterSize: 4}
	rd := InitRawData()
	for i := 0; i < 100000; i++ {
		pts := rand.Intn(6)
		rd.AddRFMRawData("test"+strconv.Itoa(i), rand.Intn(500), pts, float64(pts*rand.Intn(2000)))
	}

	println("= TestGenerateRecency Process: 1/5 : must be not error : PASS")
	rfm.GenerateRecency(rd)
	for _, e := range rfm.Recency.Cluster {
		println("Cluster: ", e.ClusterNum, ", Size: ", len(e.Points))
		println("Count: ", fmt.Sprintf("%v", e.Stat.Count), ",Min: ", fmt.Sprintf("%.2f", e.Stat.Min), ",Max: ", fmt.Sprintf("%.2f", e.Stat.Max),
			",STD: ", fmt.Sprintf("%.2f", e.Stat.STD),
			",Mean: ", fmt.Sprintf("%.2f", e.Stat.Mean),
			",P25: ", fmt.Sprintf("%.2f", e.Stat.P25),
			",P50: ", fmt.Sprintf("%.2f", e.Stat.P50),
			",P75: ", fmt.Sprintf("%.2f", e.Stat.P75))
	}
}

func BenchmarkGenerateRecency(b *testing.B) {
	rfm := &RFM{}
	rfm.SetClusterSize(4)
	rd := InitRawData()
	for i := 0; i < 100000; i++ {
		pts := rand.Intn(6)
		rd.AddRFMRawData("test"+strconv.Itoa(i), rand.Intn(500), pts, float64(pts*rand.Intn(2000)))
	}

	rfm.GenerateRecency(rd)
}
