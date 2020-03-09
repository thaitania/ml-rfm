package rfm

import "errors"

// Recency is struct for keep Recency
type Recency struct {
	NumCluster int
	Recency    []ClusterStat
}

// RecencyData is function for keep recency data
type RecencyData struct {
	count          int
	recencyDataRaw []recencyDataRaw
}

// RecencyDataRaw is function for keep raw recency data
type recencyDataRaw struct {
	userID  string
	recency int
}

// AddRecencyRawData is function for show raw data in n (row) after skip (row)
func (rd *RecencyData) AddRecencyRawData(userID string, recencyNum int) (*RecencyData, error) {
	if len(userID) == 0 {
		return rd, errors.New("field userID is required")
	}

	rd.recencyDataRaw = append(rd.recencyDataRaw, recencyDataRaw{
		userID:  userID,
		recency: recencyNum,
	})
	rd.count++
	return rd, nil
}

// GetRecencyRawDataRange is function for show raw data in n (row) after skip (row)
func (rd *RecencyData) GetRecencyRawDataRange(row int, skip int) (*RecencyData, error) {
	if len(rd.recencyDataRaw) < row {
		return rd, errors.New("(n) raw is more than data length")
	}
	if len(rd.recencyDataRaw) < row+skip {
		return rd, errors.New("(n) raw and skip is more than data length")
	}
	ct := &RecencyData{}
	ct.count = row
	for i := 0; i < row; i++ {
		ct.recencyDataRaw = append(ct.recencyDataRaw, rd.recencyDataRaw[skip+i])
	}
	return ct, nil
}
