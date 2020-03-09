package rfm

import "errors"

// RFM is struct for contain Recency
type RFM struct {
	overallStat *Stat
	recency     *Recency
	// frequency   *[]ClusterStat
	// monetary    *[]ClusterStat
	err error
}

// rfmData is function for keep RFM data
type rfmData struct {
	count int
	data  []rfmDataRaw
}

// rfmDataRaw is function for keep raw RFM data
type rfmDataRaw struct {
	userID              string
	recency             int
	purchaseTransaction int
	price               float64
}

// RFMResult is function for show data of RFM Stat struct in Cluster
func (cls *RFM) RFMResult() error {
	return nil
}

// AddRFMRawData is function for show raw data in n (row) after skip (row)
func (rd *rfmData) AddRFMRawData(userID string, recencyNum int, purchaseTransaction int, price float64) (*rfmData, error) {
	if len(userID) == 0 {
		return rd, errors.New("field userID is required")
	}

	rd.data = append(rd.data, rfmDataRaw{
		userID:              userID,
		recency:             recencyNum,
		purchaseTransaction: purchaseTransaction,
		price:               price,
	})
	rd.count++
	return rd, nil
}

// GetRFMRawData is function for show raw data in n (row) after skip (row)
func (rd *rfmData) GetRFMRawData() (*rfmData, error) {
	return rd, nil
}

// GetRFMRawDataByRange is function for show raw data in n (row) after skip (row)
func (rd *rfmData) GetRFMRawDataByRange(row int, skip int) (*rfmData, error) {
	if len(rd.data) < row {
		return rd, errors.New("(n) raw is more than data length")
	}
	if len(rd.data) < row+skip {
		return rd, errors.New("(n) raw and skip is more than data length")
	}
	ct := &rfmData{}
	ct.count = row
	for i := 0; i < row; i++ {
		ct.data = append(ct.data, rd.data[skip+i])
	}
	return ct, nil
}
