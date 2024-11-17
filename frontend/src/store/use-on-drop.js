import { ParseFilePaths } from "@wailsjs/go/services/fileService"
import { OnFileDrop, OnFileDropOff } from "@wailsjs/runtime/runtime"
import { onMounted, onUnmounted } from "vue"
import useList from "./use-list"

function useOnDrop() {
    const store = useList()
    const add = store.$add.bind(store)

    onMounted(() => {
        OnFileDrop(async (x, y, paths) => {
            const infos = await ParseFilePaths(paths).catch(() => ([]))
            add(infos);
        })
    })

    onUnmounted(() => {
        OnFileDropOff()
    })
}

export default useOnDrop