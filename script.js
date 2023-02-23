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
        this.taskList = document.querySelector("#taskList");
        this.newTaskField = document.querySelector("#addTaskField");

        this.newTaskField.addEventListener('keydown', this.newTask);
        document.querySelector('#addTaskBtn').addEventListener("click", this.newTask);
    }


    newListItem(task) {
        const item = document.createElement('div');
        item.classList.add('task-item');
        item.dataset.id = task.id;

        const info = document.createElement('div');
        info.classList.add('task-info');

        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.classList.add('task-status');
        checkbox.title = 'Mark as done';
        if(task.status) {
            checkbox.checked = true;
        }
        checkbox.dataset.id = task.id;
        checkbox.addEventListener('change', this.toggleTaskStatus)
        info.appendChild(checkbox);

        const itemTitle = document.createElement('div');
        itemTitle.classList.add('task-title');
        itemTitle.textContent = task.title;
        info.appendChild(itemTitle);

        item.appendChild(info);

        const control = document.createElement('div');
        control.classList.add('task-control');

        const btnDel = document.createElement('button');
        btnDel.classList.add('btn', 'btn-delete');
        btnDel.title = 'Delete task';
        btnDel.textContent = 'x';
        btnDel.dataset.id = task.id;
        btnDel.addEventListener('click', this.removeTask);

        control.appendChild(btnDel);
        item.appendChild(control);

        return item;
    }


    newTask = (e) => {
        switch (e.type) {
            case 'keydown':
                if (e.key !== 'Enter') { return; }
                // fallthrough

            case 'click':
                const title = this.newTaskField.value;
                if(title !== '') {
                    this.newTaskField.value = '';
                    const task = this.#ts.add(title);
                    const item = this.newListItem(task);
                    this.taskList.appendChild(item);
                }
                break;
        }
    }

    removeTask = (e) => {
        const id = e.target.dataset.id
        if(this.#ts.remove(id)) {
            const task = this.taskList.querySelector(`.task-item[data-id="${id}"]`);
            this.taskList.removeChild(task);
        }
    }

    toggleTaskStatus = (e) => {
        const id = e.target.dataset.id;
        this.#ts.toggleStatus(id);
    }
}

const app = new App();