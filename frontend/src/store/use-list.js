import { defineStore } from "pinia";
import { OpenSelectFilesDialog } from "@wailsjs/go/services/fileService";
import { formatFileSize, sumBy } from "@utils/index";

const useList = defineStore("image_list", {
    state: () => ({dataList: []}),
    getters: {
        isEmpty: (store) => {
            return store.dataList.length <= 0
        },
        // 计算统计数据
        stat: (store) => {
            const list = store.dataList
            const total = store.dataList.length
            const successCount = sumBy(list, (it) => {
                return it.result.status === 3 ? 1 : 0;
            })

            let osize = sumBy(list, (it) => +it.size)
            let nsize = sumBy(list, (it) => {
                if (it.result.success) {
                    return +it.result.size
                }
                return +it.size;
            })
            let dsize = osize - nsize;
            let dtext = formatFileSize(dsize);
            let rate = (1 - nsize / osize) * 100
            let rateText = Number(rate).toFixed(1) + '%'
            return {total, successCount, dsize, dtext, rate, rateText }
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