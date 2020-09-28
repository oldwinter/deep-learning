const Koa = require('koa');
const Router = require('koa-router');
const axios = require('axios')
const app = new Koa();
const router = new Router();

const HANGZHOU={
    "city": {
        "name": "Hangzhou",
    },
    "cnt": 3,
    "list": [
        {
            "dt": "today", 
            "temp": {
                "day": 28.89,
                "min": 21.16,
                "max": 31.94,
            },
            "weather": [
                {
                    "main": "Rain",
                    "description": "light rain",
                    "icon": "10d"
                }
            ],
            recommendation:""
        },
        {
            "dt": "tomorrow",
            "temp": {
                "day": 30.5,
                "min": 29.73,
                "max": 40.5,
            },
            "weather": [
                {
                    "main": "Clear",
                    "description": "sky is clear",
                    "icon": "01d"
                }
            ],
            recommendation:""
        },
        {
            "dt": "the day after tomorrow",
            "temp": {
                "day": 32.72,
                "min": 27.03,
                "max": 39.72,
            },
            "weather": [
                {
                    "main": "Rain",
                    "description": "light rain",
                    "icon": "10d"
                }
            ],
            recommendation:""
        }
    ]
}
const BEIJING={
    "city": {
        "name": "Beijing",
    },
    "cnt": 3,
    "list": [
        {
            "dt": "today",
            "temp": {
                "day": 32.94,
                "min": 24.48,
                "max": 34.63
            },
            "weather": [
                {
                    "main": "Clear",
                    "description": "sky is clear",
                    "icon": "01d"
                }
            ],
            recommendation:""
        },
        {
            "dt": "tomorrow",
            "temp": {
                "day": 28.67,
                "min": 24.37,
                "max": 35.62,
            },
            "weather": [
                {
                    "id": 601,
                    "main": "Snow",
                    "description": "snow",
                    "icon": "13d"
                }
            ],
            recommendation:""
        },
        {
            "dt": "the day after tomorrow",
            "temp": {
                "day": 18.72,
                "min": 17.03,
                "max": 29.72
            },
            "weather": [
                {
                    "main": "Rain",
                    "description": "light rain",
                    "icon": "10d"
                }
            ],
            recommendation:""
        }
    ]
}
const SHANGHAI={
    "city": {
        "name": "Shanghai",
    },
    "cnt": 3,
    "list": [
        {
            "dt": "today",
            "temp": {
                "day": 27.94,
                "min": 24.48,
                "max": 34.63,
            },
            "weather": [
                {
                    "main": "Clear",
                    "description": "sky is clear",
                    "icon": "01d"
                }
            ],
            recommendation:""
        },
        {
            "dt": "tomorrow",
            "temp": {
                "day": 33.89,
                "min": 32.16,
                "max": 35.94
            },
            "pressure": 1021.47,
            "humidity": 93,
            "weather": [
                {
                    "main": "Rain",
                    "description": "light rain",
                    "icon": "10d"
                }
            ],
            recommendation:""
        },
        {
            "dt": "the day after tomorrow",
            "temp": {
                "day": 34.5,
                "min": 29.73,
                "max": 40.5
            },
            "weather": [
                {
                    "main": "Clear",
                    "description": "sky is clear",
                    "icon": "01d"
                }
            ],
            recommendation:""
        }
    ]
}

function getForwardHeaders(request) {
    headers = {}
    incoming_headers = [
        'x-request-id',
        'x-b3-traceid',
        'x-b3-spanid',
        'x-b3-parentspanid',
        'x-b3-sampled',
        'x-b3-flags',
        'x-ot-span-context'
    ]

    for (idx in incoming_headers) {
        ihdr = incoming_headers[idx]
        val = request.headers[ihdr]
        if (val !== undefined && val !== '') {
            headers[ihdr] = val
            console.log("incoming: " + ihdr + ":" + val)
        }
    }
    //如果要对调用链上所有组件做基于cookie的基于内容请求的灰度发布，需要组件自己转发cookie。
    if(request.headers["cookie"]){
        headers["cookie"] = request.headers["cookie"]
    }
    return headers
}



router.get('/status', async (ctx, next) => {
    ctx.body = 'ok';
})

router.get('/weather', async (ctx, next) => {
    let locate = ctx.query.locate;
    forwardHeaders = getForwardHeaders(ctx.request)
    //本地调试用
    // let recommendation_java_url = 'http://' + 'localhost' + ':3005'  + '/activity'
    let recommendation_java_url = 'http://' + 'recommendation' + ':3005'  + '/activity'
    //获取推荐数据
    let data = "";
    if(locate =="Hangzhou"){
        data = HANGZHOU;
    }else if (locate == "Beijing"){
        data =  BEIJING;
    }else if (locate == "Shanghai"){
        data =  SHANGHAI;
    }else{
        data = HANGZHOU;
    }
    let promises = []
    for(let i=0;i<data.cnt;i++){
        promises.push(axios.get(recommendation_java_url, {
            headers: forwardHeaders,
            params: {
                temp: Math.round(data.list[i].temp.day),
                weather: data.list[i].weather[0].main
            }
        }));
    }
    try {
        let responses = await Promise.all(promises);
        responses.forEach((e,i) => {
            console.log(e.data);
            data.list[i].recommendation = e.data;
        });
        ctx.body = data;
    } catch (error) {
        //若返回500，则不会产生熔断效果
        ctx.response.status = 503;
        ctx.body = {
            errMessage:"get recommendation info failed"
        }
    }
})

app.use(router.routes()).use(router.allowedMethods());
app.listen(3002);