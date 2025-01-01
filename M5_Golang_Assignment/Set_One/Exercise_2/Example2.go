package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Constants
const (
	DEPOSIT     = "DEPOSIT"
	WITHDRAWAL  = "WITHDRAWAL"
	MIN_BALANCE = 100.0
)

type Account struct {
	ID           int
	Name         string
	Balance      float64
	Transactions []Transaction
	CreatedAt    time.Time
}

type Transaction struct {
	Type      string
	Amount    float64
	Balance   float64
	Timestamp time.Time
}

func validateAmount(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Amount must be greater than zero")
	}
	return nil
}

func findAccount(accounts []Account, id int) (*Account, error) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}
	return nil, fmt.Errorf("Account with ID %d not found", id)
}

func promptInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func addAccount(accounts *[]Account) {
	fmt.Println("\n=== Add New Account ===")
	id, _ := strconv.Atoi(promptInput("Enter Account ID: "))
	if _, err := findAccount(*accounts, id); err == nil {
		fmt.Printf("Account ID %d already exists\n", id)
		return
	}
	name := promptInput("Enter Account Holder Name: ")
	balance, _ := strconv.ParseFloat(promptInput("Enter Initial Balance: "), 64)
	if balance < MIN_BALANCE {
		fmt.Printf("Initial balance must be at least %.2f\n", MIN_BALANCE)
		return
	}
	newAccount := Account{
		ID:        id,
		Name:      name,
		Balance:   balance,
		CreatedAt: time.Now(),
		Transactions: []Transaction{
			{Type: "INITIAL_DEPOSIT", Amount: balance, Balance: balance, Timestamp: time.Now()},
		},
	}
	*accounts = append(*accounts, newAccount)
	fmt.Println("Account added successfully!")
}

func deposit(accounts *[]Account) {
	fmt.Println("\n=== Deposit Money ===")
	id, _ := strconv.Atoi(promptInput("Enter Account ID: "))
	account, err := findAccount(*accounts, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	amount, _ := strconv.ParseFloat(promptInput("Enter Deposit Amount: "), 64)
	if err := validateAmount(amount); err != nil {
		fmt.Println(err)
		return
	}
	account.Balance += amount
	account.Transactions = append(account.Transactions, Transaction{
		Type:      DEPOSIT,
		Amount:    amount,
		Balance:   account.Balance,
		Timestamp: time.Now(),
	})
	fmt.Printf("Successfully deposited %.2f. New balance: %.2f\n", amount, account.Balance)
}

func withdraw(accounts *[]Account) {
	fmt.Println("\n=== Withdraw Money ===")
	id, _ := strconv.Atoi(promptInput("Enter Account ID: "))
	account, err := findAccount(*accounts, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	amount, _ := strconv.ParseFloat(promptInput("Enter Withdrawal Amount: "), 64)
	if err := validateAmount(amount); err != nil {
		fmt.Println(err)
		return
	}
	if account.Balance-amount < MIN_BALANCE {
		fmt.Printf("Insufficient balance. Minimum balance required: %.2f\n", MIN_BALANCE)
		return
	}
	account.Balance -= amount
	account.Transactions = append(account.Transactions, Transaction{
		Type:      WITHDRAWAL,
		Amount:    amount,
		Balance:   account.Balance,
		Timestamp: time.Now(),
	})
	fmt.Printf("Successfully withdrew %.2f. New balance: %.2f\n", amount, account.Balance)
}

func viewBalance(accounts []Account) {
	fmt.Println("\n=== View Balance ===")
	id, _ := strconv.Atoi(promptInput("Enter Account ID: "))
	account, err := findAccount(accounts, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Account Holder: %s\nAccount ID: %d\nBalance: %.2f\n", account.Name, account.ID, account.Balance)
}

func viewTransactionHistory(accounts []Account) {
	fmt.Println("\n=== Transaction History ===")
	id, _ := strconv.Atoi(promptInput("Enter Account ID: "))
	account, err := findAccount(accounts, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(account.Transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}
	fmt.Println("\nTransaction History:")
	for _, t := range account.Transactions {
		fmt.Printf("%s - %s: %.2f (Balance: %.2f)\n", t.Timestamp.Format("2006-01-02 15:04:05"), t.Type, t.Amount, t.Balance)
	}
}

func mainMenu() {
	fmt.Println("\n=== Bank Transaction System ===")
	fmt.Println("1. Add Account")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. View Balance")
	fmt.Println("5. View Transaction History")
	fmt.Println("6. Exit")
}

func main() {
	accounts := []Account{}
	for {
		mainMenu()
		choice := promptInput("Enter your choice: ")
		switch choice {
		case "1":
			addAccount(&accounts)
		case "2":
			deposit(&accounts)
		case "3":
			withdraw(&accounts)
		case "4":
			viewBalance(accounts)
		case "5":
			viewTransactionHistory(accounts)
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
