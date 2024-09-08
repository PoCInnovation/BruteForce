package matcher

import (
	"bruteforce/src/models"
	"errors"
	"strings"
)

func matchContents(body []byte, criteria models.MatchCriteria) error {
	if criteria.BodyContains != "" && !strings.Contains(string(body), criteria.BodyContains) {
		return errors.New("body content mismatch")
	}
	return nil
}
