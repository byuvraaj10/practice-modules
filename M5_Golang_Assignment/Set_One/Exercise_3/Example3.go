package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Constants for validation
const (
	MinPrice = 0.01
	MaxPrice = 999999.99
	MaxStock = 999999
)

type Product struct {
	ID        int
	Name      string
	Price     float64
	Stock     int
	UpdatedAt time.Time
}

func validateProduct(p Product, inventory []Product) error {
	if p.ID <= 0 {
		return fmt.Errorf("ID must be positive and unique")
	}
	for _, item := range inventory {
		if item.ID == p.ID {
			return fmt.Errorf("ID already exists")
		}
	}
	if p.Price < MinPrice || p.Price > MaxPrice {
		return fmt.Errorf("Price must be between %.2f and %.2f", MinPrice, MaxPrice)
	}
	if p.Stock < 0 || p.Stock > MaxStock {
		return fmt.Errorf("Stock must be between 0 and %d", MaxStock)
	}
	return nil
}

func promptInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func addProduct(inventory *[]Product) {
	fmt.Println("\n=== Add New Product ===")
	id, _ := strconv.Atoi(promptInput("Enter Product ID: "))
	name := promptInput("Enter Product Name: ")
	price, _ := strconv.ParseFloat(promptInput("Enter Product Price: "), 64)
	stock, _ := strconv.Atoi(promptInput("Enter Product Stock: "))

	product := Product{
		ID:        id,
		Name:      name,
		Price:     price,
		Stock:     stock,
		UpdatedAt: time.Now(),
	}

	if err := validateProduct(product, *inventory); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	*inventory = append(*inventory, product)
	fmt.Println("Product added successfully!")
}

func updateStock(inventory *[]Product) {
	fmt.Println("\n=== Update Product Stock ===")
	id, _ := strconv.Atoi(promptInput("Enter Product ID: "))
	for i, product := range *inventory {
		if product.ID == id {
			newStock, _ := strconv.Atoi(promptInput("Enter New Stock: "))
			if newStock < 0 || newStock > MaxStock {
				fmt.Printf("Invalid stock level. Must be between 0 and %d.\n", MaxStock)
				return
			}
			(*inventory)[i].Stock = newStock
			(*inventory)[i].UpdatedAt = time.Now()
			fmt.Println("Stock updated successfully!")
			return
		}
	}
	fmt.Println("Product not found.")
}

func displayInventory(inventory []Product) {
	if len(inventory) == 0 {
		fmt.Println("\nNo products in inventory.")
		return
	}

	fmt.Println("\n=== Inventory ===")
	fmt.Printf("%-5s %-20s %-10s %-10s %-20s\n", "ID", "Name", "Price", "Stock", "Updated")
	fmt.Println(strings.Repeat("-", 70))
	for _, product := range inventory {
		fmt.Printf("%-5d %-20s %-10.2f %-10d %-20s\n",
			product.ID, product.Name, product.Price, product.Stock, product.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

func searchProduct(inventory []Product) {
	fmt.Println("\n=== Search Product ===")
	query := promptInput("Enter Product ID or Name: ")
	found := false
	for _, product := range inventory {
		if strconv.Itoa(product.ID) == query || strings.Contains(strings.ToLower(product.Name), strings.ToLower(query)) {
			fmt.Printf("Found Product: ID=%d, Name=%s, Price=%.2f, Stock=%d\n",
				product.ID, product.Name, product.Price, product.Stock)
			found = true
		}
	}
	if !found {
		fmt.Println("No matching product found.")
	}
}

func sortInventory(inventory *[]Product) {
	fmt.Println("\n=== Sort Inventory ===")
	fmt.Println("1. Sort by Price (ascending)")
	fmt.Println("2. Sort by Stock (ascending)")
	choice := promptInput("Enter choice: ")

	switch choice {
	case "1":
		sort.Slice(*inventory, func(i, j int) bool {
			return (*inventory)[i].Price < (*inventory)[j].Price
		})
	case "2":
		sort.Slice(*inventory, func(i, j int) bool {
			return (*inventory)[i].Stock < (*inventory)[j].Stock
		})
	default:
		fmt.Println("Invalid choice.")
		return
	}

	fmt.Println("Inventory sorted successfully!")
	displayInventory(*inventory)
}

func main() {
	inventory := []Product{}
	for {
		fmt.Println("\n=== Inventory Management ===")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Display Inventory")
		fmt.Println("4. Search Product")
		fmt.Println("5. Sort Inventory")
		fmt.Println("6. Exit")
		choice := promptInput("Enter choice: ")
		switch choice {
		case "1":
			addProduct(&inventory)
		case "2":
			updateStock(&inventory)
		case "3":
			displayInventory(inventory)
		case "4":
			searchProduct(inventory)
		case "5":
			sortInventory(&inventory)
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
