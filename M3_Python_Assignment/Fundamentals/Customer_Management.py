import re
from dataclasses import dataclass
from typing import List, Optional

@dataclass
class Customer:
    """Represents a customer with basic information."""
    name: str
    email: str
    phone: str

    def __post_init__(self):
        """Validates customer data after initialization."""
        if not self.name or not self.email or not self.phone:
            raise ValueError("All customer fields are required")
        if not re.match(r"[^@]+@[^@]+\.[^@]+", self.email):
            raise ValueError("Invalid email format")
        if not re.match(r"^\+?[\d\s-]{10,}$", self.phone):
            raise ValueError("Invalid phone number format")

    def display_details(self) -> str:
        """Returns a formatted string of customer details."""
        return f"Name: {self.name}\nEmail: {self.email}\nPhone: {self.phone}"

    def update_phone(self, new_phone: str) -> None:
        """Updates the customer's phone number."""
        if not re.match(r"^\+?[\d\s-]{10,}$", new_phone):
            raise ValueError("Invalid phone number format")
        self.phone = new_phone

    def update_email(self, new_email: str) -> None:
        """Updates the customer's email."""
        if not re.match(r"[^@]+@[^@]+\.[^@]+", new_email):
            raise ValueError("Invalid email format")
        self.email = new_email

class CustomerManager:
    """Manages customer records and related operations."""
    def __init__(self):
        self.customers: List[Customer] = []

    def add_customer(self, name: str, email: str, phone: str) -> None:
        """Adds a new customer to the system."""
        customer = Customer(name, email, phone)
        self.customers.append(customer)

    def find_customer(self, email: str) -> Optional[Customer]:
        """Finds a customer by their email."""
        return next((c for c in self.customers if c.email == email), None)

    def find_customer_by_name(self, name: str) -> List[Customer]:
        """Finds customers by their name (case-insensitive)."""
        return [c for c in self.customers if name.lower() in c.name.lower()]

    def remove_customer(self, email: str) -> bool:
        """Removes a customer by their email."""
        customer = self.find_customer(email)
        if customer:
            self.customers.remove(customer)
            return True
        return False

    def update_customer_details(self, email: str, new_email: Optional[str] = None, new_phone: Optional[str] = None) -> bool:
        """Updates customer details (email or phone)."""
        customer = self.find_customer(email)
        if customer:
            if new_email:
                customer.update_email(new_email)
            if new_phone:
                customer.update_phone(new_phone)
            return True
        return False

    def list_customers(self) -> List[str]:
        """Lists all customers in the system."""
        return [customer.display_details() for customer in self.customers]
