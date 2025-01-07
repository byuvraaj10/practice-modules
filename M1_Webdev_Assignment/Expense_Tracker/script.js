class BudgetManager {
  constructor() {
    this.items = JSON.parse(localStorage.getItem("budgetItems")) || [];
    this.chart = null;

    this.init();
  }

  init() {
    document.getElementById("budgetForm").addEventListener("submit", (e) => {
      e.preventDefault();
      this.addItem();
    });
    this.renderItems();
    this.renderChart();
  }

  addItem() {
    const item = document.getElementById("item").value;
    const cost = parseFloat(document.getElementById("cost").value);
    const type = document.getElementById("type").value;

    const newItem = {
      id: Date.now(),
      item,
      cost,
      type,
    };

    this.items.push(newItem);
    this.saveItems();
    this.renderItems();
    this.renderChart();
    document.getElementById("budgetForm").reset();
  }

  deleteItem(id) {
    this.items = this.items.filter((item) => item.id !== id);
    this.saveItems();
    this.renderItems();
    this.renderChart();
  }

  saveItems() {
    localStorage.setItem("budgetItems", JSON.stringify(this.items));
  }

  renderItems() {
    const budgetList = document.getElementById("budgetList");
    budgetList.innerHTML = "";

    this.items.forEach((item) => {
      const row = document.createElement("tr");
      row.innerHTML = `
        <td>${item.item}</td>
        <td>${item.type}</td>
        <td>$${item.cost.toFixed(2)}</td>
        <td>
          <button onclick="budgetManager.deleteItem(${item.id})">Delete</button>
        </td>
      `;
      budgetList.appendChild(row);
    });

    this.renderTotals();
  }

  renderTotals() {
    const totalsDiv = document.getElementById("totals");
    const totals = this.calculateTotals();

    totalsDiv.innerHTML = `
      <p>Essential: $${totals.Essential.toFixed(2)}</p>
      <p>Luxury: $${totals.Luxury.toFixed(2)}</p>
      <p>Savings: $${totals.Savings.toFixed(2)}</p>
    `;
  }

  calculateTotals() {
    return this.items.reduce(
      (totals, item) => {
        totals[item.type] += item.cost;
        return totals;
      },
      { Essential: 0, Luxury: 0, Savings: 0 }
    );
  }

  renderChart() {
    const ctx = document.getElementById("budgetChart").getContext("2d");

    const totals = this.calculateTotals();
    const data = {
      labels: ["Essential", "Luxury", "Savings"],
      datasets: [
        {
          data: [totals.Essential, totals.Luxury, totals.Savings],
          backgroundColor: ["#4CAF50", "#FFA500", "#2196F3"],
        },
      ],
    };

    if (this.chart) this.chart.destroy();

    this.chart = new Chart(ctx, {
      type: "pie",
      data,
      options: {
        responsive: true,
        plugins: {
          legend: {
            position: "bottom",
          },
        },
      },
    });
  }
}

const budgetManager = new BudgetManager();
