import {defineStore} from 'pinia'

export const useMainStore = defineStore('main', {
    state: () => {
        return {
            count: 0,
            username: 'iChaly',
        }
    },
    actions: {
        increment() {
            this.count++
        },
    }
})