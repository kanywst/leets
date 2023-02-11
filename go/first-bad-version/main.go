package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) (ans int) {
	for i := 1; i < n+1; i++ {
		if isBadVersion(i) {
			ans = i
			break
		}
	}
	return ans
}
