from datetime import datetime
from dataclasses import dataclass
from typing import List
from Customer_Management import Customer  
from Book_Management import BookManager  
from Customer_Management import CustomerManager  

@dataclass
class Transaction:
    """
    Represents a sales transaction, linking customer and book information.
    Adds transaction-specific details.
    """
    customer_name: str
    customer_email: str
    customer_phone: str
    book_title: str
    quantity_sold: int
    transaction_date: datetime = datetime.now()
    total_amount: float = 0.0

    def display_transaction(self) -> str:
        """Returns a formatted string of transaction details."""
        return (f"Transaction Date: {self.transaction_date:%Y-%m-%d %H:%M}\n"
                f"Customer: {self.customer_name}\n"
                f"Book: {self.book_title}\n"
                f"Quantity: {self.quantity_sold}\n"
                f"Total Amount: ${self.total_amount:.2f}")

class SalesManager:
    """Manages sales transactions and related operations."""
    def __init__(self, book_manager: BookManager, customer_manager: CustomerManager):
        self.transactions: List[Transaction] = []
        self.book_manager = book_manager
        self.customer_manager = customer_manager

    def create_sale(self, customer_email: str, book_title: str, quantity: int) -> Transaction:
        """Creates a new sale transaction."""
        # Find customer by email
        customer = self.customer_manager.find_customer(customer_email)
        if not customer:
            raise ValueError("Customer not found")

        # Retrieve the book from the inventory
        book = self.book_manager.get_book(book_title)
        if not book:
            raise ValueError("Book not found")

        # Validate the quantity to be sold
        if quantity <= 0:
            raise ValueError("Quantity must be positive")
        if quantity > book.quantity:
            raise ValueError(f"Error: Only {book.quantity} copies available")

        # Update book quantity after the sale
        book.update_quantity(quantity)

        # Calculate total amount for the transaction
        total_amount = book.price * quantity

        # Create a new transaction
        transaction = Transaction(
            customer_name=customer.name,
            customer_email=customer.email,
            customer_phone=customer.phone,
            book_title=book_title,
            quantity_sold=quantity,
            total_amount=total_amount
        )

        # Append the transaction to the list
        self.transactions.append(transaction)
        return transaction

    def list_transactions(self) -> List[str]:
        """Returns a list of all transactions in a formatted string."""
        return [transaction.display_transaction() for transaction in self.transactions]

    def get_transaction_summary(self) -> str:
        """Generates a summary of all transactions."""
        total_sales = sum(transaction.total_amount for transaction in self.transactions)
        total_transactions = len(self.transactions)
        return f"Total Sales: ${total_sales:.2f}\nTotal Transactions: {total_transactions}"
