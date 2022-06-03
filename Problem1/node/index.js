const express = require('express');

const app = express()
app.use(express.json())

app.listen(8080, () => {
    console.log('Server is running on port 8080');
});

app.use((req, res, next) => {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Request-Method', '*');
    res.header('Access-Control-Allow-Methods', 'OPTIONS, GET, POST');
    res.header('Access-Control-Allow-Headers', '*');

    if (req.method === 'OPTIONS') {
        res.sendStatus(200);
    } else {
        next();
    }
});

app.post('/', (req, res) => {
    const pattern = req.body.pattern;
    console.log(pattern);

    // 以下に処理を記述し、res.writeに出力内容を渡してください。
    // ===============処理記述部分==================
    // やっぱりどう頑張ってもpatternからjsonデータを取得できない
    var jsonData = '{"obj": [{"num": "4", "text": "fizz"}, {"num": "7", "text": "buzz"}, {"num": "8", "text": "hoge"}, {"num": "15", "text": "huga"}]}';
    var jsonObject = JSON.parse(jsonData);

    var result = "";
    var ans = "";
    for (var num = 1; num < 31; num++) {

        jsonObject.obj.forEach(function(row) {
            if (num % Number(row.num) == 0) {
                result += row.text;
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

    // ===========================================
    res.writeHead(200, { 'Content-Type': 'text/html' });
    // 出力結果を以下に渡してください。
    res.write(ans);
    res.end();
});