const API_BASE_URL = '/api/todo';

// DOM элементы
const todoForm = document.getElementById('todoForm');
const todoInput = document.getElementById('todoInput');
const todoList = document.getElementById('todoList');
const loading = document.getElementById('loading');
const totalCount = document.getElementById('totalCount');
const completedCount = document.getElementById('completedCount');
const remainingCount = document.getElementById('remainingCount');

// Загрузка todos при загрузке страницы
document.addEventListener('DOMContentLoaded', () => {
    loadTodos();
});

// Обработка формы добавления todo
todoForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const body = todoInput.value.trim();
    if (!body) return;

    try {
        await createTodo(body);
        todoInput.value = '';
        loadTodos();
    } catch (error) {
        showError('Не удалось создать задачу');
    }
});

// Загрузка всех todos
async function loadTodos() {
    try {
        loading.style.display = 'block';
        todoList.innerHTML = '';
        
        const response = await fetch(API_BASE_URL);
        if (!response.ok) {
            throw new Error('Ошибка загрузки');
        }
        
        const todos = await response.json();
        loading.style.display = 'none';
        
        if (todos.length === 0) {
            showEmptyState();
        } else {
            renderTodos(todos);
            updateStats(todos);
        }
    } catch (error) {
        loading.style.display = 'none';
        showError('Не удалось загрузить задачи');
    }
}

// Создание нового todo
async function createTodo(body) {
    const response = await fetch(API_BASE_URL, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ body }),
    });
    
    if (!response.ok) {
        throw new Error('Ошибка создания');
    }
}

// Обновление todo (отметить как выполненное)
async function updateTodo(id) {
    try {
        const response = await fetch(`${API_BASE_URL}/${id}`, {
            method: 'PATCH',
        });
        
        if (!response.ok) {
            throw new Error('Ошибка обновления');
        }
        
        loadTodos();
    } catch (error) {
        showError('Не удалось обновить задачу');
    }
}

// Удаление todo
async function deleteTodo(id) {
    if (!confirm('Вы уверены, что хотите удалить эту задачу?')) {
        return;
    }
    
    try {
        const response = await fetch(`${API_BASE_URL}/${id}`, {
            method: 'DELETE',
        });
        
        if (!response.ok) {
            throw new Error('Ошибка удаления');
        }
        
        loadTodos();
    } catch (error) {
        showError('Не удалось удалить задачу');
    }
}

// Отображение todos
function renderTodos(todos) {
    todoList.innerHTML = '';
    
    todos.forEach(todo => {
        const todoItem = document.createElement('div');
        todoItem.className = `todo-item ${todo.completed ? 'completed' : ''}`;
        todoItem.innerHTML = `
            <input 
                type="checkbox" 
                class="todo-checkbox" 
                ${todo.completed ? 'checked' : ''}
                onchange="updateTodo(${todo.id})"
            >
            <span class="todo-text">${escapeHtml(todo.body)}</span>
            <div class="todo-actions">
                ${!todo.completed ? `
                    <button class="btn btn-complete" onclick="updateTodo(${todo.id})">
                        Выполнено
                    </button>
                ` : ''}
                <button class="btn btn-delete" onclick="deleteTodo(${todo.id})">
                    Удалить
                </button>
            </div>
        `;
        todoList.appendChild(todoItem);
    });
}

// Обновление статистики
function updateStats(todos) {
    const total = todos.length;
    const completed = todos.filter(t => t.completed).length;
    const remaining = total - completed;
    
    totalCount.textContent = total;
    completedCount.textContent = completed;
    remainingCount.textContent = remaining;
}

// Показать пустое состояние
function showEmptyState() {
    todoList.innerHTML = `
        <div class="empty-state">
            <div class="empty-state-icon">📋</div>
            <div class="empty-state-text">Нет задач. Добавьте первую!</div>
        </div>
    `;
    updateStats([]);
}

// Показать ошибку
function showError(message) {
    const errorDiv = document.createElement('div');
    errorDiv.className = 'error-message';
    errorDiv.textContent = message;
    todoList.insertBefore(errorDiv, todoList.firstChild);
    
    setTimeout(() => {
        errorDiv.remove();
    }, 5000);
}

// Экранирование HTML
function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

// Экспорт функций для использования в onclick
window.updateTodo = updateTodo;
window.deleteTodo = deleteTodo;
