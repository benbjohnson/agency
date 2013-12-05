package agency

var types = map[string][][]byte{
	"iOS": {
		[]byte("iPhone"),
		[]byte("iPad"),
		[]byte("iPod"),
	},
	"Blackberry": {
		[]byte("iBlackberry"),
		[]byte("Blackberry"),
	},
	"webOS": {
		[]byte("WebOS"),
	},
	"Windows Phone": {
		[]byte("Windows Phone"),
	},
	"Android": {
		[]byte("Android"),
	},
	"Windows": {
		[]byte("Windows NT"),
	},
	"Mac": {
		[]byte("Mac OS"),
	},
	"Linux": {
		[]byte("Linux"),
	},
}
