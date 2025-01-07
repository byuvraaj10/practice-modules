class TaskManager {
  constructor() {
    this.tasks = JSON.parse(localStorage.getItem("taskData")) || [];
    this.taskInput = document.getElementById("newTask");
    this.addButton = document.getElementById("addButton");
    this.tasksContainer = document.getElementById("tasksContainer");
    this.taskSummary = document.getElementById("taskSummary");

    this.initialize();
  }

  initialize() {
    this.addButton.addEventListener("click", () => this.addTask());
    this.renderTasks();
  }

  addTask() {
    const taskName = this.taskInput.value.trim();
    if (taskName) {
      this.tasks.push({ name: taskName, completed: false });
      this.updateLocalStorage();
      this.renderTasks();
      this.taskInput.value = "";
    }
  }

  toggleTask(index) {
    this.tasks[index].completed = !this.tasks[index].completed;
    this.updateLocalStorage();
    this.renderTasks();
  }

  editTask(index) {
    const newName = prompt("Edit your task:", this.tasks[index].name);
    if (newName) {
      this.tasks[index].name = newName.trim();
      this.updateLocalStorage();
      this.renderTasks();
    }
  }

  deleteTask(index) {
    this.tasks.splice(index, 1);
    this.updateLocalStorage();
    this.renderTasks();
  }

  updateLocalStorage() {
    localStorage.setItem("taskData", JSON.stringify(this.tasks));
  }

  renderTasks() {
    this.tasksContainer.innerHTML = "";
    this.tasks.forEach((task, index) => {
      const taskElement = document.createElement("li");
      taskElement.className = task.completed ? "completed" : "";

      taskElement.innerHTML = `
        <span>${task.name}</span>
        <div class="actions">
          <button onclick="taskManager.toggleTask(${index})">✔</button>
          <button onclick="taskManager.editTask(${index})">✏</button>
          <button onclick="taskManager.deleteTask(${index})">✖</button>
        </div>
      `;

      this.tasksContainer.appendChild(taskElement);
    });

    const pendingTasks = this.tasks.filter((task) => !task.completed).length;
    this.taskSummary.innerText = pendingTasks
      ? `Pending Tasks: ${pendingTasks}`
      : "No tasks yet!";
  }
}

const taskManager = new TaskManager();
