# 使用SCF加API网关实现微信支付商家小票

## 小白客户安装指南

- 开通一个腾讯云账号
- 开通Serverless服务
- 在`https://cloud.tencent.com/online-service?from=console_bar`提单反馈，寻求安装帮助
- 如果你有开发经验，也可以尝试自己部署

## 研发安装指南

- 开通一个腾讯云账号
- 开通Serverless服务
- 修改`main.go`和`serverless.yaml`的内容
- 根目录下先执行`make build`，再执行`sls deploy --force`。如果你的电脑上没有make，可以查看*Makefile*文件的build内容，手动部署。

## 微信支付商家小票开发指引

`https://wx.gtimg.com/pay/download/goldplan/goldplan_developer_guideline.pdf`

## TODO
- 支持图片下载
- 添加md5校验