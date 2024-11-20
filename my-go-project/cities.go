package main

import (
	"math/rand"
	"time"
)

var cities = [100]string{
	"New York", "Tokyo", "London", "Paris", "Berlin",
	"Sydney", "Toronto", "Dubai", "Singapore", "Barcelona",
	"Moscow", "Los Angeles", "Rio de Janeiro", "Istanbul", "Seoul",
	"Mexico City", "Bangkok", "Buenos Aires", "Cairo", "Rome",
	"Amsterdam", "Hong Kong", "Madrid", "Lima", "Lisbon",
	"San Francisco", "Chicago", "Kuala Lumpur", "Mumbai", "Jakarta",
	"Vienna", "Brussels", "Athens", "Helsinki", "Oslo",
	"Stockholm", "Copenhagen", "Zurich", "Dublin", "Budapest",
	"Prague", "Warsaw", "Santiago", "Nairobi", "Manila",
	"Lisbon", "Bangalore", "Cape Town", "Tel Aviv", "Abu Dhabi",
	"Doha", "Hanoi", "Colombo", "Tunis", "Lagos",
	"Accra", "Addis Ababa", "Algiers", "Helsinki", "Havana",
	"Valencia", "Belfast", "Edinburgh", "Glasgow", "Montreal",
	"Brisbane", "Perth", "Adelaide", "Wellington", "Christchurch",
	"Osaka", "Nagoya", "Kobe", "Fukuoka", "Kyoto",
	"Seville", "Malaga", "Bilbao", "Granada", "Valencia",
	"Antwerp", "Ghent", "Bruges", "Rotterdam", "Utrecht",
	"Geneva", "Lausanne", "Lucerne", "Stuttgart", "Frankfurt",
	"Munich", "Düsseldorf", "Cologne", "Hamburg", "Leipzig",
}

func GenerateRandomCities() []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cities), func(i, j int) {
		cities[i], cities[j] = cities[j], cities[i]
	})
	return cities[:100]
}