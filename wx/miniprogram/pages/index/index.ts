// pages/index/index.ts

Page({

  /**
   * 页面的初始数据
   */
  data: {
    isPageShowing: true,

    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: true,
      subKey: '',
      layerStyle: -1,
      enableZoom: true,
      enableScroll: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },
    location: {
      latitude: 31,
      longitude: 120,
    },
    scale: 16,
    markers: [
      {
        iconPath: "/resources/car.png",
        id: 0,
        latitude: 30.30173583984375,
        longitude: 120.15481580946181,
        width: 30,
        height: 30
      },
      {
        iconPath: "/resources/car.png",
        id: 1,
        latitude: 31.30173583984375,
        longitude: 121.15481580946181,
        width: 30,
        height: 30
      },
  
    ]
  },
  onScanCliked(){
    wx.scanCode({
      success: () => {
        wx.navigateTo({
          url: '/pages/register/register',
        })
      },
      fail : () => {
        console.log(1)
      }
    })
  },
  onMyLocationTap(){
    wx.getLocation({
      type: 'gcj02',
      success: res => {
        this.setData({
            location: {
              latitude: res.latitude,
              longitude: res.longitude,
            }
        })
      },
      fail: ()=>{
        wx.showToast({
          icon: 'none',
          title: "请前往设置页授权"
        })
      }
    })
  },

  movesCars(){
    const map = wx.createMapContext("map")
    const dets = {
      latitude: 31.30173583984375,
      longitude: 121.15481580946181,
    }
    const moveCar = () => {
      dets.latitude += 0.1
      dets.longitude += 0.1
      map.translateMarker({
        destination: {
          latitude : dets.latitude,
          longitude: dets.longitude,
        },
        markerId: 0,
        autoRotate: false,
        rotate: 0,
        duration: 5000,
        animationEnd: () => {
          if(this.isPageShowing){
            moveCar()
          }
        }
      })
    }
    moveCar()
    
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
    this.setData({
      isPageShowing: true
    }) 
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {
    this.setData({
      isPageShowing: false
    }) 
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