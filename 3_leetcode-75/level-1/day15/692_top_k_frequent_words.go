package heap

/*
692. Top K Frequent Words
Medium
6.9K
322
Companies

Given an array of strings words and an integer k, return the k most frequent strings.

Return the answer sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.

Example 1:
Input: words = ["i","love","leetcode","i","love","coding"], k = 2
Output: ["i","love"]
Explanation: "i" and "love" are the two most frequent words.
Note that "i" comes before "love" due to a lower alphabetical order.

Example 2:
Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
Output: ["the","is","sunny","day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.


Constraints:
1 <= words.length <= 500
1 <= words[i].length <= 10
words[i] consists of lowercase English letters.
k is in the range [1, The number of unique words[i]]


Follow-up: Could you solve it in O(n log(k)) time and O(n) extra space?
*/

import (
	//"fmt"
	"sort"
)

func TopKFrequent(words []string, k int) []string {

	return solution_k_frequent_words(words, k)
}

func solution_k_frequent_words(words []string, k int) []string {

	m := make(map[string]int)

	for i := 0; i < len(words); i++ {

		_, ok := m[words[i]]
		if !ok {
			m[words[i]] = 1
		} else {
			m[words[i]]++
		}
	}

	t := make(map[int]struct{})
	for _, v := range m {
		t[v] = struct{}{}
	}

	var occur []int
	for k, _ := range t {
		occur = append(occur, k)
	}

	sort.Slice(occur, func(i, j int) bool {
		return occur[i] > occur[j]
	})

	var result []string
	tot := 0
	for i := 0; i < len(occur); i++ {

		temp := []string{}
		for w, o := range m {

			if occur[i] == o {
				tot++
				temp = append(temp, w)
			}

		}

		sort.Strings(temp)
		result = append(result, temp...)
	}

	return result[:k]
}

func isValid(occur []int, k int) bool {

	for _, val := range occur {
		if val == k {
			return true
		}
	}

	return false
}
