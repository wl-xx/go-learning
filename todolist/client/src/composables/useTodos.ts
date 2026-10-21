import { computed, readonly, ref } from 'vue'
import {
  clearCompletedTodos,
  createTodo,
  deleteTodo,
  fetchTodos,
  updateTodo,
} from '@/services/todo-api'
import type { Todo, TodoFilter } from '@/types/todo'

function compareTodoByIdDesc(first: Todo, second: Todo): number {
  if (/^\d+$/.test(first.id) && /^\d+$/.test(second.id)) {
    const firstId = BigInt(first.id)
    const secondId = BigInt(second.id)
    return firstId === secondId ? 0 : firstId > secondId ? -1 : 1
  }

  return second.id.localeCompare(first.id, undefined, { numeric: true })
}

export function useTodos() {
  const todos = ref<Todo[]>([])
  const filter = ref<TodoFilter>('all')
  const isLoading = ref(false)
  const errorMessage = ref('')
  const pendingRequestCount = ref(0)
  const isMutating = computed(() => pendingRequestCount.value > 0)

  const visibleTodos = computed(() => {
    let filteredTodos = todos.value

    if (filter.value === 'active') {
      filteredTodos = todos.value.filter((todo) => !todo.completed)
    }

    if (filter.value === 'completed') {
      filteredTodos = todos.value.filter((todo) => todo.completed)
    }

    return [...filteredTodos].sort(compareTodoByIdDesc)
  })

  const remainingCount = computed(() => todos.value.filter((todo) => !todo.completed).length)
  const completedCount = computed(() => todos.value.length - remainingCount.value)

  async function loadTodos() {
    isLoading.value = true
    errorMessage.value = ''

    try {
      todos.value = await fetchTodos()
    } catch (error) {
      errorMessage.value = error instanceof Error ? error.message : '加载待办失败'
    } finally {
      isLoading.value = false
    }
  }

  async function runMutation<T>(action: () => Promise<T>): Promise<T | null> {
    pendingRequestCount.value += 1
    errorMessage.value = ''

    try {
      return await action()
    } catch (error) {
      errorMessage.value = error instanceof Error ? error.message : '待办操作失败'
      return null
    } finally {
      pendingRequestCount.value -= 1
    }
  }

  async function addTodo(content: string) {
    const normalizedContent = content.trim()
    if (!normalizedContent) return

    const todo = await runMutation(() => createTodo({ content: normalizedContent }))
    if (todo) {
      todos.value = [todo, ...todos.value]
    }
  }

  async function toggleTodo(id: string) {
    const currentTodo = todos.value.find((todo) => todo.id === id)
    if (!currentTodo) return

    const todo = await runMutation(() =>
      updateTodo(id, { completed: !currentTodo.completed }),
    )
    if (todo) {
      todos.value = todos.value.map((item) => (item.id === id ? todo : item))
    }
  }

  async function removeTodo(id: string) {
    const result = await runMutation(() => deleteTodo(id))
    if (result !== null) {
      todos.value = todos.value.filter((todo) => todo.id !== id)
    }
  }

  async function clearCompleted() {
    const remainingTodos = await runMutation(clearCompletedTodos)
    if (remainingTodos) {
      todos.value = remainingTodos
    }
  }

  return {
    todos: readonly(todos),
    filter,
    isLoading: readonly(isLoading),
    isMutating,
    errorMessage: readonly(errorMessage),
    visibleTodos,
    remainingCount,
    completedCount,
    loadTodos,
    addTodo,
    toggleTodo,
    removeTodo,
    clearCompleted,
  }
}
