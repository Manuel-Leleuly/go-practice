package helpers

func GetUrl(path string) string {
	const url = "localhost:3000"
	if path != "" {
		return "http://" + url + path
	}
	return url
}
