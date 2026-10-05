import "slices"
func groupAnagrams(strs []string) [][]string {
	seen := make(map[string][]int)
	for index, str := range strs {
		bytes := []byte(str)
		slices.Sort(bytes)
		ss := string(bytes)
		if _, exist := seen[ss]; !exist {
			seen[ss] = make([]int, 0)
		}
		seen[ss] = append(seen[ss], index)
	}
	result := make([][]string, 0, len(seen))
	for _, ids := range seen {
		tmp := make([]string, 0, len(ids))
		for _, id := range ids {
			tmp = append(tmp, strs[id])
		}
		result = append(result, tmp)
	}
	return result
}
