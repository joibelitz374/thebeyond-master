package utils

import (
	"fmt"
	"strings"
	"time"
)

func TimeRemaining(until time.Time) (years, months, days, hours, minutes int) {
	now := time.Now()
	until = until.In(now.Location())

	years = until.Year() - now.Year()
	if now.AddDate(years, 0, 0).After(until) {
		years--
	}

	tmp := now.AddDate(years, 0, 0)

	for !tmp.AddDate(0, months+1, 0).After(until) {
		months++
	}
	tmp = tmp.AddDate(0, months, 0)

	duration := until.Sub(tmp)

	days = int(duration / (24 * time.Hour))
	duration -= time.Duration(days) * 24 * time.Hour

	hours = int(duration / time.Hour)
	duration -= time.Duration(hours) * time.Hour

	minutes = int(duration / time.Minute)

	return
}

func HumanizeDuration(years, months, days, hours, minutes int) string {
	parts := make([]string, 0, 4)
	if years != 0 {
		parts = append(parts, fmt.Sprintf("%d %s", years, plural("year", years)))
	}

	if months != 0 {
		parts = append(parts, fmt.Sprintf("%d %s", months, plural("month", months)))
	}

	if days != 0 {
		parts = append(parts, fmt.Sprintf("%d %s", days, plural("day", days)))
	}

	if hours != 0 {
		parts = append(parts, fmt.Sprintf("%d %s", hours, plural("hour", hours)))
	}

	if minutes != 0 {
		parts = append(parts, fmt.Sprintf("%d %s", minutes, plural("minute", minutes)))
	}

	if len(parts) == 0 {
		return "0 minutes"
	}

	return strings.Join(parts, " ")
}

func plural(unit string, n int) string {
	if n == 1 {
		return unit
	}

	return unit + "s"
}
