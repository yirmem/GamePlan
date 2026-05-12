<!-- src/components/GameSidebar/GameItem.vue -->
<template>
  <div
      class="game-item"
      :class="{
      'is-active': isActive,
      'is-editing': isEditingInline,
    }"
      @click="$emit('click')"
      @dblclick="$emit('dblclick')"
  >
    <!-- 左侧图标 -->
    <div class="game-icon" :class="{ 'icon-active': isActive }" :style="{ backgroundColor: getIconColor(game.name) }">
      {{ getFirstChar(game.name) }}
    </div>
    <!-- 游戏名称 / 内联编辑输入框 -->
    <div class="game-name-wrapper">
      <span v-if="!isEditingInline" class="game-name">{{ game.name }}</span>
      <input
          v-else
          ref="inlineInputRef"
          class="inline-input"
          :value="editValue"
          @input="editValue = $event.target.value"
          @keydown.enter="confirmInlineEdit"
          @keydown.escape="$emit('inline-edit-cancel')"
          @blur="confirmInlineEdit"
          @click.stop
      />
    </div>

    <!-- 操作按钮（hover 显示） -->
    <div v-if="!isEditingInline" class="game-actions">
      <button
          class="action-btn edit-btn"
          title="编辑"
          @click.stop="$emit('edit')"
      >
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7" />
          <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z" />
        </svg>
      </button>
      <button
          class="action-btn delete-btn"
          title="删除"
          @click.stop="$emit('delete')"
      >
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="3 6 5 6 21 6" />
          <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2" />
        </svg>
      </button>
    </div>

    <!-- 编辑中的操作按钮 -->
    <div v-else class="inline-edit-actions">
      <button class="inline-action confirm" @click.stop="confirmInlineEdit" title="确认">
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.8" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="20 6 9 17 4 12" />
        </svg>
      </button>
      <button class="inline-action cancel" @click.stop="$emit('inline-edit-cancel')" title="取消">
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.8" stroke-linecap="round">
          <line x1="18" y1="6" x2="6" y2="18" />
          <line x1="6" y1="6" x2="18" y2="18" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, nextTick, onMounted } from 'vue'

const props = defineProps({
  game: {
    type: Object,
    required: true,
  },
  isActive: {
    type: Boolean,
    default: false,
  },
  isEditingInline: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits([
  'click',
  'dblclick',
  'edit',
  'delete',
  'inline-edit-confirm',
  'inline-edit-cancel',
])

const editValue = ref('')
const inlineInputRef = ref(null)

// 进入编辑模式时同步输入框的值
watch(
    () => props.isEditingInline,
    async (editing) => {
      if (editing) {
        editValue.value = props.game.name
        await nextTick()
        inlineInputRef.value?.focus()
        inlineInputRef.value?.select()
      }
    }
)

// 确认内联编辑
function confirmInlineEdit() {
  const trimmed = editValue.value.trim()
  if (trimmed && trimmed !== props.game.name) {
    emit('inline-edit-confirm', trimmed)
  } else if (!trimmed) {
    // 名称为空则取消
    emit('inline-edit-cancel')
  } else {
    emit('inline-edit-cancel')
  }
}

// 获取首字符（英文取首字母大写，中文取第一个汉字）
const getFirstChar = (name) => {
  if (!name) return '?';
  const firstChar = name.trim().charAt(0);
  // 判断是否为中文字符（Unicode 范围 0x4E00-0x9FFF）
  const isChinese = /[\u4e00-\u9fff]/.test(firstChar);
  if (isChinese) {
    return firstChar;
  } else {
    // 英文或数字：转为大写
    return firstChar.toUpperCase();
  }
};
const getIconColor = (name) => {
  let hash = 0;
  for (let i = 0; i < name.length; i++) {
    hash = name.charCodeAt(i) + ((hash << 5) - hash);
  }
  const hue = Math.abs(hash % 360);
  return `hsl(${hue}, 70%, 60%)`;
};
</script>

<style scoped>
/* ========== 游戏项 ========== */
.game-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 9px 10px 9px 10px;
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.16s ease;
  position: relative;
  min-height: 44px;
}

.game-item:hover {
  background: #f1f3ee;
}

.game-item.is-active {
  background: #eaf0e5;
}

.game-item.is-active:hover {
  background: #e3ece0;
}

.game-item.is-editing {
  background: #f5f7f2;
  box-shadow: inset 0 0 0 1.5px #d5dfcd;
}

/* ========== 图标 ========== */
.game-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  font-size: 16px;
  background: #f3f5f0;
  flex-shrink: 0;
  transition: all 0.16s ease;
}
.game-item.is-active .game-icon,
.icon-active {
  background: #dce6d4;
}

/* ========== 名称 ========== */
.game-name-wrapper {
  flex: 1;
  min-width: 0;
}

.game-name {
  font-size: 13.5px;
  font-weight: 500;
  color: #3d4045;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  display: block;
  letter-spacing: -0.01em;
  line-height: 1.3;
}

.game-item.is-active .game-name {
  font-weight: 600;
  color: #3a4e32;
}

/* ========== 内联编辑输入框 ========== */
.inline-input {
  width: 100%;
  padding: 4px 8px;
  border: none;
  border-radius: 6px;
  font-size: 13.5px;
  font-weight: 500;
  color: #3d4045;
  background: #fff;
  outline: none;
  font-family: inherit;
  letter-spacing: -0.01em;
  box-shadow: 0 0 0 1.5px #c5d3bb;
}

/* ========== 操作按钮（hover显示） ========== */
.game-actions {
  display: flex;
  gap: 2px;
  opacity: 0;
  transition: opacity 0.14s ease;
  flex-shrink: 0;
}

.game-item:hover .game-actions,
.game-item.is-active .game-actions {
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
  background: transparent;
  cursor: pointer;
  color: #a0a89a;
  transition: all 0.14s ease;
  padding: 0;
}

.action-btn:hover {
  background: #e6eae1;
  color: #6b7264;
}

.delete-btn:hover {
  background: #f5e8e8;
  color: #c06b6b;
}

/* ========== 内联编辑操作按钮 ========== */
.inline-edit-actions {
  display: flex;
  gap: 4px;
  flex-shrink: 0;
}

.inline-action {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border: none;
  border-radius: 7px;
  cursor: pointer;
  transition: all 0.14s ease;
  padding: 0;
}

.inline-action.confirm {
  background: #dce8d4;
  color: #5a8a44;
}

.inline-action.confirm:hover {
  background: #cfe0c5;
}

.inline-action.cancel {
  background: #f0efeb;
  color: #9a9d94;
}

.inline-action.cancel:hover {
  background: #e6e5e0;
  color: #6e7068;
}
</style>