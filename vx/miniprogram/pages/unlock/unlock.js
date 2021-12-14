"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Page({
    data: {
        head_photo: '/resource/image/head_photo.png',
        share: false,
    },
    onLoad() {
        return __awaiter(this, void 0, void 0, function* () {
            const userInfo = yield getApp().globalData.userInfo;
            if (userInfo) {
                this.setData({
                    head_photo: userInfo.avatarUrl,
                });
                wx.setStorage({
                    key: "userInfo",
                    data: userInfo,
                });
            }
        });
    },
    setPermission() {
        let status = !this.data.share;
        wx.setStorageSync('share', status);
        this.setData({
            share: status
        });
    },
    storeLocalAvatar() {
    },
    getHeadPhoto() {
        wx.getUserProfile({
            desc: '展示用户信息',
            success: (res) => {
                if (res.userInfo) {
                    getApp().parseUserInfo(res.userInfo);
                }
            }
        });
    },
    scanUseCar() {
        wx.getLocation({
            type: 'gcj02',
            success: () => {
                wx.redirectTo({
                    url: '/pages/travel/travel',
                });
            },
            fail: () => {
                wx.showModal({
                    title: '系统提示',
                    content: '扫码用车必须地理位置授权,请手动开启允许权限',
                });
            }
        });
    },
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoidW5sb2NrLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsidW5sb2NrLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7QUFDQSxJQUFJLENBQUM7SUFLRCxJQUFJLEVBQUU7UUFDRixVQUFVLEVBQUUsZ0NBQWdDO1FBQzVDLEtBQUssRUFBRSxLQUFLO0tBQ2Y7SUFLSyxNQUFNOztZQUNSLE1BQU0sUUFBUSxHQUFHLE1BQU0sTUFBTSxFQUFjLENBQUMsVUFBVSxDQUFDLFFBQVEsQ0FBQztZQUNoRSxJQUFJLFFBQVEsRUFBRTtnQkFDVixJQUFJLENBQUMsT0FBTyxDQUFDO29CQUNULFVBQVUsRUFBRSxRQUFRLENBQUMsU0FBUztpQkFDakMsQ0FBQyxDQUFDO2dCQUVILEVBQUUsQ0FBQyxVQUFVLENBQUM7b0JBQ1YsR0FBRyxFQUFFLFVBQVU7b0JBQ2YsSUFBSSxFQUFFLFFBQVE7aUJBQ2pCLENBQUMsQ0FBQzthQUNOO1FBQ0wsQ0FBQztLQUFBO0lBQ0QsYUFBYTtRQUNULElBQUksTUFBTSxHQUFHLENBQUMsSUFBSSxDQUFDLElBQUksQ0FBQyxLQUFLLENBQUM7UUFFOUIsRUFBRSxDQUFDLGNBQWMsQ0FBQyxPQUFPLEVBQUUsTUFBTSxDQUFDLENBQUM7UUFFbkMsSUFBSSxDQUFDLE9BQU8sQ0FBQztZQUNULEtBQUssRUFBRSxNQUFNO1NBQ2hCLENBQUMsQ0FBQztJQUNQLENBQUM7SUFDRCxnQkFBZ0I7SUFFaEIsQ0FBQztJQUNELFlBQVk7UUFDUixFQUFFLENBQUMsY0FBYyxDQUFDO1lBQ2QsSUFBSSxFQUFFLFFBQVE7WUFDZCxPQUFPLEVBQUUsQ0FBQyxHQUFHLEVBQUUsRUFBRTtnQkFDYixJQUFJLEdBQUcsQ0FBQyxRQUFRLEVBQUU7b0JBQ2QsTUFBTSxFQUFjLENBQUMsYUFBYSxDQUFDLEdBQUcsQ0FBQyxRQUFRLENBQUMsQ0FBQTtpQkFDbkQ7WUFDTCxDQUFDO1NBQ0osQ0FBQyxDQUFBO0lBQ04sQ0FBQztJQUNELFVBQVU7UUFDTixFQUFFLENBQUMsV0FBVyxDQUFDO1lBQ1gsSUFBSSxFQUFFLE9BQU87WUFDYixPQUFPLEVBQUUsR0FBRyxFQUFFO2dCQUNWLEVBQUUsQ0FBQyxVQUFVLENBQUM7b0JBQ1YsR0FBRyxFQUFFLHNCQUFzQjtpQkFDOUIsQ0FBQyxDQUFDO1lBQ1AsQ0FBQztZQUNELElBQUksRUFBRSxHQUFHLEVBQUU7Z0JBQ1AsRUFBRSxDQUFDLFNBQVMsQ0FBQztvQkFDVCxLQUFLLEVBQUUsTUFBTTtvQkFDYixPQUFPLEVBQUUsd0JBQXdCO2lCQUNwQyxDQUFDLENBQUE7WUFDTixDQUFDO1NBQ0osQ0FBQyxDQUFDO0lBQ1AsQ0FBQztDQUNKLENBQUMsQ0FBQyIsInNvdXJjZXNDb250ZW50IjpbIi8vIG1pbmlwcm9ncmFtL3BhZ2VzL3VubG9jay91bmxvY2suanNcblBhZ2Uoe1xuXG4gICAgLyoqXG4gICAgICog6aG16Z2i55qE5Yid5aeL5pWw5o2uXG4gICAgICovXG4gICAgZGF0YToge1xuICAgICAgICBoZWFkX3Bob3RvOiAnL3Jlc291cmNlL2ltYWdlL2hlYWRfcGhvdG8ucG5nJyxcbiAgICAgICAgc2hhcmU6IGZhbHNlLFxuICAgIH0sXG5cbiAgICAvKipcbiAgICAgKiDnlJ/lkb3lkajmnJ/lh73mlbAtLeebkeWQrOmhtemdouWKoOi9vVxuICAgICAqL1xuICAgIGFzeW5jIG9uTG9hZCgpIHtcbiAgICAgICAgY29uc3QgdXNlckluZm8gPSBhd2FpdCBnZXRBcHA8SUFwcE9wdGlvbj4oKS5nbG9iYWxEYXRhLnVzZXJJbmZvO1xuICAgICAgICBpZiAodXNlckluZm8pIHtcbiAgICAgICAgICAgIHRoaXMuc2V0RGF0YSh7XG4gICAgICAgICAgICAgICAgaGVhZF9waG90bzogdXNlckluZm8uYXZhdGFyVXJsLFxuICAgICAgICAgICAgfSk7XG5cbiAgICAgICAgICAgIHd4LnNldFN0b3JhZ2Uoe1xuICAgICAgICAgICAgICAgIGtleTogXCJ1c2VySW5mb1wiLFxuICAgICAgICAgICAgICAgIGRhdGE6IHVzZXJJbmZvLFxuICAgICAgICAgICAgfSk7XG4gICAgICAgIH1cbiAgICB9LFxuICAgIHNldFBlcm1pc3Npb24oKSB7XG4gICAgICAgIGxldCBzdGF0dXMgPSAhdGhpcy5kYXRhLnNoYXJlO1xuXG4gICAgICAgIHd4LnNldFN0b3JhZ2VTeW5jKCdzaGFyZScsIHN0YXR1cyk7XG5cbiAgICAgICAgdGhpcy5zZXREYXRhKHtcbiAgICAgICAgICAgIHNoYXJlOiBzdGF0dXNcbiAgICAgICAgfSk7XG4gICAgfSxcbiAgICBzdG9yZUxvY2FsQXZhdGFyKCkge1xuICAgICAgICAvL+WktOWDj1xuICAgIH0sXG4gICAgZ2V0SGVhZFBob3RvKCkge1xuICAgICAgICB3eC5nZXRVc2VyUHJvZmlsZSh7XG4gICAgICAgICAgICBkZXNjOiAn5bGV56S655So5oi35L+h5oGvJywgLy8g5aOw5piO6I635Y+W55So5oi35Liq5Lq65L+h5oGv5ZCO55qE55So6YCU77yM5ZCO57ut5Lya5bGV56S65Zyo5by556qX5Lit77yM6K+36LCo5oWO5aGr5YaZXG4gICAgICAgICAgICBzdWNjZXNzOiAocmVzKSA9PiB7XG4gICAgICAgICAgICAgICAgaWYgKHJlcy51c2VySW5mbykge1xuICAgICAgICAgICAgICAgICAgICBnZXRBcHA8SUFwcE9wdGlvbj4oKS5wYXJzZVVzZXJJbmZvKHJlcy51c2VySW5mbylcbiAgICAgICAgICAgICAgICB9XG4gICAgICAgICAgICB9XG4gICAgICAgIH0pXG4gICAgfSxcbiAgICBzY2FuVXNlQ2FyKCkge1xuICAgICAgICB3eC5nZXRMb2NhdGlvbih7XG4gICAgICAgICAgICB0eXBlOiAnZ2NqMDInLFxuICAgICAgICAgICAgc3VjY2VzczogKCkgPT4ge1xuICAgICAgICAgICAgICAgIHd4LnJlZGlyZWN0VG8oe1xuICAgICAgICAgICAgICAgICAgICB1cmw6ICcvcGFnZXMvdHJhdmVsL3RyYXZlbCcsXG4gICAgICAgICAgICAgICAgfSk7XG4gICAgICAgICAgICB9LFxuICAgICAgICAgICAgZmFpbDogKCkgPT4ge1xuICAgICAgICAgICAgICAgIHd4LnNob3dNb2RhbCh7XG4gICAgICAgICAgICAgICAgICAgIHRpdGxlOiAn57O757uf5o+Q56S6JyxcbiAgICAgICAgICAgICAgICAgICAgY29udGVudDogJ+aJq+eggeeUqOi9puW/hemhu+WcsOeQhuS9jee9ruaOiOadgyzor7fmiYvliqjlvIDlkK/lhYHorrjmnYPpmZAnLFxuICAgICAgICAgICAgICAgIH0pXG4gICAgICAgICAgICB9XG4gICAgICAgIH0pO1xuICAgIH0sXG59KTsiXX0=