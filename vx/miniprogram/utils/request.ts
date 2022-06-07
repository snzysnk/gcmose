const domain = "http://localhost:9002/v1"

export function http(url: any, method: any, data: any) {
    return new Promise((resolve, reject) => {
        let tokenInfo = JSON.parse(wx.getStorageSync('tokenInfo'));
        let options: WechatMiniprogram.RequestOption = {
            url: domain + '/' + url,
            method: method,
            header: {Authorization: 'Bearer ' + tokenInfo.token},
            data,
            timeout: 6000,
            success: (response) => {
                if (response.statusCode >= 500) {
                    wx.showModal({
                        title: '系统提示',
                        content: '出错了',
                    });
                } else {
                    resolve(response.data)
                }
            },
            fail: reject
        }

        return wx.request(options);
    });

}