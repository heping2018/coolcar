// pages/register/register.ts
Page({

  /**
   * 页面的初始数据
   */
  data: {
    licImgURL: '',
    licNo:'',
    name: '',
    gendersIndex: 0,
    genders: ['未知','男','女','其他'],
    birthDate: '1990-01-01',
    state: 'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED'
  },
  onResubmit(){
    this.setData({
      state: 'UNSUBMITTED',
      licImgURL: '',
    })
  },
  onSubmit() {
    //TODO: submit to server
    this.setData({
      state: 'PENDING'
    })
    setTimeout(() => {
        this.onLicVerfied()
    },3000)
  },

  onLicVerfied(){
    this.setData({
      state:  'VERIFIED'
    })
    wx.redirectTo({
      url: '/pages/lock/lock'
    })

  },

  onGenderChange(e: any){
    this.setData({
      gendersIndex: e.detail.value,
    })
  },
  onBirthDateChange(e: any){
    this.setData({
      birthDate: e.detail.value,
    })
  },
  onUpLoalLic(){
    wx.chooseImage({
      success: (res) => {
        if(res.tempFilePaths.length >0){
          this.setData({
            licImgURL: res.tempFilePaths[0]
          })
          setTimeout(() => {
           this.setData({
            licNo: '9812312312',
            name: '张三',
            gendersIndex: 1,
            birthDate: '1990-09-01'
           }) 
          },1000)
        }
      }
    })
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