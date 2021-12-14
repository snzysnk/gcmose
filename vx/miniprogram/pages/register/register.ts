Page({

    /**
     * 页面的初始数据
     */
    data: {
        image: '/resource/image/register/default.jpg',
        sex: 0,
        userName: '',
        date: '',
        status: '',
        sex_range: [
            '未知',
            '男',
            '女',
        ],
        upload_status: 0,
    },
    chooseImg() {
        wx.chooseImage({
            success: (res) => {
                this.setData({
                    image: res.tempFiles[0].path,
                    userName: 'xieruixiang',
                    sex: 1,
                    date: '1994-01-13',
                    status: 'checking'
                });
            }
        });
    },
    sexChange(e: any) {
        this.setData({
            sex: e.detail.value
        });
    },
    DateChange(e: any) {
        this.setData({
            date: e.detail.value
        });
    },
    submit() {
        this.setData({status: 'loading'});
        setTimeout(() => {
            this.setData({
                status: 'ok',
            });

            wx.redirectTo({
                url: "/pages/unlock/unlock",
            });
        }, 3000);
    },
});