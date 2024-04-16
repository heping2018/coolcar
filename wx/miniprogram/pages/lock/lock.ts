// pages/lock/lock.ts
import {getUserInfo, getSetting} from "../../utils/util"

const shareLocalkey = "share_location"
Page({

  /**
   * 页面的初始数据
   */
  data: {
    avatarURL: '',
    shareLocal: false
  },
  onShareLocation(e: any){
    const shareLocation: boolean =e.detail.value
    wx.setStorageSync(shareLocalkey,shareLocation)
  },
  onGetAvatar(e: any){
    getSetting().then(res => {
      if(res.authSetting['scope.userInfo']){
        return getUserInfo()
      }
      return undefined
    }).then(res => {
      if(!res){
        return
      }
      const userInfo: WechatMiniprogram.UserInfo = res.userInfo
      userInfo.avatarUrl = e.detail.avatarUrl
      this.setData({
        avatarURL: userInfo.avatarUrl
      })
      getApp<IAppOption>().resolveUserInfo(userInfo)
    })
  },
  
  /**
   * 生命周期函数--监听页面加载
   */
  async onLoad() {
    this.setData({
      shareLocal:  wx.getStorageSync(shareLocalkey) || false
    })
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  }
})