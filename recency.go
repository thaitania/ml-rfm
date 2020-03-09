package rfm

// Recency is struct for keep Recency
type Recency struct {
	NumCluster int
	Recency    []ClusterStat
}

// GetRawData is function for show raw data in n (row) after skip (row)
func (cls *RFM) GetRawData(row int, skip int) error {
	return nil
}
