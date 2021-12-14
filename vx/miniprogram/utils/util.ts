export const formatTime = (date: Date) => {
    const year = date.getFullYear()
    const month = date.getMonth() + 1
    const day = date.getDate()
    const hour = date.getHours()
    const minute = date.getMinutes()
    const second = date.getSeconds()

    return (
        [year, month, day].map(formatNumber).join('/') +
        ' ' +
        [hour, minute, second].map(formatNumber).join(':')
    )
}

const formatNumber = (n: number) => {
    const s = n.toString()
    return s[1] ? s : '0' + s
}

export function getUserInfo(): Promise<WechatMiniprogram.GetUserProfileSuccessCallbackResult> {
    return new Promise((resolve, reject) => {
        wx.getUserProfile({
            desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
            success: resolve,
            fail: reject,
        })
    });
}


export function getSetting(): Promise<WechatMiniprogram.GetSettingSuccessCallbackResult> {
    return new Promise((resolve, reject) => {
        wx.getSetting({
            success: resolve,
            fail: reject
        })
    })
}
