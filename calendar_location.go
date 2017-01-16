package songkick

import "strconv"

//GetLocationCalendarBySKLocation ... get location calendar by SK ID
func GetLocationCalendarBySKLocation(locationID int, pageNumber int) *Calendar {
	venue := strconv.Itoa(locationID)

	requestURL := APIUrl + "metro_areas/" + venue + "/calendar.json?" + "page=" + convertPage(pageNumber) + "&apikey=" + APIKey
	var c *Calendar
	err := get(requestURL, &c)
	if err != nil {
		return nil
	}
	return c
}
