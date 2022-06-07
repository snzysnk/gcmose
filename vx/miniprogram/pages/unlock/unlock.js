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
Object.defineProperty(exports, "__esModule", { value: true });
const request_1 = require("../../utils/request");
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
        let _this = this;
        wx.getLocation({
            type: 'gcj02',
            success: () => {
                _this.createTrip();
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
    createTrip() {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield request_1.http("create/trip", "POST", {
                cart_id: 1001,
                start: {
                    longitude: 111.11,
                    latitude: 122.22
                }
            });
            console.log(response);
        });
    }
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoidW5sb2NrLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsidW5sb2NrLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7O0FBQ0EsaURBQXlDO0FBRXpDLElBQUksQ0FBQztJQUtELElBQUksRUFBRTtRQUNGLFVBQVUsRUFBRSxnQ0FBZ0M7UUFDNUMsS0FBSyxFQUFFLEtBQUs7S0FDZjtJQUtLLE1BQU07O1lBQ1IsTUFBTSxRQUFRLEdBQUcsTUFBTSxNQUFNLEVBQWMsQ0FBQyxVQUFVLENBQUMsUUFBUSxDQUFDO1lBQ2hFLElBQUksUUFBUSxFQUFFO2dCQUNWLElBQUksQ0FBQyxPQUFPLENBQUM7b0JBQ1QsVUFBVSxFQUFFLFFBQVEsQ0FBQyxTQUFTO2lCQUNqQyxDQUFDLENBQUM7Z0JBRUgsRUFBRSxDQUFDLFVBQVUsQ0FBQztvQkFDVixHQUFHLEVBQUUsVUFBVTtvQkFDZixJQUFJLEVBQUUsUUFBUTtpQkFDakIsQ0FBQyxDQUFDO2FBQ047UUFDTCxDQUFDO0tBQUE7SUFDRCxhQUFhO1FBQ1QsSUFBSSxNQUFNLEdBQUcsQ0FBQyxJQUFJLENBQUMsSUFBSSxDQUFDLEtBQUssQ0FBQztRQUU5QixFQUFFLENBQUMsY0FBYyxDQUFDLE9BQU8sRUFBRSxNQUFNLENBQUMsQ0FBQztRQUVuQyxJQUFJLENBQUMsT0FBTyxDQUFDO1lBQ1QsS0FBSyxFQUFFLE1BQU07U0FDaEIsQ0FBQyxDQUFDO0lBQ1AsQ0FBQztJQUNELGdCQUFnQjtJQUVoQixDQUFDO0lBQ0QsWUFBWTtRQUNSLEVBQUUsQ0FBQyxjQUFjLENBQUM7WUFDZCxJQUFJLEVBQUUsUUFBUTtZQUNkLE9BQU8sRUFBRSxDQUFDLEdBQUcsRUFBRSxFQUFFO2dCQUNiLElBQUksR0FBRyxDQUFDLFFBQVEsRUFBRTtvQkFDZCxNQUFNLEVBQWMsQ0FBQyxhQUFhLENBQUMsR0FBRyxDQUFDLFFBQVEsQ0FBQyxDQUFBO2lCQUNuRDtZQUNMLENBQUM7U0FDSixDQUFDLENBQUE7SUFDTixDQUFDO0lBQ0QsVUFBVTtRQUNOLElBQUksS0FBSyxHQUFHLElBQUksQ0FBQTtRQUNoQixFQUFFLENBQUMsV0FBVyxDQUFDO1lBQ1gsSUFBSSxFQUFFLE9BQU87WUFDYixPQUFPLEVBQUUsR0FBRyxFQUFFO2dCQUNULEtBQUssQ0FBQyxVQUFVLEVBQUUsQ0FBQTtnQkFDbkIsRUFBRSxDQUFDLFVBQVUsQ0FBQztvQkFDVixHQUFHLEVBQUUsc0JBQXNCO2lCQUM5QixDQUFDLENBQUM7WUFDUCxDQUFDO1lBQ0QsSUFBSSxFQUFFLEdBQUcsRUFBRTtnQkFDUCxFQUFFLENBQUMsU0FBUyxDQUFDO29CQUNULEtBQUssRUFBRSxNQUFNO29CQUNiLE9BQU8sRUFBRSx3QkFBd0I7aUJBQ3BDLENBQUMsQ0FBQTtZQUNOLENBQUM7U0FDSixDQUFDLENBQUM7SUFDUCxDQUFDO0lBQ0ssVUFBVTs7WUFDWixNQUFNLFFBQVEsR0FBRyxNQUFNLGNBQUksQ0FBQyxhQUFhLEVBQUUsTUFBTSxFQUFFO2dCQUMvQyxPQUFPLEVBQUUsSUFBSTtnQkFDYixLQUFLLEVBQUU7b0JBQ0gsU0FBUyxFQUFFLE1BQU07b0JBQ2pCLFFBQVEsRUFBRSxNQUFNO2lCQUNuQjthQUNKLENBQUMsQ0FBQztZQUNILE9BQU8sQ0FBQyxHQUFHLENBQUMsUUFBUSxDQUFDLENBQUM7UUFDMUIsQ0FBQztLQUFBO0NBQ0osQ0FBQyxDQUFDIiwic291cmNlc0NvbnRlbnQiOlsiLy8gbWluaXByb2dyYW0vcGFnZXMvdW5sb2NrL3VubG9jay5qc1xuaW1wb3J0IHtodHRwfSBmcm9tIFwiLi4vLi4vdXRpbHMvcmVxdWVzdFwiO1xuXG5QYWdlKHtcblxuICAgIC8qKlxuICAgICAqIOmhtemdoueahOWIneWni+aVsOaNrlxuICAgICAqL1xuICAgIGRhdGE6IHtcbiAgICAgICAgaGVhZF9waG90bzogJy9yZXNvdXJjZS9pbWFnZS9oZWFkX3Bob3RvLnBuZycsXG4gICAgICAgIHNoYXJlOiBmYWxzZSxcbiAgICB9LFxuXG4gICAgLyoqXG4gICAgICog55Sf5ZG95ZGo5pyf5Ye95pWwLS3nm5HlkKzpobXpnaLliqDovb1cbiAgICAgKi9cbiAgICBhc3luYyBvbkxvYWQoKSB7XG4gICAgICAgIGNvbnN0IHVzZXJJbmZvID0gYXdhaXQgZ2V0QXBwPElBcHBPcHRpb24+KCkuZ2xvYmFsRGF0YS51c2VySW5mbztcbiAgICAgICAgaWYgKHVzZXJJbmZvKSB7XG4gICAgICAgICAgICB0aGlzLnNldERhdGEoe1xuICAgICAgICAgICAgICAgIGhlYWRfcGhvdG86IHVzZXJJbmZvLmF2YXRhclVybCxcbiAgICAgICAgICAgIH0pO1xuXG4gICAgICAgICAgICB3eC5zZXRTdG9yYWdlKHtcbiAgICAgICAgICAgICAgICBrZXk6IFwidXNlckluZm9cIixcbiAgICAgICAgICAgICAgICBkYXRhOiB1c2VySW5mbyxcbiAgICAgICAgICAgIH0pO1xuICAgICAgICB9XG4gICAgfSxcbiAgICBzZXRQZXJtaXNzaW9uKCkge1xuICAgICAgICBsZXQgc3RhdHVzID0gIXRoaXMuZGF0YS5zaGFyZTtcblxuICAgICAgICB3eC5zZXRTdG9yYWdlU3luYygnc2hhcmUnLCBzdGF0dXMpO1xuXG4gICAgICAgIHRoaXMuc2V0RGF0YSh7XG4gICAgICAgICAgICBzaGFyZTogc3RhdHVzXG4gICAgICAgIH0pO1xuICAgIH0sXG4gICAgc3RvcmVMb2NhbEF2YXRhcigpIHtcbiAgICAgICAgLy/lpLTlg49cbiAgICB9LFxuICAgIGdldEhlYWRQaG90bygpIHtcbiAgICAgICAgd3guZ2V0VXNlclByb2ZpbGUoe1xuICAgICAgICAgICAgZGVzYzogJ+WxleekuueUqOaIt+S/oeaBrycsIC8vIOWjsOaYjuiOt+WPlueUqOaIt+S4quS6uuS/oeaBr+WQjueahOeUqOmAlO+8jOWQjue7reS8muWxleekuuWcqOW8ueeql+S4re+8jOivt+iwqOaFjuWhq+WGmVxuICAgICAgICAgICAgc3VjY2VzczogKHJlcykgPT4ge1xuICAgICAgICAgICAgICAgIGlmIChyZXMudXNlckluZm8pIHtcbiAgICAgICAgICAgICAgICAgICAgZ2V0QXBwPElBcHBPcHRpb24+KCkucGFyc2VVc2VySW5mbyhyZXMudXNlckluZm8pXG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgfVxuICAgICAgICB9KVxuICAgIH0sXG4gICAgc2NhblVzZUNhcigpIHtcbiAgICAgICAgbGV0IF90aGlzID0gdGhpc1xuICAgICAgICB3eC5nZXRMb2NhdGlvbih7XG4gICAgICAgICAgICB0eXBlOiAnZ2NqMDInLFxuICAgICAgICAgICAgc3VjY2VzczogKCkgPT4ge1xuICAgICAgICAgICAgICAgICBfdGhpcy5jcmVhdGVUcmlwKClcbiAgICAgICAgICAgICAgICB3eC5yZWRpcmVjdFRvKHtcbiAgICAgICAgICAgICAgICAgICAgdXJsOiAnL3BhZ2VzL3RyYXZlbC90cmF2ZWwnLFxuICAgICAgICAgICAgICAgIH0pO1xuICAgICAgICAgICAgfSxcbiAgICAgICAgICAgIGZhaWw6ICgpID0+IHtcbiAgICAgICAgICAgICAgICB3eC5zaG93TW9kYWwoe1xuICAgICAgICAgICAgICAgICAgICB0aXRsZTogJ+ezu+e7n+aPkOekuicsXG4gICAgICAgICAgICAgICAgICAgIGNvbnRlbnQ6ICfmiavnoIHnlKjovablv4XpobvlnLDnkIbkvY3nva7mjojmnYMs6K+35omL5Yqo5byA5ZCv5YWB6K645p2D6ZmQJyxcbiAgICAgICAgICAgICAgICB9KVxuICAgICAgICAgICAgfVxuICAgICAgICB9KTtcbiAgICB9LFxuICAgIGFzeW5jIGNyZWF0ZVRyaXAoKSB7XG4gICAgICAgIGNvbnN0IHJlc3BvbnNlID0gYXdhaXQgaHR0cChcImNyZWF0ZS90cmlwXCIsIFwiUE9TVFwiLCB7XG4gICAgICAgICAgICBjYXJ0X2lkOiAxMDAxLFxuICAgICAgICAgICAgc3RhcnQ6IHtcbiAgICAgICAgICAgICAgICBsb25naXR1ZGU6IDExMS4xMSxcbiAgICAgICAgICAgICAgICBsYXRpdHVkZTogMTIyLjIyXG4gICAgICAgICAgICB9XG4gICAgICAgIH0pO1xuICAgICAgICBjb25zb2xlLmxvZyhyZXNwb25zZSk7XG4gICAgfVxufSk7Il19