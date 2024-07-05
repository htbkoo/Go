// https://www.codewars.com/kata/52742f58faf5485cae000b9a/train/go

/*
Your task in order to complete this Kata is to write a function which formats a duration, given as a number of seconds, in a human-friendly way.

The function must accept a non-negative integer. If it is zero, it just returns "now". Otherwise, the duration is expressed as a combination of years, days, hours, minutes and seconds.

It is much easier to understand with an example:

* For seconds = 62, your function should return
    "1 minute and 2 seconds"
* For seconds = 3662, your function should return
    "1 hour, 1 minute and 2 seconds"
For the purpose of this Kata, a year is 365 days and a day is 24 hours.

Note that spaces are important.

Detailed rules
The resulting expression is made of components like 4 seconds, 1 year, etc. In general, a positive integer and one of the valid units of time, separated by a space. The unit of time is used in plural if the integer is greater than 1.

The components are separated by a comma and a space (", "). Except the last component, which is separated by " and ", just like it would be written in English.

A more significant units of time will occur before than a least significant one. Therefore, 1 second and 1 year is not correct, but 1 year and 1 second is.

Different components have different unit of times. So there is not repeated units like in 5 seconds and 1 second.

A component will not appear at all if its value happens to be zero. Hence, 1 minute and 0 seconds is not valid, but it should be just 1 minute.

A unit of time must be used "as much as possible". It means that the function should not return 61 seconds, but 1 minute and 1 second instead. Formally, the duration specified by of a component must not be greater than any valid more significant unit of time.

*/

package kata

import (
	"fmt"
	"strings"
)

type TimePart struct {
	val  int64
	unit string
}

func FormatDuration(seconds int64) string {
	// Your code here
	if seconds == 0 {
		return "now"
	}

	numSecondsInMinute := int64(60)
	numMinutesInHour := int64(60)
	numHoursInDay := int64(24)
	numDaysInYear := int64(365)

	ss := seconds % numSecondsInMinute
	minutes := seconds / numSecondsInMinute
	mm := minutes % numMinutesInHour
	hours := minutes / numMinutesInHour
	hh := hours % numHoursInDay
	days := hours / numHoursInDay
	dd := days % numDaysInYear
	yy := days / numDaysInYear

	parts := []TimePart{
		{yy, "year"},
		{dd, "day"},
		{hh, "hour"},
		{mm, "minute"},
		{ss, "second"},
	}

	includedParts := make([]string, 0)
	for _, part := range parts {
		if part.val == 0 {
			continue
		}

		if part.val == 1 {
			includedParts = append(includedParts, fmt.Sprintf("%d %s", part.val, part.unit))
		} else {
			includedParts = append(includedParts, fmt.Sprintf("%d %ss", part.val, part.unit))
		}
	}

	// TODO: handle this programming error case
	//    if len(includedParts) == 0

	if len(includedParts) == 1 {
		return includedParts[0]
	} else {
		firstPart := strings.Join(includedParts[:len(includedParts)-1], ", ")
		tailPart := includedParts[len(includedParts)-1]

		return fmt.Sprintf("%s and %s", firstPart, tailPart)
	}
}
