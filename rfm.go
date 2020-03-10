package rfm

import (
	"errors"
	"sort"

	"github.com/montanaflynn/stats"
	"github.com/muesli/kmeans"
)

// RFM is struct for contain Recency
type RFM struct {
	ClusterSize int
	OverallStat *Stat
	Recency     *Recency
	// frequency   *[]ClusterStat
	// monetary    *[]ClusterStat
	err error
}

// DataRawSet is function for keep RFM data
type DataRawSet struct {
	count int
	data  []DataRaw
}

// DataRaw is function for keep raw RFM data
type DataRaw struct {
	userID              string
	recency             int
	purchaseTransaction int
	price               float64
}

// // RFMResult is function for show data of RFM Stat struct in Cluster
// func (rfm *RFM) RFMResult() error {
// 	return nil
// }

// InitRawData is function for create DataRawSet
func InitRawData() *DataRawSet {
	return &DataRawSet{}
}

// SetClusterSize is function for edit cluster size
func (rfm *RFM) SetClusterSize(n int) (*RFM, error) {
	if n == 0 {
		return rfm, errors.New("cluster should be n (num cluster) > 0")
	}
	rfm.ClusterSize = n
	return rfm, nil
}

// GenerateRecency is function for process recency
func (rfm *RFM) GenerateRecency(rd *DataRawSet) (*RFM, error) {
	if rd.count != len(rd.data) {
		return rfm, errors.New("count and data length not valid")
	}

	// Create Recency Stat
	re := &Recency{}
	rstat := &Stat{Max: 0, Min: 0, Count: -1, STD: -1, Mean: -1, P25: -1, P50: -1, P75: -1}
	// println(fmt.Sprintf("%v", rstat.Max))
	for _, rd := range rd.data {
		rstat.Count++
		if float64(rd.recency) > rstat.Max {
			rstat.Max = float64(rd.recency)
		}

		if float64(rd.recency) < rstat.Min {
			rstat.Min = float64(rd.recency)
		}
	}
	mxmn := rstat.Max - rstat.Min
	var d kmeans.Points
	for _, e := range rd.data {
		d = append(d, kmeans.Point{
			0.00,
			(float64(e.recency) - rstat.Min) / mxmn,
		})
	}

	km := kmeans.New()
	clusters, err := km.Partition(d, rfm.ClusterSize)
	if err != nil {
		return rfm, err
	}
	for _, c := range clusters {
		cs := ClusterStat{}
		cs.PointsCenter = &Point{X: c.Center[0], Y: c.Center[1]}
		st := Stat{Max: (c.Points[0][1] * rstat.Max), Min: (c.Points[0][1] * rstat.Max)}
		dtl := []float64{}
		for _, e := range c.Points {
			v := e[1] * rstat.Max
			dtl = append(dtl, v)
			cs.Points = append(cs.Points, Point{
				X:     e[0],
				Y:     e[1],
				Value: (e[1] * rstat.Max),
			})

			if v > st.Max {
				st.Max = v
			}
			if v < st.Min {
				st.Min = v
			}
			st.Count++
		}
		cs.Stat = &st
		cs.Stat.Mean, _ = stats.Mean(dtl)
		cs.Stat.STD, _ = stats.StandardDeviation(dtl)
		cs.Stat.P25, _ = stats.Percentile(dtl, 25)
		cs.Stat.P50, _ = stats.Percentile(dtl, 50)
		cs.Stat.P75, _ = stats.Percentile(dtl, 75)
		// cs.ClusterNum = i
		re.ClusterSize++
		re.Cluster = append(re.Cluster, cs)
	}

	sort.Slice(re.Cluster, func(i, j int) bool {
		return re.Cluster[i].Stat.Min < re.Cluster[j].Stat.Min
	})
	rfm.Recency = re

	for i := range rfm.Recency.Cluster {
		rfm.Recency.Cluster[i].ClusterNum = i
		// println(fmt.Sprintf("%v", e.Points))
	}

	return rfm, nil
}

// AddRFMRawData is function for show raw data in n (row) after skip (row)
func (rd *DataRawSet) AddRFMRawData(userID string, recencyNum int, purchaseTransaction int, price float64) (*DataRawSet, error) {
	if len(userID) == 0 {
		return rd, errors.New("field userID is required")
	}

	rd.data = append(rd.data, DataRaw{
		userID:              userID,
		recency:             recencyNum,
		purchaseTransaction: purchaseTransaction,
		price:               price,
	})
	rd.count++
	return rd, nil
}

// GetRFMRawData is function for show raw data in n (row) after skip (row)
func (rd *DataRawSet) GetRFMRawData() (*DataRawSet, error) {
	return rd, nil
}

// GetRFMRawDataByRange is function for show raw data in n (row) after skip (row)
func (rd *DataRawSet) GetRFMRawDataByRange(row int, skip int) (*DataRawSet, error) {
	if len(rd.data) < row {
		return rd, errors.New("(n) raw is more than data length")
	}
	if len(rd.data) < row+skip {
		return rd, errors.New("(n) raw and skip is more than data length")
	}
	ct := &DataRawSet{}
	ct.count = row
	for i := 0; i < row; i++ {
		ct.data = append(ct.data, rd.data[skip+i])
	}
	return ct, nil
}
