import { defineStore } from "pinia";
import { OpenSelectFilesDialog } from "@wailsjs/go/services/fileService";

const useList = defineStore("image_list", {
    state: () => ({dataList: [], currentId: ''}),
    getters: {
        current: (stat) => {
            return stat.dataList.find(it => it.id === stat.currentId)
        },
        isEmpty: (stat) => {
            return stat.dataList.length <= 0
        }
    },
    actions: {
        select(id) {
            if (id) {
                this.currentId = id
            }
        },
        clear() {
            this.dataList = []
        },
        $add(list) {
            this.dataList.push(...list)
        },
        async openSelectImages() {
            const resp = await OpenSelectFilesDialog()
            if (resp.success) {
                this.$add(resp.data);
            }
        }
    }
})

export default useList