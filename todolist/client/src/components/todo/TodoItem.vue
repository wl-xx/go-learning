<script setup lang="ts">
import { computed } from 'vue'
import type { Todo } from '@/types/todo'

const props = defineProps<{
  todo: Todo
  isDisabled: boolean
}>()

const emit = defineEmits<{
  toggle: [id: string]
  remove: [id: string]
}>()

const itemClasses = computed(() => ({
  'todo-item': true,
  'todo-item--completed': props.todo.completed,
}))
</script>

<template>
  <li :class="itemClasses">
    <button
      class="check-button"
      type="button"
      :aria-label="todo.completed ? '标记为未完成' : '标记为已完成'"
      :aria-pressed="todo.completed"
      :disabled="isDisabled"
      @click="emit('toggle', todo.id)"
    >
      <svg v-if="todo.completed" viewBox="0 0 16 16" aria-hidden="true">
        <path d="m3.2 8.3 3 3 6.6-6.7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.1" />
      </svg>
    </button>
    <span class="todo-content">{{ todo.content }}</span>
    <button class="remove-button" type="button" aria-label="删除待办" :disabled="isDisabled" @click="emit('remove', todo.id)">
      <svg viewBox="0 0 16 16" aria-hidden="true">
        <path d="m4.4 4.4 7.2 7.2m0-7.2-7.2 7.2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.8" />
      </svg>
    </button>
  </li>
</template>

<style scoped>
.todo-item {
  display: flex;
  align-items: center;
  gap: 14px;
  min-height: 63px;
  padding: 12px 13px 12px 17px;
  border: 1px solid transparent;
  border-radius: 13px;
  background: rgba(255, 255, 255, 0.55);
  transition: background 160ms ease, border-color 160ms ease;
}

.todo-item:hover {
  border-color: var(--line);
  background: var(--paper);
}

.check-button {
  display: grid;
  width: 22px;
  height: 22px;
  flex: 0 0 22px;
  place-items: center;
  padding: 0;
  border: 1.5px solid #b8b9c9;
  border-radius: 50%;
  background: transparent;
  color: #fff;
  cursor: pointer;
  transition: border-color 160ms ease, background 160ms ease;
}

.check-button:disabled,
.remove-button:disabled {
  cursor: not-allowed;
}

.check-button svg,
.remove-button svg {
  width: 15px;
  height: 15px;
}

.todo-item--completed .check-button {
  border-color: var(--accent);
  background: var(--accent);
}

.todo-content {
  min-width: 0;
  flex: 1;
  color: var(--ink);
  font-size: 0.98rem;
  line-height: 1.5;
  overflow-wrap: anywhere;
  transition: color 160ms ease, text-decoration-color 160ms ease;
}

.todo-item--completed .todo-content {
  color: #a2a3b0;
  text-decoration: line-through;
  text-decoration-color: #c5c6cf;
}

.remove-button {
  display: grid;
  width: 30px;
  height: 30px;
  flex: 0 0 30px;
  place-items: center;
  padding: 0;
  border: 0;
  border-radius: 8px;
  background: transparent;
  color: #b3b4bf;
  cursor: pointer;
  opacity: 0;
  transition: opacity 160ms ease, color 160ms ease, background 160ms ease;
}

.todo-item:hover .remove-button,
.remove-button:focus-visible {
  opacity: 1;
}

.remove-button:hover {
  background: #fff0ed;
  color: #c35b4b;
}

@media (hover: none) {
  .remove-button {
    opacity: 1;
  }
}
</style>
