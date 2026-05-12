<!-- src/components/GameSidebar/GameSidebar.vue -->
<template>
  <aside class="sidebar">
    <!-- 头部标题区 -->
    <div class="sidebar-header">
      <span class="header-title">游戏列表</span>
      <span class="header-count">{{ gameList.length }}</span>
    </div>

    <!-- 分割线 -->
    <div class="divider"></div>

    <!-- 游戏列表 -->
    <div class="game-list" ref="listRef">
      <GameItem
          v-for="game in gameList"
          :key="game.id"
          :game="game"
          :is-active="selectedGameId === game.id"
          :is-editing-inline="inlineEditingId === game.id"
          @click="handleSelect(game)"
          @edit="handleOpenEdit(game)"
          @delete="handleDelete(game)"
          @inline-edit-confirm="(name) => handleInlineEditConfirm(game.id, name)"
          @inline-edit-cancel="cancelInlineEdit"
          @dblclick="handleDblClick(game)"
      />

      <!-- 空状态 -->
      <div v-if="gameList.length === 0" class="empty-state">
        <div class="empty-icon">🎮</div>
        <p class="empty-text">还没有游戏</p>
        <p class="empty-hint">点击下方按钮添加</p>
      </div>
    </div>

    <!-- 底部操作区 -->
    <div class="sidebar-footer" v-if="props.showTrigger === 1">
      <button class="add-btn" @click="handleOpenAdd" title="添加游戏">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
          <line x1="12" y1="5" x2="12" y2="19" />
          <line x1="5" y1="12" x2="19" y2="12" />
        </svg>
        <span>添加游戏</span>
      </button>
    </div>

    <!-- 编辑/添加弹窗 -->
    <GameEditModal
        v-model:visible="modalVisible"
        :mode="modalMode"
        :game="editingGame"
        @confirm="handleModalConfirm"
        @cancel="handleModalCancel"
    />

    <!-- 删除确认弹窗 -->
    <Teleport to="body">
      <Transition name="fade-scale">
        <div v-if="deleteConfirmVisible" class="confirm-overlay" @click.self="deleteConfirmVisible = false">
          <div class="confirm-dialog">
            <div class="confirm-icon">⚠️</div>
            <p class="confirm-title">确认删除</p>
            <p class="confirm-message">
              确定要删除游戏<br /><strong>「{{ deletingGame?.name }}」</strong> 吗？
            </p>
            <p class="confirm-sub-message">此操作不可撤销</p>
            <div class="confirm-actions">
              <button class="confirm-btn cancel" @click="deleteConfirmVisible = false">取消</button>
              <button class="confirm-btn danger" @click="confirmDelete">确认删除</button>
            </div>
          </div>
        </div>
      </Transition>
    </Teleport>
  </aside>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import GameItem from './GameSidebar/GameItem.vue'
import GameEditModal from './GameSidebar/GameEditModal.vue'
import {GameService} from '../../bindings/gameplan/game'

// ==================== 数据定义 ====================
// 后期可替换为接口获取的数据
const gameList = reactive([])

onMounted(()=>{
  getData()
})

const getData = ()=>{
  GameService.GetGameList().then((res)=>{
    gameList.splice(0, gameList.length, ...res)
    if(res.length !== 0){
      selectedGameId.value = res[0].id
    }
  })
}

// 当前选中的游戏 ID
const selectedGameId = ref(null)

// 内联编辑状态
const inlineEditingId = ref(null)

// Modal 相关状态
const modalVisible = ref(false)
const modalMode = ref('add') // 'add' | 'edit'
const editingGame = ref(null)

// 删除确认相关状态
const deleteConfirmVisible = ref(false)
const deletingGame = ref(null)

// 列表容器引用
const listRef = ref(null)

// ==================== 操作方法 ====================
// 选择游戏
function handleSelect(game) {
  
  if (inlineEditingId.value) return // 编辑中不切换
  selectedGameId.value = game.id
}

// 双击进入内联编辑
function handleDblClick(game) {
  inlineEditingId.value = game.id
}

// 内联编辑确认
function handleInlineEditConfirm(gameId, newName) {
  const game = gameList.find(g => g.id === gameId)
  if (game && newName.trim()) {
    game.name = newName.trim()
  }
  inlineEditingId.value = null
}

// 取消内联编辑
function cancelInlineEdit() {
  inlineEditingId.value = null
}

// 打开添加弹窗
function handleOpenAdd() {
  modalMode.value = 'add'
  editingGame.value = null
  modalVisible.value = true
}

// 打开编辑弹窗
function handleOpenEdit(game) {
  modalMode.value = 'edit'
  editingGame.value = { ...game }
  modalVisible.value = true
}

// Modal 确认
function handleModalConfirm(data) {
  if (modalMode.value === 'add') {
    GameService.AddGame({id:0,name:data.name}).then((res)=>{
        getData()
        modalVisible.value = false
        editingGame.value = null
    })
    selectedGameId.value = newGame.id
  } else if (modalMode.value === 'edit' && editingGame.value) {
    const game = gameList.find(g => g.id === editingGame.value.id)
    if (game) {
      GameService.UpdateGame({id:game.id,name:data.name}).then((res)=>{
        getData()
        modalVisible.value = false
        editingGame.value = null
      })
    }
  }
  modalVisible.value = false
  editingGame.value = null
}

