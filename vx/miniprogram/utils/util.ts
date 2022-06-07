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

export const formatTimeTwo = (n: number, format: string) => {
    let formateArr = ['Y', 'M', 'D', 'h', 'm', 's'];
    let returnArr: any = [];

    let date = new Date(n);
    returnArr.push(date.getFullYear());
    returnArr.push(formatNumber(date.getMonth() + 1));
    returnArr.push(formatNumber(date.getDate()));

    returnArr.push(formatNumber(date.getHours()));
    returnArr.push(formatNumber(date.getMinutes()));
    returnArr.push(formatNumber(date.getSeconds()));

    for (let i = 0; i < returnArr.length; i++) {
        format = format.replace(formateArr[i], returnArr[i]);
    }
    return format;
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
