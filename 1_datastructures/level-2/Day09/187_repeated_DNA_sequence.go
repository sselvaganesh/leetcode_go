package string

/*
187. Repeated DNA Sequences
Medium
2.9K
478
Companies

The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.

For example, "ACGAATTCCG" is a DNA sequence.
When studying DNA, it is useful to identify repeated sequences within the DNA.

Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in any order.

Example 1:
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]

Example 2:
Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]


Constraints:
1 <= s.length <= 105
s[i] is either 'A', 'C', 'G', or 'T'.
*/

func FindRepeatedDnaSequences(s string) []string {

	m := make(map[string]int)
	var dna []byte
	var res []string

	for i := 0; i < len(s); i++ {

		if i < 9 {
			dna = append(dna, s[i])
			continue
		}
		dna = append(dna, s[i])

		if _, ok := m[string(dna)]; !ok {
			m[string(dna)] = 1
		} else {
			m[string(dna)]++
		}

		dna = dna[1:]
	}

	for dnaSeq, occur := range m {
		if occur > 1 {
			res = append(res, dnaSeq)
		}
	}

	return res
}
