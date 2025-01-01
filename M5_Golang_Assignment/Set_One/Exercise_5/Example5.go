package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ClimateData struct {
	City    string
	AvgTemp float64
	AvgRain float64
}

func displayMenu() {
	fmt.Println("\n=== Climate Data Analysis Menu ===")
	fmt.Println("1. Display all cities")
	fmt.Println("2. Show highest and lowest temperatures")
	fmt.Println("3. Calculate average rainfall")
	fmt.Println("4. Filter cities by rainfall threshold")
	fmt.Println("5. Search for a city")
	fmt.Println("6. Exit")
	fmt.Print("\nEnter your choice (1-6): ")
}

func displayAllCities(cities []ClimateData) {
	fmt.Println("\n=== All Cities Data ===")
	fmt.Printf("%-15s %-15s %-15s\n", "City", "Temperature (°C)", "Rainfall (mm)")
	fmt.Println(strings.Repeat("-", 45))
	for _, city := range cities {
		fmt.Printf("%-15s %-15.2f %-15.2f\n", city.City, city.AvgTemp, city.AvgRain)
	}
}

func findTemperatureExtremes(cities []ClimateData) (ClimateData, ClimateData) {
	highestCity := cities[0]
	lowestCity := cities[0]
	for _, city := range cities {
		if city.AvgTemp > highestCity.AvgTemp {
			highestCity = city
		}
		if city.AvgTemp < lowestCity.AvgTemp {
			lowestCity = city
		}
	}
	return highestCity, lowestCity
}

func calculateAverageRainfall(cities []ClimateData) float64 {
	if len(cities) == 0 {
		return 0
	}
	var totalRainfall float64
	for _, city := range cities {
		totalRainfall += city.AvgRain
	}
	return totalRainfall / float64(len(cities))
}

func filterCitiesByRainfall(cities []ClimateData, threshold float64) []ClimateData {
	filteredCities := []ClimateData{}
	for _, city := range cities {
		if city.AvgRain > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

func promptThreshold() float64 {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nEnter the rainfall threshold (mm): ")
		input, _ := reader.ReadString('\n')
		threshold, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err == nil && threshold >= 0 {
			return threshold
		}
		fmt.Println("Invalid input. Please enter a non-negative number.")
	}
}

func searchCity(cities []ClimateData, name string) (*ClimateData, error) {
	for _, city := range cities {
		if strings.EqualFold(city.City, name) {
			return &city, nil
		}
	}
	return nil, fmt.Errorf("city '%s' not found", name)
}

func getMenuChoice() int {
	for {
		reader := bufio.NewReader(os.Stdin)
		choiceStr, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
		if err == nil && choice >= 1 && choice <= 6 {
			return choice
		}
		fmt.Println("Invalid choice. Please enter a number between 1 and 6.")
	}
}

func main() {
	cities := []ClimateData{
		{"New York", 22.5, 50.0},
		{"London", 15.0, 30.0},
		{"Tokyo", 25.0, 40.0},
		{"Sydney", 28.0, 60.0},
	}

	for {
		displayMenu()
		choice := getMenuChoice()
		switch choice {
		case 1:
			displayAllCities(cities)
		case 2:
			highest, lowest := findTemperatureExtremes(cities)
			fmt.Printf("\nHighest Temperature: %s (%.2f °C)\n", highest.City, highest.AvgTemp)
			fmt.Printf("Lowest Temperature: %s (%.2f °C)\n", lowest.City, lowest.AvgTemp)
		case 3:
			avgRainfall := calculateAverageRainfall(cities)
			fmt.Printf("\nAverage Rainfall: %.2f mm\n", avgRainfall)
		case 4:
			threshold := promptThreshold()
			filteredCities := filterCitiesByRainfall(cities, threshold)
			if len(filteredCities) > 0 {
				fmt.Printf("\nCities with rainfall above %.2f mm:\n", threshold)
				for _, city := range filteredCities {
					fmt.Printf("%s: %.2f mm\n", city.City, city.AvgRain)
				}
			} else {
				fmt.Printf("\nNo cities found with rainfall above %.2f mm\n", threshold)
			}
		case 5:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("\nEnter city name: ")
			cityName, _ := reader.ReadString('\n')
			cityName = strings.TrimSpace(cityName)
			city, err := searchCity(cities, cityName)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\nCity: %s\nTemperature: %.2f °C\nRainfall: %.2f mm\n", city.City, city.AvgTemp, city.AvgRain)
			}
		case 6:
			fmt.Println("\nThank you for using the Climate Data Analysis Program! Goodbye!")
			return
		}
	}
}
