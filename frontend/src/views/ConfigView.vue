<template>
  <t-card :bordered="false" class="card">
    <t-tabs v-model="tabValue">
      <t-tab-panel value="system" label="基础" :destroy-on-hide="false">
        <template #panel>
        <t-form ref="systemForm" :data="systemFromData" :rules="systemRules" :colon="true" :label-width="0" class="form" @submit="onSystem">
          <t-form-item name="domain">
            <t-input v-model="systemFromData.domain" clearable placeholder="请输入域名">
              <template #prefix-icon>
                <logo-chrome-icon />
              </template>
            </t-input>
          </t-form-item>
          <t-form-item name="port">
            <t-input v-model="systemFromData.port"  type="number" clearable placeholder="请输入端口号">
              <template #prefix-icon>
                <lock-on-icon />
              </template>
            </t-input>
          </t-form-item>
          <t-form-item name="sslCert">
            <t-input v-model="systemFromData.sslCert" placeholder="请输入ssl证书 包含服务域名以及*.服务域名的证书" @click="selectFile('cert')">
              <template #prefix-icon>
                <certificate-icon />
              </template>
            </t-input>
            <input
                type="file"
                ref="sslCertInput"
                style="display: none"
                @change="handleFileChange('cert',$event)"
            />
          </t-form-item>
          <t-form-item name="sslKey">
            <t-input v-model="systemFromData.sslKey"  placeholder="请输入ssl密钥 包含服务域名以及*.服务域名的证书" @click="selectFile('key')">
              <template #prefix-icon>
                <certificate1-icon />
              </template>
            </t-input>
            <input
                type="file"
                ref="sslKeyInput"
                style="display: none"
                @change="handleFileChange('key',$event)"
            />
          </t-form-item>
          <t-form-item>
            <t-button theme="primary" type="submit" block>设置</t-button>
          </t-form-item>
        </t-form>
        </template>
      </t-tab-panel>
      <t-tab-panel value="code" label="代码补全" :destroy-on-hide="false">
        <template #panel>
          <t-form ref="codeForm" :data="codeFromData" :rules="codeRules" :colon="true" :label-width="0"  class="form" @submit="onCodeCompletion">
            <t-form-item name="api">
              <t-input v-model="codeFromData.api" clearable placeholder="请输入api 包含http或者https">
                <template #prefix-icon>
                  <api-icon />
                </template>
              </t-input>
            </t-form-item>

            <t-form-item name="sk">
              <t-input v-model="codeFromData.sk" clearable placeholder="请输入sk sk-xxxxxxxx">
                <template #prefix-icon>
                  <key-icon />
                </template>
              </t-input>
            </t-form-item>
            <t-form-item name="adapter">
              <t-select placeholder="请选择api对应的转发适配器" v-model="codeFromData.adapter">
                <template #prefixIcon>
                  <application-icon/>
                </template>
                <t-option key="deepseek" label="DeepSeek" value="deepseek" />
                <t-option key="openai" label="OpenAi" value="openai" />
              </t-select>
            </t-form-item>
            <t-form-item name="model">
              <t-input v-model="codeFromData.model" clearable placeholder="请输入model名称">
                <template #prefix-icon>
                  <control-platform-icon/>
                </template>
              </t-input>
            </t-form-item>
            <t-form-item name="max_tokens">
              <t-input v-model="codeFromData.maxTokens" type="number" placeholder="请输入最大token数">
                <template #prefix-icon>
                  <calculation1-icon />
                </template>
              </t-input>
            </t-form-item>
            <t-form-item>
              <t-button theme="primary" type="submit" block>设置</t-button>
            </t-form-item>
          </t-form>
        </template>
      </t-tab-panel>
      <t-tab-panel value="chat" label="聊天补全" :destroy-on-hide="false">
        <template #panel>
          <t-form ref="formForm" :data="chatFromData" :rules="codeRules" :colon="true" :label-width="0"  class="form" @submit="onChatCompletion">
            <t-form-item name="api">
              <t-input v-model="chatFromData.api" clearable placeholder="请输入api 包含http或者https">
                <template #prefix-icon>
                  <api-icon />
                </template>
              </t-input>
            </t-form-item>

            <t-form-item name="sk">
              <t-input v-model="chatFromData.sk" clearable placeholder="请输入sk sk-xxxxxxxx">
                <template #prefix-icon>
                  <key-icon />
                </template>
              </t-input>
            </t-form-item>
            <t-form-item name="adapter">
              <t-select placeholder="请选择api对应的转发适配器" v-model="chatFromData.adapter">
                <template #prefixIcon>
                  <application-icon/>
                </template>
                <t-option key="deepseek" label="DeepSeek" value="deepseek" />
                <t-option key="openai" label="OpenAi" value="openai" />
              </t-select>
            </t-form-item>
            <t-form-item name="model">
              <t-input v-model="chatFromData.model" clearable placeholder="请输入model名称">
                <template #prefix-icon>
                  <control-platform-icon/>
                </template>
              </t-input>
            </t-form-item>
            <t-form-item name="max_tokens">
              <t-input v-model="chatFromData.maxTokens" type="number"  placeholder="请输入最大token数">
                <template #prefix-icon>
                  <calculation1-icon />
                </template>
              </t-input>
            </t-form-item>
            <t-form-item>
              <t-button theme="primary" type="submit" block>设置</t-button>
            </t-form-item>
          </t-form>
        </template>
      </t-tab-panel>
    </t-tabs>
  </t-card>
