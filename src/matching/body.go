package matcher

import (
	"errors"
	"strings"
)

func matchContents(body []byte, criteria MatchCriteria) (bool, error) {
	if criteria.BodyContains != "" && !strings.Contains(string(body), criteria.BodyContains) {
		return false, errors.New("body content mismatch")
	}
	return true, nil
}
