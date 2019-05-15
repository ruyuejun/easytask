
## 请求结果说明

所有的请求结果都必定包含三个字段：
```
{
  "status": 1,              # 网络响应状态
  "code": 5001,             # 数据响应结果状态值
  "msg": "数据库错误"        # 数据响应结果状态值对应说明
  "data": [                 # 接口的真实返回结果封装在data可变数组中
    {
        "veryCode":4021
    }
   ]         
}
```

## 账户系统

#### 获取验证码
```
接口地址：/account/getVerifyCode
接口参数：
        tel                 必选，手机号码
返回结果：
        verifyCode          验证码 
接口说明：
```


#### 用户登录
```
接口地址：/account/login
接口参数：
        way                 必选，登录方式，默认0(验证码登录) 1（账户登录） 2（小程序登录） 3第三方登录
        tel                 way为0时必选，手机
        verifyCode          way为0时必选，验证码
        code                way为2时必选，小程序用户code
返回结果：
        uid                 用户id
        appid               way为2时用户唯一标识
        tel                 way为2时解密出来的手机号     
接口说明：登录时如果发现用户未注册，则自动注册           
```

#### 用户注册
```
接口地址：/account/reister
接口参数：
        way                 必选，登录方式，默认0(验证码注册) 2（小程序注册） 3第三方注册
        tel                 way为0时必选，手机
        verifyCode          way为0时必选，验证码
        code                way为2时必选，小程序用户code
返回结果：
        uid                 用户id
        appid               way为2时用户唯一标识
        tel                 way为2时解密出来的手机号     
```

#### 用户退出
```
接口地址：/account/logout
接口参数：
        uid                 必选，用户id
```

#### 刷新Token
```
接口地址：/account/refreshToken
接口参数：
        way                 必选，登录方式，默认0(普通登录刷新token) 2（小程序刷新token）
        uid                 必选，用户id
返回结果：
        token               新token              
```

