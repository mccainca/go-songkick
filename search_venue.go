package songkick

import "net/url"

//Venues ...
type Venues struct {
	ResultsPage struct {
		Results struct {
			Venue []struct {
				ID          int    `json:"id"`
				DisplayName string `json:"displayName"`
				City        struct {
					URI         string `json:"uri"`
					DisplayName string `json:"displayName"`
					Country     struct {
						DisplayName string `json:"displayName"`
					} `json:"country"`
					ID int `json:"id"`
				} `json:"city"`
				MetroArea struct {
					URI         string `json:"uri"`
					DisplayName string `json:"displayName"`
					Country     struct {
						DisplayName string `json:"displayName"`
					} `json:"country"`
					ID int `json:"id"`
				} `json:"metroArea"`
				URI         string  `json:"uri"`
				Street      string  `json:"street"`
				Zip         string  `json:"zip"`
				Lat         float64 `json:"lat"`
				Lng         float64 `json:"lng"`
				Phone       string  `json:"phone"`
				Website     string  `json:"website"`
				Capacity    int     `json:"capacity"`
				Description string  `json:"description"`
			} `json:"venue"`
		} `json:"results"`
		TotalEntries int    `json:"totalEntries"`
		PerPage      int    `json:"perPage"`
		Page         int    `json:"page"`
		Status       string `json:"status"`
	} `json:"resultsPage"`
}

//GetVenues ... Get venues
func GetVenues(venueName string, pageNumber int) *Venues {

	requestURL := APIUrl + "search/venues.json?query=" + url.QueryEscape(venueName) + "&page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var e *Venues
	err := get(requestURL, &e)
	if err != nil {
		return nil
	}
	return e
}
