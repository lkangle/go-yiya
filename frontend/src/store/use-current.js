import { defineStore } from "pinia"

const useCurrent = defineStore("image-current", {
    state:() => ({ current: {id:"", result:{}} }),
    getters: {
        // 是否可以撤销压缩
        canUndo() {
            const status = this.current.result?.status;
            return status === 3
        },
        // 是否可以重新压缩
        canRedo() {
            const status = this.current.result?.status;
            const code = this.current.result?.code;
            return status === 5 || (status === 4 && code !== 1055)
        }
    },
    actions: {
        updateCurrent(item) {
            this.current = item
        }
    }
})

export default useCurrent