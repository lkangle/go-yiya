import { defineStore } from "pinia"

const useCurrent = defineStore("image-current", {
    state:() => ({ current: {id:""} }),
    actions: {
        updateCurrent(item) {
            Object.assign(this.current, item)
        }
    }
})

export default useCurrent