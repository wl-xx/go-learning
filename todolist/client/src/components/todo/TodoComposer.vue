<script setup lang="ts">
import { shallowRef, useTemplateRef } from 'vue'

const emit = defineEmits<{
  add: [content: string]
}>()

defineProps<{
  isSubmitting: boolean
}>()

const content = shallowRef('')
const inputRef = useTemplateRef<HTMLInputElement>('todo-input')

function submitTodo() {
  const nextContent = content.value.trim()
  if (!nextContent) {
    inputRef.value?.focus()
    return
  }

  emit('add', nextContent)
  content.value = ''
  inputRef.value?.focus()
}
</script>

<template>
  <form class="composer" @submit.prevent="submitTodo">
    <span class="composer-mark" aria-hidden="true">+</span>
    <input
      id="todo-input"
      ref="todo-input"
      v-model="content"
      class="composer-input"
      type="text"
      :disabled="isSubmitting"
      autocomplete="off"
      placeholder="接下来要完成什么？"
      aria-label="新的待办事项"
    />
    <button class="add-button" type="submit" :disabled="isSubmitting">
      {{ isSubmitting ? '添加中…' : '添加' }}
    </button>
  </form>
</template>

<style scoped>
.composer {
  display: flex;
  align-items: center;
  gap: 13px;
  padding: 9px 10px 9px 18px;
  border: 1px solid var(--line);
  border-radius: 16px;
  background: var(--paper);
  box-shadow: 0 11px 28px rgba(41, 44, 79, 0.07);
}

.composer:focus-within {
  border-color: #9b9cea;
  box-shadow: 0 0 0 4px rgba(91, 92, 226, 0.1), 0 11px 28px rgba(41, 44, 79, 0.07);
}

.composer-mark {
  color: var(--accent);
  font-size: 1.55rem;
  font-weight: 400;
}

.composer-input {
  min-width: 0;
  flex: 1;
  border: 0;
  outline: 0;
  background: transparent;
  color: var(--ink);
  font: inherit;
  font-size: 1rem;
}

.composer-input::placeholder {
  color: #a7a8b6;
}

.add-button {
  padding: 11px 17px;
  border: 0;
  border-radius: 10px;
  background: var(--accent);
  color: #fff;
  cursor: pointer;
  font: inherit;
  font-size: 0.88rem;
  font-weight: 750;
  transition: background 160ms ease, transform 160ms ease;
}

.add-button:hover {
  background: #4849c5;
  transform: translateY(-1px);
}

.add-button:disabled,
.composer-input:disabled {
  cursor: not-allowed;
  opacity: 0.55;
}
</style>
