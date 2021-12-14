// index.ts
// 获取应用实例


import {Route} from "../../utils/route";

Page({
    data: {
        motto: 'Hello World',

        avatar: '/resource/image/home.png',
        latitude: 31,
        longitude: 121,
        scale: 10,
        userInfo: {},
        hasUserInfo: false,
        canIUse: wx.canIUse('button.open-3type.getUserInfo'),
        canIUseGetUserProfile: false,
        // canIUseOpenData: wx.canIUse('open-data.type.userAvatarUrl') && wx.canIUse('open-data.type.userNickName') // 如需尝试获取用户信息可改为false
        canIUseOpenData: false,
    },
    // 事件处理函数
    bindViewTap() {

    },
    async onLoad() {
        const userInfo = await getApp<IAppOption>().globalData.userInfo;
        if (userInfo) {
            this.setData({
                avatar: userInfo.avatarUrl
            })
        }
    },
    confirmLocation() {
        wx.getLocation({
            type: 'gcj02',
            success: (res) => {
                console.log('success', res);
                this.setData({
                    longitude: res.longitude,
                    latitude: res.latitude
                })
            },
            fail: console.log
        });
    },
    getCarByScan() {
        wx.navigateTo({
            // url: "/pages/register/register",
            url: Route.register(2)
        });

        // wx.scanCode({
        //     success: (res) => {
        //         console.log(res);
        //         wx.navigateTo({
        //             // url: "/pages/register/register",
        //             url: Route.register(2)
        //         });
        //     },
        //     fail: console.error
        // });
    },
    goHome() {
        wx.navigateTo({url: "/pages/home/home"});
    },
    // async onLoad() {
    //     const userInfo = await getApp<IAppOption>().globalData.userInfo;
    //     if (userInfo) {
    //         this.setData({
    //             hasUserInfo:true,
    //             avatarUrl: userInfo.avatarUrl,
    //             nickName: userInfo.nickName
    //         })
    //     }
    // },
    getUserProfile() {
        wx.getUserProfile({
            desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
            success: (res) => {
                if (res.userInfo) {
                    getApp<IAppOption>().parseUserInfo(res.userInfo)
                }
            }
        })
        // const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo;
        // getApp<IAppOption>().parseUserInfo(userInfo);
    }
})
