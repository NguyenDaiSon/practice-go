/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	badVersion, firstVersion, lastVersion := -1, 1, n

	for firstVersion <= lastVersion {
		version := firstVersion + (lastVersion-firstVersion)/2
		if isBadVersion(version) {
			badVersion = version
			lastVersion = version - 1
		} else {
			firstVersion = version + 1
		}
	}
    
	return badVersion
}
