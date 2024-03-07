<script setup lang="ts">
import { ref } from "vue"
import { useSocketStore } from "../stores/socket"
import router from "../router"

const server_address = ref<string>("ws://localhost:3000/ws")
const username = ref<string>("User")

const socketStore = useSocketStore()

function connect() {
  console.log(server_address.value, username.value)
  socketStore.server_address = server_address.value
  socketStore.username = username.value

  if (socketStore.server_address && socketStore.username) {
    router.push("/chat")
  }
}
</script>

<template>
  <div class="mx-auto py-10 px-5 md:px-10">
    <h1 class="font-semibold text-blue-400 text-3xl mb-10">
      Connect to the chat
    </h1>
    <div class="mb-3">
      <label for="server-address">Server Address</label>
      <input type="text" id="server-address" v-model="server_address" />
    </div>

    <div class="mb-3">
      <label for="username">Username</label>
      <input type="text" id="username" v-model="username" />
    </div>

    <button @click="connect">Connect</button>
  </div>
</template>

<style scoped>
label {
  display: block;
  @apply text-neutral-400 mb-2 text-lg;
}
</style>
