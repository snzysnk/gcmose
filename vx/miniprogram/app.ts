// app.ts

let resolveUserInfo: any;

App<IAppOption>({
    globalData: {
        userInfo: new Promise((resolve) => {
            resolveUserInfo = resolve;
        })
    },
    onLaunch() {
        // 展示本地存储能力
        // const logs = wx.getStorageSync('logs') || []
        // logs.unshift(Date.now())
        //
        // // const userInfo = wx.getStorageSync('userInfo') || false
        //
        // // if (userInfo) {
        // //     console.log(getApp().globalData.userInfo);
        // // }
        //
        //
        // // 登录
        wx.login({
            success: res => {
                wx.request({
                    url: 'http://localhost:9002/v1/auth/login',
                    method: 'POST',
                    data: {code: res.code},
                    timeout: 6000,
                    success: (response) => {
                        let data = response.data as any;
                        console.log(data.Token);
                    },
                    fail: console.log
                });
            },
        })
    },
    parseUserInfo(e: WechatMiniprogram.UserInfo) {
        resolveUserInfo(e)
    }
})