package constant

var (
	//root = "/mnt/c/Users/zen/Videos/test/data"
	root = "/Volumes/Ventoy/翻译完成 等待合并/Hitomi"
)

func SetRoot(s string) {
	root = s
}
func GetRoot() string {
	return root
}
