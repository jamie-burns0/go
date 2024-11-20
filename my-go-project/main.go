package main

import (
	"fmt"
)

func main() {
	cities := GenerateRandomCities()
	fmt.Println("Random Global Cities:", cities)
}

func GenerateRandomCities() []string {
	cityNames := []string{
		"New York", "Tokyo", "London", "Paris", "Berlin",
		"Sydney", "Toronto", "Dubai", "Singapore", "Barcelona",
		"Moscow", "Los Angeles", "Rio de Janeiro", "Istanbul", "Seoul",
		"Mexico City", "Bangkok", "Buenos Aires", "Cairo", "Rome",
		"Amsterdam", "Hong Kong", "Madrid", "Lima", "Kuala Lumpur",
		"Lisbon", "Mumbai", "Jakarta", "Santiago", "Athens",
		"Stockholm", "Helsinki", "Oslo", "Copenhagen", "Brussels",
		"Vienna", "Zurich", "Dublin", "Prague", "Budapest",
		"Manila", "Nairobi", "Cape Town", "Lagos", "Accra",
		"Tel Aviv", "Bangalore", "Karachi", "Chennai", "Colombo",
		"Addis Ababa", "Tunis", "Algiers", "Hanoi", "Ho Chi Minh City",
		"Seville", "Valencia", "Malmo", "Bordeaux", "Marseille",
		"Montreal", "Calgary", "Vancouver", "Edmonton", "Ottawa",
		"Brisbane", "Perth", "Adelaide", "Gold Coast", "Wellington",
		"Auckland", "Christchurch", "Dakar", "Lima", "Helsinki",
		"Oslo", "Copenhagen", "Amsterdam", "Brussels", "Lisbon",
		"Stockholm", "Vienna", "Zurich", "Budapest", "Prague",
		"Belgrade", "Sarajevo", "Sofia", "Skopje", "Tbilisi",
		"Yerevan", "Baku", "Astana", "Tashkent", "Dushanbe",
		"Almaty", "Bishkek", "Ulaanbaatar", "Hanoi", "Vientiane",
		"Phnom Penh", "Yangon", "Dhaka", "Kathmandu", "Thimphu",
		"Malé", "Suva", "Port Moresby", "Nuku'alofa", "Apia",
		"Funafuti", "Tarawa", "Honiara", "Port Vila", "Nauru",
	}

	return cityNames
}