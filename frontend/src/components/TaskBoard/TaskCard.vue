<!-- src/components/TaskBoard/TaskCard.vue -->
<template>
  <div class="task-card" :style="taskCardStyle">
    <div class="card-header">
      <h4 class="task-name">{{ task.name }}</h4>
      <span class="game-badge" v-if="!props.isMini">{{ gameName }}</span>
    </div>

    <div class="card-meta" >
      <div class="meta-item" >
        <span class="meta-label">周期</span>
        <span class="meta-value">{{ periodText }}</span>
      </div>
      <div class="meta-item" v-if="!isMini">
        <span class="meta-label">类型</span>
        <span class="meta-value">{{ task.isRepeat ? '重复' : '单次' }}</span>
      </div>
      <div class="meta-item" v-if="!isMini">
        <span class="meta-label">开始</span>
        <span class="meta-value">{{ task.startDate }}</span>
      </div>
      <div class="meta-item">
        <span class="meta-label">完成</span>
        <n-switch size="small" @update:value="changeStatusValue" :value="task.status || 0"
        :checked-value="1" 
        :unchecked-value="0">
        <template #checked>
          完成
        </template>
        <template #unchecked>
          未完成
        </template>
      </n-switch>
      </div>
    </div>

    <p v-if="task.content && !isMini" class="task-content">{{ task.content }}</p>

    <!-- 悬停显示的操作按钮 -->
    <div class="card-actions" v-if="!props.isMini">
      <button class="action-btn edit-btn" title="编辑" @click.stop="$emit('edit')">
        <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7" />
          <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z" />
        </svg>
      </button>
      <button class="action-btn delete-btn" title="删除" @click.stop="$emit('delete')">
        <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="3 6 5 6 21 6" />
          <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed, watch,ref  } from 'vue'
import { TaskService } from '../../../bindings/gameplan/game'

const props = defineProps({
  task: { type: Object, required: true },
  gameName: { type: String, default: '' },
  isMini: { type: Boolean, default: false },
})

defineEmits(['edit', 'delete'])

const periodText = computed(() => {
  const day = props.task.checkDay
  switch (day) {
    case 1: return '每天'
    case 7: return '每周'
    case 14: return '双周'
    case 30: return '每月'
    default: return `每${day}天`
  }
})

const changeStatusValue = (val) => {
  props.task.status = val
  TaskService.UpdateTask(props.task).then(res=>{

  })
}

const taskCardStyle=ref({
      background: "#fffefb",
      border: "1px solid #eef0eb",
      borderRadius: "12px",
      padding: "16px 16px 14px",
      position: "relative",
      transition: "all 0.18s ease",
      cursor: "default",
      display: "flex",
      flexDirection: "column",
      gap: "10px"
})

const changeTaskCardStyle = () => {
  if(props.isMini){
    taskCardStyle.value = {
      border: "1px solid #eef0eb",
      borderRadius: "12px",
      padding: "2px 12px 2px",
      gap:"0px"
    }
  }else{
    taskCardStyle.value = {
      background: "#fffefb",
      border: "1px solid #eef0eb",
      borderRadius: "12px",
      padding: "16px 16px 14px",
      position: "relative",
      transition: "all 0.18s ease",
      cursor: "default",
      display: "flex",
      flexDirection: "column",
      gap: "10px"
    }
  }
}

watch(() => props.isMini, (newVal, oldVal) => {
  changeTaskCardStyle()
});
</script>

<style scoped>

.task-card:hover {
  border-color: #d5dbd0;
  box-shadow: 0 4px 14px rgba(110, 110, 100, 0.06);
  transform: translateY(-1px);
}

/* 头部：名称 + 游戏标签 */
.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
}

.task-name {
  font-size: 14px;
  font-weight: 600;
  color: #3d4045;
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  letter-spacing: -0.01em;
}

.game-badge {
  font-size: 11px;
  font-weight: 500;
  color: #7a9a6b;
  background: #eaf0e5;
  padding: 2px 8px;
  border-radius: 12px;
  white-space: nowrap;
  flex-shrink: 0;
}

/* 元信息区 */
.card-meta {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.meta-label {
  font-size: 11px;
  font-weight: 500;
  color: #b0b5aa;
  text-transform: uppercase;
  letter-spacing: 0.02em;
}

.meta-value {
  font-size: 12.5px;
  font-weight: 500;
  color: #5c6058;
}

/* 任务内容简介 */
.task-content {
  font-size: 12.5px;
  color: #8b9086;
  margin: 0;
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* 操作按钮（hover 才显示） */
.card-actions {
  position: absolute;
  top: 10px;
  right: 12px;
  display: flex;
  gap: 4px;
  opacity: 0;
  transition: opacity 0.15s ease;
}

.task-card:hover .card-actions {
  opacity: 1;
}

.action-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border: none;
  border-radius: 7px;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(4px);
  color: #a0a89a;
  cursor: pointer;
  transition: all 0.14s ease;
  padding: 0;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
}

.action-btn:hover {
  background: #f1f3ee;
  color: #6b7264;
}

.delete-btn:hover {
  background: #f5e8e8;
  color: #c06b6b;
}
</style>