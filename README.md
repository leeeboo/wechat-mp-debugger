# wechat-mp-debugger
A debug tool for wechat public account

## Usage

go build
wechat-mp-debugger Url Token FromUserName ToUserName Text

* Url : The url which set in WeChat Admin used to receive WeChat message.
* Token : The token which set in WeChat Admin used to generate a signature for communication between your app and WeChat.
* FromUserName : A fans openid.
* ToUserName : WeChat ID.
* Text : This tool can only send text message to wechat public account now.
