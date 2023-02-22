class Generator {
    #idCounter = 0;

    generate() {
        const timestamp = Date.now().toString(36);
        this.#idCounter++;
        return `t${this.#idCounter.toString(36)}-${timestamp}`;
    }
}

class Task {
    constructor(id, title, status = false) {
        this.id = id;
        this.title = title;
        this.status = status;
    }
}

class TaskStore {
    #idGen = new Generator();
    #store = new Map();

    constructor() {
        this.#store = new Map();
    }

    get storage() {
        const store = [];
        for (const task of this.#store.values()) {
            store.push(task);
        }

        return store;
    }

    add(title) {
        if (title.length !== 0) {
            const id = this.#idGen.generate();
            const task = new Task(id, title);
            this.#store.set(id, task);

            return task;
        }
    }

    remove(id) {
        return this.#store.delete(id);
    }

    toggleStatus(id) {
        if (this.#store.has(id)) {
            const task = this.#store.get(id);
            task.status = !task.status;
        }
    }
}

class App {
    #ts = new TaskStore();

    constructor() {
        this.taskList = document.getElementById("taskList");
        document.getElementById("newTaskField").addEventListener('keydown', this.newTask);
    }


    newListItem(task) {
        const item = document.createElement('div');
        item.innerHTML = `<span>${task.title}</span>`
        item.classList.add('task-item');
        item.dataset.id = task.id;

        const delBtn = document.createElement('button');
        delBtn.textContent = 'x';
        delBtn.onclick = this.removeTask;
        item.appendChild(delBtn);

        return item;
    }


    newTask = (e) => {
        if (e.key === 'Enter') {
            const title = e.target.value;
            if (title.length !== 0) {
                const task = this.#ts.add(title);
                const item = this.newListItem(task);
                this.taskList.appendChild(item);
            }
            e.target.value = '';
        } else if (e.key === 'Esc') {
            e.target.blur();
        }
    }

    removeTask = (e) => {
        const id = e.target.parentElement.dataset.id;
        if (this.#ts.remove(id)) {
            this.taskList.removeChild(e.target.parentElement);
        }
    }
}

const app = new App();