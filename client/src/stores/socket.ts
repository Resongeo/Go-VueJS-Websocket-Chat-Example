import { defineStore } from "pinia"

export const useSocketStore = defineStore("socket", {
  state: () => {
    return {
      server_address: null as string | null,
      username: null as string | null,
    }
  },
})
