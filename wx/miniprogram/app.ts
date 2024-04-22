import { getSetting, getUserInfo } from "./utils/wxapi"
let resolveUserInfo: (value: WechatMiniprogram.UserInfo | PromiseLike<WechatMiniprogram.UserInfo>) => void ;
let rejectUserInfo: (reason?: any) => void;
// app.ts
App<IAppOption>({
  globalData: {
    userInfo : new Promise((resolve,reject) => {
      resolveUserInfo  = resolve
      rejectUserInfo = reject
    })
  },
  onLaunch() {
    //登录
    wx.login({
      success: res => {
        console.log(res.code)
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
        wx.request({
          url: 'http://192.168.2.1:8080/v1/auth/login',
          method: 'POST',
          data:  {
            code: res.code,
          },
          fail: console.log
          
        })
      },
    })
    getSetting().then(res => {
      console.log(res)
      if(res.authSetting['scope.userInfo']){
        return getUserInfo()
      }
      return undefined
    }).then(res => {
      if(!res){
        return
      }
      // 赋值给全局变量
      // this.globalData.userInfo = res?.userInfo
      // 通知页面 获取用户信息
      // if(this.userInfoReadyCallback){
      //   this.userInfoReadyCallback(res)
      // }
      //通知页面 获取用户信息  ----then 参数----
      resolveUserInfo(res.userInfo)
    }).catch(rejectUserInfo)
  },
  resolveUserInfo(userInfo:WechatMiniprogram.UserInfo) {
    resolveUserInfo(userInfo)
  }
})