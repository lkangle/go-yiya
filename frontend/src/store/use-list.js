import { defineStore } from "pinia";
import { OpenSelectFilesDialog } from "@wailsjs/go/services/fileService";
import { formatFileSize, lessRate, sumBy } from "@utils/index";

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
            let rate = 1 - nsize / osize
            let rateText = lessRate(osize, nsize)

            let dsize = osize - nsize;
            let dtext = formatFileSize(dsize);
            if (dsize < 0) {
                dtext = `+${dtext}`
            }

            return {total, successCount, rate, rateText, dsize, dtext }
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