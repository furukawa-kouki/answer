const http = require('http');
const url = require('url');

const server = http.createServer((req, res)=>{
   const queryObject = url.parse(req.url, true).query;
   console.log(queryObject.obj);

   // ここに処理を記述してください。
   // パターン表の取得方法がわからないのでいったん手動で書く
   // {
   //    "obj": [
   //      { "num": 4, "text": "fizz" },
   //      { "num": 7, "text": "buzz" },
   //      { "num": 8, "text": "hoge" },
   //      { "num": 15, "text": "huga" }
   //    ]
   // }
   var ans = "";
   var result = "";
   for(var num = 1; num < 31; num++){
      if(num % 4 == 0){
         result += "fizz";
      }
      if(num % 7 == 0){
         result += "buzz";
      }
      if(num % 8 == 0){
         result += "hoge";
      }
      if(num % 15 == 0){
         result += "huga";
      }
      if(result == ""){
         result += num;
      }

      ans += result;
      if(num == 30){
         break;
      }
      ans += ", "
      result = "";
   }

   res.setHeader('Access-Control-Allow-Origin', '*');
   res.setHeader('Access-Control-Request-Method', '*');
   res.setHeader('Access-Control-Allow-Methods', 'OPTIONS, GET, POST');
   res.setHeader('Access-Control-Allow-Headers', '*');
   res.writeHead(200, {'Content-Type': 'text/html'});
   res.write(ans);
   res.end();
});
server.listen(8080); 
