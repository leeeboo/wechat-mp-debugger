# wechat-mp-debugger
A debug tool for wechat public account

## Usage

go install
wechat-mp-debugger Url Token FromUserName ToUserName Text

Example:
wechat-mp-debugger http://wxdev.imlibo.com/weixin asdj8234001 osd3Kt0DeyrNde3EuFb0oHs93NeU gh_3v0a11ece332 Hello!

* Url : The url which set in WeChat Admin used to receive WeChat message.
* Token : The token which set in WeChat Admin used to generate a signature for communication between your app and WeChat.
* FromUserName : A fans openid.
* ToUserName : Original ID.
* Text : This tool can only send text message to wechat public account now.
