<template>
  <t-layout class="layout-main">
    <t-header>
      <t-head-menu theme="light" v-model="menu" height="120px">
        <t-menu-item value="start" @click="goStart">启动</t-menu-item>
        <t-menu-item value="config" @click="goConfig">配置</t-menu-item>
          <template #operations>
            <t-tooltip content="显示运行日志" placemen="left">
            <t-button variant="text" shape="square" @click="logVisible=true">
            <template #icon><system-log-icon/></template>
            </t-button>
            </t-tooltip>
          </template>
      </t-head-menu>
    </t-header>
    <t-content>
      <t-drawer
          v-model:visible="logVisible"
          :show-overlay="false"
          cancelBtn="关闭日志"
          :confirm-btn="null"
          header="运行日志"
          placement="bottom"
      >
        <div class="logArea" ref="logArea">
          {{logConetent}}
        </div>
      </t-drawer>
      <router-view/>
    </t-content>
  </t-layout>
</template>

<script setup lang="ts">
import {useRouter} from 'vue-router';
import {onMounted, ref} from "vue";
import {SystemLogIcon} from "tdesign-icons-vue-next";
import {EventsOn} from "../../wailsjs/runtime";

const route = useRouter();
const logVisible = ref(false);
const logConetent = ref("");
const menu = ref("start");
const logArea = ref<HTMLDivElement|null>(null);

const goStart = () => {
  route.push('/start');
};

const goConfig = () => {
  route.push('/config');
};

const showLog = ()=>{
  let logLine = 0;
  EventsOn("log", (msg:string) => {
     if (logLine>500){
        logConetent.value = "";
        logLine = 0;
     }
     logConetent.value += msg;
     if(logArea.value){
       logArea.value.scrollTop = logArea.value.scrollHeight;
     }
      logLine++;
  });
}

//跟随系统设置暗黑模式
// 检测系统颜色模式
const prefersDarkMode = window.matchMedia('(prefers-color-scheme: dark)');

// 根据系统颜色模式设置页面样式
const setTheme = (isDarkMode:boolean)=> {
  if (isDarkMode) {
    document.documentElement.setAttribute('theme-mode', 'dark');
  } else {
    document.documentElement.removeAttribute('theme-mode');
  }
}
// 监听系统颜色模式变化
prefersDarkMode.addEventListener('change', (event) => {
  setTheme(event.matches);
});

onMounted(() => {
  goStart();
  setTheme(prefersDarkMode.matches);
  showLog();
});
</script>

<style scoped>
.layout-main {
  height: 100vh;
}
.logArea {
  background-color: black;
  color: white;
  height: 100%;
  overflow-y: auto;
  font-family: monospace;
  white-space: pre-wrap;
  text-align: left;
  padding-left: 2em;
}
</style>