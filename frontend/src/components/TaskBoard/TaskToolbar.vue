<!-- src/components/TaskBoard/TaskToolbar.vue -->
<template>
  <div class="task-toolbar">
    <div class="toolbar-left" v-if="!props.isMini">

      <!-- 周期筛选 tabs -->
      <div class="filter-tabs">
        <button
          v-for="tab in filterTabs"
          :key="tab.key"
          class="filter-tab"
          :class="{ active: activeFilter === tab.key }"
          @click="handleTabClick(tab.key)"
        >
          {{ tab.label }}
        </button>
        <!-- 自定义天数输入 -->
        <div v-if="activeFilter === 'custom'" class="custom-days-input">
          <input
            type="number"
            min="1"
            placeholder="天数"
            :value="customDays"
            @input="updateCustomDays"
            @click.stop
          />
          <span class="days-unit">天</span>
        </div>
      </div>
    </div>

    <div class="toolbar-right">
      <!-- 工具按钮 -->
      <button v-if="!props.isMini" class="tool-btn" title="添加任务" @click="$emit('add-task')">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
          <line x1="12" y1="5" x2="12" y2="19" />
          <line x1="5" y1="12" x2="19" y2="12" />
        </svg>
      </button>
      <button v-if="!props.isMini" class="tool-btn" title="分享" @click="$emit('share')">
        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="18" cy="5" r="3" />
          <circle cx="6" cy="12" r="3" />
          <circle cx="18" cy="19" r="3" />
          <line x1="8.59" y1="13.51" x2="15.42" y2="17.49" />
          <line x1="15.41" y1="6.51" x2="8.59" y2="10.49" />
        </svg>
      </button>
      <button v-if="!props.isMini" class="tool-btn" title="导入" @click="$emit('import')">
        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
          <polyline points="17 8 12 3 7 8" />
          <line x1="12" y1="3" x2="12" y2="15" />
        </svg>
      </button>
      <span v-if="props.isMini" style="margin-right: 30px;">未完成任务</span>
      <button class="tool-btn" @click="$emit('mini')" title="迷你模式">
        <!-- 使用 SVG 画一个经典的最小化图标（一条横线） -->
        <svg
          width="16"
          height="16"
          viewBox="0 0 16 16"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <rect x="2" y="7" width="12" height="2" fill="currentColor" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { TaskService } from '../../../bindings/gameplan/game'

const props = defineProps({
  activeFilter: String,
  customDays: [Number, null],
  isMini: [Boolean, false],
})

const emit = defineEmits([
  'update:activeFilter',
  'update:customDays',
  'add-task',
  'share',
  'import',
  'mini'
])

const filterTabs = [
  { key: 'all', label: '全部' },
  { key: 'daily', label: '每日' },
  { key: 'weekly', label: '每周' },
  { key: 'biweekly', label: '双周' },
  { key: 'monthly', label: '月常' },
  { key: 'custom', label: '自定义' },
]

function handleTabClick(key) {
  emit('update:activeFilter', key)
}

function updateCustomDays(event) {
  const val = parseInt(event.target.value, 10)
  if (!isNaN(val) && val > 0) {
    emit('update:customDays', val)
  } else {
    emit('update:customDays', null)
  }
}
</script>

<style scoped>
.task-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 20px;
  background: #f9faf7;
  border-bottom: 1px solid #eef0eb;
  flex-shrink: 0;
  gap: 16px;
}

.toolbar-left {
  display: flex;
  align-items: center;
  gap: 14px;
  flex-wrap: wrap;
}

.add-task-btn {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  border: none;
  border-radius: 10px;
  background: #dce6d4;
  color: #4a6e3a;
  font-size: 13px;
  font-weight: 550;
  cursor: pointer;
  transition: all 0.16s ease;
  white-space: nowrap;
  font-family: inherit;
  letter-spacing: -0.01em;
}

.add-task-btn:hover {
  background: #cfe0c5;
  color: #3a5c2c;
}

.add-task-btn:active {
  transform: scale(0.97);
}

.filter-tabs {
  display: flex;
  align-items: center;
  gap: 2px;
  background: #f1f3ee;
  padding: 3px;
  border-radius: 9px;
}

.filter-tab {
  padding: 5px 14px;
  border: none;
  background: transparent;
  border-radius: 7px;
  font-size: 12.5px;
  font-weight: 500;
  color: #7b8076;
  cursor: pointer;
  transition: all 0.15s ease;
  white-space: nowrap;
  font-family: inherit;
}

.filter-tab:hover {
  color: #4a5244;
  background: #e6e8e2;
}

.filter-tab.active {
  background: #ffffff;
  color: #3a4e32;
  font-weight: 600;
  box-shadow: 0 1px 3px rgba(0,0,0,0.04);
}

.custom-days-input {
  display: flex;
  align-items: center;
  margin-left: 4px;
}

.custom-days-input input {
  width: 50px;
  padding: 4px 8px;
  border: 1px solid #d5dbd0;
  border-radius: 6px;
  font-size: 12.5px;
  outline: none;
  font-family: inherit;
  background: #fff;
  transition: border-color 0.15s;
}

.custom-days-input input:focus {
  border-color: #b5c9a8;
  box-shadow: 0 0 0 2px rgba(165, 185, 150, 0.2);
}

.days-unit {
  font-size: 12px;
  color: #8b9086;
  margin-left: 5px;
}

.toolbar-right {
  display: flex;
  align-items: center;
  gap: 4px;
}

.tool-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border: none;
  border-radius: 8px;
  background: transparent;
  color: #a0a89a;
  cursor: pointer;
  transition: all 0.14s ease;
}

.tool-btn:hover {
  background: #f1f3ee;
  color: #6b7264;
}

.tool-btn:active {
  background: #e6e9e1;
}
</style>