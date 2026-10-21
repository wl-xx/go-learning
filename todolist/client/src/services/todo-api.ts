import type {
  ApiResponse,
  CreateTodoPayload,
  Todo,
  UpdateTodoPayload,
} from '@/types/todo'

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(`/api${path}`, {
    headers: {
      'Content-Type': 'application/json',
      ...init?.headers,
    },
    ...init,
  })

  let payload: ApiResponse<T>

  try {
    payload = (await response.json()) as ApiResponse<T>
  } catch {
    throw new Error('服务端返回了无法识别的响应')
  }

  if (!response.ok || !payload.success) {
    throw new Error(payload.msg || '待办操作失败')
  }

  return payload.data
}

export function fetchTodos(signal?: AbortSignal): Promise<Todo[]> {
  return request<Todo[]>('/list', { signal })
}

export function createTodo(payload: CreateTodoPayload): Promise<Todo> {
  return request<Todo>('/todo', {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}

export function updateTodo(id: string, payload: UpdateTodoPayload): Promise<Todo> {
  return request<Todo>(`/todo/${encodeURIComponent(id)}`, {
    method: 'PATCH',
    body: JSON.stringify(payload),
  })
}

export async function deleteTodo(id: string): Promise<void> {
  await request<null>(`/todo/${encodeURIComponent(id)}`, {
    method: 'DELETE',
  })
}

export function clearCompletedTodos(): Promise<Todo[]> {
  return request<Todo[]>('/completed', {
    method: 'DELETE',
  })
}
