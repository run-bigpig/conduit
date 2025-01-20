import {
    SetChatCompletionConfig,
    SetCodeCompletionConfig,
    SetSystemConfig,
    Start, Stop
} from "../../wailsjs/go/gui/GuiHandler";
import {NotifyPlugin} from "tdesign-vue-next";
import {Ref} from "vue";

export interface SystemConfig {
    port: number,
    domain: string,
    sslCert: string,
    sslKey: string
}

export function setSystemConfig(conf: SystemConfig) {
    SetSystemConfig(conf).then((res) => {
        if (res.code==0){
            NotifyPlugin.success({"title":"系统设置", "content":res.msg})
            return
        }
        NotifyPlugin('error', { title: '系统设置', "content":res.msg })
    })
}

export interface CompletionConfig {
    api: string
    sk: string
    model: string
    adapter: string
    maxTokens: number
    locale?: string
}

export function setCodeCompletionConfig(conf: CompletionConfig) {
    SetCodeCompletionConfig(conf).then((res) => {
        if (res.code==0){
            NotifyPlugin.success({"title":"代码补全设置", "content":res.msg})
            return
        }
        NotifyPlugin('error', { title: '代码补全设置', "content":res.msg })
    })
}

export function setChatCompletionConfig(conf: CompletionConfig) {
    SetChatCompletionConfig(conf).then((res) => {
        if (res.code==0){
            NotifyPlugin.success({"title":"对话补全设置", "content":res.msg})
            return
        }
        NotifyPlugin('error', { title: '对话补全设置', "content":res.msg })
    })
}

export function start(running:Ref<boolean>){
    Start().then((res) => {
        if (res.code==0){
            running.value = true
            NotifyPlugin.success({"title":"启动代理", "content":res.msg})
            return
        }
        NotifyPlugin('error', { title: '启动代理', "content":res.msg })
    })
}

export function stop(running:Ref<boolean>){
    Stop().then((res) => {
        if (res.code==0){
            running.value = false
            NotifyPlugin.success({"title":"停止代理", "content":res.msg})
            return
        }
        NotifyPlugin('error', { title: '停止代理', "content":res.msg })

    })
}