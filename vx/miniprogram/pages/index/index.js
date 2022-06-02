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
const route_1 = require("../../utils/route");
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
        canIUseOpenData: false,
    },
    bindViewTap() {
    },
    onLoad() {
        return __awaiter(this, void 0, void 0, function* () {
            const userInfo = yield getApp().globalData.userInfo;
            if (userInfo) {
                this.setData({
                    avatar: userInfo.avatarUrl
                });
            }
        });
    },
    confirmLocation() {
        wx.getLocation({
            type: 'gcj02',
            success: (res) => {
                console.log('success', res);
                this.setData({
                    longitude: res.longitude,
                    latitude: res.latitude
                });
            },
            fail: console.log
        });
    },
    getCarByScan() {
        wx.navigateTo({
            url: route_1.Route.register(2)
        });
    },
    goHome() {
        wx.navigateTo({ url: "/pages/home/home" });
    },
    getUserProfile() {
        wx.getUserProfile({
            desc: '展示用户信息',
            success: (res) => {
                if (res.userInfo) {
                    getApp().parseUserInfo(res.userInfo);
                }
            }
        });
    }
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiaW5kZXguanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJpbmRleC50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7Ozs7Ozs7OztBQUlBLDZDQUF3QztBQUV4QyxJQUFJLENBQUM7SUFDRCxJQUFJLEVBQUU7UUFDRixLQUFLLEVBQUUsYUFBYTtRQUVwQixNQUFNLEVBQUUsMEJBQTBCO1FBQ2xDLFFBQVEsRUFBRSxFQUFFO1FBQ1osU0FBUyxFQUFFLEdBQUc7UUFDZCxLQUFLLEVBQUUsRUFBRTtRQUNULFFBQVEsRUFBRSxFQUFFO1FBQ1osV0FBVyxFQUFFLEtBQUs7UUFDbEIsT0FBTyxFQUFFLEVBQUUsQ0FBQyxPQUFPLENBQUMsK0JBQStCLENBQUM7UUFDcEQscUJBQXFCLEVBQUUsS0FBSztRQUU1QixlQUFlLEVBQUUsS0FBSztLQUN6QjtJQUVELFdBQVc7SUFFWCxDQUFDO0lBQ0ssTUFBTTs7WUFDUixNQUFNLFFBQVEsR0FBRyxNQUFNLE1BQU0sRUFBYyxDQUFDLFVBQVUsQ0FBQyxRQUFRLENBQUM7WUFDaEUsSUFBSSxRQUFRLEVBQUU7Z0JBQ1YsSUFBSSxDQUFDLE9BQU8sQ0FBQztvQkFDVCxNQUFNLEVBQUUsUUFBUSxDQUFDLFNBQVM7aUJBQzdCLENBQUMsQ0FBQTthQUNMO1FBQ0wsQ0FBQztLQUFBO0lBQ0QsZUFBZTtRQUNYLEVBQUUsQ0FBQyxXQUFXLENBQUM7WUFDWCxJQUFJLEVBQUUsT0FBTztZQUNiLE9BQU8sRUFBRSxDQUFDLEdBQUcsRUFBRSxFQUFFO2dCQUNiLE9BQU8sQ0FBQyxHQUFHLENBQUMsU0FBUyxFQUFFLEdBQUcsQ0FBQyxDQUFDO2dCQUM1QixJQUFJLENBQUMsT0FBTyxDQUFDO29CQUNULFNBQVMsRUFBRSxHQUFHLENBQUMsU0FBUztvQkFDeEIsUUFBUSxFQUFFLEdBQUcsQ0FBQyxRQUFRO2lCQUN6QixDQUFDLENBQUE7WUFDTixDQUFDO1lBQ0QsSUFBSSxFQUFFLE9BQU8sQ0FBQyxHQUFHO1NBQ3BCLENBQUMsQ0FBQztJQUNQLENBQUM7SUFDRCxZQUFZO1FBQ1IsRUFBRSxDQUFDLFVBQVUsQ0FBQztZQUVWLEdBQUcsRUFBRSxhQUFLLENBQUMsUUFBUSxDQUFDLENBQUMsQ0FBQztTQUN6QixDQUFDLENBQUM7SUFZUCxDQUFDO0lBQ0QsTUFBTTtRQUNGLEVBQUUsQ0FBQyxVQUFVLENBQUMsRUFBQyxHQUFHLEVBQUUsa0JBQWtCLEVBQUMsQ0FBQyxDQUFDO0lBQzdDLENBQUM7SUFXRCxjQUFjO1FBQ1YsRUFBRSxDQUFDLGNBQWMsQ0FBQztZQUNkLElBQUksRUFBRSxRQUFRO1lBQ2QsT0FBTyxFQUFFLENBQUMsR0FBRyxFQUFFLEVBQUU7Z0JBQ2IsSUFBSSxHQUFHLENBQUMsUUFBUSxFQUFFO29CQUNkLE1BQU0sRUFBYyxDQUFDLGFBQWEsQ0FBQyxHQUFHLENBQUMsUUFBUSxDQUFDLENBQUE7aUJBQ25EO1lBQ0wsQ0FBQztTQUNKLENBQUMsQ0FBQTtJQUNOLENBQUM7Q0FDSixDQUFDLENBQUEiLCJzb3VyY2VzQ29udGVudCI6WyIvLyBpbmRleC50c1xuLy8g6I635Y+W5bqU55So5a6e5L6LXG5cblxuaW1wb3J0IHtSb3V0ZX0gZnJvbSBcIi4uLy4uL3V0aWxzL3JvdXRlXCI7XG5cblBhZ2Uoe1xuICAgIGRhdGE6IHtcbiAgICAgICAgbW90dG86ICdIZWxsbyBXb3JsZCcsXG5cbiAgICAgICAgYXZhdGFyOiAnL3Jlc291cmNlL2ltYWdlL2hvbWUucG5nJyxcbiAgICAgICAgbGF0aXR1ZGU6IDMxLFxuICAgICAgICBsb25naXR1ZGU6IDEyMSxcbiAgICAgICAgc2NhbGU6IDEwLFxuICAgICAgICB1c2VySW5mbzoge30sXG4gICAgICAgIGhhc1VzZXJJbmZvOiBmYWxzZSxcbiAgICAgICAgY2FuSVVzZTogd3guY2FuSVVzZSgnYnV0dG9uLm9wZW4tM3R5cGUuZ2V0VXNlckluZm8nKSxcbiAgICAgICAgY2FuSVVzZUdldFVzZXJQcm9maWxlOiBmYWxzZSxcbiAgICAgICAgLy8gY2FuSVVzZU9wZW5EYXRhOiB3eC5jYW5JVXNlKCdvcGVuLWRhdGEudHlwZS51c2VyQXZhdGFyVXJsJykgJiYgd3guY2FuSVVzZSgnb3Blbi1kYXRhLnR5cGUudXNlck5pY2tOYW1lJykgLy8g5aaC6ZyA5bCd6K+V6I635Y+W55So5oi35L+h5oGv5Y+v5pS55Li6ZmFsc2VcbiAgICAgICAgY2FuSVVzZU9wZW5EYXRhOiBmYWxzZSxcbiAgICB9LFxuICAgIC8vIOS6i+S7tuWkhOeQhuWHveaVsFxuICAgIGJpbmRWaWV3VGFwKCkge1xuXG4gICAgfSxcbiAgICBhc3luYyBvbkxvYWQoKSB7XG4gICAgICAgIGNvbnN0IHVzZXJJbmZvID0gYXdhaXQgZ2V0QXBwPElBcHBPcHRpb24+KCkuZ2xvYmFsRGF0YS51c2VySW5mbztcbiAgICAgICAgaWYgKHVzZXJJbmZvKSB7XG4gICAgICAgICAgICB0aGlzLnNldERhdGEoe1xuICAgICAgICAgICAgICAgIGF2YXRhcjogdXNlckluZm8uYXZhdGFyVXJsXG4gICAgICAgICAgICB9KVxuICAgICAgICB9XG4gICAgfSxcbiAgICBjb25maXJtTG9jYXRpb24oKSB7XG4gICAgICAgIHd4LmdldExvY2F0aW9uKHtcbiAgICAgICAgICAgIHR5cGU6ICdnY2owMicsXG4gICAgICAgICAgICBzdWNjZXNzOiAocmVzKSA9PiB7XG4gICAgICAgICAgICAgICAgY29uc29sZS5sb2coJ3N1Y2Nlc3MnLCByZXMpO1xuICAgICAgICAgICAgICAgIHRoaXMuc2V0RGF0YSh7XG4gICAgICAgICAgICAgICAgICAgIGxvbmdpdHVkZTogcmVzLmxvbmdpdHVkZSxcbiAgICAgICAgICAgICAgICAgICAgbGF0aXR1ZGU6IHJlcy5sYXRpdHVkZVxuICAgICAgICAgICAgICAgIH0pXG4gICAgICAgICAgICB9LFxuICAgICAgICAgICAgZmFpbDogY29uc29sZS5sb2dcbiAgICAgICAgfSk7XG4gICAgfSxcbiAgICBnZXRDYXJCeVNjYW4oKSB7XG4gICAgICAgIHd4Lm5hdmlnYXRlVG8oe1xuICAgICAgICAgICAgLy8gdXJsOiBcIi9wYWdlcy9yZWdpc3Rlci9yZWdpc3RlclwiLFxuICAgICAgICAgICAgdXJsOiBSb3V0ZS5yZWdpc3RlcigyKVxuICAgICAgICB9KTtcblxuICAgICAgICAvLyB3eC5zY2FuQ29kZSh7XG4gICAgICAgIC8vICAgICBzdWNjZXNzOiAocmVzKSA9PiB7XG4gICAgICAgIC8vICAgICAgICAgY29uc29sZS5sb2cocmVzKTtcbiAgICAgICAgLy8gICAgICAgICB3eC5uYXZpZ2F0ZVRvKHtcbiAgICAgICAgLy8gICAgICAgICAgICAgLy8gdXJsOiBcIi9wYWdlcy9yZWdpc3Rlci9yZWdpc3RlclwiLFxuICAgICAgICAvLyAgICAgICAgICAgICB1cmw6IFJvdXRlLnJlZ2lzdGVyKDIpXG4gICAgICAgIC8vICAgICAgICAgfSk7XG4gICAgICAgIC8vICAgICB9LFxuICAgICAgICAvLyAgICAgZmFpbDogY29uc29sZS5lcnJvclxuICAgICAgICAvLyB9KTtcbiAgICB9LFxuICAgIGdvSG9tZSgpIHtcbiAgICAgICAgd3gubmF2aWdhdGVUbyh7dXJsOiBcIi9wYWdlcy9ob21lL2hvbWVcIn0pO1xuICAgIH0sXG4gICAgLy8gYXN5bmMgb25Mb2FkKCkge1xuICAgIC8vICAgICBjb25zdCB1c2VySW5mbyA9IGF3YWl0IGdldEFwcDxJQXBwT3B0aW9uPigpLmdsb2JhbERhdGEudXNlckluZm87XG4gICAgLy8gICAgIGlmICh1c2VySW5mbykge1xuICAgIC8vICAgICAgICAgdGhpcy5zZXREYXRhKHtcbiAgICAvLyAgICAgICAgICAgICBoYXNVc2VySW5mbzp0cnVlLFxuICAgIC8vICAgICAgICAgICAgIGF2YXRhclVybDogdXNlckluZm8uYXZhdGFyVXJsLFxuICAgIC8vICAgICAgICAgICAgIG5pY2tOYW1lOiB1c2VySW5mby5uaWNrTmFtZVxuICAgIC8vICAgICAgICAgfSlcbiAgICAvLyAgICAgfVxuICAgIC8vIH0sXG4gICAgZ2V0VXNlclByb2ZpbGUoKSB7XG4gICAgICAgIHd4LmdldFVzZXJQcm9maWxlKHtcbiAgICAgICAgICAgIGRlc2M6ICflsZXnpLrnlKjmiLfkv6Hmga8nLCAvLyDlo7DmmI7ojrflj5bnlKjmiLfkuKrkurrkv6Hmga/lkI7nmoTnlKjpgJTvvIzlkI7nu63kvJrlsZXnpLrlnKjlvLnnqpfkuK3vvIzor7fosKjmhY7loavlhplcbiAgICAgICAgICAgIHN1Y2Nlc3M6IChyZXMpID0+IHtcbiAgICAgICAgICAgICAgICBpZiAocmVzLnVzZXJJbmZvKSB7XG4gICAgICAgICAgICAgICAgICAgIGdldEFwcDxJQXBwT3B0aW9uPigpLnBhcnNlVXNlckluZm8ocmVzLnVzZXJJbmZvKVxuICAgICAgICAgICAgICAgIH1cbiAgICAgICAgICAgIH1cbiAgICAgICAgfSlcbiAgICB9XG59KVxuIl19