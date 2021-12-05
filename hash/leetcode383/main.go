package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	hash := make([]int, 26)
	for _, v := range magazine {
		hash[v-'a'] = hash[v-'a'] + 1
	}

	for _, v := range ransomNote {
		hash[v-'a'] = hash[v-'a'] - 1
		if hash[v-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
