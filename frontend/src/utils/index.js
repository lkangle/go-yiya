export function debounce(fn, delay) {
    let timeout
    return function (...args) {
        clearTimeout(timeout)

        timeout = setTimeout(() => {
            fn(...args)
        }, delay)
    }
}

export function range(start, end) {
    return new Array(end-start).fill(0).map((_, index) => index+start)
}

export function formatFileSize(bytes) {
    bytes = +bytes
    if (!bytes) {
        return '0B'
    }
    const sizes = ["B", "KB", "MB", "GB", "TB"];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    
    // 计算对应单位的值并保留两位小数
    const formattedSize = (bytes / Math.pow(1024, i)).toFixed(2);
    
    return `${formattedSize}${sizes[i]}`;
}