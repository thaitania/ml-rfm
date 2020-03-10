package rfm

// Recency is struct for keep Recency
type Recency struct {
	ClusterSize int
	Cluster     []ClusterStat
	Stat        *Stat
}

// // RecencyData is function for keep recency data
// type RecencyData struct {
// 	count          int
// 	recencyDataRaw []recencyDataRaw
// }
