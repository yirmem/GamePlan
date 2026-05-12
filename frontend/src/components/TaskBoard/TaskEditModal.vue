<!-- src/components/TaskBoard/TaskEditModal.vue -->
<template>
  <Teleport to="body">
    <Transition name="modal-transition">
      <div v-if="visible" class="modal-overlay" @click.self="return">
        <div class="modal-panel">
          <div class="modal-header">
            <h3 class="modal-title">{{ isEditMode ? '编辑任务' : '添加任务' }}</h3>
            <button class="modal-close-btn" @click="handleCancel">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round">
                <line x1="18" y1="6" x2="6" y2="18" />
                <line x1="6" y1="6" x2="18" y2="18" />
              </svg>
            </button>
          </div>

          <div class="modal-body">
            <!-- 任务名称 -->
            <label class="field-label">任务名称</label>
            <input v-model="form.name" class="input" placeholder="输入任务名称..." maxlength="50" />

            <!-- 关联游戏（可选） -->
            <label class="field-label">关联游戏</label>
            <select v-model="form.gameId" class="input select">
              <option v-for="g in gameList" :key="g.id" :value="g.id">{{ g.name }}</option>
            </select>

            <!-- 周期天数 -->
            <label class="field-label">周期（天数）</label>
            <input v-model.number="form.checkDay" type="number" min="1" class="input" placeholder="例如：1 表示每天" />

            <!-- 是否重复 -->
            <label class="field-label">任务类型</label>
            <div class="switch-row">
              <button
                :class="['switch-option', { active: form.isRepeat === 0 }]"
                @click="form.isRepeat = 0"
              >单次</button>
              <button
                :class="['switch-option', { active: form.isRepeat === 1 }]"
                @click="form.isRepeat = 1"
              >重复</button>
            </div>

            <!-- 开始日期 -->
            <label class="field-label">开始日期</label>
            <n-date-picker 
              v-model:formatted-value="form.startDate"
              format="yyyy-MM-dd HH:mm:ss"
              type="datetime"
              clearable
            />

            <!-- 内容/奖励 -->
            <label class="field-label">内容说明</label>
            <textarea v-model="form.content" class="input textarea" placeholder="描述任务详情或奖励..." rows="3" maxlength="500"></textarea>
          </div>

          <div class="modal-footer">
            <button class="modal-btn cancel-btn" @click="handleCancel">取消</button>
            <button class="modal-btn confirm-btn" :disabled="!form.name.trim()" @click="handleConfirm">
              {{ isEditMode ? '保存' : '添加' }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { ref, reactive, watch, computed, nextTick } from 'vue'
import { TaskService } from '../../../bindings/gameplan/game'
import { Task } from '../../../bindings/gameplan/model'

const props = defineProps({
  visible: Boolean,
  mode: { type: String, default: 'add' }, // 'add' | 'edit'
  task: { type: Object, default: null },
  gameList: { type: Array, default: () => [] },
  defaultGameId: { type: [Number, null], default: null },
})

const emit = defineEmits(['update:visible', 'confirm', 'cancel'])

const isEditMode = computed(() => props.mode === 'edit')

const form = reactive({
  name: '',
  gameId: null,
  checkDay: 1,
  isRepeat: 0,
  startDate: '2026-01-01 01:00:00',
  content: '',
})

// 监听 visible 变化，填充表单
watch(() => props.visible, (v) => {
  if (v) {
    if (props.mode === 'edit' && props.task) {
      form.name = props.task.name
      form.gameId = props.task.gameId
      form.checkDay = props.task.checkDay
      form.isRepeat = props.task.isRepeat
      form.startDate = props.task.startDate || ''
      form.content = props.task.content || ''
    } else {
      form.name = ''
      form.checkDay = 1
      form.isRepeat = 0
      form.gameId = props.defaultGameId ?? (props.gameList[0]?.id ?? 1)
      form.startDate = new Date().toLocaleString('sv-SE').replace(' ', ' ');
      form.content = ''
    }
  }
})

function handleConfirm() {
  if (!form.name.trim()) return
  if(isEditMode.value === true){
  }else{
  }
  emit('confirm', { ...form })
}

function handleCancel() {
  emit('update:visible', false)
  emit('cancel')
}
</script>

<style scoped>
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

.modal-panel {
  background: #fffefb;
  border-radius: 18px;
  width: 420px;
  max-width: 94vw;
  box-shadow: 0 10px 40px rgba(60, 60, 50, 0.2), 0 2px 10px rgba(60, 60, 50, 0.08);
  overflow: hidden;
}

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
}

.modal-close-btn:hover { background: #f2f3ef; color: #6b6e66; }

.modal-body {
  padding: 6px 20px 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.field-label {
  font-size: 12px;
  font-weight: 600;
  color: #8b9086;
  text-transform: uppercase;
  letter-spacing: 0.03em;
  margin-bottom: -8px;
}

.input {
  width: 100%;
  padding: 9px 12px;
  border: 1.5px solid #e8ebe6;
  border-radius: 10px;
  font-size: 13.5px;
  outline: none;
  background: #fafbf8;
  font-family: inherit;
  transition: border-color 0.15s, box-shadow 0.15s;
  box-sizing: border-box;
}

.input:focus {
  border-color: #c5d3bb;
  box-shadow: 0 0 0 3px rgba(165, 185, 150, 0.15);
  background: #fff;
}

.select {
  appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg width='12' height='8' viewBox='0 0 12 8' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M1 1.5L6 6.5L11 1.5' stroke='%238b9086' stroke-width='1.5' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 12px center;
  padding-right: 32px;
}

.textarea {
  resize: vertical;
  min-height: 70px;
}

.switch-row {
  display: flex;
  gap: 0;
  background: #f1f3ee;
  border-radius: 9px;
  padding: 3px;
}

.switch-option {
  flex: 1;
  padding: 7px 0;
  background: transparent;
  border: none;
  border-radius: 7px;
  font-size: 13px;
  font-weight: 500;
  color: #7b8076;
  cursor: pointer;
  transition: all 0.15s;
  font-family: inherit;
}

.switch-option.active {
  background: #fff;
  color: #3a4e32;
  font-weight: 600;
  box-shadow: 0 1px 3px rgba(0,0,0,0.04);
}

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
  transition: all 0.15s;
  font-family: inherit;
}

.cancel-btn { background: #f0f1ed; color: #6b6e66; }
.cancel-btn:hover { background: #e6e8e2; }

.confirm-btn { background: #dce6d4; color: #4a6e3a; }
.confirm-btn:hover:not(:disabled) { background: #cfe0c5; color: #3a5c2c; }
.confirm-btn:disabled { background: #f0f1ed; color: #c0c5ba; cursor: not-allowed; }

.confirm-btn:not(:disabled):active,
.cancel-btn:active { transform: scale(0.96); }

/* 动画 */
.modal-transition-enter-active { transition: all 0.22s ease-out; }
.modal-transition-leave-active { transition: all 0.16s ease-in; }
.modal-transition-enter-from { opacity: 0; }
.modal-transition-enter-from .modal-panel { transform: scale(0.92) translateY(10px); opacity: 0; }
.modal-transition-leave-to { opacity: 0; }
.modal-transition-leave-to .modal-panel { transform: scale(0.92) translateY(10px); opacity: 0; }
</style>