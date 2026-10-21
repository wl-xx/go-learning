<script setup lang="ts">
import TodoItem from './TodoItem.vue'
import type { Todo } from '@/types/todo'

defineProps<{
  todos: readonly Todo[]
  isLoading: boolean
  isMutating: boolean
}>()

const emit = defineEmits<{
  toggle: [id: string]
  remove: [id: string]
}>()
</script>

<template>
  <section class="list-section" aria-live="polite" aria-label="待办清单">
    <div v-if="isLoading" class="loading-state">
      <span class="loading-dot"></span>
      正在整理你的待办事项…
    </div>

    <TransitionGroup v-else-if="todos.length" name="todo" tag="ul" class="todo-list">
      <TodoItem
        v-for="todo in todos"
        :key="todo.id"
        :todo="todo"
        :is-disabled="isMutating"
        @toggle="emit('toggle', $event)"
        @remove="emit('remove', $event)"
      />
    </TransitionGroup>

    <div v-else class="empty-state">
      <span class="empty-icon" aria-hidden="true">✦</span>
      <strong>清单已经很轻盈</strong>
      <span>给自己一点留白，或添加下一件想完成的事。</span>
    </div>
  </section>
</template>

<style scoped>
.list-section {
  margin-top: 24px;
}

.todo-list {
  display: grid;
  gap: 9px;
  margin: 0;
  padding: 0;
  list-style: none;
}

.loading-state,
.empty-state {
  display: grid;
  place-items: center;
  min-height: 210px;
  border: 1px dashed var(--line);
  border-radius: 16px;
  color: var(--muted);
  text-align: center;
}

.loading-state {
  display: flex;
  gap: 10px;
  font-size: 0.92rem;
}

.loading-dot {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: var(--accent);
  animation: pulse 1s ease-in-out infinite alternate;
}

.empty-state {
  align-content: center;
  gap: 9px;
  padding: 24px;
  font-size: 0.88rem;
}

.empty-state strong {
  color: var(--ink);
  font-size: 1rem;
}

.empty-icon {
  color: var(--accent);
  font-size: 1.8rem;
}

.todo-enter-active,
.todo-leave-active {
  transition: all 220ms ease;
}

.todo-enter-from,
.todo-leave-to {
  opacity: 0;
  transform: translateY(9px);
}

@keyframes pulse {
  from { opacity: 0.35; transform: scale(0.82); }
  to { opacity: 1; transform: scale(1.1); }
}
</style>
