import { routing } from "../../utils/routing"

// pages/mytrips/mytrips.ts
Page({

  /**
   * 页面的初始数据
   */
  data: {
    promotionItems: [
      {
          img: 'https://img.mukewang.com/5f7301d80001fdee18720764.jpg',
          promotionID: 1,
      },            
      {
          img: 'https://img.mukewang.com/5f6805710001326c18720764.jpg',
          promotionID: 2,
      },
      {
          img: 'https://img.mukewang.com/5f6173b400013d4718720764.jpg',
          promotionID: 3,
      },
      {
          img: 'https://img.mukewang.com/5f7141ad0001b36418720764.jpg',
          promotionID: 4,
      },
  ],

  avatarURL: '',
  tripsHeight: 0,
  navCount: 0,
  mainScroll: '',
  navSel: '',
  navScroll: '',
  },

  onRegisiterTap(){
    wx.navigateTo({
      url: routing.mytrips(),
    })
  },

  onPromotionItemTap(e: any) {
    const promotionID:number = e.currentTarget.dataset.promotionId
    if (promotionID) {
        console.log('claiming promotion', promotionID)
    }
 },
  onGetUserInfo(e: any){
    const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo
    if(userInfo){
      getApp<IAppOption>().resolveUserInfo(userInfo)
      this.setData({
        avatarURL: userInfo.avatarUrl
      })
    }
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {

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