// Modal 取消
function handleModalCancel() {
  modalVisible.value = false
  editingGame.value = null
}

// 点击删除按钮
function handleDelete(game) {
  deletingGame.value = game
  deleteConfirmVisible.value = true
}

// 确认删除
function confirmDelete() {
  
  if (deletingGame.value) {
    const index = gameList.findIndex(g => g.id === deletingGame.value.id)
    
    GameService.DelGame({id:deletingGame.value.id}).then((res)=>{
    
    })
    if (index !== -1) {
      gameList.splice(index, 1)
      if (selectedGameId.value === deletingGame.value.id) {
        selectedGameId.value = gameList.length > 0 ? gameList[0]?.id ?? null : null
      }
    }
  }
  deleteConfirmVisible.value = false
  deletingGame.value = null
}

// ==================== 暴露方法供父组件调用 ====================
defineExpose({
  gameList,
  selectedGameId,
})

const props = defineProps({
  showTrigger:Number
})

</script>

<style scoped>
/* ========== 侧边栏容器 ========== */
.sidebar {
  display: flex;
  flex-direction: column;
  width: 200px;
  height: 100vh;
  background: #f9faf7;
  border-right: 1px solid #eef0eb;
  user-select: none;
  font-family:
      -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC',
      'Microsoft YaHei', sans-serif;
}

/* ========== 头部 ========== */
.sidebar-header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 18px 16px 14px;
  flex-shrink: 0;
}

.header-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 34px;
  height: 34px;
  border-radius: 10px;
  background: #eaf0e5;
  color: #7a9a6b;
  flex-shrink: 0;
}

.header-title {
  font-size: 15px;
  font-weight: 600;
  color: #3d4045;
  letter-spacing: -0.01em;
}

.header-count {
  margin-left: auto;
  font-size: 11px;
  font-weight: 500;
  color: #a0a69c;
  background: #eef0eb;
  padding: 2px 9px;
  border-radius: 20px;
  min-width: 20px;
  text-align: center;
}

/* ========== 分割线 ========== */
.divider {
  height: 1px;
  background: #eef0eb;
  margin: 0 12px;
  flex-shrink: 0;
}

/* ========== 游戏列表 ========== */
.game-list {
  flex: 1;
  overflow-y: auto;
  padding: 8px 8px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

/* 自定义滚动条 */
.game-list::-webkit-scrollbar {
  width: 4px;
}
.game-list::-webkit-scrollbar-track {
  background: transparent;
}
.game-list::-webkit-scrollbar-thumb {
  background: #d5dbd0;
  border-radius: 10px;
}
.game-list::-webkit-scrollbar-thumb:hover {
  background: #c0c7ba;
}

/* ========== 空状态 ========== */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  text-align: center;
  flex: 1;
}

.empty-icon {
  font-size: 40px;
  margin-bottom: 12px;
  opacity: 0.7;
}

.empty-text {
  font-size: 14px;
  color: #8b9086;
  margin: 0 0 4px;
  font-weight: 500;
}

.empty-hint {
  font-size: 12px;
  color: #b5bab0;
  margin: 0;
}

/* ========== 底部操作区 ========== */
.sidebar-footer {
  flex-shrink: 0;
  padding: 10px 12px 14px;
}

.add-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 7px;
  width: 100%;
  padding: 10px 14px;
  border: none;
  border-radius: 10px;
  background: #eaf0e5;
  color: #6b8a5e;
  font-size: 13px;
  font-weight: 550;
  cursor: pointer;
  transition: all 0.18s ease;
  letter-spacing: -0.01em;
  font-family: inherit;
}

.add-btn:hover {
  background: #dee8d7;
  color: #567748;
}

.add-btn:active {
  background: #d3e0ca;
  transform: scale(0.97);
}

.add-btn svg {
  flex-shrink: 0;
}

/* ========== 删除确认弹窗 ========== */
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
  box-shadow:
      0 8px 32px rgba(80, 80, 70, 0.18),
      0 2px 8px rgba(80, 80, 70, 0.08);
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

.confirm-btn.cancel:hover {
  background: #e6e8e2;
}

.confirm-btn.danger {
  background: #f5e6e6;
  color: #b85c5c;
}

.confirm-btn.danger:hover {
  background: #f0d8d8;
  color: #a34a4a;
}

.confirm-btn:active {
  transform: scale(0.96);
}

/* ========== 过渡动画 ========== */
.fade-scale-enter-active {
  transition: all 0.2s ease-out;
}
.fade-scale-leave-active {
  transition: all 0.16s ease-in;
}
.fade-scale-enter-from {
  opacity: 0;
}
.fade-scale-enter-from .confirm-dialog {
  transform: scale(0.9);
  opacity: 0;
}
.fade-scale-leave-to {
  opacity: 0;
}
.fade-scale-leave-to .confirm-dialog {
  transform: scale(0.9);
  opacity: 0;
}
</style>