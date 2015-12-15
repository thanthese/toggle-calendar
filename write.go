package main

import (
	"fmt"
	"time"
)

func printIrregular(recs recs, today time.Time) (out string) {
	if len(recs) == 0 {
		return
	}
	min, max := recs[0].date, recs[len(recs)-1].date
	date := min
	tomorrow := today.AddDate(0, 0, 1)
	for !date.After(max) {
		if date.Equal(tomorrow) {
			out += "\n"
		}
		if !date.Equal(tomorrow) && date.After(today) && time.Monday == date.Weekday() {
			out += "\n"
		}
		ms := matchingDates(date, recs)
		if date.After(today) {
			out += prettyDate(date) + "\n"
			for _, m := range ms {
				out += prettyDescOnly(m.desc) + "\n"
			}
		} else {
			for _, m := range ms {
				out += prettyRegular(m) + "\n"
			}
		}
		date = date.AddDate(0, 0, 1)
	}
	return
}

func matchingDates(date time.Time, recs recs) (matches recs) {
	for _, r := range recs {
		if r.date == date {
			matches = append(matches, r)
		}
	}
	return
}

func printRegular(recs recs, today time.Time) (out string) {
	foundTomorrow := false
	for _, r := range recs {
		if r.date.After(today) && !foundTomorrow {
			foundTomorrow = true
			out += "\n"
		}
		out += prettyRegular(r) + "\n"
	}
	return
}

func prettyRegular(r rec) string {
	return prettyDate(r.date) + " " + r.desc
}

func prettyDescOnly(s string) string {
	return "    " + s
}

func prettyDate(d time.Time) string {
	return fmt.Sprintf("%02d.%02d.%02d%c",
		d.Year()-2000,
		int(d.Month()),
		d.Day(),
		"umtwrfs"[d.Weekday()])
}