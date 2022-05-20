package sol

func Run(commands *[]string, values *[][]string) []string {
	dict := Constructor()
	c := *commands
	v := *values
	cLen := len(c)
	result := []string{"null"}
	for idx := 1; idx < cLen; idx++ {
		command := c[idx]
		datum := v[idx][0]
		switch command {
		case "addWord":
			dict.AddWord(datum)
			result = append(result, "null")
		case "search":
			if dict.Search(datum) {
				result = append(result, "true")
			} else {
				result = append(result, "false")
			}
		}
	}
	return result
}
