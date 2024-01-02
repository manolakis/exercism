package gigasecond

import "time"

// AddGigasecond calculates the moment when someone has lived for 10^9 seconds..
func AddGigasecond(t time.Time) time.Time {
	const gigaSecond = time.Duration(1e9) * time.Second

	return t.Add(gigaSecond)
}
