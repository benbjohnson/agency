package agency

var categories = map[string][][]byte{
	"Tablet": {
		[]byte("iPad"),
		[]byte("Android"),
	},
	"Mobile": {
		[]byte("iMobile"),
		[]byte("Mobile"),
		[]byte("Blackberry"),
		[]byte("webOS"),
		[]byte("IEMobile"),
	},
}
