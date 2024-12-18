import { CompressImage } from "@wailsjs/go/services/fileService"
import useOptions from "./use-options"
import pLimit from 'p-limit'
import { withError } from "@utils/index"

const limit = pLimit(5)

// status
// 1-等待压缩
// 2-压缩中
// 3-压缩成功
// 4-压缩失败
// 5-被还原了
const useCompress = () => {
    const optStore = useOptions()

    const emitCompress = async (item, onProgress) => {
        await limit(async () => {
            onProgress({status: 2})

            let [r, err] = await withError(CompressImage(item, optStore.data))
            if (err != null || !r.success) {
                onProgress({ ...r, status: 4, message: r.message||err.message})
            } else {
                onProgress({...r, status: 3})
            }
        })
    }
    return emitCompress
}

export default useCompress