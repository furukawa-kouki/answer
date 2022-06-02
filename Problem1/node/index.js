const http = require('http');
const url = require('url');

const server = http.createServer((req, res) => {
    const queryObject = url.parse(req.url, true).query;
    console.log(queryObject.obj);

    // ここに処理を記述してください。
    // パターン表の取得方法がわからないのでいったん手動で書く
    var obj = [
        ["4", "fizz"],
        ["7", "buzz"],
        ["8", "hoge"],
        ["15", "huga"]
    ]



    var result = "";
    var ans = "";
    for (var num = 1; num < 31; num++) {

        obj.forEach(function(row) {
            if (num % Number(row[0]) == 0) {
                result += row[1];
            }
        });

        if (result == "") {
            result = String(num);
        }

        ans += result;
        if (num == 30) {
            break;
        }
        ans += ", "
        result = "";
    }

    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Access-Control-Request-Method', '*');
    res.setHeader('Access-Control-Allow-Methods', 'OPTIONS, GET, POST');
    res.setHeader('Access-Control-Allow-Headers', '*');
    res.writeHead(200, { 'Content-Type': 'text/html' });
    res.write(ans);
    res.end();
});
server.listen(8080);
