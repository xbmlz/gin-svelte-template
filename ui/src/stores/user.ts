import { writable } from "svelte/store"

interface UserState {
    token: string
    username: string
}

export const userStore = writable<UserState>({
    token: "",
    username: ""
})