package matcher

import "strings"

func matchContents(body []byte, criteria MatchCriteria) (bool, string) {
	if criteria.BodyContains != "" && !strings.Contains(string(body), criteria.BodyContains) {
		return false, "body content mismatch"
	}
	return true, "body content matches"
}
