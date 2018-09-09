/*

Package epochtime implements a type similar to time.Time, but expects
input in the form of a Unix epoch time in a string. This is meant to
be used in structs representing JSON input.

Adapted from https://stackoverflow.com/a/43431284.

*/
package epochtime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// EpochTime has a time.Time as its underlying type.
type EpochTime time.Time

// UnmarshalJSON converts epoch from JSON string to EpochTime.
func (et *EpochTime) UnmarshalJSON(data []byte) error {
	t := strings.Trim(string(data), `"`)
	sec, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return err
	}
	*et = EpochTime(time.Unix(sec, 0))
	return nil
}

// Add returns the EpochTime t+d.
func (et EpochTime) Add(d time.Duration) EpochTime {
	return EpochTime(time.Time(et).Add(d))
}

func (et EpochTime) String() string {
	return fmt.Sprintf("%v", time.Time(et).Unix())
}
