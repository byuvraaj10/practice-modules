from typing import List, Optional, Dict
from dataclasses import dataclass
import re

@dataclass
class Book:
    """
    Represents a book in the inventory with essential attributes and methods.
    """
    title: str
    author: str
    price: float
    quantity: int

    def __post_init__(self):
        """Validates the book data after initialization."""
        if self.price <= 0:
            raise ValueError("Price must be a positive number")
        if self.quantity < 0:
            raise ValueError("Quantity cannot be negative")
        if not self.title or not self.author:
            raise ValueError("Title and author cannot be empty")

    def display_details(self) -> str:
        """Returns a formatted string of book details."""
        return f"Title: {self.title}\nAuthor: {self.author}\nPrice: ${self.price:.2f}\nQuantity: {self.quantity}"

    def update_quantity(self, sold_quantity: int) -> None:
        """Updates book quantity after a sale."""
        if sold_quantity > self.quantity:
            raise ValueError(f"Error: Only {self.quantity} copies available")
        self.quantity -= sold_quantity

class BookManager:
    """Manages the book inventory and related operations."""
    
    def __init__(self):
        self.books: Dict[str, Book] = {}

    def add_book(self, title: str, author: str, price: float, quantity: int) -> None:
        """Adds a new book to the inventory."""
        title = title.strip()
        if title in self.books:
            raise ValueError("Book already exists in inventory")
        self.books[title] = Book(title, author, price, quantity)

    def search_book(self, query: str) -> List[Book]:
        """Searches for books by title or author."""
        query = query.lower().strip()
        found_books = [book for book in self.books.values() 
                       if query in book.title.lower() or query in book.author.lower()]
        if not found_books:
            raise ValueError(f"No books found for '{query}'")
        return found_books

    def get_book(self, title: str) -> Optional[Book]:
        """Retrieves a book by its exact title."""
        book = self.books.get(title.strip())
        if not book:
            raise ValueError(f"Book '{title}' not found in inventory")
        return book

    def list_books(self) -> List[str]:
        """Lists all books in the inventory."""
        return [book.display_details() for book in self.books.values()]

    def filter_books_by_price(self, min_price: float = 0, max_price: float = float('inf')) -> List[Book]:
        """Filters books by price range."""
        return [book for book in self.books.values() 
                if min_price <= book.price <= max_price]

    def sort_books_by_title(self, reverse: bool = False) -> List[Book]:
        """Sorts books alphabetically by title."""
        return sorted(self.books.values(), key=lambda book: book.title, reverse=reverse)

    def sort_books_by_price(self, reverse: bool = False) -> List[Book]:
        """Sorts books by price."""
        return sorted(self.books.values(), key=lambda book: book.price, reverse=reverse)

    def update_book_quantity(self, title: str, sold_quantity: int) -> None:
        """Updates quantity for a specific book after a sale."""
        book = self.get_book(title)
        book.update_quantity(sold_quantity)
