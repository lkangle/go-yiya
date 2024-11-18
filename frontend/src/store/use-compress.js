import { CompressImage } from "@wailsjs/go/services/fileService"
import useOptions from "./use-options"
import pLimit from 'p-limit'
import { withError } from "@utils/index"

const limit = pLimit(5)

const useCompress = () => {
    const optStore = useOptions()

    const emitCompress = async (item, onProgress) => {
        limit(async () => {
            onProgress({status: 2})

            let [r, err] = await withError(CompressImage(item.path, optStore.data))
            if (err != null || !r.success) {
                onProgress({ status: 4, message: r.message||err.message})
            } else {
                onProgress({...r, status: 3})
            }
        })
    }
    return emitCompress
}

export default useCompress