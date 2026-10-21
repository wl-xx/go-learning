<script setup lang="ts">
import { onMounted } from 'vue'
import TodoComposer from './TodoComposer.vue'
import TodoFooter from './TodoFooter.vue'
import TodoList from './TodoList.vue'
import { useTodos } from '@/composables/useTodos'
import type { TodoFilter } from '@/types/todo'

const {
  filter,
  isLoading,
  isMutating,
  errorMessage,
  visibleTodos,
  remainingCount,
  completedCount,
  loadTodos,
  addTodo,
  toggleTodo,
  removeTodo,
  clearCompleted,
} = useTodos()

function changeFilter(nextFilter: TodoFilter) {
  filter.value = nextFilter
}

onMounted(loadTodos)
</script>

<template>
  <main class="todo-page">
    <section class="todo-shell" aria-labelledby="app-title">
      <header class="hero">
        <div>
          <p class="eyebrow">FOCUS LIST</p>
          <h1 id="app-title" class="hero-title">今天，完成一件重要的事。</h1>
          <p class="hero-description">把注意力留给真正值得完成的任务。</p>
        </div>
        <div class="date-card" aria-label="今日日期">
          <span class="date-weekday">TODAY</span>
          <strong class="date-day">{{ new Date().getDate() }}</strong>
          <span class="date-month">{{ new Intl.DateTimeFormat('zh-CN', { month: 'long' }).format(new Date()) }}</span>
        </div>
      </header>

      <TodoComposer :is-submitting="isMutating" @add="addTodo" />

      <p v-if="isMutating" class="sync-note" role="status">
        正在同步待办事项…
      </p>

      <div v-if="errorMessage" class="error-state" role="alert">
        <span>{{ errorMessage }}</span>
        <button class="text-button" type="button" @click="loadTodos">重新加载</button>
      </div>

      <TodoList
        :todos="visibleTodos"
        :is-loading="isLoading"
        :is-mutating="isMutating"
        @toggle="toggleTodo"
        @remove="removeTodo"
      />

      <TodoFooter
        :filter="filter"
        :remaining-count="remainingCount"
        :completed-count="completedCount"
        :is-mutating="isMutating"
        @change-filter="changeFilter"
        @clear-completed="clearCompleted"
      />
    </section>
  </main>
</template>

<style scoped>
.todo-page {
  min-height: 100vh;
  padding: 72px 24px;
}

.todo-shell {
  width: min(100%, 760px);
  margin: 0 auto;
}

.hero {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 32px;
  margin-bottom: 38px;
}

.eyebrow {
  margin: 0 0 12px;
  color: var(--accent);
  font-size: 0.74rem;
  font-weight: 800;
  letter-spacing: 0.16em;
}

.hero-title {
  max-width: 560px;
  margin: 0;
  color: var(--ink);
  font-size: clamp(2rem, 5vw, 3.45rem);
  letter-spacing: -0.06em;
  line-height: 1.08;
}

.hero-description {
  margin: 17px 0 0;
  color: var(--muted);
  font-size: 1rem;
}

.date-card {
  display: grid;
  min-width: 90px;
  padding: 15px 12px 13px;
  border: 1px solid var(--line);
  border-radius: 18px;
  background: rgba(255, 255, 255, 0.74);
  color: var(--muted);
  text-align: center;
}

.date-weekday {
  color: var(--accent);
  font-size: 0.62rem;
  font-weight: 800;
  letter-spacing: 0.12em;
}

.date-day {
  margin: 2px 0;
  color: var(--ink);
  font-size: 2.25rem;
  line-height: 1;
}

.date-month {
  font-size: 0.78rem;
}

.sync-note {
  margin: 18px 2px -4px;
  color: #916c20;
  font-size: 0.82rem;
}

.error-state {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  margin-top: 22px;
  padding: 14px 16px;
  border: 1px solid #f3c5bd;
  border-radius: 12px;
  background: #fff5f2;
  color: #a74836;
  font-size: 0.9rem;
}

.text-button {
  border: 0;
  background: transparent;
  color: inherit;
  cursor: pointer;
  font: inherit;
  font-weight: 700;
  text-decoration: underline;
  text-underline-offset: 3px;
}

@media (max-width: 560px) {
  .todo-page {
    padding: 42px 17px;
  }

  .hero {
    gap: 18px;
  }

  .date-card {
    min-width: 76px;
  }
}
</style>
