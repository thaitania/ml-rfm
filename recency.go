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
func (cls *RecencyData) AddRecencyRawData(userID string, recencyNum int) (*RecencyData, error) {
	if len(userID) == 0 {
		return cls, errors.New("field userID is required")
	}

	cls.recencyDataRaw = append(cls.recencyDataRaw, recencyDataRaw{
		userID:  userID,
		recency: recencyNum,
	})
	cls.count++
	return cls, nil
}

// GetRecencyRawData is function for show raw data in n (row) after skip (row)
func (cls *RecencyData) GetRecencyRawData(row int, skip int) (*RecencyData, error) {
	if len(cls.recencyDataRaw) < row {
		return cls, errors.New("(n) raw is more than data length")
	}
	if len(cls.recencyDataRaw) < row+skip {
		return cls, errors.New("(n) raw and skip is more than data length")
	}
	return cls, nil
}
