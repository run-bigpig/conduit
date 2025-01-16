<template>
  <t-card class="card" :bordered="false">
    <t-button v-if="running==false" theme="primary" size="large" variant="base" shape="round" class="button" @click="toggleProxy(true)">开启代理</t-button>
    <t-button v-else theme="danger" size="large" variant="base" shape="round" class="button" @click="toggleProxy(false)">关闭代理</t-button>
  </t-card>
</template>

<script setup lang="ts">
import {onMounted, ref} from 'vue';
import {start,stop} from "@/api/gui";
import {GetStatus} from "../../wailsjs/go/gui/GuiHandler";
import {NotifyPlugin} from "tdesign-vue-next";
const running = ref(false);

const toggleProxy = (on:boolean) => {
  GetStatus().then((res) => {
    if (on && !Boolean(res.data)) {
      start(running);
    } else if (!on && Boolean(res.data)) {
      stop(running);
    }else {
      NotifyPlugin.error({title: "提示", content: "操作失败"})
    }
  });
};
 onMounted(() => {
   GetStatus().then((res) => {
     running.value = Boolean(res.data);
   })
 });

</script>

<style scoped>
.card {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.button {
  display: flex;
}
</style>