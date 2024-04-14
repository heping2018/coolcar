// index.ts


// 获取应用实例
const app = getApp<IAppOption>()
const defaultAvatarUrl = 'https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0'

Component({
  data: {
    motto: 'Hello World from',
    userInfo: {
      avatarUrl: defaultAvatarUrl,
      nickName: '',
    },
    hasUserInfo: false,
    canIUseGetUserProfile: wx.canIUse('getUserProfile'),
    canIUseNicknameComp: wx.canIUse('input.type.nickname'),
  },
  methods: {
    getUserInfo(e: any){
      console.log(e)
      const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo
      app.resolveUserInfo(userInfo)
      this.setData({
        userInfo: e.detail.userInfo,
        hasUserInfo: true,
      })
    },
    updataMotoData(){  
      let count = 0
      let shouldStop = false
      setTimeout(() => {
        shouldStop = true
      }, 10000);
     const update = ()=> {
      count ++
      if(!shouldStop){
      this.setData({
        motto: "updateMoto data:" + count,
        },
        () => {
          update()
        }
      )}
    }
    update()
  },
    // 事件处理函数
    bindViewTap() {
      // 一层一层的栈
      wx.navigateTo({
        url: '../logs/logs',
        
      })
    },
    onLoad() {
      try {
        //this.updataMotoData()
      } catch (error) {
      }
      app.globalData.userInfo?.then(userInfo => {
        this.setData({
               userInfo,
               hasUserInfo: true,
            })
      })
      // if(app.globalData.userInfo){
      //   // 优先于加载完成 app.ts(onLaunch)
      //   this.setData({
      //     userInfo: app.globalData.userInfo,
      //     hasUserInfo: true,
      //   })
      // }else{
      //   // 滞后于加载 app.ts(onLaunch)
      //   wx.getUserInfo({
      //     success: res => {
      //       app.globalData.userInfo = res.userInfo
      //       this.setData({
      //         userInfo: app.globalData.userInfo,
      //         hasUserInfo: true,
      //       })
      //     }  
      //   })
      // }
    
    },
    onChooseAvatar(e: any) {
      const { avatarUrl } = e.detail
      const { nickName } = this.data.userInfo
      this.setData({
        "userInfo.avatarUrl": avatarUrl,
        hasUserInfo: nickName && avatarUrl && avatarUrl !== defaultAvatarUrl,
      })
    },
    onInputChange(e: any) {
      const nickName = e.detail.value
      const { avatarUrl } = this.data.userInfo
      this.setData({
        "userInfo.nickName": nickName,
        hasUserInfo: nickName && avatarUrl && avatarUrl !== defaultAvatarUrl,
      })
    },
    getUserProfile() {
      // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认，开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
      wx.getUserProfile({
        desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
        success: (res) => {
          console.log(res)
          this.setData({
            userInfo: res.userInfo,
            hasUserInfo: true
          })
        }
      })
    },
  },
})
