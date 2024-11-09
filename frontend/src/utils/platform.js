import { GetAppInfo, GetIsAllowsOnTop } from '@wailsjs/go/services/systimeService'
import { Environment } from '@wailsjs/runtime/runtime.js'

let os = ''

let info = {}

export async function loadEnvironment() {
    const env = await Environment()
    info = await GetAppInfo()
    info.onTop = await GetIsAllowsOnTop()
    os = env.platform
}

export function isMacOS() {
    return os === 'darwin'
}

export function isWindows() {
    return os === 'windows'
}

export function getAppName() {
    return info.name || 'dev'
}

export function isAllowsOnTop() {
    return info?.onTop || false
}
