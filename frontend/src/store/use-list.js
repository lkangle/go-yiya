import create from "zustand-vue";

const useList = create((set, get) => {
    return {
        data: [],
        clear() {
            set({ data: [] })
        }
    }
})

export default useList