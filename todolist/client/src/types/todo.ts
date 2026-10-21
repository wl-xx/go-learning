export interface Todo {
  id: string
  content: string
  completed: boolean
}

export interface CreateTodoPayload {
  content: string
}

export interface UpdateTodoPayload {
  completed: boolean
}

export type TodoFilter = 'all' | 'active' | 'completed'

export interface ApiResponse<T> {
  success: boolean
  data: T
  msg?: string
}
