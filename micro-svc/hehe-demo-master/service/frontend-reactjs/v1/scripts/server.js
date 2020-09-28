const express = require('express')
const bodyParser = require('body-parser')
const proxy = require('http-proxy-middleware')
const querystring = require('querystring')

const app = express()

// 设置端口
app.set('port', 3000);

const proxyConfig = [
  
    {
      "url":"/data/*",
      "target": "http://forecast:3002",
      "pathRewrite": {'^/data' : ''},
      "changeOrigin": true,
      "secure": false
    },
    {
      "url":"/advertisement/*",
      "target": "http://advertisement:3003",
      "pathRewrite": {'^/advertisement' : ''},
      "changeOrigin": true,
      "secure": false
    },
     {
      "url": "/recommendation/*",
      "target": "http://recommendation:3005",
      "pathRewrite": {'^/recommendation' : ''},
      "changeOrigin": true,
      "secure": false
    }
  
];
// make http proxy middleware setting
const createProxySetting = function (item) {
  return {
    target: item.target,
    pathRewrite: item.pathRewrite,
    secure: false,
    changeOrigin: true,
    headers: {
      Accept: 'application/json',
      'X-Requested-With': 'XMLHttpRequest'
    },
    
    onProxyReq: function (proxyReq, req) {
      if (req.method === 'POST' && req.body) {
        const bodyData = querystring.stringify(req.body)
        proxyReq.write(bodyData)
      }
    }
  }
}

// 静态页面
// 这里一般设置你的静态资源路径
app.use('/', express.static('build'));
app.use('/forecast/:city', express.static('build'));
app.use('/maintenanced', express.static('build'));
app.use('/dashboard', express.static('build'));

// parse application/json
app.use(bodyParser.json())

// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: false }))

// proxy
proxyConfig.forEach(function (item) {
   app.use(item.url, proxy(createProxySetting(item)))
})

// 监听端口
app.listen(app.get('port'), () => {
  console.log(`server running @${app.get('port')}`);
});