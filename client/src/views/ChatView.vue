<script setup lang="ts">
import { ref, onMounted } from "vue"
import { useSocketStore } from "../stores/socket"
import router from "../router"

const socketStore = useSocketStore()
const connection_ready = ref<boolean>(false)
const messages = ref<string[]>([])
const message = ref<string>("")
let socket = null as WebSocket | null

onMounted(() => {
  if (socketStore.server_address === null || socketStore.username === null) {
    router.push("/")
  }

  socket = new WebSocket(socketStore.server_address as string)
  socket.onopen = onSocketOpen
  socket.onmessage = onSocketMessage
  socket.onclose = onSocketErrorOrClose
  socket.onerror = onSocketErrorOrClose
})

function onSocketOpen() {
  connection_ready.value = true
}

function onSocketMessage(event: MessageEvent) {
  const msg: string = event.data as string
  messages.value.push(msg)
}

function onSocketErrorOrClose() {
  connection_ready.value = false
  socketStore.server_address = null
  socketStore.username = null
  socket = null

  router.push("/")
}

function myMessage(msg: string) {
  return socketStore.username === msg.split(":")[0]
}

function sendMessage() {
  if (message.value === "") {
    return
  }

  socket?.send(`${socketStore.username}:${message.value}`)
  message.value = ""
}
</script>

<template>
  <div v-if="!connection_ready">
    <div
      class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-20 h-20 text-4xl animate-pulse"
    >
      Connecting...
    </div>
  </div>
  <div v-else>
    <input type="text" v-model="message" />
    <button @click="sendMessage" @keyup.enter="sendMessage">Send</button>
    <ul class="p-2 text-lg text-neutral-300">
      <li v-for="msg in messages">
        <div class="flex justify-end items-center" v-if="myMessage(msg)">
          <div class="bg-blue-500 px-2 py-1 rounded-lg w-fit mb-2">
            {{ msg.split(":")[1] }}
          </div>
        </div>
        <div class="flex justify-start items-center" v-else>
          <div class="bg-[#233654] px-2 py-1 rounded-lg w-fit mb-2">
            {{ msg.split(":")[1] }}
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<style scoped>
button {
  margin-top: 1rem !important;
}
</style>
