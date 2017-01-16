package songkick

import (
	"fmt"
	"strconv"
	"time"
)

//GetArtistCalendarByID ... get artists calendar by SK ID
func GetArtistCalendarByID(artistID int, minDate time.Time, maxDate time.Time, sortOrder string, pageNumber int) *Calendar {
	artist := strconv.Itoa(artistID)
	if !maxDate.IsZero() && !minDate.IsZero() {
		dateRange = fmt.Sprintf("&min_date=%s&max_date=%s", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"))
	}
	if !(sortOrder == "asc" || sortOrder == "desc") {
		sortOrder = "asc"
	}
	requestURL := APIUrl + "artists/" + artist + "/calendar.json?" + "page=" + convertPage(pageNumber) + dateRange + "&order=" + sortOrder + "&apikey=" + APIKey
	var c *Calendar
	err := get(requestURL, &c)
	if err != nil {
		return nil
	}
	return c
}

//GetArtistCalendarByMBID ... get artists calendar by MusicBrainz ID
func GetArtistCalendarByMBID(MBID int, minDate time.Time, maxDate time.Time, sortOrder string, pageNumber int) *Calendar {
	artist := fmt.Sprintf("mbid:%s", strconv.Itoa(MBID))
	if !maxDate.IsZero() && !minDate.IsZero() {
		dateRange = fmt.Sprintf("&min_date=%s&max_date=%s", minDate.Format("2006-01-02"), maxDate.Format("2006-01-02"))
	}
	if !(sortOrder == "asc" || sortOrder == "desc") {
		sortOrder = "asc"
	}
	requestURL := APIUrl + "artists/" + artist + "/calendar.json?" + "page=" + convertPage(pageNumber) + dateRange + "&order=" + sortOrder + "&apikey=" + APIKey
	var c *Calendar
	err := get(requestURL, &c)
	if err != nil {
		return nil
	}
	return c
}
