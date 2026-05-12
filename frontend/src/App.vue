<script setup>
import { ref } from "vue";
import GameSidebar from "./components/GameSidebar.vue";
import TaskBoard from "./components/TaskBoard/TaskBoard.vue";
import { AppService } from "../bindings/gameplan";

const showTrigger = ref(1)
const changeShowTrigger = (val) => {
  showTrigger.value = val
}

const sidebarRef = ref(null)

const isMini = ref(false)
const handleMini = ()=>{
  isMini.value = !isMini.value
  if(isMini.value){
    AppService.ChangeWindowOptions('mini').then(res=>{

    })
  }else{
    AppService.ChangeWindowOptions('noMini').then(res=>{

    })
  }
}

</script>

<template>
    <div style="">
      <n-layout>
  <!--      <n-layout-header>颐和园路</n-layout-header>-->
        <n-layout has-sider>
          <n-layout-sider v-if="!isMini"
            collapse-mode="width" 
            :collapsed-width="80" 
            :width="201"
            @after-enter="changeShowTrigger(1)"
            @after-leave="changeShowTrigger(0)"
            show-trigger="arrow-circle" >
            <GameSidebar ref="sidebarRef" :showTrigger="showTrigger"></GameSidebar>
          </n-layout-sider>
          <n-layout-content>
            <TaskBoard 
            :game-list="sidebarRef?.gameList"
            :selected-game-id="sidebarRef?.selectedGameId"
            :isMini="isMini"
            @mini="handleMini"
            ></TaskBoard>
          </n-layout-content>
        </n-layout>
      </n-layout>
    </div>
</template>

<style scoped>

</style>
