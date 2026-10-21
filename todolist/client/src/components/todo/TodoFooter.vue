<script setup lang="ts">
import type { TodoFilter } from '@/types/todo'

defineProps<{
  filter: TodoFilter
  remainingCount: number
  completedCount: number
  isMutating: boolean
}>()

const emit = defineEmits<{
  changeFilter: [filter: TodoFilter]
  clearCompleted: []
}>()

const filters: Array<{ value: TodoFilter; label: string }> = [
  { value: 'all', label: '全部' },
  { value: 'active', label: '待完成' },
  { value: 'completed', label: '已完成' },
]
</script>

<template>
  <footer class="todo-footer">
    <p class="summary">
      <strong>{{ remainingCount }}</strong>
      件待完成
    </p>

    <div class="filters" role="group" aria-label="筛选待办">
      <button
        v-for="option in filters"
        :key="option.value"
        class="filter-button"
        :class="{ 'filter-button--active': filter === option.value }"
        type="button"
        :aria-pressed="filter === option.value"
        :disabled="isMutating"
        @click="emit('changeFilter', option.value)"
      >
        {{ option.label }}
      </button>
    </div>

    <button
      class="clear-button"
      type="button"
      :disabled="completedCount === 0 || isMutating"
      @click="emit('clearCompleted')"
    >
      清除已完成
    </button>
  </footer>
</template>

<style scoped>
.todo-footer {
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  align-items: center;
  gap: 16px;
  margin-top: 21px;
  padding: 0 4px;
  color: var(--muted);
  font-size: 0.82rem;
}

.summary {
  margin: 0;
}

.summary strong {
  color: var(--ink);
}

.filters {
  display: flex;
  gap: 2px;
  padding: 3px;
  border-radius: 9px;
  background: rgba(221, 221, 232, 0.6);
}

.filter-button,
.clear-button {
  border: 0;
  background: transparent;
  color: inherit;
  cursor: pointer;
  font: inherit;
}

.filter-button {
  padding: 6px 9px;
  border-radius: 6px;
  transition: color 140ms ease, background 140ms ease, box-shadow 140ms ease;
}

.filter-button--active {
  background: #fff;
  color: var(--ink);
  box-shadow: 0 1px 4px rgba(50, 51, 80, 0.1);
  font-weight: 700;
}

.clear-button {
  justify-self: end;
  color: #747582;
  text-decoration: underline;
  text-decoration-color: transparent;
  text-underline-offset: 3px;
  transition: color 140ms ease, text-decoration-color 140ms ease;
}

.clear-button:hover:not(:disabled) {
  color: var(--ink);
  text-decoration-color: currentColor;
}

.clear-button:disabled {
  cursor: not-allowed;
  opacity: 0.38;
}

@media (max-width: 560px) {
  .todo-footer {
    grid-template-columns: 1fr auto;
  }

  .filters {
    grid-row: 2;
    grid-column: 1 / -1;
    justify-content: center;
  }
}
</style>
