package rfm

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// TestAddRecencyRawData is function for test add recency raw data
func TestAddRecencyRawData(t *testing.T) {
	rd := &RecencyData{}
	for i := 0; i < 1024; i++ {
		rd.AddRecencyRawData("test"+strconv.Itoa(i), rand.Intn(500))
	}

	raw, err := rd.GetRecencyRawDataRange(5, 0)
	if err != nil {
		t.Errorf(err.Error())
	}
	println("--- row = 5 and skip = 0 ---", "row: ", len(raw.recencyDataRaw))
	println(fmt.Sprintf("%v", raw.recencyDataRaw))

	raw, err = rd.GetRecencyRawDataRange(5, 100)
	if err != nil {
		t.Errorf(err.Error())
	}
	println("--- row = 5 and skip = 100 ---", "row: ", len(raw.recencyDataRaw))
	println(fmt.Sprintf("%v", raw.recencyDataRaw))

	raw, err = rd.GetRecencyRawDataRange(5, 1021)
	if err == nil {
		println(fmt.Sprintf("%v", raw.recencyDataRaw))
		t.Errorf("it's shoud be error: out of index")
	} else {
		println("--- row = 5 and skip = 1021 ---", "row: ", len(raw.recencyDataRaw))
		println("it's shoud be error: out of index")
	}
}