</template>
<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import {
  ApiIcon,
  LockOnIcon,
  KeyIcon,
  ControlPlatformIcon,
  ApplicationIcon, Calculation1Icon, LogoChromeIcon, CertificateIcon, Certificate1Icon
} from "tdesign-icons-vue-next";
import {setChatCompletionConfig, setCodeCompletionConfig, setSystemConfig} from "@/api/gui";
import {FormProps} from "tdesign-vue-next";
import {GetConfig} from "../../wailsjs/go/gui/GuiHandler";
const sslCertInput = ref<HTMLInputElement | null>(null);
const sslKeyInput = ref<HTMLInputElement | null>(null);
const tabValue = ref("system")
const systemFromData = ref({
  port:443,
  domain:'',
  sslCert:'',
  sslKey:''
})
const codeFromData = ref({
  api:'',
  sk:'',
  model:'',
  adapter:'',
  maxTokens: 1024
})
const chatFromData = ref({
  api:'',
  sk:'',
  model:'',
  adapter:'',
  maxTokens: 1024
})

const selectFile = (key:string)=>{
  if (key=="cert"){
    if (sslCertInput.value){
      sslCertInput.value.click()
    }
  }else{
    if (sslKeyInput.value){
      sslKeyInput.value.click()
    }
  }
}
// 处理文件选择事件
const handleFileChange = (key:string,event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files.length > 0) {
    const file = target.files[0];
    if (file){
      const reader = new FileReader();
      reader.onload = (e) => {
        if (key=="cert"){
          if (e.target){
            systemFromData.value.sslCert = e.target.result as string;
          }
        }else {
          if (e.target) {
            systemFromData.value.sslKey = e.target.result as string;
          }
        }
      };
      reader.readAsText(file);
    }
  }
};

const systemRules: FormProps['rules'] = {
  port: [
    {
      required: true,
      message: '端口必填',
      type: 'error',
      trigger: 'blur',
    },
    {
      number: true,
      message: '请输入数字',
      type: 'warning',
    },
    {
      required: true,
      message: '端口必填',
      type: 'error',
      trigger: 'change',
    },
  ],
  domain: [
    {
      required: true,
      message: '域名必填',
      type: 'error',
      trigger: 'blur',
    },
    {
      required: true,
      message: '域名必填',
      type: 'error',
      trigger: 'change',
    },
  ],
  sslCert: [
    {
      required: true,
      message: '证书必填',
      type: 'error',
      trigger: 'blur',
    },
    {
      required: true,
      message: '证书必填',
      type: 'error',
      trigger: 'change',
    },
  ],
  sslKey: [
    {
      required: true,
      message: '密钥必填',
      type: 'error',
      trigger: 'blur',
    },
    {
      required: true,
      message: '密钥必填',
      type: 'error',
      trigger: 'change',
    },
  ],
}

const onSystem : FormProps['onSubmit'] = ({ validateResult, firstError}) =>{
  if (validateResult === true){
    setSystemConfig(systemFromData.value)
  }
}

const codeRules: FormProps['rules'] = {
  api: [
    {
      required: true,
      message: 'api必填',
      type: 'error',
      trigger: 'blur',
    },
    {
      required: true,
      message: 'api必填',
      type: 'error',
      trigger: 'change',
    },
  ],
  sk: [
    {
      required: true,
      message: 'sk必填',
      type: 'error',
      trigger: 'blur',
    },
    {
      required: true,
      message: 'sk必填',
      type: 'error',
      trigger: 'change',
    },
  ],
  model: [
    {
      required: true,
      message: 'model必填',
      type: 'error',
      trigger: 'blur',
    },
    {
      required: true,
      message: 'model必填',
      type: 'error',
      trigger: 'change',
    },
  ],
  adapter: [
    {
      required: true,
      message: 'adapter必填',
      type: 'error',
      trigger: 'blur',
    },
    {
      required: true,
      message: 'adapter必填',
    }
  ],
  maxTokens: [
    {
      required: true,
      message: 'maxTokens必填',
      type: 'error',
      trigger: 'blur',
    },
    {
      required: true,
      message: 'maxTokens必填',
      type: 'error',
      trigger: 'change',
    },
  ],
}

const onCodeCompletion = ()=>{
  setCodeCompletionConfig(codeFromData.value)
}

const onChatCompletion = ()=>{
  setChatCompletionConfig(chatFromData.value)
}

const getConfig = ()=>{
 GetConfig().then((res)=>{
   if (res.code==0){
     systemFromData.value.port = res.data.port
     systemFromData.value.domain = res.data.domain
     systemFromData.value.sslCert = res.data.sslCert
     systemFromData.value.sslKey = res.data.sslKey
     codeFromData.value.api = res.data.code.api
     codeFromData.value.sk = res.data.code.sk
     codeFromData.value.model = res.data.code.model
     codeFromData.value.adapter = res.data.code.adapter
     codeFromData.value.maxTokens = res.data.code.maxTokens
     chatFromData.value.api = res.data.chat.api
     chatFromData.value.sk = res.data.chat.sk
     chatFromData.value.model = res.data.chat.model
     chatFromData.value.adapter = res.data.chat.adapter
     chatFromData.value.maxTokens = res.data.chat.maxTokens
   }
})
}
onMounted(()=>{
  getConfig()
})
</script>

<style scoped>
.form {
  margin-top: 20px;
}
.card {
  width: 100%;
  height: 100%;
}
</style>
