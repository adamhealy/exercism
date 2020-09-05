// Package space implements a function to return the age of a
// person on the different planets of the solar system
package space

import "math"

// Planet is a string type for planet names
type Planet string

// Age calculates a persons age on a give planet based on their earth age in seconds
func Age(seconds float64, planet Planet) float64 {
	const earthYearSecs float64 = 31557600
	var planetAge float64 = 0

	if planet == "Earth" {
		planetAge = seconds / earthYearSecs
	} else if planet == "Mercury" {
		planetAge = seconds / earthYearSecs / 0.2408467
	} else if planet == "Venus" {
		planetAge = seconds / earthYearSecs / 0.61519726
	} else if planet == "Mars" {
		planetAge = seconds / earthYearSecs / 1.8808158
	} else if planet == "Jupiter" {
		planetAge = seconds / earthYearSecs / 11.862615
	} else if planet == "Saturn" {
		planetAge = seconds / earthYearSecs / 29.447498
	} else if planet == "Uranus" {
		planetAge = seconds / earthYearSecs / 84.016846
	} else if planet == "Neptune" {
		planetAge = seconds / earthYearSecs / 164.79132
	}

	return math.Round(planetAge*100) / 100
}
