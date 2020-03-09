package rfm

import (
	"math/rand"
	"strconv"
	"testing"
)

// TestHello is function for test hello
func TestAddRecencyRawData(t *testing.T) {
	rd := &RecencyData{}
	for i := 0; i < 1024; i++ {
		rd.AddRecencyRawData("test"+strconv.Itoa(i), rand.Intn(500))
	}

	raw, err := rd.GetRecencyRawData(5, 0)
	if err != nil {
		t.Errorf(err.Error())
	}
	println("--- row = 5 and skip = 0 ---", "row: ", len(raw.recencyDataRaw))

	raw, err = rd.GetRecencyRawData(5, 100)
	if err != nil {
		t.Errorf(err.Error())
	}
	println("--- row = 5 and skip = 100 ---", "row: ", len(raw.recencyDataRaw))

	raw, err = rd.GetRecencyRawData(5, 1021)
	if err == nil {
		t.Errorf("This function shoud be error")
	}
	println("--- row = 5 and skip = 1021 ---", "row: ", len(raw.recencyDataRaw))
}
