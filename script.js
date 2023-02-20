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
