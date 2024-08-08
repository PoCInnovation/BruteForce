package matcher

import (
	"bruteforce/src/models"
	"errors"
	"strings"
)

func matchContents(body []byte, criteria models.MatchCriteria) (bool, error) {
	if criteria.BodyContains != "" && !strings.Contains(string(body), criteria.BodyContains) {
		return false, errors.New("body content mismatch")
	}
	return true, nil
}
