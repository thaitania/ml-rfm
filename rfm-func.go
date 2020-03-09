package rfm

// RFM is struct for contain Recency
type RFM struct {
	overallStat *Stat
	recency     *Recency
	// frequency   *[]ClusterStat
	// monetary    *[]ClusterStat
	err error
}

// RFMResult is function for show data of RFM Stat struct in Cluster
func (cls *RFM) RFMResult() error {
	return nil
}

// GetRecency is function for show recency data
func (cls *RFM) GetRecency() error {
	return nil
}

// GetFrequency is function for show frequency data
func (cls *RFM) GetFrequency() error {
	return nil
}

// GetMonetary is function for show monetary data
func (cls *RFM) GetMonetary() error {
	return nil
}
