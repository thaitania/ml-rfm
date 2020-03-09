package rfm

// Stat is struct for keep statistic data
type Stat struct {
	Count int     `json:"count"`
	Mean  float64 `json:"mean"`
	STD   float64 `json:"std"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	P25   float64 `json:"percent_25"`
	P50   float64 `json:"percent_50"`
	P75   float64 `json:"percent_75"`
}

// ClusterStat is struct for keep statistic data
type ClusterStat struct {
	ClusterNum int
	Stat       *Stat
}
