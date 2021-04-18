package main

type Dashboard struct {
	Status string `json:"status"`
	Data   struct {
		Statistics []struct {
			Time             int     `json:"time"`
			ReportedHashrate int     `json:"reportedHashrate"`
			CurrentHashrate  float64 `json:"currentHashrate"`
			ValidShares      int     `json:"validShares"`
			InvalidShares    int     `json:"invalidShares"`
			StaleShares      int     `json:"staleShares"`
			ActiveWorkers    int     `json:"activeWorkers"`
		} `json:"statistics"`
		Workers []struct {
			Worker           string  `json:"worker"`
			Time             int     `json:"time"`
			LastSeen         int     `json:"lastSeen"`
			ReportedHashrate int     `json:"reportedHashrate"`
			CurrentHashrate  float64 `json:"currentHashrate"`
			ValidShares      int     `json:"validShares"`
			InvalidShares    int     `json:"invalidShares"`
			StaleShares      int     `json:"staleShares"`
		} `json:"workers"`
		CurrentStatistics struct {
			Time             int     `json:"time"`
			LastSeen         int     `json:"lastSeen"`
			ReportedHashrate int     `json:"reportedHashrate"`
			CurrentHashrate  float64 `json:"currentHashrate"`
			ValidShares      int     `json:"validShares"`
			InvalidShares    int     `json:"invalidShares"`
			StaleShares      int     `json:"staleShares"`
			ActiveWorkers    int     `json:"activeWorkers"`
			Unpaid           int64   `json:"unpaid"`
		} `json:"currentStatistics"`
		Settings struct {
			Email     interface{} `json:"email"`
			Monitor   int         `json:"monitor"`
			MinPayout int64       `json:"minPayout"`
		} `json:"settings"`
	} `json:"data"`
}

type SMSRequestBody struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	To        string `json:"to"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}
