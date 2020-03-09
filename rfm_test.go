package rfm

import (
	"math/rand"
	"strconv"
	"testing"
)

// TestAddRFMData is function for test add recency raw data
func TestAddRFMData(t *testing.T) {
	println("========== TestAddRFMData ==========")
	rd := &rfmData{}
	for i := 0; i < 1024; i++ {
		pts := rand.Intn(6)
		rd.AddRFMRawData("test"+strconv.Itoa(i), rand.Intn(500), pts, float64(pts*rand.Intn(2000)))
	}
	println("= TestAddRFMData Process: 1/5 : must be not error : PASS")

	_, err := rd.GetRFMRawDataByRange(5, 0)
	if err != nil {
		t.Errorf("= TestAddRFMData Process: 2/5 : must be not error : ERROR :" + " " + err.Error())
	} else {
		println("= TestAddRFMData Process: 2/5 : must be not error : PASS")
	}

	_, err = rd.GetRFMRawDataByRange(5, 100)
	if err != nil {
		t.Errorf("= TestAddRFMData Process: 3/5 : must be not error : ERROR" + " " + err.Error())
	} else {
		println("= TestAddRFMData Process: 3/5 : must be not error : PASS")
	}

	_, err = rd.GetRFMRawDataByRange(5, 1021)
	if err == nil {
		t.Errorf("= TestAddRFMData Process: 4/5 : must be error : ERROR")
	} else {
		println("= TestAddRFMData Process: 4/5 : must be error : PASS")
	}

	_, err = rd.GetRFMRawData()
	if err != nil {
		t.Errorf("= TestAddRFMData Process: 5/5 : must be not error : ERROR")
	} else {
		println("= TestAddRFMData Process: 5/5 : must be not error : PASS")
	}
}

func BenchmarkAddRFMData(b *testing.B) {
	rd := &rfmData{}
	for i := 0; i < b.N; i++ {
		pts := rand.Intn(6)
		rd.AddRFMRawData("test"+strconv.Itoa(i), rand.Intn(500), pts, float64(pts*rand.Intn(2000)))
	}
}
