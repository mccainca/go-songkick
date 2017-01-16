package songkick

import "strconv"

//GetVenueCalendarByID ... get venue calendar by SK ID
func GetVenueCalendarByID(venueID int, pageNumber int) *Calendar {
	venue := strconv.Itoa(venueID)

	requestURL := APIUrl + "venues/" + venue + "/calendar.json?" + "page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var c *Calendar
	err := get(requestURL, &c)
	if err != nil {
		return nil
	}
	return c
}
