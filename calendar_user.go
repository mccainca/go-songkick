package songkick

import "net/url"

//UserTrackedArtists ...
type UserTrackedArtists struct {
	ResultsPage struct {
		Status  string `json:"status"`
		Results struct {
			CalendarEntry []struct {
				Reason struct {
					TrackedArtist []struct {
						Identifier  []interface{} `json:"identifier"`
						OnTourUntil string        `json:"onTourUntil"`
						URI         string        `json:"uri"`
						ID          int           `json:"id"`
						DisplayName string        `json:"displayName"`
					} `json:"trackedArtist"`
				} `json:"reason"`
				Event struct {
					Type  string `json:"type"`
					Venue struct {
						Lng       float64 `json:"lng"`
						URI       string  `json:"uri"`
						ID        int     `json:"id"`
						MetroArea struct {
							URI     string `json:"uri"`
							Country struct {
								DisplayName string `json:"displayName"`
							} `json:"country"`
							ID    int `json:"id"`
							State struct {
								DisplayName string `json:"displayName"`
							} `json:"state"`
							DisplayName string `json:"displayName"`
						} `json:"metroArea"`
						Lat         float64 `json:"lat"`
						DisplayName string  `json:"displayName"`
					} `json:"venue"`
					Status      string `json:"status"`
					Performance []struct {
						Artist struct {
							Identifier  []interface{} `json:"identifier"`
							URI         string        `json:"uri"`
							ID          int           `json:"id"`
							DisplayName string        `json:"displayName"`
						} `json:"artist"`
						Billing      string `json:"billing"`
						ID           int    `json:"id"`
						BillingIndex int    `json:"billingIndex"`
						DisplayName  string `json:"displayName"`
					} `json:"performance"`
					Popularity float64 `json:"popularity"`
					Start      struct {
						Time     string `json:"time"`
						Date     string `json:"date"`
						Datetime string `json:"datetime"`
					} `json:"start"`
					AgeRestriction interface{} `json:"ageRestriction"`
					Location       struct {
						Lng  float64 `json:"lng"`
						City string  `json:"city"`
						Lat  float64 `json:"lat"`
					} `json:"location"`
					URI         string `json:"uri"`
					ID          int    `json:"id"`
					DisplayName string `json:"displayName"`
				} `json:"event"`
				CreatedAt string `json:"createdAt"`
			} `json:"calendarEntry"`
		} `json:"results"`
		PerPage      int `json:"perPage"`
		Page         int `json:"page"`
		TotalEntries int `json:"totalEntries"`
	} `json:"resultsPage"`
}

//GetUserCalendarTrackedArtists ... get user calendar by tracked artists
func GetUserCalendarTrackedArtists(userName string, sortOrder string, pageNumber int) *UserTrackedArtists {
	if !(sortOrder == "asc" || sortOrder == "desc") {
		sortOrder = "asc"
	}
	requestURL := APIUrl + "users/" + url.QueryEscape(userName) + "/calendar.json?" + "reason=tracked_artist&order=" + sortOrder + "&page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var c *UserTrackedArtists
	err := get(requestURL, &c)
	if err != nil {
		return nil
	}
	return c
}

//GetUserCalendarAttendance ... get user calendar by attendance
func GetUserCalendarAttendance(userName string, sortOrder string, pageNumber int) *UserTrackedArtists {
	if !(sortOrder == "asc" || sortOrder == "desc") {
		sortOrder = "asc"
	}
	requestURL := APIUrl + "users/" + url.QueryEscape(userName) + "/calendar.json?" + "reason=attendance&order=" + sortOrder + "&page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var c *UserTrackedArtists
	err := get(requestURL, &c)
	if err != nil {
		return nil
	}
	return c
}
