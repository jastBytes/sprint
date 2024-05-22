package sprint

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// See: http://en.wikipedia.org/wiki/Binary_prefix
const (
	Ki = 1024
	Mi = 1024 * Ki
	Gi = 1024 * Mi
	Ti = 1024 * Gi
	Pi = 1024 * Ti
)

var (
	binaryAbbrs = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi", "Yi"}
	binaryMap   = map[byte]int64{'k': Ki, 'm': Mi, 'g': Gi, 't': Ti, 'p': Pi}
)

// Humanize returns a string representation of the given number in power of two.
func Humanize(number int64) string {
	sizef := float64(number)
	for _, unit := range binaryAbbrs {
		if math.Abs(sizef) < Ki {
			return fmt.Sprintf("%0.0f%s", sizef, unit)
		}
		sizef /= Ki
	}
	return fmt.Sprintf("%0.0f", sizef)
}

// Dehumanize parses a string representation of a number in power of two and returns as number, e.g. 1, 5Mi, 10Gi.
func Dehumanize(humanizedNumber string) (int64, error) {
	sep := strings.LastIndexAny(humanizedNumber, "01234567890 ")
	if sep == -1 {
		// There should be at least a digit.
		return -1, fmt.Errorf("invalid: '%s'", humanizedNumber)
	}
	var num, sfx string
	if humanizedNumber[sep] != ' ' {
		num = humanizedNumber[:sep+1]
		sfx = humanizedNumber[sep+1:]
	} else {
		// Omit the space separator.
		num = humanizedNumber[:sep]
		sfx = humanizedNumber[sep+1:]
	}

	size, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return -1, err
	}

	// reject negative sizes.
	if size < 0 {
		return -1, fmt.Errorf("invalid: '%s'", humanizedNumber)
	}

	if len(sfx) == 0 {
		return int64(size), nil
	}

	// Process the suffix.
	if len(sfx) > 3 { // Too long.
		return -1, fmt.Errorf("invalid: '%s'", sfx)
	}
	sfx = strings.ToLower(sfx)
	// Trivial case: no suffix.
	if sfx[0] == ' ' {
		if len(sfx) > 1 { // no extra characters allowed after b.
			return -1, fmt.Errorf("invalid: '%s'", sfx)
		}
		return int64(size), nil
	}
	// A suffix from the map.
	if mul, ok := binaryMap[sfx[0]]; ok {
		size *= float64(mul)
	} else {
		return -1, fmt.Errorf("invalid suffix: '%s'", sfx)
	}

	return int64(size), nil
}
