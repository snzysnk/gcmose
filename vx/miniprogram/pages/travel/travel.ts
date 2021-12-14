// miniprogram/pages/travel/travel.js

const addZero = (time: number) => {
    return time > 9 ? time : '0' + time;
}

const changeFee = (money: number) => {
    let fee = (money / 100);
    return fee.toFixed(2);
}

const formTime = (time: number) => {
    let second = time % 60;
    let overTime = time - second;
    let minutes = overTime / 60;
    let minute = minutes % 60;
    let hour = (minutes - minute) / 60;

    let s = addZero(second);
    let m = addZero(minute);
    let h = addZero(hour);

    return `${h}:${m}:${s}`;
}
Page({
    timer: undefined as number | undefined,
    /**
     * 页面的初始数据
     */
    data: {
        latitude: 31,
        longitude: 121,
        scale: 10,
        money: '0.00',
        time: '00:00:00',
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function () {
        this.setTimer();
    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady: function () {

    },

    setTimer() {
        let timeCount = 0;
        let fee = 0;
        this.timer = setInterval(() => {
            this.setData({
                time: formTime(timeCount++),
                money: changeFee((fee++)*0.7)
            })
        }, 1000);
    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow: function () {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide: function () {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload: function () {
        if (this.timer) {
            clearInterval(this.timer)
        }
    },
});