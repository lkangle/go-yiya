import create from "zustand-vue";
import { EventsOn, OnFileDrop } from '@wailsjs/runtime/runtime';
import { OpenSelectFilesDialog } from "@wailsjs/go/services/fileService";

const useList = create((set, get) => {
    return {
        data: [],
        clear() {
            set({ data: [] })
        },
        add(files) {
            let list = get().data || [];
            set({data: [...list, ...files]})
        }
    }
})

OnFileDrop((x, y, paths) => {
    console.log("drop paths", x, y, paths)

    useList.getState().add(paths)
})

export async function onSelectFiles() {
    const resp = await OpenSelectFilesDialog()
    const paths = resp.data

    console.log("select paths", paths)
    useList.getState().add(paths)
}

export default useList