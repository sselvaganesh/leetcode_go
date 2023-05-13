package graph

/*
997. Find the Town Judge
Easy
5.6K
445
Companies

In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:

The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi. If a trust relationship does not exist in trust array, then such a trust relationship does not exist.

Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.

Example 1:
Input: n = 2, trust = [[1,2]]
Output: 2

Example 2:
Input: n = 3, trust = [[1,3],[2,3]]
Output: 3

Example 3:
Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1


Constraints:
1 <= n <= 1000
0 <= trust.length <= 104
trust[i].length == 2
All the pairs of trust are unique.
ai != bi
1 <= ai, bi <= n
*/

func FindJudge(n int, trust [][]int) int {

	type people map[int]struct{}

	all, judgeMap := make(map[int]struct{}), make(map[int]people)
	for i := 0; i < n; i++ {
		all[i+1] = struct{}{}
	}

	for _, pair := range trust {
		delete(all, pair[0])

		if p, ok := judgeMap[pair[1]]; !ok {
			p1 := make(people)
			p1[pair[0]] = struct{}{}
			judgeMap[pair[1]] = p1
		} else {
			p[pair[0]] = struct{}{}
			judgeMap[pair[1]] = p
		}

	}

	if len(all) != 1 {
		return -1
	}

	var judge int
	for k, _ := range all {
		judge = k
	}

	if len(judgeMap[judge]) != (n - 1) {
		return -1
	}

	return judge

}

func findJudgeOptimal(n int, trust [][]int) int {

	// trustedPersons[i] - number of people who trust person i
	trusted := make([]int, n+1)
	// personTrusts[i] - person i trusts this number of people
	personTrusts := make([]int, n+1)

	for _, t := range trust {
		a, b := t[0], t[1]
		personTrusts[a]++
		trusted[b]++
	}

	for i := 1; i <= n; i++ {
		if personTrusts[i] == 0 && trusted[i] == n-1 {
			return i
		}
	}
	return -1
}
