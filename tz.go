package fuzzytime

type TZ struct {
	Name   string
	Offset string
	Locale string
}

var tzTable = map[string][]TZ{
	"ACDT": {{"ACDT", "UTC+10:30", ""}}, //Australian Central Daylight Time
	"ACST": {{"ACST", "UTC+09:30", ""}}, //Australian Central Standard Time
	"ACT":  {{"ACT", "UTC+08", ""}},     //ASEAN Common Time
	"ADT":  {{"ADT", "UTC-03", ""}},     //Atlantic Daylight Time
	"AEDT": {{"AEDT", "UTC+11", ""}},    //Australian Eastern Daylight Time
	"AEST": {{"AEST", "UTC+10", ""}},    //Australian Eastern Standard Time
	"AFT":  {{"AFT", "UTC+04:30", ""}},  //Afghanistan Time
	"AKDT": {{"AKDT", "UTC-08", ""}},    //Alaska Daylight Time
	"AKST": {{"AKST", "UTC-09", ""}},    //Alaska Standard Time
	"AMST": {{"AMST", "UTC+05", ""}},    //Armenia Summer Time
	"AMT":  {{"AMT", "UTC+04", ""}},     //Armenia Time
	"ART":  {{"ART", "UTC-03", ""}},     //Argentina Time
	"AST": {{"AST", "UTC+03", ""}, //"Arab Standard Time (Kuwait, Riyadh)"
		{"AST", "UTC+04", ""},  //"Arabian Standard Time (Abu Dhabi, Muscat)"
		{"AST", "UTC+03", ""},  //Arabic Standard Time (Baghdad)
		{"AST", "UTC-04", ""}}, //Atlantic Standard Time
	"AWDT":  {{"AWDT", "UTC+09", ""}},  //Australian Western Daylight Time
	"AWST":  {{"AWST", "UTC+08", ""}},  //Australian Western Standard Time
	"AZOST": {{"AZOST", "UTC-01", ""}}, //Azores Standard Time
	"AZT":   {{"AZT", "UTC+04", ""}},   //Azerbaijan Time
	"BDT":   {{"BDT", "UTC+08", ""}},   //Brunei Time
	"BIOT":  {{"BIOT", "UTC+06", ""}},  //British Indian Ocean Time
	"BIT":   {{"BIT", "UTC-12", ""}},   //Baker Island Time
	"BOT":   {{"BOT", "UTC-04", ""}},   //Bolivia Time
	"BRT":   {{"BRT", "UTC-03", ""}},   //Brasilia Time
	"BST": {{"BST", "UTC+06", ""}, //Bangladesh Standard Time
		{"BST", "UTC+01", ""}}, //British Summer Time (British Standard Time from Feb 1968 to Oct 1971)
	"BTT":   {{"BTT", "UTC+06", ""}},      //Bhutan Time
	"CAT":   {{"CAT", "UTC+02", ""}},      //Central Africa Time
	"CCT":   {{"CCT", "UTC+06:30", ""}},   //Cocos Islands Time
	"CDT":   {{"CDT", "UTC-05", ""}},      //Central Daylight Time (North America)
	"CEDT":  {{"CEDT", "UTC+02", ""}},     //Central European Daylight Time
	"CEST":  {{"CEST", "UTC+02", ""}},     //Central European Summer Time (Cf. HAEC)
	"CET":   {{"CET", "UTC+01", ""}},      //Central European Time
	"CHADT": {{"CHADT", "UTC+13:45", ""}}, //Chatham Daylight Time
	"CHAST": {{"CHAST", "UTC+12:45", ""}}, //Chatham Standard Time
	"CIST":  {{"CIST", "UTC-08", ""}},     //Clipperton Island Standard Time
	"CKT":   {{"CKT", "UTC-10", ""}},      //Cook Island Time
	"CLST":  {{"CLST", "UTC-03", ""}},     //Chile Summer Time
	"CLT":   {{"CLT", "UTC-04", ""}},      //Chile Standard Time
	"COST":  {{"COST", "UTC-04", ""}},     //Colombia Summer Time
	"COT":   {{"COT", "UTC-05", ""}},      //Colombia Time
	"CST": {{"CST", "UTC-06", ""}, //Central Standard Time (North America)
		{"CST", "UTC+08", ""},       //China Standard Time
		{"CST", "UTC+09:30", "au"}}, //Central Standard Time (Australia)
	"CT":   {{"CT", "UTC+08", ""}},   //China Time
	"CVT":  {{"CVT", "UTC-01", ""}},  //Cape Verde Time
	"CXT":  {{"CXT", "UTC+07", ""}},  //Christmas Island Time
	"CHST": {{"CHST", "UTC+10", ""}}, //Chamorro Standard Time
	"DFT":  {{"DFT", "UTC+01", ""}},  //AIX specific equivalent of Central European Time
	"EAST": {{"EAST", "UTC-06", ""}}, //Easter Island Standard Time
	"EAT":  {{"EAT", "UTC+03", ""}},  //East Africa Time
	"ECT": {{"ECT", "UTC-04", ""}, //Eastern Caribbean Time (does not recognise DST)
		{"ECT", "UTC-05", ""}}, //Ecuador Time
	"EDT":  {{"EDT", "UTC-04", ""}},  //Eastern Daylight Time (North America)
	"EEDT": {{"EEDT", "UTC+03", ""}}, //Eastern European Daylight Time
	"EEST": {{"EEST", "UTC+03", ""}}, //Eastern European Summer Time
	"EET":  {{"EET", "UTC+02", ""}},  //Eastern European Time
	"EST":  {{"EST", "UTC-05", ""}},  //Eastern Standard Time (North America)
	"FET":  {{"FET", "UTC+03", ""}},  //Further-eastern_European_Time
	"FJT":  {{"FJT", "UTC+12", ""}},  //Fiji Time
	"FKST": {{"FKST", "UTC-03", ""}}, //Falkland Islands Summer Time
	"FKT":  {{"FKT", "UTC-04", ""}},  //Falkland Islands Time
	"GALT": {{"GALT", "UTC-06", ""}}, //Galapagos Time
	"GET":  {{"GET", "UTC+04", ""}},  //Georgia Standard Time
	"GFT":  {{"GFT", "UTC-03", ""}},  //French Guiana Time
	"GILT": {{"GILT", "UTC+12", ""}}, //Gilbert Island Time
	"GIT":  {{"GIT", "UTC-09", ""}},  //Gambier Island Time
	"GMT":  {{"GMT", "UTC", ""}},     //Greenwich Mean Time
	"GST": {{"GST", "UTC-02", ""}, //South Georgia and the South Sandwich Islands
		{"GST", "UTC+04", ""}}, //Gulf Standard Time
	"GYT":  {{"GYT", "UTC-04", ""}},     //Guyana Time
	"HADT": {{"HADT", "UTC-09", ""}},    //Hawaii-Aleutian Daylight Time
	"HAEC": {{"HAEC", "UTC+02", ""}},    //Heure Avancée d'Europe Centrale francised name for CEST
	"HAST": {{"HAST", "UTC-10", ""}},    //Hawaii-Aleutian Standard Time
	"HKT":  {{"HKT", "UTC+08", ""}},     //Hong Kong Time
	"HMT":  {{"HMT", "UTC+05", ""}},     //Heard and McDonald Islands Time
	"HST":  {{"HST", "UTC-10", ""}},     //Hawaii Standard Time
	"ICT":  {{"ICT", "UTC+07", ""}},     //Indochina Time
	"IDT":  {{"IDT", "UTC+03", ""}},     //Israeli Daylight Time
	"IRKT": {{"IRKT", "UTC+08", ""}},    //Irkutsk Time
	"IRST": {{"IRST", "UTC+03:30", ""}}, //Iran Standard Time
	"IST": {{"IST", "UTC+05:30", ""}, //Indian Standard Time
		{"IST", "UTC+01", "ie"}, //Irish Summer Time
		{"IST", "UTC+02", ""}},  //Israel Standard Time
	"JST":  {{"JST", "UTC+09", ""}},     //Japan Standard Time
	"KRAT": {{"KRAT", "UTC+07", ""}},    //Krasnoyarsk Time
	"KST":  {{"KST", "UTC+09", ""}},     //Korea Standard Time
	"LHST": {{"LHST", "UTC+10:30", ""}}, //Lord Howe Standard Time
	"LINT": {{"LINT", "UTC+14", ""}},    //Line Islands Time
	"MAGT": {{"MAGT", "UTC+11", ""}},    //Magadan Time
	"MDT":  {{"MDT", "UTC-06", ""}},     //Mountain Daylight Time (North America)
	"MET":  {{"MET", "UTC+01", ""}},     //Middle European Time Same zone as CET
	"MEST": {{"MEST", "UTC+02", ""}},    //Middle European Saving Time Same zone as CEST
	"MIT":  {{"MIT", "UTC-09:30", ""}},  //Marquesas Islands Time
	"MSK":  {{"MSK", "UTC+04", ""}},     //Moscow Time
	"MST": {{"MST", "UTC+08", ""}, //Malaysian Standard Time
		{"MST", "UTC-07", ""},     //Mountain Standard Time (North America)
		{"MST", "UTC+06:30", ""}}, //Myanmar Standard Time
	"MUT":  {{"MUT", "UTC+04", ""}},    //Mauritius Time
	"MYT":  {{"MYT", "UTC+08", ""}},    //Malaysia Time
	"NDT":  {{"NDT", "UTC-02:30", ""}}, //Newfoundland Daylight Time
	"NFT":  {{"NFT", "UTC+11:30", ""}}, //Norfolk Time[1]
	"NPT":  {{"NPT", "UTC+05:45", ""}}, //Nepal Time
	"NST":  {{"NST", "UTC-03:30", ""}}, //Newfoundland Standard Time
	"NT":   {{"NT", "UTC-03:30", ""}},  //Newfoundland Time
	"NZDT": {{"NZDT", "UTC+13", ""}},   //New Zealand Daylight Time
	"NZST": {{"NZST", "UTC+12", ""}},   //New Zealand Standard Time
	"OMST": {{"OMST", "UTC+06", ""}},   //Omsk Time
	"PDT":  {{"PDT", "UTC-07", ""}},    //Pacific Daylight Time (North America)
	"PETT": {{"PETT", "UTC+12", ""}},   //Kamchatka Time
	"PHOT": {{"PHOT", "UTC+13", ""}},   //Phoenix Island Time
	"PKT":  {{"PKT", "UTC+05", ""}},    //Pakistan Standard Time
	"PST": {{"PST", "UTC-08", ""}, //Pacific Standard Time (North America)
		{"PST", "UTC+08", ""}}, //Philippine Standard Time
	"RET":  {{"RET", "UTC+04", ""}},    //Réunion Time
	"SAMT": {{"SAMT", "UTC+04", ""}},   //Samara Time
	"SAST": {{"SAST", "UTC+02", ""}},   //South African Standard Time
	"SBT":  {{"SBT", "UTC+11", ""}},    //Solomon Islands Time
	"SCT":  {{"SCT", "UTC+04", ""}},    //Seychelles Time
	"SGT":  {{"SGT", "UTC+08", ""}},    //Singapore Time
	"SLT":  {{"SLT", "UTC+05:30", ""}}, //Sri Lanka Time
	"SST": {{"SST", "UTC-11", ""}, //Samoa Standard Time
		{"SST", "UTC+08", ""}}, //Singapore Standard Time
	"TAHT": {{"TAHT", "UTC-10", ""}},   //Tahiti Time
	"THA":  {{"THA", "UTC+07", ""}},    //Thailand Standard Time
	"UTC":  {{"UTC", "UTC", ""}},       //Coordinated Universal Time
	"UYST": {{"UYST", "UTC-02", ""}},   //Uruguay Summer Time
	"UYT":  {{"UYT", "UTC-03", ""}},    //Uruguay Standard Time
	"VET":  {{"VET", "UTC-04:30", ""}}, //Venezuelan Standard Time
	"VLAT": {{"VLAT", "UTC+10", ""}},   //Vladivostok Time
	"WAT":  {{"WAT", "UTC+01", ""}},    //West Africa Time
	"WEDT": {{"WEDT", "UTC+01", ""}},   //Western European Daylight Time
	"WEST": {{"WEST", "UTC+01", ""}},   //Western European Summer Time
	"WET":  {{"WET", "UTC", ""}},       //Western European Time
	"WST":  {{"WST", "UTC+08", ""}},    //Western Standard Time
	"YAKT": {{"YAKT", "UTC+09", ""}},   //Yakutsk Time
	"YEKT": {{"YEKT", "UTC+05", ""}},   //Yekaterinburg Time
}
