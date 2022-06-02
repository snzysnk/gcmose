// miniprogram/pages/unlock/unlock.js
Page({

    /**
     * 页面的初始数据
     */
    data: {
        head_photo: '/resource/image/head_photo.png',
        share: false,
    },

    /**
     * 生命周期函数--监听页面加载
     */
    async onLoad() {
        const userInfo = await getApp<IAppOption>().globalData.userInfo;
        if (userInfo) {
            this.setData({
                head_photo: userInfo.avatarUrl,
            });

            wx.setStorage({
                key: "userInfo",
                data: userInfo,
            });
        }
    },
    setPermission() {
        let status = !this.data.share;

        wx.setStorageSync('share', status);

        this.setData({
            share: status
        });
    },
    storeLocalAvatar() {
        //头像
    },
    getHeadPhoto() {
        wx.getUserProfile({
            desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
            success: (res) => {
                if (res.userInfo) {
                    getApp<IAppOption>().parseUserInfo(res.userInfo)
                }
            }
        })
    },
    scanUseCar() {
        wx.getLocation({
            type: 'gcj02',
            success: () => {
                let tokenInfo = JSON.parse(wx.getStorageSync('tokenInfo'));
                wx.request({
                    url: 'http://localhost:9002/v1/create/trip',
                    method: 'POST',
                    header: {Authorization: 'Bearer ' + tokenInfo.token},
                    data: {
                        cart_id: 1001,
                        start: {
                            longitude: 111.11,
                            latitude: 122.22
                        }
                    },
                    timeout: 6000,
                    success: (response) => {
                        let data = response.data as any;
                        console.log(data.trip_id);
                    },
                    fail: console.log
                });
            },
            fail: () => {
                wx.showModal({
                    title: '系统提示',
                    content: '扫码用车必须地理位置授权,请手动开启允许权限',
                })
            }
        });
    },
});