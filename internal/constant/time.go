package constant

const TimeZones = `
[
  {
    "name": "Dateline Standard Time",
    "windowsDaylightDisplayName": "Dateline Daylight Time",
    "windowsDisplayName": "(UTC-12:00) International Date Line West",
    "windowsId": "Dateline Standard Time",
    "windowsStandardName": "Dateline Standard Time",
    "timezoneId": "Etc/GMT+12"
  },
  {
    "name": "UTC-11",
    "windowsDaylightDisplayName": "UTC-11",
    "windowsDisplayName": "(UTC-11:00) Coordinated Universal Time-11",
    "windowsId": "UTC-11",
    "windowsStandardName": "UTC-11",
    "timezoneId": "Etc/GMT+11"
  },
  {
    "name": "Aleutian Standard Time",
    "windowsDaylightDisplayName": "Aleutian Daylight Time",
    "windowsDisplayName": "(UTC-10:00) Aleutian Islands",
    "windowsId": "Aleutian Standard Time",
    "windowsStandardName": "Aleutian Standard Time",
    "timezoneId": "America/Adak"
  },
  {
    "name": "Hawaiian Standard Time",
    "windowsDaylightDisplayName": "Hawaiian Daylight Time",
    "windowsDisplayName": "(UTC-10:00) Hawaii",
    "windowsId": "Hawaiian Standard Time",
    "windowsStandardName": "Hawaiian Standard Time",
    "timezoneId": "Pacific/Honolulu"
  },
  {
    "name": "Marquesas Standard Time",
    "windowsDaylightDisplayName": "Marquesas Daylight Time",
    "windowsDisplayName": "(UTC-09:30) Marquesas Islands",
    "windowsId": "Marquesas Standard Time",
    "windowsStandardName": "Marquesas Standard Time",
    "timezoneId": "Pacific/Marquesas"
  },
  {
    "name": "Alaskan Standard Time",
    "windowsDaylightDisplayName": "Alaskan Daylight Time",
    "windowsDisplayName": "(UTC-09:00) Alaska",
    "windowsId": "Alaskan Standard Time",
    "windowsStandardName": "Alaskan Standard Time",
    "timezoneId": "America/Anchorage"
  },
  {
    "name": "UTC-09",
    "windowsDaylightDisplayName": "UTC-09",
    "windowsDisplayName": "(UTC-09:00) Coordinated Universal Time-09",
    "windowsId": "UTC-09",
    "windowsStandardName": "UTC-09",
    "timezoneId": "Etc/GMT+9"
  },
  {
    "name": "Pacific Standard Time (Mexico)",
    "windowsDaylightDisplayName": "Pacific Daylight Time (Mexico)",
    "windowsDisplayName": "(UTC-08:00) Baja California",
    "windowsId": "Pacific Standard Time (Mexico)",
    "windowsStandardName": "Pacific Standard Time (Mexico)",
    "timezoneId": "America/Tijuana"
  },
  {
    "name": "UTC-08",
    "windowsDaylightDisplayName": "UTC-08",
    "windowsDisplayName": "(UTC-08:00) Coordinated Universal Time-08",
    "windowsId": "UTC-08",
    "windowsStandardName": "UTC-08",
    "timezoneId": "Etc/GMT+8"
  },
  {
    "name": "Pacific Standard Time",
    "windowsDaylightDisplayName": "Pacific Daylight Time",
    "windowsDisplayName": "(UTC-08:00) Pacific Time (US & Canada)",
    "windowsId": "Pacific Standard Time",
    "windowsStandardName": "Pacific Standard Time",
    "timezoneId": "America/Los_Angeles"
  },
  {
    "name": "US Mountain Standard Time",
    "windowsDaylightDisplayName": "US Mountain Daylight Time",
    "windowsDisplayName": "(UTC-07:00) Arizona",
    "windowsId": "US Mountain Standard Time",
    "windowsStandardName": "US Mountain Standard Time",
    "timezoneId": "America/Phoenix"
  },
  {
    "name": "Mountain Standard Time (Mexico)",
    "windowsDaylightDisplayName": "Mountain Daylight Time (Mexico)",
    "windowsDisplayName": "(UTC-07:00) La Paz, Mazatlan",
    "windowsId": "Mountain Standard Time (Mexico)",
    "windowsStandardName": "Mountain Standard Time (Mexico)",
    "timezoneId": "America/Chihuahua"
  },
  {
    "name": "Mountain Standard Time",
    "windowsDaylightDisplayName": "Mountain Daylight Time",
    "windowsDisplayName": "(UTC-07:00) Mountain Time (US & Canada)",
    "windowsId": "Mountain Standard Time",
    "windowsStandardName": "Mountain Standard Time",
    "timezoneId": "America/Denver"
  },
  {
    "name": "Yukon Standard Time",
    "windowsDaylightDisplayName": "Yukon Daylight Time",
    "windowsDisplayName": "(UTC-07:00) Yukon",
    "windowsId": "Yukon Standard Time",
    "windowsStandardName": "Yukon Standard Time",
    "timezoneId": "America/Whitehorse"
  },
  {
    "name": "Central America Standard Time",
    "windowsDaylightDisplayName": "Central America Daylight Time",
    "windowsDisplayName": "(UTC-06:00) Central America",
    "windowsId": "Central America Standard Time",
    "windowsStandardName": "Central America Standard Time",
    "timezoneId": "America/Guatemala"
  },
  {
    "name": "Central Standard Time",
    "windowsDaylightDisplayName": "Central Daylight Time",
    "windowsDisplayName": "(UTC-06:00) Central Time (US & Canada)",
    "windowsId": "Central Standard Time",
    "windowsStandardName": "Central Standard Time",
    "timezoneId": "America/Chicago"
  },
  {
    "name": "Easter Island Standard Time",
    "windowsDaylightDisplayName": "Easter Island Daylight Time",
    "windowsDisplayName": "(UTC-06:00) Easter Island",
    "windowsId": "Easter Island Standard Time",
    "windowsStandardName": "Easter Island Standard Time",
    "timezoneId": "Pacific/Easter"
  },
  {
    "name": "Central Standard Time (Mexico)",
    "windowsDaylightDisplayName": "Central Daylight Time ({ Mexico)",
    "windowsDisplayName": "(UTC-06:00) Guadalajara, Mexico City, Monterrey",
    "windowsId": "Central Standard Time (Mexico)",
    "windowsStandardName": "Central Standard Time (Mexico)",
    "timezoneId": "America/Mexico_City"
  },
  {
    "name": "Canada Central Standard Time",
    "windowsDaylightDisplayName": "Canada Central Daylight Time",
    "windowsDisplayName": "(UTC-06:00) Saskatchewan",
    "windowsId": "Canada Central Standard Time",
    "windowsStandardName": "Canada Central Standard Time",
    "timezoneId": "America/Regina"
  },
  {
    "name": "SA Pacific Standard Time",
    "windowsDaylightDisplayName": "SA Pacific Daylight Time",
    "windowsDisplayName": "(UTC-05:00) Bogota, Lima, Quito, Rio Branco",
    "windowsId": "SA Pacific Standard Time",
    "windowsStandardName": "SA Pacific Standard Time",
    "timezoneId": "America/Bogota"
  },
  {
    "name": "Eastern Standard Time (Mexico)",
    "windowsDaylightDisplayName": "Eastern Daylight Time (Mexico)",
    "windowsDisplayName": "(UTC-05:00) Chetumal",
    "windowsId": "Eastern Standard Time (Mexico)",
    "windowsStandardName": "Eastern Standard Time (Mexico)",
    "timezoneId": "America/Cancun"
  },
  {
    "name": "Eastern Standard Time",
    "windowsDaylightDisplayName": "Eastern Daylight Time",
    "windowsDisplayName": "(UTC-05:00) Eastern Time (US & Canada)",
    "windowsId": "Eastern Standard Time",
    "windowsStandardName": "Eastern Standard Time",
    "timezoneId": "America/New_York"
  },
  {
    "name": "Haiti Standard Time",
    "windowsDaylightDisplayName": "Haiti Daylight Time",
    "windowsDisplayName": "(UTC-05:00) Haiti",
    "windowsId": "Haiti Standard Time",
    "windowsStandardName": "Haiti Standard Time",
    "timezoneId": "America/Port-au-Prince"
  },
  {
    "name": "Cuba Standard Time",
    "windowsDaylightDisplayName": "Cuba Daylight Time",
    "windowsDisplayName": "(UTC-05:00) Havana",
    "windowsId": "Cuba Standard Time",
    "windowsStandardName": "Cuba Standard Time",
    "timezoneId": "America/Havana"
  },
  {
    "name": "US Eastern Standard Time",
    "windowsDaylightDisplayName": "US Eastern Daylight Time",
    "windowsDisplayName": "(UTC-05:00) Indiana (East)",
    "windowsId": "US Eastern Standard Time",
    "windowsStandardName": "US Eastern Standard Time",
    "timezoneId": "America/Indianapolis"
  },
  {
    "name": "Turks And Caicos Standard Time",
    "windowsDaylightDisplayName": "Turks and Caicos Daylight Time",
    "windowsDisplayName": "(UTC-05:00) Turks and Caicos",
    "windowsId": "Turks And Caicos Standard Time",
    "windowsStandardName": "Turks and Caicos Standard Time",
    "timezoneId": "America/Grand_Turk"
  },
  {
    "name": "Paraguay Standard Time",
    "windowsDaylightDisplayName": "Paraguay Daylight Time",
    "windowsDisplayName": "(UTC-04:00) Asuncion",
    "windowsId": "Paraguay Standard Time",
    "windowsStandardName": "Paraguay Standard Time",
    "timezoneId": "America/Asuncion"
  },
  {
    "name": "Atlantic Standard Time",
    "windowsDaylightDisplayName": "Atlantic Daylight Time",
    "windowsDisplayName": "(UTC-04:00) Atlantic Time (Canada)",
    "windowsId": "Atlantic Standard Time",
    "windowsStandardName": "Atlantic Standard Time",
    "timezoneId": "America/Halifax"
  },
  {
    "name": "Venezuela Standard Time",
    "windowsDaylightDisplayName": "Venezuela Daylight Time",
    "windowsDisplayName": "(UTC-04:00) Caracas",
    "windowsId": "Venezuela Standard Time",
    "windowsStandardName": "Venezuela Standard Time",
    "timezoneId": "America/Caracas"
  },
  {
    "name": "Central Brazilian Standard Time",
    "windowsDaylightDisplayName": "Central Brazilian Daylight Time",
    "windowsDisplayName": "(UTC-04:00) Cuiaba",
    "windowsId": "Central Brazilian Standard Time",
    "windowsStandardName": "Central Brazilian Standard Time",
    "timezoneId": "America/Cuiaba"
  },
  {
    "name": "SA Western Standard Time",
    "windowsDaylightDisplayName": "SA Western Daylight Time",
    "windowsDisplayName": "(UTC-04:00) Georgetown, La Paz, Manaus, San Juan",
    "windowsId": "SA Western Standard Time",
    "windowsStandardName": "SA Western Standard Time",
    "timezoneId": "America/La_Paz"
  },
  {
    "name": "Pacific SA Standard Time",
    "windowsDaylightDisplayName": "Pacific SA Daylight Time",
    "windowsDisplayName": "(UTC-04:00) Santiago",
    "windowsId": "Pacific SA Standard Time",
    "windowsStandardName": "Pacific SA Standard Time",
    "timezoneId": "America/Santiago"
  },
  {
    "name": "Newfoundland Standard Time",
    "windowsDaylightDisplayName": "Newfoundland Daylight Time",
    "windowsDisplayName": "(UTC-03:30) Newfoundland",
    "windowsId": "Newfoundland Standard Time",
    "windowsStandardName": "Newfoundland Standard Time",
    "timezoneId": "America/St_Johns"
  },
  {
    "name": "Tocantins Standard Time",
    "windowsDaylightDisplayName": "Tocantins Daylight Time",
    "windowsDisplayName": "(UTC-03:00) Araguaina",
    "windowsId": "Tocantins Standard Time",
    "windowsStandardName": "Tocantins Standard Time",
    "timezoneId": "America/Araguaina"
  },
  {
    "name": "E. South America Standard Time",
    "windowsDaylightDisplayName": "E. South America Daylight Time",
    "windowsDisplayName": "(UTC-03:00) Brasilia",
    "windowsId": "E. South America Standard Time",
    "windowsStandardName": "E. South America Standard Time",
    "timezoneId": "America/Sao_Paulo"
  },
  {
    "name": "SA Eastern Standard Time",
    "windowsDaylightDisplayName": "SA Eastern Daylight Time",
    "windowsDisplayName": "(UTC-03:00) Cayenne, Fortaleza",
    "windowsId": "SA Eastern Standard Time",
    "windowsStandardName": "SA Eastern Standard Time",
    "timezoneId": "America/Cayenne"
  },
  {
    "name": "Argentina Standard Time",
    "windowsDaylightDisplayName": "Argentina Daylight Time",
    "windowsDisplayName": "(UTC-03:00) City of Buenos Aires",
    "windowsId": "Argentina Standard Time",
    "windowsStandardName": "Argentina Standard Time",
    "timezoneId": "America/Buenos_Aires"
  },
  {
    "name": "Greenland Standard Time",
    "windowsDaylightDisplayName": "Greenland Daylight Time",
    "windowsDisplayName": "(UTC-03:00) Greenland",
    "windowsId": "Greenland Standard Time",
    "windowsStandardName": "Greenland Standard Time",
    "timezoneId": "America/Godthab"
  },
  {
    "name": "Montevideo Standard Time",
    "windowsDaylightDisplayName": "Montevideo Daylight Time",
    "windowsDisplayName": "(UTC-03:00) Montevideo",
    "windowsId": "Montevideo Standard Time",
    "windowsStandardName": "Montevideo Standard Time",
    "timezoneId": "America/Montevideo"
  },
  {
    "name": "Magallanes Standard Time",
    "windowsDaylightDisplayName": "Magallanes Daylight Time",
    "windowsDisplayName": "(UTC-03:00) Punta Arenas",
    "windowsId": "Magallanes Standard Time",
    "windowsStandardName": "Magallanes Standard Time",
    "timezoneId": "America/Punta_Arenas"
  },
  {
    "name": "Saint Pierre Standard Time",
    "windowsDaylightDisplayName": "Saint Pierre Daylight Time",
    "windowsDisplayName": "(UTC-03:00) Saint Pierre and Miquelon",
    "windowsId": "Saint Pierre Standard Time",
    "windowsStandardName": "Saint Pierre Standard Time",
    "timezoneId": "America/Miquelon"
  },
  {
    "name": "Bahia Standard Time",
    "windowsDaylightDisplayName": "Bahia Daylight Time",
    "windowsDisplayName": "(UTC-03:00) Salvador",
    "windowsId": "Bahia Standard Time",
    "windowsStandardName": "Bahia Standard Time",
    "timezoneId": "America/Bahia"
  },
  {
    "name": "UTC-02",
    "windowsDaylightDisplayName": "UTC-02",
    "windowsDisplayName": "(UTC-02:00) Coordinated Universal Time-02",
    "windowsId": "UTC-02",
    "windowsStandardName": "UTC-02",
    "timezoneId": "Etc/GMT+2"
  },
  {
    "name": "Azores Standard Time",
    "windowsDaylightDisplayName": "Azores Daylight Time",
    "windowsDisplayName": "(UTC-01:00) Azores",
    "windowsId": "Azores Standard Time",
    "windowsStandardName": "Azores Standard Time",
    "timezoneId": "Atlantic/Azores"
  },
  {
    "name": "Cape Verde Standard Time",
    "windowsDaylightDisplayName": "Cabo Verde Daylight Time",
    "windowsDisplayName": "(UTC-01:00) Cabo Verde Is.",
    "windowsId": "Cape Verde Standard Time",
    "windowsStandardName": "Cabo Verde Standard Time",
    "timezoneId": "Atlantic/Cape_Verde"
  },
  {
    "name": "UTC",
    "windowsDaylightDisplayName": "Coordinated Universal Time",
    "windowsDisplayName": "(UTC) Coordinated Universal Time",
    "windowsId": "UTC",
    "windowsStandardName": "Coordinated Universal Time",
    "timezoneId": "Etc/UTC"
  },
  {
    "name": "GMT Standard Time",
    "windowsDaylightDisplayName": "GMT Daylight Time",
    "windowsDisplayName": "(UTC+00:00) Dublin, Edinburgh, Lisbon, London",
    "windowsId": "GMT Standard Time",
    "windowsStandardName": "GMT Standard Time",
    "timezoneId": "Europe/London"
  },
  {
    "name": "Greenwich Standard Time",
    "windowsDaylightDisplayName": "Greenwich Daylight Time",
    "windowsDisplayName": "(UTC+00:00) Monrovia, Reykjavik",
    "windowsId": "Greenwich Standard Time",
    "windowsStandardName": "Greenwich Standard Time",
    "timezoneId": "Atlantic/Reykjavik"
  },
  {
    "name": "Sao Tome Standard Time",
    "windowsDaylightDisplayName": "Sao Tome Daylight Time",
    "windowsDisplayName": "(UTC+00:00) Sao Tome",
    "windowsId": "Sao Tome Standard Time",
    "windowsStandardName": "Sao Tome Standard Time",
    "timezoneId": "Africa/Sao_Tome"
  },
  {
    "name": "Morocco Standard Time",
    "windowsDaylightDisplayName": "Morocco Daylight Time",
    "windowsDisplayName": "(UTC+01:00) Casablanca",
    "windowsId": "Morocco Standard Time",
    "windowsStandardName": "Morocco Standard Time",
    "timezoneId": "Africa/Casablanca"
  },
  {
    "name": "W. Europe Standard Time",
    "windowsDaylightDisplayName": "W. Europe Daylight Time",
    "windowsDisplayName": "(UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna",
    "windowsId": "W. Europe Standard Time",
    "windowsStandardName": "W. Europe Standard Time",
    "timezoneId": "Europe/Berlin"
  },
  {
    "name": "Central Europe Standard Time",
    "windowsDaylightDisplayName": "Central Europe { Daylight Time",
    "windowsDisplayName": "(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague",
    "windowsId": "Central Europe Standard Time",
    "windowsStandardName": "Central Europe Standard Time",
    "timezoneId": "Europe/Budapest"
  },
  {
    "name": "Romance Standard Time",
    "windowsDaylightDisplayName": "Romance Daylight Time",
    "windowsDisplayName": "(UTC+01:00) Brussels, Copenhagen, Madrid, Paris",
    "windowsId": "Romance Standard Time",
    "windowsStandardName": "Romance Standard Time",
    "timezoneId": "Europe/Paris"
  },
  {
    "name": "Central European Standard Time",
    "windowsDaylightDisplayName": "Central European Daylight { Time",
    "windowsDisplayName": "(UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb",
    "windowsId": "Central European Standard Time",
    "windowsStandardName": "Central European Standard Time",
    "timezoneId": "Europe/Warsaw"
  },
  {
    "name": "W. Central Africa Standard Time",
    "windowsDaylightDisplayName": "W. Central Africa Daylight Time",
    "windowsDisplayName": "(UTC+01:00) West Central Africa",
    "windowsId": "W. Central Africa Standard Time",
    "windowsStandardName": "W. Central Africa Standard Time",
    "timezoneId": "Africa/Lagos"
  },
  {
    "name": "GTB Standard Time",
    "windowsDaylightDisplayName": "GTB Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Athens, Bucharest",
    "windowsId": "GTB Standard Time",
    "windowsStandardName": "GTB Standard Time",
    "timezoneId": "Europe/Bucharest"
  },
  {
    "name": "Middle East Standard Time",
    "windowsDaylightDisplayName": "Middle East Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Beirut",
    "windowsId": "Middle East Standard Time",
    "windowsStandardName": "Middle East Standard Time",
    "timezoneId": "Asia/Beirut"
  },
  {
    "name": "Egypt Standard Time",
    "windowsDaylightDisplayName": "Egypt Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Cairo",
    "windowsId": "Egypt Standard Time",
    "windowsStandardName": "Egypt Standard Time",
    "timezoneId": "Africa/Cairo"
  },
  {
    "name": "E. Europe Standard Time",
    "windowsDaylightDisplayName": "E. Europe Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Chisinau",
    "windowsId": "E. Europe Standard Time",
    "windowsStandardName": "E. Europe Standard Time",
    "timezoneId": "Europe/Chisinau"
  },
  {
    "name": "Syria Standard Time",
    "windowsDaylightDisplayName": "Syria Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Damascus",
    "windowsId": "Syria Standard Time",
    "windowsStandardName": "Syria Standard Time",
    "timezoneId": "Asia/Damascus"
  },
  {
    "name": "West Bank Standard Time",
    "windowsDaylightDisplayName": "West Bank Gaza Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Gaza, Hebron",
    "windowsId": "West Bank Standard Time",
    "windowsStandardName": "West Bank Gaza Standard Time",
    "timezoneId": "Asia/Hebron"
  },
  {
    "name": "South Africa Standard Time",
    "windowsDaylightDisplayName": "South Africa Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Harare, Pretoria",
    "windowsId": "South Africa Standard Time",
    "windowsStandardName": "South Africa Standard Time",
    "timezoneId": "Africa/Johannesburg"
  },
  {
    "name": "FLE Standard Time",
    "windowsDaylightDisplayName": "FLE Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius",
    "windowsId": "FLE Standard Time",
    "windowsStandardName": "FLE Standard Time",
    "timezoneId": "Europe/Kiev"
  },
  {
    "name": "Israel Standard Time",
    "windowsDaylightDisplayName": "Jerusalem Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Jerusalem",
    "windowsId": "Israel Standard Time",
    "windowsStandardName": "Jerusalem Standard Time",
    "timezoneId": "Asia/Jerusalem"
  },
  {
    "name": "South Sudan Standard Time",
    "windowsDaylightDisplayName": "South Sudan Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Juba",
    "windowsId": "South Sudan Standard Time",
    "windowsStandardName": "South Sudan Standard Time",
    "timezoneId": "Africa/Juba"
  },
  {
    "name": "Kaliningrad Standard Time",
    "windowsDaylightDisplayName": "Russia TZ 1 Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Kaliningrad",
    "windowsId": "Kaliningrad Standard Time",
    "windowsStandardName": "Russia TZ 1 Standard Time",
    "timezoneId": "Europe/Kaliningrad"
  },
  {
    "name": "Sudan Standard Time",
    "windowsDaylightDisplayName": "Sudan Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Khartoum",
    "windowsId": "Sudan Standard Time",
    "windowsStandardName": "Sudan Standard Time",
    "timezoneId": "Africa/Khartoum"
  },
  {
    "name": "Libya Standard Time",
    "windowsDaylightDisplayName": "Libya Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Tripoli",
    "windowsId": "Libya Standard Time",
    "windowsStandardName": "Libya Standard Time",
    "timezoneId": "Africa/Tripoli"
  },
  {
    "name": "Namibia Standard Time",
    "windowsDaylightDisplayName": "Namibia Daylight Time",
    "windowsDisplayName": "(UTC+02:00) Windhoek",
    "windowsId": "Namibia Standard Time",
    "windowsStandardName": "Namibia Standard Time",
    "timezoneId": "Africa/Windhoek"
  },
  {
    "name": "Jordan Standard Time",
    "windowsDaylightDisplayName": "Jordan Daylight Time",
    "windowsDisplayName": "(UTC+03:00) Amman",
    "windowsId": "Jordan Standard Time",
    "windowsStandardName": "Jordan Standard Time",
    "timezoneId": "Asia/Amman"
  },
  {
    "name": "Arabic Standard Time",
    "windowsDaylightDisplayName": "Arabic Daylight Time",
    "windowsDisplayName": "(UTC+03:00) Baghdad",
    "windowsId": "Arabic Standard Time",
    "windowsStandardName": "Arabic Standard Time",
    "timezoneId": "Asia/Baghdad"
  },
  {
    "name": "Turkey Standard Time",
    "windowsDaylightDisplayName": "Turkey Daylight Time",
    "windowsDisplayName": "(UTC+03:00) Istanbul",
    "windowsId": "Turkey Standard Time",
    "windowsStandardName": "Turkey Standard Time",
    "timezoneId": "Europe/Istanbul"
  },
  {
    "name": "Arab Standard Time",
    "windowsDaylightDisplayName": "Arab Daylight Time",
    "windowsDisplayName": "(UTC+03:00) Kuwait, Riyadh",
    "windowsId": "Arab Standard Time",
    "windowsStandardName": "Arab Standard Time",
    "timezoneId": "Asia/Riyadh"
  },
  {
    "name": "Belarus Standard Time",
    "windowsDaylightDisplayName": "Belarus Daylight Time",
    "windowsDisplayName": "(UTC+03:00) Minsk",
    "windowsId": "Belarus Standard Time",
    "windowsStandardName": "Belarus Standard Time",
    "timezoneId": "Europe/Minsk"
  },
  {
    "name": "Russian Standard Time",
    "windowsDaylightDisplayName": "Russia TZ 2 Daylight Time",
    "windowsDisplayName": "(UTC+03:00) Moscow, St. Petersburg",
    "windowsId": "Russian Standard Time",
    "windowsStandardName": "Russia TZ 2 Standard Time",
    "timezoneId": "Europe/Moscow"
  },
  {
    "name": "E. Africa Standard Time",
    "windowsDaylightDisplayName": "E. Africa Daylight Time",
    "windowsDisplayName": "(UTC+03:00) Nairobi",
    "windowsId": "E. Africa Standard Time",
    "windowsStandardName": "E. Africa Standard Time",
    "timezoneId": "Africa/Nairobi"
  },
  {
    "name": "Volgograd Standard Time",
    "windowsDaylightDisplayName": "Volgograd Daylight Time",
    "windowsDisplayName": "(UTC+03:00) Volgograd",
    "windowsId": "Volgograd Standard Time",
    "windowsStandardName": "Volgograd Standard Time",
    "timezoneId": "Europe/Volgograd"
  },
  {
    "name": "Iran Standard Time",
    "windowsDaylightDisplayName": "Iran Daylight Time",
    "windowsDisplayName": "(UTC+03:30) Tehran",
    "windowsId": "Iran Standard Time",
    "windowsStandardName": "Iran Standard Time",
    "timezoneId": "Asia/Tehran"
  },
  {
    "name": "Arabian Standard Time",
    "windowsDaylightDisplayName": "Arabian Daylight Time",
    "windowsDisplayName": "(UTC+04:00) Abu Dhabi, Muscat",
    "windowsId": "Arabian Standard Time",
    "windowsStandardName": "Arabian Standard Time",
    "timezoneId": "Asia/Dubai"
  },
  {
    "name": "Astrakhan Standard Time",
    "windowsDaylightDisplayName": "Astrakhan Daylight Time",
    "windowsDisplayName": "(UTC+04:00) Astrakhan, Ulyanovsk",
    "windowsId": "Astrakhan Standard Time",
    "windowsStandardName": "Astrakhan Standard Time",
    "timezoneId": "Europe/Astrakhan"
  },
  {
    "name": "Azerbaijan Standard Time",
    "windowsDaylightDisplayName": "Azerbaijan Daylight Time",
    "windowsDisplayName": "(UTC+04:00) Baku",
    "windowsId": "Azerbaijan Standard Time",
    "windowsStandardName": "Azerbaijan Standard Time",
    "timezoneId": "Asia/Baku"
  },
  {
    "name": "Russia Time Zone 3",
    "windowsDaylightDisplayName": "Russia TZ 3 Daylight Time",
    "windowsDisplayName": "(UTC+04:00) Izhevsk, Samara",
    "windowsId": "Russia Time Zone 3",
    "windowsStandardName": "Russia TZ 3 Standard Time",
    "timezoneId": "Europe/Samara"
  },
  {
    "name": "Mauritius Standard Time",
    "windowsDaylightDisplayName": "Mauritius Daylight Time",
    "windowsDisplayName": "(UTC+04:00) Port Louis",
    "windowsId": "Mauritius Standard Time",
    "windowsStandardName": "Mauritius Standard Time",
    "timezoneId": "Indian/Mauritius"
  },
  {
    "name": "Saratov Standard Time",
    "windowsDaylightDisplayName": "Saratov Daylight Time",
    "windowsDisplayName": "(UTC+04:00) Saratov",
    "windowsId": "Saratov Standard Time",
    "windowsStandardName": "Saratov Standard Time",
    "timezoneId": "Europe/Saratov"
  },
  {
    "name": "Georgian Standard Time",
    "windowsDaylightDisplayName": "Georgian Daylight Time",
    "windowsDisplayName": "(UTC+04:00) Tbilisi",
    "windowsId": "Georgian Standard Time",
    "windowsStandardName": "Georgian Standard Time",
    "timezoneId": "Asia/Tbilisi"
  },
  {
    "name": "Caucasus Standard Time",
    "windowsDaylightDisplayName": "Caucasus Daylight Time",
    "windowsDisplayName": "(UTC+04:00) Yerevan",
    "windowsId": "Caucasus Standard Time",
    "windowsStandardName": "Caucasus Standard Time",
    "timezoneId": "Asia/Yerevan"
  },
  {
    "name": "Afghanistan Standard Time",
    "windowsDaylightDisplayName": "Afghanistan Daylight Time",
    "windowsDisplayName": "(UTC+04:30) Kabul",
    "windowsId": "Afghanistan Standard Time",
    "windowsStandardName": "Afghanistan Standard Time",
    "timezoneId": "Asia/Kabul"
  },
  {
    "name": "West Asia Standard Time",
    "windowsDaylightDisplayName": "West Asia Daylight Time",
    "windowsDisplayName": "(UTC+05:00) Ashgabat, Tashkent",
    "windowsId": "West Asia Standard Time",
    "windowsStandardName": "West Asia Standard Time",
    "timezoneId": "Asia/Tashkent"
  },
  {
    "name": "Ekaterinburg Standard Time",
    "windowsDaylightDisplayName": "Russia TZ 4 Daylight Time",
    "windowsDisplayName": "(UTC+05:00) Ekaterinburg",
    "windowsId": "Ekaterinburg Standard Time",
    "windowsStandardName": "Russia TZ 4 Standard Time",
    "timezoneId": "Asia/Yekaterinburg"
  },
  {
    "name": "Pakistan Standard Time",
    "windowsDaylightDisplayName": "Pakistan Daylight Time",
    "windowsDisplayName": "(UTC+05:00) Islamabad, Karachi",
    "windowsId": "Pakistan Standard Time",
    "windowsStandardName": "Pakistan Standard Time",
    "timezoneId": "Asia/Karachi"
  },
  {
    "name": "Qyzylorda Standard Time",
    "windowsDaylightDisplayName": "Qyzylorda Daylight Time",
    "windowsDisplayName": "(UTC+05:00) Qyzylorda",
    "windowsId": "Qyzylorda Standard Time",
    "windowsStandardName": "Qyzylorda Standard Time",
    "timezoneId": "Asia/Qyzylorda"
  },
  {
    "name": "India Standard Time",
    "windowsDaylightDisplayName": "India Daylight Time",
    "windowsDisplayName": "(UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi",
    "windowsId": "India Standard Time",
    "windowsStandardName": "India Standard Time",
    "timezoneId": "Asia/Calcutta"
  },
  {
    "name": "Sri Lanka Standard Time",
    "windowsDaylightDisplayName": "Sri Lanka Daylight Time",
    "windowsDisplayName": "(UTC+05:30) Sri Jayawardenepura",
    "windowsId": "Sri Lanka Standard Time",
    "windowsStandardName": "Sri Lanka Standard Time",
    "timezoneId": "Asia/Colombo"
  },
  {
    "name": "Nepal Standard Time",
    "windowsDaylightDisplayName": "Nepal Daylight Time",
    "windowsDisplayName": "(UTC+05:45) Kathmandu",
    "windowsId": "Nepal Standard Time",
    "windowsStandardName": "Nepal Standard Time",
    "timezoneId": "Asia/Katmandu"
  },
  {
    "name": "Central Asia Standard Time",
    "windowsDaylightDisplayName": "Central Asia Daylight Time",
    "windowsDisplayName": "(UTC+06:00) Astana",
    "windowsId": "Central Asia Standard Time",
    "windowsStandardName": "Central Asia Standard Time",
    "timezoneId": "Asia/Almaty"
  },
  {
    "name": "Bangladesh Standard Time",
    "windowsDaylightDisplayName": "Bangladesh Daylight Time",
    "windowsDisplayName": "(UTC+06:00) Dhaka",
    "windowsId": "Bangladesh Standard Time",
    "windowsStandardName": "Bangladesh Standard Time",
    "timezoneId": "Asia/Dhaka"
  },
  {
    "name": "Omsk Standard Time",
    "windowsDaylightDisplayName": "Omsk Daylight Time",
    "windowsDisplayName": "(UTC+06:00) Omsk",
    "windowsId": "Omsk Standard Time",
    "windowsStandardName": "Omsk Standard Time",
    "timezoneId": "Asia/Omsk"
  },
  {
    "name": "Myanmar Standard Time",
    "windowsDaylightDisplayName": "Myanmar Daylight Time",
    "windowsDisplayName": "(UTC+06:30) Yangon (Rangoon)",
    "windowsId": "Myanmar Standard Time",
    "windowsStandardName": "Myanmar Standard Time",
    "timezoneId": "Asia/Rangoon"
  },
  {
    "name": "SE Asia Standard Time",
    "windowsDaylightDisplayName": "SE Asia Daylight Time",
    "windowsDisplayName": "(UTC+07:00) Bangkok, Hanoi, Jakarta",
    "windowsId": "SE Asia Standard Time",
    "windowsStandardName": "SE Asia Standard Time",
    "timezoneId": "Asia/Bangkok"
  },
  {
    "name": "Altai Standard Time",
    "windowsDaylightDisplayName": "Altai Daylight Time",
    "windowsDisplayName": "(UTC+07:00) Barnaul, Gorno-Altaysk",
    "windowsId": "Altai Standard Time",
    "windowsStandardName": "Altai Standard Time",
    "timezoneId": "Asia/Barnaul"
  },
  {
    "name": "W. Mongolia Standard Time",
    "windowsDaylightDisplayName": "W. Mongolia Daylight Time",
    "windowsDisplayName": "(UTC+07:00) Hovd",
    "windowsId": "W. Mongolia Standard Time",
    "windowsStandardName": "W. Mongolia Standard Time",
    "timezoneId": "Asia/Hovd"
  },
  {
    "name": "North Asia Standard Time",
    "windowsDaylightDisplayName": "Russia TZ 6 Daylight Time",
    "windowsDisplayName": "(UTC+07:00) Krasnoyarsk",
    "windowsId": "North Asia Standard Time",
    "windowsStandardName": "Russia TZ 6 Standard Time",
    "timezoneId": "Asia/Krasnoyarsk"
  },
  {
    "name": "N. Central Asia Standard Time",
    "windowsDaylightDisplayName": "Novosibirsk Daylight Time",
    "windowsDisplayName": "(UTC+07:00) Novosibirsk",
    "windowsId": "N. Central Asia Standard Time",
    "windowsStandardName": "Novosibirsk Standard Time",
    "timezoneId": "Asia/Novosibirsk"
  },
  {
    "name": "Tomsk Standard Time",
    "windowsDaylightDisplayName": "Tomsk Daylight Time",
    "windowsDisplayName": "(UTC+07:00) Tomsk",
    "windowsId": "Tomsk Standard Time",
    "windowsStandardName": "Tomsk Standard Time",
    "timezoneId": "Asia/Tomsk"
  },
  {
    "name": "China Standard Time",
    "windowsDaylightDisplayName": "China Daylight Time",
    "windowsDisplayName": "(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi",
    "windowsId": "China Standard Time",
    "windowsStandardName": "China Standard Time",
    "timezoneId": "Asia/Shanghai"
  },
  {
    "name": "North Asia East Standard Time",
    "windowsDaylightDisplayName": "Russia TZ 7 Daylight Time",
    "windowsDisplayName": "(UTC+08:00) Irkutsk",
    "windowsId": "North Asia East Standard Time",
    "windowsStandardName": "Russia TZ 7 Standard Time",
    "timezoneId": "Asia/Irkutsk"
  },
  {
    "name": "Singapore Standard Time",
    "windowsDaylightDisplayName": "Malay Peninsula Daylight Time",
    "windowsDisplayName": "(UTC+08:00) Kuala Lumpur, Singapore",
    "windowsId": "Singapore Standard Time",
    "windowsStandardName": "Malay Peninsula Standard Time",
    "timezoneId": "Asia/Singapore"
  },
  {
    "name": "W. Australia Standard Time",
    "windowsDaylightDisplayName": "W. Australia Daylight Time",
    "windowsDisplayName": "(UTC+08:00) Perth",
    "windowsId": "W. Australia Standard Time",
    "windowsStandardName": "W. Australia Standard Time",
    "timezoneId": "Australia/Perth"
  },
  {
    "name": "Taipei Standard Time",
    "windowsDaylightDisplayName": "Taipei Daylight Time",
    "windowsDisplayName": "(UTC+08:00) Taipei",
    "windowsId": "Taipei Standard Time",
    "windowsStandardName": "Taipei Standard Time",
    "timezoneId": "Asia/Taipei"
  },
  {
    "name": "Ulaanbaatar Standard Time",
    "windowsDaylightDisplayName": "Ulaanbaatar Daylight Time",
    "windowsDisplayName": "(UTC+08:00) Ulaanbaatar",
    "windowsId": "Ulaanbaatar Standard Time",
    "windowsStandardName": "Ulaanbaatar Standard Time",
    "timezoneId": "Asia/Ulaanbaatar"
  },
  {
    "name": "Aus Central W. Standard Time",
    "windowsDaylightDisplayName": "Aus Central W. Daylight Time",
    "windowsDisplayName": "(UTC+08:45) Eucla",
    "windowsId": "Aus Central W. Standard Time",
    "windowsStandardName": "Aus Central W. Standard Time",
    "timezoneId": "Australia/Eucla"
  },
  {
    "name": "Transbaikal Standard Time",
    "windowsDaylightDisplayName": "Transbaikal Daylight Time",
    "windowsDisplayName": "(UTC+09:00) Chita",
    "windowsId": "Transbaikal Standard Time",
    "windowsStandardName": "Transbaikal Standard Time",
    "timezoneId": "Asia/Chita"
  },
  {
    "name": "Tokyo Standard Time",
    "windowsDaylightDisplayName": "Tokyo Daylight Time",
    "windowsDisplayName": "(UTC+09:00) Osaka, Sapporo, Tokyo",
    "windowsId": "Tokyo Standard Time",
    "windowsStandardName": "Tokyo Standard Time",
    "timezoneId": "Asia/Tokyo"
  },
  {
    "name": "North Korea Standard Time",
    "windowsDaylightDisplayName": "North Korea Daylight Time",
    "windowsDisplayName": "(UTC+09:00) Pyongyang",
    "windowsId": "North Korea Standard Time",
    "windowsStandardName": "North Korea Standard Time",
    "timezoneId": "Asia/Pyongyang"
  },
  {
    "name": "Korea Standard Time",
    "windowsDaylightDisplayName": "Korea Daylight Time",
    "windowsDisplayName": "(UTC+09:00) Seoul",
    "windowsId": "Korea Standard Time",
    "windowsStandardName": "Korea Standard Time",
    "timezoneId": "Asia/Seoul"
  },
  {
    "name": "Yakutsk Standard Time",
    "windowsDaylightDisplayName": "Russia TZ 8 Daylight Time",
    "windowsDisplayName": "(UTC+09:00) Yakutsk",
    "windowsId": "Yakutsk Standard Time",
    "windowsStandardName": "Russia TZ 8 Standard Time",
    "timezoneId": "Asia/Yakutsk"
  },
  {
    "name": "Cen. Australia Standard Time",
    "windowsDaylightDisplayName": "Cen. Australia Daylight Time",
    "windowsDisplayName": "(UTC+09:30) Adelaide",
    "windowsId": "Cen. Australia Standard Time",
    "windowsStandardName": "Cen. Australia Standard Time",
    "timezoneId": "Australia/Adelaide"
  },
  {
    "name": "AUS Central Standard Time",
    "windowsDaylightDisplayName": "AUS Central Daylight Time",
    "windowsDisplayName": "(UTC+09:30) Darwin",
    "windowsId": "AUS Central Standard Time",
    "windowsStandardName": "AUS Central Standard Time",
    "timezoneId": "Australia/Darwin"
  },
  {
    "name": "E. Australia Standard Time",
    "windowsDaylightDisplayName": "E. Australia Daylight Time",
    "windowsDisplayName": "(UTC+10:00) Brisbane",
    "windowsId": "E. Australia Standard Time",
    "windowsStandardName": "E. Australia Standard Time",
    "timezoneId": "Australia/Brisbane"
  },
  {
    "name": "AUS Eastern Standard Time",
    "windowsDaylightDisplayName": "AUS Eastern Daylight Time",
    "windowsDisplayName": "(UTC+10:00) Canberra, Melbourne, Sydney",
    "windowsId": "AUS Eastern Standard Time",
    "windowsStandardName": "AUS Eastern Standard Time",
    "timezoneId": "Australia/Sydney"
  },
  {
    "name": "West Pacific Standard Time",
    "windowsDaylightDisplayName": "West Pacific Daylight Time",
    "windowsDisplayName": "(UTC+10:00) Guam, Port Moresby",
    "windowsId": "West Pacific Standard Time",
    "windowsStandardName": "West Pacific Standard Time",
    "timezoneId": "Pacific/Port_Moresby"
  },
  {
    "name": "Tasmania Standard Time",
    "windowsDaylightDisplayName": "Tasmania Daylight Time",
    "windowsDisplayName": "(UTC+10:00) Hobart",
    "windowsId": "Tasmania Standard Time",
    "windowsStandardName": "Tasmania Standard Time",
    "timezoneId": "Australia/Hobart"
  },
  {
    "name": "Vladivostok Standard Time",
    "windowsDaylightDisplayName": "Russia TZ 9 Daylight Time",
    "windowsDisplayName": "(UTC+10:00) Vladivostok",
    "windowsId": "Vladivostok Standard Time",
    "windowsStandardName": "Russia TZ 9 Standard Time",
    "timezoneId": "Asia/Vladivostok"
  },
  {
    "name": "Lord Howe Standard Time",
    "windowsDaylightDisplayName": "Lord Howe Daylight Time",
    "windowsDisplayName": "(UTC+10:30) Lord Howe Island",
    "windowsId": "Lord Howe Standard Time",
    "windowsStandardName": "Lord Howe Standard Time",
    "timezoneId": "Australia/Lord_Howe"
  },
  {
    "name": "Bougainville Standard Time",
    "windowsDaylightDisplayName": "Bougainville Daylight Time",
    "windowsDisplayName": "(UTC+11:00) Bougainville Island",
    "windowsId": "Bougainville Standard Time",
    "windowsStandardName": "Bougainville Standard Time",
    "timezoneId": "Pacific/Bougainville"
  },
  {
    "name": "Russia Time Zone 10",
    "windowsDaylightDisplayName": "Russia TZ 10 Daylight Time",
    "windowsDisplayName": "(UTC+11:00) Chokurdakh",
    "windowsId": "Russia Time Zone 10",
    "windowsStandardName": "Russia TZ 10 Standard Time",
    "timezoneId": "Asia/Srednekolymsk"
  },
  {
    "name": "Magadan Standard Time",
    "windowsDaylightDisplayName": "Magadan Daylight Time",
    "windowsDisplayName": "(UTC+11:00) Magadan",
    "windowsId": "Magadan Standard Time",
    "windowsStandardName": "Magadan Standard Time",
    "timezoneId": "Asia/Magadan"
  },
  {
    "name": "Norfolk Standard Time",
    "windowsDaylightDisplayName": "Norfolk Daylight Time",
    "windowsDisplayName": "(UTC+11:00) Norfolk Island",
    "windowsId": "Norfolk Standard Time",
    "windowsStandardName": "Norfolk Standard Time",
    "timezoneId": "Pacific/Norfolk"
  },
  {
    "name": "Sakhalin Standard Time",
    "windowsDaylightDisplayName": "Sakhalin Daylight Time",
    "windowsDisplayName": "(UTC+11:00) Sakhalin",
    "windowsId": "Sakhalin Standard Time",
    "windowsStandardName": "Sakhalin Standard Time",
    "timezoneId": "Asia/Sakhalin"
  },
  {
    "name": "Central Pacific Standard Time",
    "windowsDaylightDisplayName": "Central Pacific Daylight Time",
    "windowsDisplayName": "(UTC+11:00) Solomon Is., New Caledonia",
    "windowsId": "Central Pacific Standard Time",
    "windowsStandardName": "Central Pacific Standard Time",
    "timezoneId": "Pacific/Guadalcanal"
  },
  {
    "name": "Russia Time Zone 11",
    "windowsDaylightDisplayName": "Russia TZ 11 Daylight Time",
    "windowsDisplayName": "(UTC+12:00) Anadyr, Petropavlovsk-Kamchatsky",
    "windowsId": "Russia Time Zone 11",
    "windowsStandardName": "Russia TZ 11 Standard Time",
    "timezoneId": "Asia/Kamchatka"
  },
  {
    "name": "New Zealand Standard Time",
    "windowsDaylightDisplayName": "New Zealand Daylight Time",
    "windowsDisplayName": "(UTC+12:00) Auckland, Wellington",
    "windowsId": "New Zealand Standard Time",
    "windowsStandardName": "New Zealand Standard Time",
    "timezoneId": "Pacific/Auckland"
  },
  {
    "name": "UTC+12",
    "windowsDaylightDisplayName": "UTC+12",
    "windowsDisplayName": "(UTC+12:00) Coordinated Universal Time+12",
    "windowsId": "UTC+12",
    "windowsStandardName": "UTC+12",
    "timezoneId": "Etc/GMT-12"
  },
  {
    "name": "Fiji Standard Time",
    "windowsDaylightDisplayName": "Fiji Daylight Time",
    "windowsDisplayName": "(UTC+12:00) Fiji",
    "windowsId": "Fiji Standard Time",
    "windowsStandardName": "Fiji Standard Time",
    "timezoneId": "Pacific/Fiji"
  },
  {
    "name": "Russia Time Zone 11",
    "windowsDaylightDisplayName": "Kamchatka Daylight Time",
    "windowsDisplayName": "(UTC+12:00) Petropavlovsk-Kamchatsky - Old",
    "windowsId": "Kamchatka Standard Time",
    "windowsStandardName": "Kamchatka Standard Time",
    "timezoneId": "Asia/Kamchatka"
  },
  {
    "name": "Chatham Islands Standard Time",
    "windowsDaylightDisplayName": "Chatham Islands Daylight Time",
    "windowsDisplayName": "(UTC+12:45) Chatham Islands",
    "windowsId": "Chatham Islands Standard Time",
    "windowsStandardName": "Chatham Islands Standard Time",
    "timezoneId": "Pacific/Chatham"
  },
  {
    "name": "UTC+13",
    "windowsDaylightDisplayName": "UTC+13",
    "windowsDisplayName": "(UTC+13:00) Coordinated Universal Time+13",
    "windowsId": "UTC+13",
    "windowsStandardName": "UTC+13",
    "timezoneId": "Etc/GMT-13"
  },
  {
    "name": "Tonga Standard Time",
    "windowsDaylightDisplayName": "Tonga Daylight Time",
    "windowsDisplayName": "(UTC+13:00) Nuku'alofa",
    "windowsId": "Tonga Standard Time",
    "windowsStandardName": "Tonga Standard Time",
    "timezoneId": "Pacific/Tongatapu"
  },
  {
    "name": "Samoa Standard Time",
    "windowsDaylightDisplayName": "Samoa Daylight Time",
    "windowsDisplayName": "(UTC+13:00) Samoa",
    "windowsId": "Samoa Standard Time",
    "windowsStandardName": "Samoa Standard Time",
    "timezoneId": "Pacific/Apia"
  },
  {
    "name": "Line Islands Standard Time",
    "windowsDaylightDisplayName": "Line Islands Daylight Time",
    "windowsDisplayName": "(UTC+14:00) Kiritimati Island",
    "windowsId": "Line Islands Standard Time",
    "windowsStandardName": "Line Islands Standard Time",
    "timezoneId": "Pacific/Kiritimati"
  }
]
`
