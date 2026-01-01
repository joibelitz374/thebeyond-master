package utils

import (
	"fmt"
	"strings"
	"time"
)

func TimeRemaining(until time.Time) (years, months, days, hours, minutes int) {
	now := time.Now()
	until = until.In(now.Location())

	if !now.Before(until) {
		return 0, 0, 0, 0, 0
	}

	years = until.Year() - now.Year()
	months = int(until.Month() - now.Month())
	days = until.Day() - now.Day()

	if days < 0 {
		months--
		firstDayOfCurrentMonth := time.Date(until.Year(), until.Month(), 1, 0, 0, 0, 0, until.Location())
		firstDayOfPreviousMonth := firstDayOfCurrentMonth.AddDate(0, -1, 0)
		daysInPreviousMonth := firstDayOfCurrentMonth.Sub(firstDayOfPreviousMonth).Hours() / 24
		days += int(daysInPreviousMonth)
	}

	if months < 0 {
		years--
		months += 12
	}

	tmp := now.AddDate(years, months, days)
	duration := until.Sub(tmp)
	hours = int(duration.Hours())
	duration -= time.Duration(hours) * time.Hour
	minutes = int(duration.Minutes())

	if minutes < 0 {
		minutes = 0
	}
	if hours < 0 {
		hours = 0
	}

	return years, months, days, hours, minutes
}

func HumanizeDuration(years, months, days, hours, minutes int) string {
	parts := make([]string, 0, 5)

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
	if n == 1 || n == -1 {
		return unit
	}
	return unit + "s"
}
