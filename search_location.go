package songkick

import (
	"fmt"
	"net/url"
)

//Locations ...
type Locations struct {
	ResultsPage struct {
		Results struct {
			Location []struct {
				City struct {
					DisplayName string `json:"displayName"`
					Country     struct {
						DisplayName string `json:"displayName"`
					} `json:"country"`
					Lng float64 `json:"lng"`
					Lat float64 `json:"lat"`
				} `json:"city"`
				MetroArea struct {
					URI         string `json:"uri"`
					DisplayName string `json:"displayName"`
					Country     struct {
						DisplayName string `json:"displayName"`
					} `json:"country"`
					ID  int     `json:"id"`
					Lng float64 `json:"lng"`
					Lat float64 `json:"lat"`
				} `json:"metroArea"`
			} `json:"location"`
		} `json:"results"`
		TotalEntries int    `json:"totalEntries"`
		PerPage      int    `json:"perPage"`
		Page         int    `json:"page"`
		Status       string `json:"status"`
	} `json:"resultsPage"`
}

//GetLocationsByName ... Get locations and metro areas matching criteria
func GetLocationsByName(locationName string, pageNumber int) *Locations {

	requestURL := APIUrl + "search/locations.json?query=" + url.QueryEscape(locationName) + "&page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var l *Locations
	err := get(requestURL, &l)
	if err != nil {
		return nil
	}
	return l
}

//GetLocationsByGeoLocation ... Get locations and metro areas by latitude and longitude
func GetLocationsByGeoLocation(latitude float32, longitude float32, pageNumber int) *Locations {
	loc := fmt.Sprintf("geo:%f,%f", latitude, longitude)

	requestURL := APIUrl + "search/locations.json?location=" + loc + "&page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var l *Locations
	err := get(requestURL, &l)
	if err != nil {
		return nil
	}
	return l
}

//GetLocationsByIPLocation ... Get locations and metro areas by IP Address
func GetLocationsByIPLocation(ipAddress string, pageNumber int) *Locations {
	loc := fmt.Sprintf("ip:%s", ipAddress)

	requestURL := APIUrl + "search/locations.json?location=" + loc + "&page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var l *Locations
	err := get(requestURL, &l)
	if err != nil {
		return nil
	}
	return l
}

//GetLocationsByClientIPLocation ... Get locations and metro areas by client IP
func GetLocationsByClientIPLocation(pageNumber int) *Locations {
	loc := "clientip"

	requestURL := APIUrl + "search/locations.json?location=" + loc + "&page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var l *Locations
	err := get(requestURL, &l)
	if err != nil {
		return nil
	}
	return l
}
