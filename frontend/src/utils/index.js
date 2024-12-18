export function debounce(fn, delay, leading = false) {
    let timer = null;
    let hasLeadingCalled = false;

    return function (...args) {
        const context = this;

        if (leading && !hasLeadingCalled) {
            // 如果需要立即执行，并且尚未执行过
            fn.apply(context, args);
            hasLeadingCalled = true;
        }

        clearTimeout(timer); // 清除上一次的定时器

        timer = setTimeout(() => {
            if (!leading) {
                // 如果不是 leading 模式，延迟结束后调用
                fn.apply(context, args);
            }
            hasLeadingCalled = false; // 重置为初始状态
        }, delay);
    };
}


export function range(start, end) {
    return new Array(end-start).fill(0).map((_, index) => index+start)
}

export function formatFileSize(bytes) {
    bytes = Math.abs(+bytes)
    if (!bytes) {
        return '0B'
    }
    const sizes = ["B", "KB", "MB", "GB", "TB"];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    
    // 计算对应单位的值并保留两位小数
    const formattedSize = (bytes / Math.pow(1024, i)).toFixed(2);
    
    return `${formattedSize}${sizes[i]}`;
}

export const withError = async (promiseFn) => {
    let result;
    try {
        let promise;
        if (typeof promiseFn === 'function') {
            promise = promiseFn()
        } else {
            promise = promiseFn
        }

        const data = await promise
        return [data, null]
    } catch(err) {
        result = [null, err]
    }

    return result;
}

export const lessRate = (old, now) => {
    let osize = old.size || old || 1;
    let nsize = now.size || now || 0;

    let r = (1 - nsize / osize) * -100;
    let t = Number(r).toFixed(1) + '%'
    if (r>0) {
        return '+' + t;
    }
    return t;
}

export const sumBy = (list, get) => {
    const s = (list||[]).reduce((sum, it) => {
        let v = get(it)
        return sum + v
    }, 0);
    return s
}