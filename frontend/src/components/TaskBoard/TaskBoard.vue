<!-- src/components/TaskBoard/TaskBoard.vue -->
<template>
  <div class="task-board">
    <!-- 工具栏 -->
    <TaskToolbar
      :active-filter="activeFilter"
      :custom-days="customDays"
      @update:active-filter="activeFilter = $event"
      @update:custom-days="customDays = $event"
      @add-task="openAddModal"
      @share="handleShare"
      @import="handleImport"
      @mini="handleMini"
      @check="handleCheck"
      :isMini="props.isMini"
    />

    <!-- 任务卡片列表 -->
    <div class="task-list" ref="listRef">
      <TransitionGroup name="card-list" tag="div" class="card-grid">
        <TaskCard
          v-for="task in filteredTasks"
          :key="task.id"
          :task="task"
          :game-name="getGameName(task.gameId)"
          @edit="openEditModal(task)"
          @delete="handleDeleteClick(task)"
          :isMini="props.isMini"
        />
      </TransitionGroup>

      <!-- 空状态 -->
      <div v-if="filteredTasks.length === 0" class="empty-state">
        <div class="empty-icon">📋</div>
        <p class="empty-text">暂无任务</p>
        <p class="empty-hint">点击「添加任务」开始记录</p>
      </div>
    </div>

    <!-- 添加/编辑弹窗 -->
    <TaskEditModal
      v-model:visible="editModalVisible"
      :mode="editMode"
      :task="editingTask"
      :game-list="gameList"
      :default-game-id="selectedGameId"
      @confirm="handleModalConfirm"
      @cancel="editModalVisible = false"
    />

    <!-- 删除确认小窗 -->
    <Teleport to="body">
      <Transition name="fade-scale">
        <div v-if="deleteConfirmVisible" class="confirm-overlay" @click.self="deleteConfirmVisible = false">
          <div class="confirm-dialog">
            <div class="confirm-icon">⚠️</div>
            <p class="confirm-title">确认删除</p>
            <p class="confirm-message">确定要删除任务<br /><strong>「{{ deletingTask?.name }}」</strong>吗？</p>
            <p class="confirm-sub-message">此操作不可撤销</p>
            <div class="confirm-actions">
              <button class="confirm-btn cancel" @click="deleteConfirmVisible = false">取消</button>
              <button class="confirm-btn danger" @click="confirmDelete">确认删除</button>
            </div>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted,watch, render } from 'vue'
import TaskToolbar from './TaskToolbar.vue'
import TaskCard from './TaskCard.vue'
import TaskEditModal from './TaskEditModal.vue'
import { TaskService } from '../../../bindings/gameplan/game'
import { Events } from '@wailsio/runtime'

onMounted(()=>{
  getData()
})

Events.On('task_reset', () => {
  getData()
})

const getData = (status) => {

  let param = {
    gameId:props.selectedGameId
  }
  if(status !== undefined){
    param['status'] = status
  }

  TaskService.GetTaskList(param).then((res)=>{
    taskList.splice(0, taskList.length, ...res[0])
  })
}

const props = defineProps({
  // 游戏列表，用于显示游戏名称（与左侧共用数据源）
  gameList: {
    type: Array,
    default: () => [],
  },
  // 当前选中的游戏ID（可传，新增任务时默认选中）
  selectedGameId: {
    type: [Number, null],
    default: null,
  },
  isMini:{
    type:Boolean,
    default:false
  }
})

// ---------- 任务数据----------
const taskList = reactive([])

// ---------- 筛选状态 ----------
const activeFilter = ref('all') // 'all' | 'daily' | 'weekly' | 'biweekly' | 'monthly' | 'custom'
const customDays = ref(null) // 自定义天数

// 筛选后的任务列表
const filteredTasks = computed(() => {
  if (activeFilter.value === 'all') return taskList

  let targetDays = null
  switch (activeFilter.value) {
    case 'daily': targetDays = 1; break
    case 'weekly': targetDays = 7; break
    case 'biweekly': targetDays = 14; break
    case 'monthly': targetDays = 30; break
    case 'custom':
      targetDays = customDays.value
      if (!targetDays || targetDays < 1) return []
      break
  }

  return taskList.filter(task => task.checkDay === targetDays)
})

// 根据gameId获取游戏名称
function getGameName(gameId) {
  const game = props.gameList.find(g => g.id === gameId)
  return game ? game.name : '未知游戏'
}

// ---------- 编辑模态框 ----------
const editModalVisible = ref(false)
const editMode = ref('add') // 'add' | 'edit'
const editingTask = ref(null)

function openAddModal() {
  editMode.value = 'add'
  editingTask.value = null
  editModalVisible.value = true
}

function openEditModal(task) {
  editMode.value = 'edit'
  editingTask.value = { ...task }
  editModalVisible.value = true
}

function handleModalConfirm(formData) {
  if (editMode.value === 'add') {
    formData.id=0
    TaskService.AddTask(formData).then((res)=>{
      getData()
    })
  } else if (editMode.value === 'edit' && editingTask.value) {
    const task = taskList.find(t => t.id === editingTask.value.id)
    if (task) {
      formData.id = task.id
      formData.status = task.status
      TaskService.UpdateTask(formData).then((res)=>{
        getData()
      })
    }
  }
  editModalVisible.value = false
  editingTask.value = null
}

