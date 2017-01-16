package songkick

import (
	"fmt"
	"net/url"
	"time"
)

var artist string
var dateRange string

//Events ...
type Events struct {
	ResultsPage struct {
		Status  string `json:"status"`
		Results struct {
			Event []struct {
				Type        string  `json:"type"`
				Status      string  `json:"status"`
				Popularity  float64 `json:"popularity"`
				DisplayName string  `json:"displayName"`
				Series      struct {
					DisplayName string `json:"displayName"`
				} `json:"series,omitempty"`
				Start struct {
					Datetime interface{} `json:"datetime"`
					Time     interface{} `json:"time"`
					Date     string      `json:"date"`
				} `json:"start"`
				AgeRestriction interface{} `json:"ageRestriction"`
				End            struct {
					Datetime interface{} `json:"datetime"`
					Time     interface{} `json:"time"`
					Date     string      `json:"date"`
				} `json:"end,omitempty"`
				Location struct {
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
		PerPage        int `json:"perPage"`
		Page           int `json:"page"`
		TotalEntries   int `json:"totalEntries"`
		ClientLocation struct {
			IP          string  `json:"ip"`
			Lat         float64 `json:"lat"`
			Lng         float64 `json:"lng"`
			MetroAreaID int     `json:"metroAreaId"`
		} `json:"clientLocation"`
	} `json:"resultsPage"`
}

//GetEventsByArtist ... Get events by artist name
func GetEventsByArtist(artistName string, minDate time.Time, maxDate time.Time, pageNumber int) *Events {
	if !maxDate.IsZero() && !minDate.IsZero() {
		dateRange = fmt.Sprintf("&min_date=%s&max_date=%s", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"))
	}
	requestURL := APIUrl + "events.json?artist_name=" + url.QueryEscape(artistName) + dateRange + "&page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var e *Events
	err := get(requestURL, &e)
	if err != nil {
		return nil
	}
	return e
}

//GetEventsBySKLocation ... lookup location by SK Id
func GetEventsBySKLocation(location int, artistName string, minDate time.Time, maxDate time.Time, pageNumber int) *Events {
	loc := fmt.Sprintf("sk:%d", location)
	if artistName != "" {
		artist = fmt.Sprintf("&artist_name=%s", url.QueryEscape(artistName))
	}
	if !maxDate.IsZero() && !minDate.IsZero() {
		dateRange = fmt.Sprintf("&min_date=%s&max_date=%s", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"))
	}
	requestURL := APIUrl + "events.json?location=" + url.QueryEscape(loc) + "&page=" + convertPage(pageNumber) + artist + dateRange + "&apikey=" + APIKey
	var e *Events
	err := get(requestURL, &e)
	if err != nil {
		return nil
	}
	return e
}

//GetEventsByGeoLocation ... lookup location by latitude and longitude
func GetEventsByGeoLocation(latitude float32, longitude float32, artistName string, minDate time.Time, maxDate time.Time, pageNumber int) *Events {
	loc := fmt.Sprintf("geo:%f,%f", latitude, longitude)
	if artistName != "" {
		artist = fmt.Sprintf("&artist_name=%s", url.QueryEscape(artistName))
	}
	if !maxDate.IsZero() && !minDate.IsZero() {
		dateRange = fmt.Sprintf("&min_date=%s&max_date=%s", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"))
	}
	requestURL := APIUrl + "events.json?location=" + url.QueryEscape(loc) + "&page=" + convertPage(pageNumber) + artist + dateRange + "&apikey=" + APIKey
	var e *Events
	err := get(requestURL, &e)
	if err != nil {
		return nil
	}
	return e
}

//GetEventsByIPLocation ... lookup location by IP Address
func GetEventsByIPLocation(location string, artistName string, minDate time.Time, maxDate time.Time, pageNumber int) *Events {
	loc := fmt.Sprintf("ip:%s", location)
	if artistName != "" {
		artist = fmt.Sprintf("&artist_name=%s", url.QueryEscape(artistName))
	}
	if !maxDate.IsZero() && !minDate.IsZero() {
		dateRange = fmt.Sprintf("&min_date=%s&max_date=%s", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"))
	}
	requestURL := APIUrl + "events.json?location=" + url.QueryEscape(loc) + "&page=" + convertPage(pageNumber) + artist + dateRange + "&apikey=" + APIKey
	var e *Events
	err := get(requestURL, &e)
	if err != nil {
		return nil
	}
	return e
}

//GetEventsByClientIPLocation ... lookup location by Client IP Address. Useful for purely client side implementations.
func GetEventsByClientIPLocation(artistName string, minDate time.Time, maxDate time.Time, pageNumber int) *Events {
	loc := "clientip"
	if artistName != "" {
		artist = fmt.Sprintf("&artist_name=%s", url.QueryEscape(artistName))
	}
	if !maxDate.IsZero() && !minDate.IsZero() {
		dateRange = fmt.Sprintf("&min_date=%s&max_date=%s", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"))
	}
	requestURL := APIUrl + "events.json?location=" + url.QueryEscape(loc) + "&page=" + convertPage(pageNumber) + artist + dateRange + "&apikey=" + APIKey
	var e *Events
	err := get(requestURL, &e)
	if err != nil {
		return nil
	}
	return e
}
