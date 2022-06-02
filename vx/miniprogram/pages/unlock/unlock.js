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
                let tokenInfo = JSON.parse(wx.getStorageSync('tokenInfo'));
                wx.request({
                    url: 'http://localhost:9002/v1/create/trip',
                    method: 'POST',
                    header: { Authorization: 'Bearer ' + tokenInfo.token },
                    data: {
                        cart_id: 1001,
                        start: {
                            longitude: 111.11,
                            latitude: 122.22
                        }
                    },
                    timeout: 6000,
                    success: (response) => {
                        let data = response.data;
                        console.log(data.trip_id);
                    },
                    fail: console.log
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
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoidW5sb2NrLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsidW5sb2NrLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7QUFDQSxJQUFJLENBQUM7SUFLRCxJQUFJLEVBQUU7UUFDRixVQUFVLEVBQUUsZ0NBQWdDO1FBQzVDLEtBQUssRUFBRSxLQUFLO0tBQ2Y7SUFLSyxNQUFNOztZQUNSLE1BQU0sUUFBUSxHQUFHLE1BQU0sTUFBTSxFQUFjLENBQUMsVUFBVSxDQUFDLFFBQVEsQ0FBQztZQUNoRSxJQUFJLFFBQVEsRUFBRTtnQkFDVixJQUFJLENBQUMsT0FBTyxDQUFDO29CQUNULFVBQVUsRUFBRSxRQUFRLENBQUMsU0FBUztpQkFDakMsQ0FBQyxDQUFDO2dCQUVILEVBQUUsQ0FBQyxVQUFVLENBQUM7b0JBQ1YsR0FBRyxFQUFFLFVBQVU7b0JBQ2YsSUFBSSxFQUFFLFFBQVE7aUJBQ2pCLENBQUMsQ0FBQzthQUNOO1FBQ0wsQ0FBQztLQUFBO0lBQ0QsYUFBYTtRQUNULElBQUksTUFBTSxHQUFHLENBQUMsSUFBSSxDQUFDLElBQUksQ0FBQyxLQUFLLENBQUM7UUFFOUIsRUFBRSxDQUFDLGNBQWMsQ0FBQyxPQUFPLEVBQUUsTUFBTSxDQUFDLENBQUM7UUFFbkMsSUFBSSxDQUFDLE9BQU8sQ0FBQztZQUNULEtBQUssRUFBRSxNQUFNO1NBQ2hCLENBQUMsQ0FBQztJQUNQLENBQUM7SUFDRCxnQkFBZ0I7SUFFaEIsQ0FBQztJQUNELFlBQVk7UUFDUixFQUFFLENBQUMsY0FBYyxDQUFDO1lBQ2QsSUFBSSxFQUFFLFFBQVE7WUFDZCxPQUFPLEVBQUUsQ0FBQyxHQUFHLEVBQUUsRUFBRTtnQkFDYixJQUFJLEdBQUcsQ0FBQyxRQUFRLEVBQUU7b0JBQ2QsTUFBTSxFQUFjLENBQUMsYUFBYSxDQUFDLEdBQUcsQ0FBQyxRQUFRLENBQUMsQ0FBQTtpQkFDbkQ7WUFDTCxDQUFDO1NBQ0osQ0FBQyxDQUFBO0lBQ04sQ0FBQztJQUNELFVBQVU7UUFDTixFQUFFLENBQUMsV0FBVyxDQUFDO1lBQ1gsSUFBSSxFQUFFLE9BQU87WUFDYixPQUFPLEVBQUUsR0FBRyxFQUFFO2dCQUNWLElBQUksU0FBUyxHQUFHLElBQUksQ0FBQyxLQUFLLENBQUMsRUFBRSxDQUFDLGNBQWMsQ0FBQyxXQUFXLENBQUMsQ0FBQyxDQUFDO2dCQUMzRCxFQUFFLENBQUMsT0FBTyxDQUFDO29CQUNQLEdBQUcsRUFBRSxzQ0FBc0M7b0JBQzNDLE1BQU0sRUFBRSxNQUFNO29CQUNkLE1BQU0sRUFBRSxFQUFDLGFBQWEsRUFBRSxTQUFTLEdBQUcsU0FBUyxDQUFDLEtBQUssRUFBQztvQkFDcEQsSUFBSSxFQUFFO3dCQUNGLE9BQU8sRUFBRSxJQUFJO3dCQUNiLEtBQUssRUFBRTs0QkFDSCxTQUFTLEVBQUUsTUFBTTs0QkFDakIsUUFBUSxFQUFFLE1BQU07eUJBQ25CO3FCQUNKO29CQUNELE9BQU8sRUFBRSxJQUFJO29CQUNiLE9BQU8sRUFBRSxDQUFDLFFBQVEsRUFBRSxFQUFFO3dCQUNsQixJQUFJLElBQUksR0FBRyxRQUFRLENBQUMsSUFBVyxDQUFDO3dCQUNoQyxPQUFPLENBQUMsR0FBRyxDQUFDLElBQUksQ0FBQyxPQUFPLENBQUMsQ0FBQztvQkFDOUIsQ0FBQztvQkFDRCxJQUFJLEVBQUUsT0FBTyxDQUFDLEdBQUc7aUJBQ3BCLENBQUMsQ0FBQztZQUNQLENBQUM7WUFDRCxJQUFJLEVBQUUsR0FBRyxFQUFFO2dCQUNQLEVBQUUsQ0FBQyxTQUFTLENBQUM7b0JBQ1QsS0FBSyxFQUFFLE1BQU07b0JBQ2IsT0FBTyxFQUFFLHdCQUF3QjtpQkFDcEMsQ0FBQyxDQUFBO1lBQ04sQ0FBQztTQUNKLENBQUMsQ0FBQztJQUNQLENBQUM7Q0FDSixDQUFDLENBQUMiLCJzb3VyY2VzQ29udGVudCI6WyIvLyBtaW5pcHJvZ3JhbS9wYWdlcy91bmxvY2svdW5sb2NrLmpzXG5QYWdlKHtcblxuICAgIC8qKlxuICAgICAqIOmhtemdoueahOWIneWni+aVsOaNrlxuICAgICAqL1xuICAgIGRhdGE6IHtcbiAgICAgICAgaGVhZF9waG90bzogJy9yZXNvdXJjZS9pbWFnZS9oZWFkX3Bob3RvLnBuZycsXG4gICAgICAgIHNoYXJlOiBmYWxzZSxcbiAgICB9LFxuXG4gICAgLyoqXG4gICAgICog55Sf5ZG95ZGo5pyf5Ye95pWwLS3nm5HlkKzpobXpnaLliqDovb1cbiAgICAgKi9cbiAgICBhc3luYyBvbkxvYWQoKSB7XG4gICAgICAgIGNvbnN0IHVzZXJJbmZvID0gYXdhaXQgZ2V0QXBwPElBcHBPcHRpb24+KCkuZ2xvYmFsRGF0YS51c2VySW5mbztcbiAgICAgICAgaWYgKHVzZXJJbmZvKSB7XG4gICAgICAgICAgICB0aGlzLnNldERhdGEoe1xuICAgICAgICAgICAgICAgIGhlYWRfcGhvdG86IHVzZXJJbmZvLmF2YXRhclVybCxcbiAgICAgICAgICAgIH0pO1xuXG4gICAgICAgICAgICB3eC5zZXRTdG9yYWdlKHtcbiAgICAgICAgICAgICAgICBrZXk6IFwidXNlckluZm9cIixcbiAgICAgICAgICAgICAgICBkYXRhOiB1c2VySW5mbyxcbiAgICAgICAgICAgIH0pO1xuICAgICAgICB9XG4gICAgfSxcbiAgICBzZXRQZXJtaXNzaW9uKCkge1xuICAgICAgICBsZXQgc3RhdHVzID0gIXRoaXMuZGF0YS5zaGFyZTtcblxuICAgICAgICB3eC5zZXRTdG9yYWdlU3luYygnc2hhcmUnLCBzdGF0dXMpO1xuXG4gICAgICAgIHRoaXMuc2V0RGF0YSh7XG4gICAgICAgICAgICBzaGFyZTogc3RhdHVzXG4gICAgICAgIH0pO1xuICAgIH0sXG4gICAgc3RvcmVMb2NhbEF2YXRhcigpIHtcbiAgICAgICAgLy/lpLTlg49cbiAgICB9LFxuICAgIGdldEhlYWRQaG90bygpIHtcbiAgICAgICAgd3guZ2V0VXNlclByb2ZpbGUoe1xuICAgICAgICAgICAgZGVzYzogJ+WxleekuueUqOaIt+S/oeaBrycsIC8vIOWjsOaYjuiOt+WPlueUqOaIt+S4quS6uuS/oeaBr+WQjueahOeUqOmAlO+8jOWQjue7reS8muWxleekuuWcqOW8ueeql+S4re+8jOivt+iwqOaFjuWhq+WGmVxuICAgICAgICAgICAgc3VjY2VzczogKHJlcykgPT4ge1xuICAgICAgICAgICAgICAgIGlmIChyZXMudXNlckluZm8pIHtcbiAgICAgICAgICAgICAgICAgICAgZ2V0QXBwPElBcHBPcHRpb24+KCkucGFyc2VVc2VySW5mbyhyZXMudXNlckluZm8pXG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgfVxuICAgICAgICB9KVxuICAgIH0sXG4gICAgc2NhblVzZUNhcigpIHtcbiAgICAgICAgd3guZ2V0TG9jYXRpb24oe1xuICAgICAgICAgICAgdHlwZTogJ2djajAyJyxcbiAgICAgICAgICAgIHN1Y2Nlc3M6ICgpID0+IHtcbiAgICAgICAgICAgICAgICBsZXQgdG9rZW5JbmZvID0gSlNPTi5wYXJzZSh3eC5nZXRTdG9yYWdlU3luYygndG9rZW5JbmZvJykpO1xuICAgICAgICAgICAgICAgIHd4LnJlcXVlc3Qoe1xuICAgICAgICAgICAgICAgICAgICB1cmw6ICdodHRwOi8vbG9jYWxob3N0OjkwMDIvdjEvY3JlYXRlL3RyaXAnLFxuICAgICAgICAgICAgICAgICAgICBtZXRob2Q6ICdQT1NUJyxcbiAgICAgICAgICAgICAgICAgICAgaGVhZGVyOiB7QXV0aG9yaXphdGlvbjogJ0JlYXJlciAnICsgdG9rZW5JbmZvLnRva2VufSxcbiAgICAgICAgICAgICAgICAgICAgZGF0YToge1xuICAgICAgICAgICAgICAgICAgICAgICAgY2FydF9pZDogMTAwMSxcbiAgICAgICAgICAgICAgICAgICAgICAgIHN0YXJ0OiB7XG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgbG9uZ2l0dWRlOiAxMTEuMTEsXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgbGF0aXR1ZGU6IDEyMi4yMlxuICAgICAgICAgICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgICAgICAgICB9LFxuICAgICAgICAgICAgICAgICAgICB0aW1lb3V0OiA2MDAwLFxuICAgICAgICAgICAgICAgICAgICBzdWNjZXNzOiAocmVzcG9uc2UpID0+IHtcbiAgICAgICAgICAgICAgICAgICAgICAgIGxldCBkYXRhID0gcmVzcG9uc2UuZGF0YSBhcyBhbnk7XG4gICAgICAgICAgICAgICAgICAgICAgICBjb25zb2xlLmxvZyhkYXRhLnRyaXBfaWQpO1xuICAgICAgICAgICAgICAgICAgICB9LFxuICAgICAgICAgICAgICAgICAgICBmYWlsOiBjb25zb2xlLmxvZ1xuICAgICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfSxcbiAgICAgICAgICAgIGZhaWw6ICgpID0+IHtcbiAgICAgICAgICAgICAgICB3eC5zaG93TW9kYWwoe1xuICAgICAgICAgICAgICAgICAgICB0aXRsZTogJ+ezu+e7n+aPkOekuicsXG4gICAgICAgICAgICAgICAgICAgIGNvbnRlbnQ6ICfmiavnoIHnlKjovablv4XpobvlnLDnkIbkvY3nva7mjojmnYMs6K+35omL5Yqo5byA5ZCv5YWB6K645p2D6ZmQJyxcbiAgICAgICAgICAgICAgICB9KVxuICAgICAgICAgICAgfVxuICAgICAgICB9KTtcbiAgICB9LFxufSk7Il19