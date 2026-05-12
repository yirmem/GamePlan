<!-- src/components/GameSidebar/GameEditModal.vue -->
<template>
  <Teleport to="body">
    <Transition name="modal-transition">
      <div v-if="visible" class="modal-overlay" @click.self="handleCancel">
        <div class="modal-panel">
          <!-- 标题栏 -->
          <div class="modal-header">
            <h3 class="modal-title">{{ isEditMode ? '编辑游戏' : '添加游戏' }}</h3>
            <button class="modal-close-btn" @click="handleCancel" title="关闭">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
                <line x1="18" y1="6" x2="6" y2="18" />
                <line x1="6" y1="6" x2="18" y2="18" />
              </svg>
            </button>
          </div>

          <!-- 内容区 -->
          <div class="modal-body">
            <!-- 游戏名称 -->
            <label class="field-label">游戏名称</label>
            <input
                ref="nameInputRef"
                v-model="formName"
                class="name-input"
                type="text"
                placeholder="输入游戏名称..."
                maxlength="30"
                @keydown.enter="handleConfirm"
            />

            <!-- 图标选择 -->
            <!-- <label class="field-label">选择图标</label>
            <div class="emoji-grid">
              <button
                  v-for="emoji in emojiList"
                  :key="emoji"
                  class="emoji-option"
                  :class="{ selected: formIcon === emoji }"
                  @click="formIcon = emoji"
                  type="button"
              >
                {{ emoji }}
              </button>
            </div>-->
          </div>

          <!-- 底部按钮 -->
          <div class="modal-footer">
            <button class="modal-btn cancel-btn" @click="handleCancel">取消</button>
            <button
                class="modal-btn confirm-btn"
                :disabled="!formName.trim()"
                @click="handleConfirm"
            >
              {{ isEditMode ? '保存' : '添加' }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { ref, watch, computed, nextTick, onMounted } from 'vue'
import { GameService } from '@/bindings/gameplan/game'

const props = defineProps({
  visible: {
    type: Boolean,
    default: false,
  },
  mode: {
    type: String,
    default: 'add', // 'add' | 'edit'
  },
  game: {
    type: Object,
    default: null,
  },
})

const emit = defineEmits(['update:visible', 'confirm', 'cancel'])

// 预设游戏图标列表
const emojiList = ['🎮', '✨', '🚂', '🗡️', '🏹', '🔮', '⚔️', '🛡️', '🎯', '🌟', '💎', '🔥', '🌙', '🎪', '👾', '🕹️', '🎲', '🏰', '🐉', '🧙']

const formName = ref('')
const formIcon = ref('🎮')
const nameInputRef = ref(null)

const isEditMode = computed(() => props.mode === 'edit')

// 监听 visible 变化，初始化表单
watch(
    () => props.visible,
    async (v) => {
      if (v) {
        if (props.mode === 'edit' && props.game) {
          formName.value = props.game.name || ''
          formIcon.value = props.game.icon || '🎮'
        } else {
          formName.value = ''
          // 随机预选一个图标
          formIcon.value = emojiList[Math.floor(Math.random() * emojiList.length)]
        }
        await nextTick()
        nameInputRef.value?.focus()
        nameInputRef.value?.select()
      }
    }
)

function handleConfirm() {
  
console.log(props.game);
  if (!formName.value.trim()) return
  if(isEditMode.value === true){
  }else{
  }
  emit('confirm', {
    name: formName.value.trim(),
    icon: formIcon.value,
  })
}

function handleCancel() {
  emit('update:visible', false)
  emit('cancel')
}
</script>

<style scoped>
/* ========== 遮罩层 ========== */
.modal-overlay {
  position: fixed;
  inset: 0;
  z-index: 1500;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(55, 55, 50, 0.3);
  backdrop-filter: blur(5px);
  -webkit-backdrop-filter: blur(5px);
}

/* ========== 弹窗面板 ========== */
.modal-panel {
  background: #fffefb;
  border-radius: 18px;
  width: 380px;
  max-width: 92vw;
  box-shadow:
      0 10px 40px rgba(60, 60, 50, 0.2),
      0 2px 10px rgba(60, 60, 50, 0.08);
  overflow: hidden;
}

/* ========== 标题栏 ========== */
.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px 12px;
}

.modal-title {
  font-size: 15px;
  font-weight: 650;
  color: #3d4045;
  margin: 0;
  letter-spacing: -0.01em;
}

.modal-close-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 30px;
  height: 30px;
  border: none;
  border-radius: 8px;
  background: transparent;
  color: #b0b5aa;
  cursor: pointer;
  transition: all 0.14s ease;
  padding: 0;
}

.modal-close-btn:hover {
  background: #f2f3ef;
  color: #6b6e66;
}

/* ========== 内容区 ========== */
.modal-body {
  padding: 6px 20px 16px;
}

.field-label {
  display: block;
  font-size: 12px;
  font-weight: 600;
  color: #8b9086;
  margin-bottom: 8px;
  text-transform: uppercase;
  letter-spacing: 0.03em;
}

.name-input {
  width: 100%;
  padding: 10px 14px;
  border: 1.5px solid #e8ebe6;
  border-radius: 10px;
  font-size: 14px;
  color: #3d4045;
  outline: none;
  transition: border-color 0.16s ease, box-shadow 0.16s ease;
  background: #fafbf8;
  font-family: inherit;
  letter-spacing: -0.01em;
  box-sizing: border-box;
}

.name-input:focus {
  border-color: #c5d3bb;
  box-shadow: 0 0 0 3px rgba(165, 185, 150, 0.15);
  background: #fff;
}

.name-input::placeholder {
  color: #c0c5ba;
}

/* ========== 图标选择 ========== */
.emoji-grid {
  display: grid;
  grid-template-columns: repeat(10, 1fr);
  gap: 6px;
}

.emoji-option {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  aspect-ratio: 1;
  border: 1.5px solid transparent;
  border-radius: 10px;
  background: #f7f8f4;
  font-size: 18px;
  cursor: pointer;
  transition: all 0.14s ease;
  padding: 0;
}

.emoji-option:hover {
  background: #eef0ea;
  border-color: #dde0d7;
}

.emoji-option.selected {
  background: #e4ece0;
  border-color: #b5c9a8;
  box-shadow: 0 0 0 2px rgba(150, 175, 135, 0.2);
}

/* ========== 底部按钮 ========== */
.modal-footer {
  display: flex;
  gap: 10px;
  padding: 14px 20px 18px;
  justify-content: flex-end;
}

.modal-btn {
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

.cancel-btn {
  background: #f0f1ed;
  color: #6b6e66;
}

.cancel-btn:hover {
  background: #e6e8e2;
}

.confirm-btn {
  background: #dce6d4;
  color: #4a6e3a;
}

.confirm-btn:hover {
  background: #cfe0c5;
  color: #3a5c2c;
}

.confirm-btn:disabled {
  background: #f0f1ed;
  color: #c0c5ba;
  cursor: not-allowed;
}

.confirm-btn:not(:disabled):active,
.cancel-btn:active {
  transform: scale(0.96);
}

/* ========== 过渡动画 ========== */
.modal-transition-enter-active {
  transition: all 0.22s ease-out;
}
.modal-transition-leave-active {
  transition: all 0.16s ease-in;
}
.modal-transition-enter-from {
  opacity: 0;
}
.modal-transition-enter-from .modal-panel {
  transform: scale(0.92) translateY(10px);
  opacity: 0;
}
.modal-transition-leave-to {
  opacity: 0;
}
.modal-transition-leave-to .modal-panel {
  transform: scale(0.92) translateY(10px);
  opacity: 0;
}
</style>