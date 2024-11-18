import { debounce } from "@utils/index"
import { GetCompressOptions, UpdateCompressOptions } from "@wailsjs/go/services/appService"
import { defineStore } from "pinia"
import { reactive } from "vue"

const useOptions = defineStore("compress-options",  () => {
    const data = reactive({
        quality: 1,
        override: true,
        newSuffix: ''
    })

    GetCompressOptions().then((resp) => {
        if(resp.success) {
            Object.assign(data, resp.data)
        }
    })

    const save = debounce((data) => {
        UpdateCompressOptions(data)
    }, 500);

    return { data, save }
})

export default useOptions