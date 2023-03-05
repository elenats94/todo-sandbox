class TaskStore {
    #api = '';

    constructor(apiUrl) {
        this.#api = apiUrl;
    }

    async list() {
        const response = await fetch(this.#api, {
            credentials: "include"
        });
        if (response.ok) {
            return await response.json();
        }
    }

    async get(id) {
        const response = await fetch(`${this.#api}/${id}`, {
            credentials: "include"
        });
        if (response.ok) {
            return await response.json();
        }

        const data = await response.json();
        throw new Error(data.error);
    }

    async add(title) {
        const response = await fetch(`${this.#api}`, {
            method: "POST",
            credentials: "include",
            body: JSON.stringify({title: title})
        });
        if (response.ok) {
            return await response.json();
        }
        const data = await response.json();
        throw new Error(data.error);
    }

    async remove(id) {
        const response = await fetch(`${this.#api}/${id}`, {
            method: "DELETE",
            credentials: "include"
        });
        if (response.ok) {
            return await response.json()
        }

        const data = await response.json()
        throw new Error(data.error)
    }

    async toggleStatus(id) {
        const response = await fetch(`${this.#api}/${id}`, {
            method: "PATCH",
            credentials: "include"
        });
        if (response.ok) {
            return await response.json();
        }

        const data = await response.json();
        throw new Error(data.error);
    }

    async edit(id, title) {
        const response = await fetch(`${this.#api}/${id}`, {
            method: "PUT",
            credentials: "include",
            body: JSON.stringify({title: title})
        })

        if (response.ok) {
            return await response.json();
        }

        const data = await response.json();
        throw new Error(data.error);
    }
}

class App {
    #ts = null;

    constructor(apiUrl) {
        this.#ts = new TaskStore(apiUrl)

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
        if (task.status) {
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

        const btnEdit = document.createElement('button');
        btnEdit.classList.add('btn', 'btn-edit');
        btnEdit.title = 'Edit task';
        btnEdit.textContent = 'edit'
        btnEdit.dataset.id = task.id;
        btnEdit.addEventListener('click', this.showEditField);
        control.appendChild(btnEdit);

        const btnDel = document.createElement('button');
        btnDel.classList.add('btn', 'btn-delete');
        btnDel.title = 'Delete task';
        btnDel.textContent = 'delete';
        btnDel.dataset.id = task.id;
        btnDel.addEventListener('click', this.removeTask);

        control.appendChild(btnDel);
        item.appendChild(control);

        return item;
    }

    getListItem(id) {
        return this.taskList.querySelector(`.task-item[data-id="${id}"]`);
    }

    listTasks() {
        this.#ts.list()
            .then(tasks => {
                for (const task of tasks) {
                    const item = this.newListItem(task);
                    this.taskList.append(item);
                }
            })
            .catch(err => console.log('Error occured:', err.message));
    }

    newTask = (e) => {
        switch (e.type) {
            case 'keydown':
                if (e.key !== 'Enter') {
                    return;
                }
            // fallthrough

            case 'click':
                const title = this.newTaskField.value;
                if (title !== '') {
                    this.#ts.add(title)
                        .then(task => {
                            this.newTaskField.value = '';
                            const item = this.newListItem(task);
                            this.taskList.appendChild(item);
                        })
                        .catch(err => console.log('Error occurred:', err.message));
                }
        }
    }

    removeTask = (e) => {
        this.#ts.remove(e.target.dataset.id)
            .then(task => {
                const item = this.getListItem(task.id);
                this.taskList.removeChild(item);
            })
            .catch(err => console.log('Error occurred:', err.message));
    }

    toggleTaskStatus = (e) => {
        this.#ts.toggleStatus(e.target.dataset.id)
            .then(task => {
                this.getListItem(task.id).querySelector(`.btn-edit`).hidden = e.target.checked;
            })
            .catch(err => console.log('Error occurred:', err.message));
    }

    showEditField = (e) => {
        this.#ts.get(e.target.dataset.id)
            .then(task => {
                const editField = document.createElement('input');
                editField.type = 'text';
                editField.classList.add('text-field');
                editField.dataset.id = task.id;
                editField.value = task.title;
                editField.addEventListener('keydown', this.editTask);

                const item = this.getListItem(task.id);

                item.replaceChildren(editField);
                editField.select();
            })
            .catch(err => console.log('Error occurred:', err.message));
    }

    editTask = (e) => {
        const id = e.target.dataset.id;

        switch (e.key) {
            case 'Enter':
                this.#ts.edit(id, e.target.value)
                    .then(task => {
                        this.taskList.replaceChild(this.newListItem(task), e.target.parentElement);
                    })
                    .catch(err => console.log('Error occurred:', err.message));
                break;
            case 'Esc':
            case 'Escape':
                this.#ts.get(id)
                    .then(task => {
                        this.taskList.replaceChild(this.newListItem(task), e.target.parentElement)
                    })
                    .catch(err => console.log('Error occurred:', err.message));
        }
    }
}

document.addEventListener('DOMContentLoaded', () => {
    const app = new App('http://localhost:8080/tasks');
    app.listTasks();
});