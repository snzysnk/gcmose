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
    onLoad() {
        var _a, _b, _c;
        return __awaiter(this, void 0, void 0, function* () {
            let response = yield request_1.http("profile/data", "POST", {});
            if (response.profile.status == 3) {
                wx.redirectTo({
                    url: "/pages/unlock/unlock",
                });
            }
            else {
                this.setData({
                    path: (_a = response.profile.path) !== null && _a !== void 0 ? _a : '/resource/image/register/default.jpg',
                    userName: (_b = response.profile.name) !== null && _b !== void 0 ? _b : "",
                    sex: (_c = response.profile.sex) !== null && _c !== void 0 ? _c : 0,
                    date: '1994-01-13',
                    status: '',
                });
            }
        });
    },
    chooseImg() {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield request_1.http("profile/uploadUrl", "GET", {});
            const uploadUrl = response.url;
            wx.chooseImage({
                success: (res) => {
                    const data = wx.getFileSystemManager().readFileSync(res.tempFilePaths[0]);
                    wx.request({
                        url: uploadUrl,
                        method: "PUT",
                        header: { "content-type": "image/png" },
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
                    });
                }
            });
        });
    },
    sexChange(e) {
        this.setData({
            sex: e.detail.value
        });
    },
    DateChange(e) {
        this.setData({
            date: e.detail.value
        });
    },
    onUnload() {
        clearInterval(this.data.timer);
    },
    submit() {
        return __awaiter(this, void 0, void 0, function* () {
            this.setData({ status: 'loading' });
            let name = this.data.userName;
            let sex = this.data.sex;
            let birth = (new Date(this.data.date)).getTime();
            let data = { name, sex, birth };
            yield request_1.http("profile/check", "POST", data);
            let that = this;
            this.data.timer = setInterval(function () {
                that.getProfile();
            }, 1000);
        });
    },
    getProfile() {
        return __awaiter(this, void 0, void 0, function* () {
            let response = yield request_1.http("profile/data", "POST", {});
            if (response.profile.status == 3) {
                wx.redirectTo({
                    url: "/pages/unlock/unlock",
                });
            }
            else {
                this.setData({ "status": "" });
            }
        });
    },
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoicmVnaXN0ZXIuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJyZWdpc3Rlci50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOzs7Ozs7Ozs7OztBQUFBLGlEQUF5QztBQUV6QyxJQUFJLENBQUM7SUFLRCxJQUFJLEVBQUU7UUFDRixLQUFLLEVBQUUsc0NBQXNDO1FBQzdDLEdBQUcsRUFBRSxDQUFDO1FBQ04sUUFBUSxFQUFFLEVBQUU7UUFDWixLQUFLLEVBQUUsQ0FBQztRQUNSLElBQUksRUFBRSxFQUFFO1FBQ1IsTUFBTSxFQUFFLEVBQUU7UUFDVixTQUFTLEVBQUU7WUFDUCxJQUFJO1lBQ0osR0FBRztZQUNILEdBQUc7U0FDTjtRQUNELGFBQWEsRUFBRSxDQUFDO0tBQ25CO0lBQ0ssTUFBTTs7O1lBQ1IsSUFBSSxRQUFRLEdBQVEsTUFBTSxjQUFJLENBQUMsY0FBYyxFQUFFLE1BQU0sRUFBRSxFQUFFLENBQUMsQ0FBQTtZQUMxRCxJQUFJLFFBQVEsQ0FBQyxPQUFPLENBQUMsTUFBTSxJQUFJLENBQUMsRUFBRTtnQkFDOUIsRUFBRSxDQUFDLFVBQVUsQ0FBQztvQkFDVixHQUFHLEVBQUUsc0JBQXNCO2lCQUM5QixDQUFDLENBQUM7YUFDTjtpQkFBTTtnQkFDSCxJQUFJLENBQUMsT0FBTyxDQUFDO29CQUNULElBQUksRUFBRSxNQUFBLFFBQVEsQ0FBQyxPQUFPLENBQUMsSUFBSSxtQ0FBSSxzQ0FBc0M7b0JBQ3JFLFFBQVEsRUFBRSxNQUFBLFFBQVEsQ0FBQyxPQUFPLENBQUMsSUFBSSxtQ0FBSSxFQUFFO29CQUNyQyxHQUFHLEVBQUUsTUFBQSxRQUFRLENBQUMsT0FBTyxDQUFDLEdBQUcsbUNBQUksQ0FBQztvQkFDOUIsSUFBSSxFQUFFLFlBQVk7b0JBQ2xCLE1BQU0sRUFBRSxFQUFFO2lCQUNiLENBQUMsQ0FBQzthQUNOOztLQUNKO0lBRUssU0FBUzs7WUFDWCxNQUFNLFFBQVEsR0FBUSxNQUFNLGNBQUksQ0FBQyxtQkFBbUIsRUFBRSxLQUFLLEVB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QUFDIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHtodHRwfSBmcm9tIFwiLi4vLi4vdXRpbHMvcmVxdWVzdFwiO1xuXG5QYWdlKHtcblxuICAgIC8qKlxuICAgICAqIOmhtemdoueahOWIneWni+aVsOaNrlxuICAgICAqL1xuICAgIGRhdGE6IHtcbiAgICAgICAgaW1hZ2U6ICcvcmVzb3VyY2UvaW1hZ2UvcmVnaXN0ZXIvZGVmYXVsdC5qcGcnLFxuICAgICAgICBzZXg6IDAsXG4gICAgICAgIHVzZXJOYW1lOiAnJyxcbiAgICAgICAgdGltZXI6IDAsXG4gICAgICAgIGRhdGU6ICcnLFxuICAgICAgICBzdGF0dXM6ICcnLFxuICAgICAgICBzZXhfcmFuZ2U6IFtcbiAgICAgICAgICAgICfmnKrnn6UnLFxuICAgICAgICAgICAgJ+eUtycsXG4gICAgICAgICAgICAn5aWzJyxcbiAgICAgICAgXSxcbiAgICAgICAgdXBsb2FkX3N0YXR1czogMCxcbiAgICB9LFxuICAgIGFzeW5jIG9uTG9hZCgpIHtcbiAgICAgICAgbGV0IHJlc3BvbnNlOiBhbnkgPSBhd2FpdCBodHRwKFwicHJvZmlsZS9kYXRhXCIsIFwiUE9TVFwiLCB7fSlcbiAgICAgICAgaWYgKHJlc3BvbnNlLnByb2ZpbGUuc3RhdHVzID09IDMpIHtcbiAgICAgICAgICAgIHd4LnJlZGlyZWN0VG8oe1xuICAgICAgICAgICAgICAgIHVybDogXCIvcGFnZXMvdW5sb2NrL3VubG9ja1wiLFxuICAgICAgICAgICAgfSk7XG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgICB0aGlzLnNldERhdGEoe1xuICAgICAgICAgICAgICAgIHBhdGg6IHJlc3BvbnNlLnByb2ZpbGUucGF0aCA/PyAnL3Jlc291cmNlL2ltYWdlL3JlZ2lzdGVyL2RlZmF1bHQuanBnJyxcbiAgICAgICAgICAgICAgICB1c2VyTmFtZTogcmVzcG9uc2UucHJvZmlsZS5uYW1lID8/IFwiXCIsXG4gICAgICAgICAgICAgICAgc2V4OiByZXNwb25zZS5wcm9maWxlLnNleCA/PyAwLFxuICAgICAgICAgICAgICAgIGRhdGU6ICcxOTk0LTAxLTEzJyxcbiAgICAgICAgICAgICAgICBzdGF0dXM6ICcnLFxuICAgICAgICAgICAgfSk7XG4gICAgICAgIH1cbiAgICB9XG4gICAgLFxuICAgIGFzeW5jIGNob29zZUltZygpIHtcbiAgICAgICAgY29uc3QgcmVzcG9uc2U6IGFueSA9IGF3YWl0IGh0dHAoXCJwcm9maWxlL3VwbG9hZFVybFwiLCBcIkdFVFwiLCB7fSlcbiAgICAgICAgY29uc3QgdXBsb2FkVXJsOiBzdHJpbmcgPSByZXNwb25zZS51cmxcblxuICAgICAgICB3eC5jaG9vc2VJbWFnZSh7XG4gICAgICAgICAgICBzdWNjZXNzOiAocmVzKSA9PiB7XG4gICAgICAgICAgICAgICAgY29uc3QgZGF0YSA9IHd4LmdldEZpbGVTeXN0ZW1NYW5hZ2VyKCkucmVhZEZpbGVTeW5jKHJlcy50ZW1wRmlsZVBhdGhzWzBdKVxuICAgICAgICAgICAgICAgIHd4LnJlcXVlc3Qoe1xuICAgICAgICAgICAgICAgICAgICB1cmw6IHVwbG9hZFVybCxcbiAgICAgICAgICAgICAgICAgICAgbWV0aG9kOiBcIlBVVFwiLFxuICAgICAgICAgICAgICAgICAgICBoZWFkZXI6IHtcImNvbnRlbnQtdHlwZVwiOiBcImltYWdlL3BuZ1wifSxcbiAgICAgICAgICAgICAgICAgICAgZGF0YTogZGF0YSxcbiAgICAgICAgICAgICAgICAgICAgc3VjY2VzczogKCkgPT4ge1xuICAgICAgICAgICAgICAgICAgICAgICAgdGhpcy5zZXREYXRhKHtcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBpbWFnZTogcmVzLnRlbXBGaWxlc1swXS5wYXRoLFxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIHVzZXJOYW1lOiAneGllcnVpeGlhbmcnLFxuICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNleDogMSxcbiAgICAgICAgICAgICAgICAgICAgICAgICAgICBkYXRlOiAnMTk5NC0wMS0xMycsXG4gICAgICAgICAgICAgICAgICAgICAgICAgICAgc3RhdHVzOiAnY2hlY2tpbmcnLFxuICAgICAgICAgICAgICAgICAgICAgICAgfSk7XG4gICAgICAgICAgICAgICAgICAgIH0sXG4gICAgICAgICAgICAgICAgICAgIGZhaWw6IGNvbnNvbGUubG9nXG4gICAgICAgICAgICAgICAgfSlcbiAgICAgICAgICAgIH1cbiAgICAgICAgfSk7XG4gICAgfSxcbiAgICBzZXhDaGFuZ2UoZTogYW55KSB7XG4gICAgICAgIHRoaXMuc2V0RGF0YSh7XG4gICAgICAgICAgICBzZXg6IGUuZGV0YWlsLnZhbHVlXG4gICAgICAgIH0pO1xuICAgIH0sXG4gICAgRGF0ZUNoYW5nZShlOiBhbnkpIHtcbiAgICAgICAgdGhpcy5zZXREYXRhKHtcbiAgICAgICAgICAgIGRhdGU6IGUuZGV0YWlsLnZhbHVlXG4gICAgICAgIH0pO1xuICAgIH0sXG4gICAgb25VbmxvYWQoKSB7XG4gICAgICAgIGNsZWFySW50ZXJ2YWwodGhpcy5kYXRhLnRpbWVyKVxuICAgIH0sXG4gICAgYXN5bmMgc3VibWl0KCkge1xuICAgICAgICB0aGlzLnNldERhdGEoe3N0YXR1czogJ2xvYWRpbmcnfSk7XG4gICAgICAgIGxldCBuYW1lID0gdGhpcy5kYXRhLnVzZXJOYW1lO1xuICAgICAgICBsZXQgc2V4ID0gdGhpcy5kYXRhLnNleDtcbiAgICAgICAgbGV0IGJpcnRoID0gKG5ldyBEYXRlKHRoaXMuZGF0YS5kYXRlKSkuZ2V0VGltZSgpO1xuICAgICAgICBsZXQgZGF0YSA9IHtuYW1lLCBzZXgsIGJpcnRofVxuICAgICAgICBhd2FpdCBodHRwKFwicHJvZmlsZS9jaGVja1wiLCBcIlBPU1RcIiwgZGF0YSlcbiAgICAgICAgbGV0IHRoYXQgPSB0aGlzXG4gICAgICAgIHRoaXMuZGF0YS50aW1lciA9IHNldEludGVydmFsKGZ1bmN0aW9uICgpIHtcbiAgICAgICAgICAgIHRoYXQuZ2V0UHJvZmlsZSgpXG4gICAgICAgIH0sIDEwMDApXG5cbiAgICAgICAgLy8gcGF0aDogcmVzcG9uc2UucHJvZmlsZS5wYXRoID8/ICcvcmVzb3VyY2UvaW1hZ2UvcmVnaXN0ZXIvZGVmYXVsdC5qcGcnLFxuICAgICAgICAvLyAgICAgdXNlck5hbWU6IHJlc3BvbnNlLnByb2ZpbGUubmFtZSA/PyBcIlwiLFxuICAgICAgICAvLyAgICAgc2V4OiByZXNwb25zZS5wcm9maWxlLnNleCA/PyAwLFxuICAgICAgICAvLyAgICAgZGF0ZTogJzE5OTQtMDEtMTMnLFxuICAgICAgICAvLyAgICAgc3RhdHVzOiAnJyxcblxuICAgICAgICAvLyBkYXRhWydiaXJ0aCddID1cbiAgICAgICAgLy8gc2V0VGltZW91dCgoKSA9PiB7XG4gICAgICAgIC8vICAgICB0aGlzLnNldERhdGEoe1xuICAgICAgICAvLyAgICAgICAgIHN0YXR1czogJ29rJyxcbiAgICAgICAgLy8gICAgIH0pO1xuICAgICAgICAvL1xuICAgICAgICAvLyAgICAgd3gucmVkaXJlY3RUbyh7XG4gICAgICAgIC8vICAgICAgICAgdXJsOiBcIi9wYWdlcy91bmxvY2svdW5sb2NrXCIsXG4gICAgICAgIC8vICAgICB9KTtcbiAgICAgICAgLy8gfSwgMzAwMCk7XG4gICAgfSxcbiAgICBhc3luYyBnZXRQcm9maWxlKCkge1xuICAgICAgICBsZXQgcmVzcG9uc2U6IGFueSA9IGF3YWl0IGh0dHAoXCJwcm9maWxlL2RhdGFcIiwgXCJQT1NUXCIsIHt9KVxuICAgICAgICBpZiAocmVzcG9uc2UucHJvZmlsZS5zdGF0dXMgPT0gMykge1xuICAgICAgICAgICAgd3gucmVkaXJlY3RUbyh7XG4gICAgICAgICAgICAgICAgdXJsOiBcIi9wYWdlcy91bmxvY2svdW5sb2NrXCIsXG4gICAgICAgICAgICB9KTtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgIHRoaXMuc2V0RGF0YSh7XCJzdGF0dXNcIjpcIlwifSlcbiAgICAgICAgfVxuICAgIH0sXG59KTsiXX0=