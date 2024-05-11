package constant

var (
	//root = "/mnt/c/Users/zen/Videos/test/data"
	root = "/data"
)

func SetRoot(s string) {
	root = s
}
func GetRoot() string {
	return root
}
