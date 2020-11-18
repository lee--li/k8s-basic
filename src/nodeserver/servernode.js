//express_demo.js 文件
var express = require('express');
var app = express();
 
var port = 8087;
if(process.env.PORT) {
    port = process.env.PORT
}

app.get('/', function (req, res) {
   res.send('Hello World nodejs');
})
 
var server = app.listen(port, function () {
 
  var host = server.address().address
  var port = server.address().port
 
  console.log("应用实例，访问地址为 http://%s:%s", host, port)
 
})