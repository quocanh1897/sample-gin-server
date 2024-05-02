package dto

type SystemTimeZone struct {
	Id          string `json:"value"`
	WindowsId   string `json:"windowsId"`
	Description string `json:"text"`
}

type TimeZone struct {
	Name                       string `json:"name"`
	WindowsDaylightDisplayName string `json:"windowsDaylightDisplayName"`
	WindowsDisplayName         string `json:"windowsDisplayName"`
	WindowsId                  string `json:"windowsId"`
	WindowsStandardName        string `json:"windowsStandardName"`
	TimezoneId                 string `json:"timezoneId"`
}
