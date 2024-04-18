export namespace routing{
    export interface DrivingOpts {
        trip_id:  string
    }
    export function drving(o: DrivingOpts){
        return `/page/driving/driving?trip_id=${o.trip_id}`
    }
    export interface LockOpts {
        car_id: string
    }
    export function lock(o: LockOpts){
        return `/pages/lock/lock?car_id=${o.car_id}`
    }
    export interface registerOpts{
            redirect?: string
    }
    export interface RegsiterParams {
        redirectUrl: string
    } 
    export function register(p?: RegsiterParams){
        const page = '/pages/register/register'
        if(!p){
            return page
        }
        return `${page}?redirect=${encodeURIComponent(p.redirectUrl)}`
    }
    export function mytrips(){
        return '/pages/register/register'
    }
}