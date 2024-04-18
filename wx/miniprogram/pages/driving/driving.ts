// pages/driving/driving.ts

import { routing } from "../../utils/routing"

const initialLat = 31.30173583984375
const initialLng = 121.15481580946181
const centPerSec = 0.7

function formatFee(cents: number){
  return (cents / 100).toFixed(2)
}


function formatDuration(sec: Number){
  const padString = (n: number) => {
   return n<10? '0' + n.toFixed(0):n.toFixed(0)
  }
  const h = Math.floor(sec/3600)
  sec -= 3600 *h
  const m = Math.floor(sec / 60)
  sec -= 60*m
  const s = Math.floor(sec)
  return `${padString(h)}:${padString(m)}:${padString(s)}`
}

Page({

  /**
   * 页面的初始数据
   */
  data: {
    timer: undefined as number | undefined,
    location: {
        latitude: initialLat,
        longitude: initialLng,
    },
    scale: 12,
    elapsed: '00:00:00',
    fee: '0.00',
    markers: [
        {
            iconPath: "/resources/car.png",
            id: 0,
            latitude: initialLat,
            longitude: initialLng,
            width: 20,
            height: 20,
        },
    ],
},

  setupLocationUpdator(){
    wx.startLocationUpdate({
       fail: console.error
      }
    ),
    wx.onLocationChange(loc => {
      console.log('location:',loc)
      this.setData({
        location: {
          latitude: loc.latitude,
          longitude: loc.longitude
        }
      })
    })
  },
   setUpTimer(){
    let elapsedSec = 0
    let cents = 0
    this.timer = setInterval(() => {
      elapsedSec++
      cents+=centPerSec
      this.setData({
        elapsed: formatDuration(elapsedSec),
        fee: formatFee(cents),
      })
    },1000)
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt: Record<'trip_id',string>) {
    const o: routing.DrivingOpts = opt
    console.log("current tripId,", o.trip_id )
    this.setupLocationUpdator()
    this.setUpTimer()

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
   wx.stopLocationUpdate()
   if(this.timer){
    clearInterval(this.timer)
   }
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