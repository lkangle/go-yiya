import { GetAppInfo, GetIsAlwaysOnTop } from '@wailsjs/go/services/appService'
import { Environment } from '@wailsjs/runtime/runtime.js'

let os = ''
let info = {}

export async function loadEnvironment() {
    const [env, fo, onTop] = await Promise.all([Environment(), GetAppInfo(), GetIsAlwaysOnTop()])
    
    Object.assign(info, fo, env, {
        onTop,
    });
    os = env.platform
    console.log("[load]", info)
}

export const isDev = () => {
    return info.buildType === 'dev'
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

export function getAppVersion() {
    return info.version || '0.0.0'
}

export function isAlwaysOnTop() {
    return info?.onTop || false
}