// ---------- 删除 ----------
const deleteConfirmVisible = ref(false)
const deletingTask = ref(null)

function handleDeleteClick(task) {
  deletingTask.value = task
  deleteConfirmVisible.value = true
}

function confirmDelete() {
  if (deletingTask.value) {
    const index = taskList.findIndex(t => t.id === deletingTask.value.id)
    if (index !== -1) {
      taskList.splice(index, 1)
    }
    TaskService.DelTask(deletingTask.value).then(res=>{
    
    })
  }
  deleteConfirmVisible.value = false
  deletingTask.value = null
}

// ---------- 工具按钮 ----------
function handleShare() {
  TaskService.ShareTask({
    gameId:props.selectedGameId
  }).then(res=>{

  })
}

function handleImport() {
  TaskService.ImportTask({
    gameId:props.selectedGameId
  }).then(res=>{

  })
}

watch(() => props.isMini, (newVal, oldVal) => {
  if(props.isMini){
    TaskService.GetTaskList({
    gameId:props.selectedGameId,status:-1}).then((res)=>{
      taskList.splice(0, taskList.length, ...res[0])
    })
  }else{
    TaskService.GetTaskList({
    gameId:props.selectedGameId,status:0}).then((res)=>{
      taskList.splice(0, taskList.length, ...res[0])
    })
  }
});

watch(() => props.selectedGameId, (newVal, oldVal) => {
  if(props.selectedGameId === null){
    return
  }
  TaskService.GetTaskList({
    gameId:props.selectedGameId
    }).then((res)=>{
      taskList.splice(0, taskList.length, ...res[0])
    })
  });

const emit = defineEmits([
  'mini'
])

function handleMini() {
  emit('mini',null)
}

const finish = ref(false)

function handleCheck() {
  finish.value = !finish.value
  if(!finish.value){
    getData()
  }else{
    getData(-1)
  }
}
</script>

<style scoped>
.task-board {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: #fdfefc;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Microsoft YaHei', sans-serif;
}

/* ========== 任务列表 ========== */
.task-list {
  flex: 1;
  overflow-y: auto;
  padding: 6px 20px 20px;
}

.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 4px;
}

/* 列表进出动画 */
.card-list-enter-active {
  transition: all 0.25s ease-out;
}
.card-list-leave-active {
  transition: all 0.15s ease-in;
}
.card-list-enter-from {
  opacity: 0;
  transform: scale(0.95);
}
.card-list-leave-to {
  opacity: 0;
  transform: scale(0.9);
}

/* ========== 空状态 ========== */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 60px 20px;
  text-align: center;
  grid-column: 1 / -1;
}

.empty-icon {
  font-size: 44px;
  margin-bottom: 14px;
  opacity: 0.6;
}

.empty-text {
  font-size: 15px;
  color: #8b9086;
  margin: 0 0 6px;
  font-weight: 500;
}

.empty-hint {
  font-size: 13px;
  color: #b5bab0;
  margin: 0;
}

/* ========== 删除确认弹窗（与左侧相同风格）========== */
.confirm-overlay {
  position: fixed;
  inset: 0;
  z-index: 2000;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(60, 60, 55, 0.35);
  backdrop-filter: blur(4px);
  -webkit-backdrop-filter: blur(4px);
}

.confirm-dialog {
  background: #fffefb;
  border-radius: 18px;
  padding: 28px 28px 22px;
  width: 320px;
  text-align: center;
  box-shadow: 0 8px 32px rgba(80, 80, 70, 0.18), 0 2px 8px rgba(80, 80, 70, 0.08);
}

.confirm-icon {
  font-size: 36px;
  margin-bottom: 10px;
}

.confirm-title {
  font-size: 16px;
  font-weight: 600;
  color: #3d4045;
  margin: 0 0 8px;
}

.confirm-message {
  font-size: 13.5px;
  color: #5c6058;
  margin: 0 0 4px;
  line-height: 1.5;
}

.confirm-message strong {
  color: #3d4045;
}

.confirm-sub-message {
  font-size: 11.5px;
  color: #b0b5aa;
  margin: 0 0 18px;
}

.confirm-actions {
  display: flex;
  gap: 10px;
  justify-content: center;
}

.confirm-btn {
  padding: 9px 22px;
  border-radius: 10px;
  border: none;
  font-size: 13px;
  font-weight: 550;
  cursor: pointer;
  transition: all 0.15s ease;
  font-family: inherit;
  letter-spacing: -0.01em;
}

.confirm-btn.cancel {
  background: #f0f1ed;
  color: #6b6e66;
}

.confirm-btn.cancel:hover { background: #e6e8e2; }

.confirm-btn.danger {
  background: #f5e6e6;
  color: #b85c5c;
}

.confirm-btn.danger:hover { background: #f0d8d8; color: #a34a4a; }

.confirm-btn:active { transform: scale(0.96); }

.fade-scale-enter-active { transition: all 0.2s ease-out; }
.fade-scale-leave-active { transition: all 0.16s ease-in; }
.fade-scale-enter-from { opacity: 0; }
.fade-scale-enter-from .confirm-dialog { transform: scale(0.9); opacity: 0; }
.fade-scale-leave-to { opacity: 0; }
.fade-scale-leave-to .confirm-dialog { transform: scale(0.9); opacity: 0; }
</style>