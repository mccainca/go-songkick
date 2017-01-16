package songkick

//Calendar ...
type Calendar struct {
	ResultsPage struct {
		Status  string `json:"status"`
		Results struct {
			Event []struct {
				Type        string  `json:"type"`
				Status      string  `json:"status"`
				Popularity  float64 `json:"popularity"`
				DisplayName string  `json:"displayName"`
				Start       struct {
					Datetime string `json:"datetime"`
					Time     string `json:"time"`
					Date     string `json:"date"`
				} `json:"start"`
				AgeRestriction interface{} `json:"ageRestriction"`
				Location       struct {
					City string  `json:"city"`
					Lat  float64 `json:"lat"`
					Lng  float64 `json:"lng"`
				} `json:"location"`
				URI         string `json:"uri"`
				ID          int    `json:"id"`
				Performance []struct {
					BillingIndex int `json:"billingIndex"`
					Artist       struct {
						DisplayName string `json:"displayName"`
						URI         string `json:"uri"`
						Identifier  []struct {
							Href string `json:"href"`
							Mbid string `json:"mbid"`
						} `json:"identifier"`
						ID int `json:"id"`
					} `json:"artist"`
					DisplayName string `json:"displayName"`
					ID          int    `json:"id"`
					Billing     string `json:"billing"`
				} `json:"performance"`
				Venue struct {
					MetroArea struct {
						DisplayName string `json:"displayName"`
						URI         string `json:"uri"`
						Country     struct {
							DisplayName string `json:"displayName"`
						} `json:"country"`
						ID    int `json:"id"`
						State struct {
							DisplayName string `json:"displayName"`
						} `json:"state"`
					} `json:"metroArea"`
					DisplayName string  `json:"displayName"`
					Lat         float64 `json:"lat"`
					Lng         float64 `json:"lng"`
					URI         string  `json:"uri"`
					ID          int     `json:"id"`
				} `json:"venue"`
			} `json:"event"`
		} `json:"results"`
		PerPage      int `json:"perPage"`
		Page         int `json:"page"`
		TotalEntries int `json:"totalEntries"`
	} `json:"resultsPage"`
}
