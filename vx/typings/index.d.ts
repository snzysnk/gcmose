/// <reference path="./types/index.d.ts" />

interface IAppOption {
    globalData: {
        // userInfo?: WechatMiniprogram.UserInfo,
        userInfo: Promise<WechatMiniprogram.UserInfo>
        // }
        userInfoReadyCallback?: WechatMiniprogram.GetUserInfoSuccessCallback,
    },

    parseUserInfo(e: WechatMiniprogram.UserInfo): void
    setTokenToStorage(token:string, aging:number):void
}




