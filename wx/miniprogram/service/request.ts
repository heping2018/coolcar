
namespace Coolcar{
    const serverAddr = 'http://localhost:8080'
    const authData = {
        token: "",
        expiryMs: 0
    }

    export interface AuthOption{
        attachAuthHeader: boolean
    }

    interface RequestOption<REQ,RES> {
        method: "GET" | "POST" | "PUT" | "DELETE"
        path: string
        data: REQ
        respMarshaller: (r: object) => RES
    }

    export async function Login(){
        const wxResp = await wxLogin()
        const resp = await sendRequest({
            method: "POST",
            path: '/v1/auth/login',
            data: 
        })
        authData.token = resp.token!
        authData.expiryMs = Date.now() + resp.expiryMs!

    }

    function wxLogin(): Promise {
        return new Promise((resolve,reject) => {
            wx.login({
                success: resolve,
                fail: reject
            }) 
        })
    }


    function sendRequest<REQ,RES>(o :RequestOption<REQ,RES>, a?: AuthOption):Promise<RES> {
        const authOpt = a || {
            attachAuthHeader: true,
        }
        return new Promise((resolve,reject) => {
            const header: Record<string,any> = {}
            if (authOpt.attachAuthHeader && authData.expiryMs >= Date.now() && authData.token){
                header.authorization = "Bearer"+authData.token
            }
            wx.request({
                url: serverAddr + o.path,
                method: o.method,
                data: o.data,
                header,
                success: res => {
                    resolve(o.respMarshaller((res.data)))
                },
                fail: reject
            })
        })
    }
    

}