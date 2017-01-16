package songkick

import "net/url"

//Artists ...
type Artists struct {
	ResultsPage struct {
		Results struct {
			Artist []struct {
				URI         string `json:"uri"`
				DisplayName string `json:"displayName"`
				ID          int    `json:"id"`
				OnTourUntil string `json:"onTourUntil"`
			} `json:"artist"`
		} `json:"results"`
		TotalEntries int    `json:"totalEntries"`
		PerPage      int    `json:"perPage"`
		Page         int    `json:"page"`
		Status       string `json:"status"`
	} `json:"resultsPage"`
}

//GetArtists ... Get artists matching criteria
func GetArtists(artistName string, pageNumber int) *Artists {

	requestURL := APIUrl + "search/artists.json?query=" + url.QueryEscape(artistName) + "&page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var e *Artists
	err := get(requestURL, &e)
	if err != nil {
		return nil
	}
	return e
}
