import {http} from "../../utils/request";

Page({

    /**
     * 页面的初始数据
     */
    data: {
        image: '/resource/image/register/default.jpg',
        sex: 0,
        userName: '',
        timer: 0,
        date: '',
        status: '',
        sex_range: [
            '未知',
            '男',
            '女',
        ],
        upload_status: 0,
    },
    async onLoad() {
        let response: any = await http("profile/data", "POST", {})
        if (response.profile.status == 3) {
            wx.redirectTo({
                url: "/pages/unlock/unlock",
            });
        } else {
            this.setData({
                path: response.profile.path ?? '/resource/image/register/default.jpg',
                userName: response.profile.name ?? "",
                sex: response.profile.sex ?? 0,
                date: '1994-01-13',
                status: '',
            });
        }
    }
    ,
    async chooseImg() {
        const response: any = await http("profile/uploadUrl", "GET", {})
        const uploadUrl: string = response.url

        wx.chooseImage({
            success: (res) => {
                const data = wx.getFileSystemManager().readFileSync(res.tempFilePaths[0])
                wx.request({
                    url: uploadUrl,
                    method: "PUT",
                    header: {"content-type": "image/png"},
                    data: data,
                    success: () => {
                        this.setData({
                            image: res.tempFiles[0].path,
                            userName: 'xieruixiang',
                            sex: 1,
                            date: '1994-01-13',
                            status: 'checking',
                        });
                    },
                    fail: console.log
                })
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
    onUnload() {
        clearInterval(this.data.timer)
    },
    async submit() {
        this.setData({status: 'loading'});
        let name = this.data.userName;
        let sex = this.data.sex;
        let birth = (new Date(this.data.date)).getTime();
        let data = {name, sex, birth}
        await http("profile/check", "POST", data)
        let that = this
        this.data.timer = setInterval(function () {
            that.getProfile()
        }, 1000)

        // path: response.profile.path ?? '/resource/image/register/default.jpg',
        //     userName: response.profile.name ?? "",
        //     sex: response.profile.sex ?? 0,
        //     date: '1994-01-13',
        //     status: '',

        // data['birth'] =
        // setTimeout(() => {
        //     this.setData({
        //         status: 'ok',
        //     });
        //
        //     wx.redirectTo({
        //         url: "/pages/unlock/unlock",
        //     });
        // }, 3000);
    },
    async getProfile() {
        let response: any = await http("profile/data", "POST", {})
        if (response.profile.status == 3) {
            wx.redirectTo({
                url: "/pages/unlock/unlock",
            });
        } else {
            this.setData({"status":""})
        }
    },
});