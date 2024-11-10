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