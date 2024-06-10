package constant

var (
	//root = "/mnt/c/Users/zen/Videos/test/data"
	root = "E:\\pikpak\\亀頭責"
)

func SetRoot(s string) {
	root = s
}
func GetRoot() string {
	return root
}
