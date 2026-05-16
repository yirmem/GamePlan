<template>
    <div>

    </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted,watch } from 'vue'
import { TaskService } from '../../../bindings/gameplan/game'

// ---------- 任务数据----------
const taskList = reactive([])

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

onMounted(()=>{
  getData()
})

</script>

<style scoped>

</style>