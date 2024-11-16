import { defineStore } from "pinia";
import { OnFileDrop } from '@wailsjs/runtime/runtime';
import { OpenSelectFilesDialog, GetBaseInfo } from "@wailsjs/go/services/fileService";
import { onMounted } from "vue";

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
        $add(data) {
            this.dataList.push(...data)
        },
        async openSelectImages() {
            const resp = await OpenSelectFilesDialog()
            if (resp.success) {
                this.$add(resp.data);
            }
        }
    }
})

export function useOnFile() {
    const store = useList()

    onMounted(() => {
        OnFileDrop(async (_, __, paths) => {
            const list = await Promise.all(paths.map((p) => GetBaseInfo(p)))
            store.$add(list)
        })
    })
}

export default useList