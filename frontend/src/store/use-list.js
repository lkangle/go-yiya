import { defineStore } from "pinia";
import { OpenSelectFilesDialog } from "@wailsjs/go/services/fileService";

const useList = defineStore("image_list", {
    state: () => ({dataList: []}),
    getters: {
        isEmpty: (stat) => {
            return stat.dataList.length <= 0
        }
    },
    actions: {
        clear() {
            this.dataList = []
        },
        $add(list) {
            const rlist = list.map(it => ({
                ...it, result: {status: 1}
            }))
            this.dataList.push(...rlist)
        },
        async openSelectImages() {
            const resp = await OpenSelectFilesDialog()
            if (resp.success) {
                this.$add(resp.data);
            }
        },
    }
})

export default useList