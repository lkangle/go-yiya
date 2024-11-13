import create from "zustand-vue";
import { EventsOn, OnFileDrop } from '@wailsjs/runtime/runtime';
import { OpenSelectFilesDialog, GetBaseInfo } from "@wailsjs/go/services/fileService";

const useList = create((set, get) => {
    return {
        data: [],
        currentId: "",
        // 当前选中的
        current: {},
        select(current) {
            if (!current?.id) {
                return
            }
            set({current, currentId: current.id})
        },
        
        clear() {
            set({ data: [] })
        },
        _add(files) {
            let list = get().data || [];
            set({data: [...list, ...files]})
        }
    }
})

OnFileDrop(async (_, __, paths) => {
    const list = await Promise.all(paths.map((p) => GetBaseInfo(p)))
    useList.getState()._add(list)
})

export async function onSelectFiles() {
    const resp = await OpenSelectFilesDialog()
    const paths = resp.data
    useList.getState()._add(paths)
}

export default useList