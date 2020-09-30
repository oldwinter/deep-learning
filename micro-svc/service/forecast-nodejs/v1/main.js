const Koa = require('koa');
const Router = require('koa-router');
const app = new Koa();
const router = new Router();

const HANGZHOU = {
    "city": {
        "name": "Hangzhou",
    },
    "cnt": 3,
    "list": [{
            "dt": "today",
            "temp": {
                "day": 28.89,
                "min": 21.16,
                "max": 31.94,
            },
            "weather": [{
                "main": "Rain",
                "description": "light rain",
                "icon": "10d"
            }],
        },
        {
            "dt": "tomorrow",
            "temp": {
                "day": 30.5,
                "min": 29.73,
                "max": 40.5,
            },
            "weather": [{
                "main": "Clear",
                "description": "sky is clear",
                "icon": "01d"
            }],
        },
        {
            "dt": "the day after tomorrow",
            "temp": {
                "day": 32.72,
                "min": 27.03,
                "max": 39.72,
            },
            "weather": [{
                "main": "Rain",
                "description": "light rain",
                "icon": "10d"
            }],
        }
    ]
}
const BEIJING = {
    "city": {
        "name": "Beijing",
    },
    "cnt": 3,
    "list": [{
            "dt": "today",
            "temp": {
                "day": 32.94,
                "min": 24.48,
                "max": 34.63
            },
            "weather": [{
                "main": "Clear",
                "description": "sky is clear",
                "icon": "01d"
            }],
        },
        {
            "dt": "tomorrow",
            "temp": {
                "day": 28.67,
                "min": 24.37,
                "max": 35.62,
            },
            "weather": [{
                "id": 601,
                "main": "Snow",
                "description": "snow",
                "icon": "13d"
            }]
        },
        {
            "dt": "the day after tomorrow",
            "temp": {
                "day": 18.72,
                "min": 17.03,
                "max": 29.72
            },
            "weather": [{
                "main": "Rain",
                "description": "light rain",
                "icon": "10d"
            }]
        }
    ]
}
const SHANGHAI = {
    "city": {
        "name": "Shanghai",
    },
    "cnt": 3,
    "list": [{
            "dt": "today",
            "temp": {
                "day": 27.94,
                "min": 24.48,
                "max": 34.63,
            },
            "weather": [{
                "main": "Clear",
                "description": "sky is clear",
                "icon": "01d"
            }]
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
            "weather": [{
                "main": "Rain",
                "description": "light rain",
                "icon": "10d"
            }]
        },
        {
            "dt": "the day after tomorrow",
            "temp": {
                "day": 34.5,
                "min": 29.73,
                "max": 40.5
            },
            "weather": [{
                "main": "Clear",
                "description": "sky is clear",
                "icon": "01d"
            }]
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
    return headers
}

router.get('/status', async (ctx, next) => {
    ctx.body = 'ok';
})

router.get('/weather', async (ctx, next) => {
    let locate = ctx.query.locate;
    console.log(locate);
    forwardHeaders = getForwardHeaders(ctx.request)
    if (locate == "Hangzhou") {
        ctx.body = HANGZHOU;
    } else if (locate == "Beijing") {
        ctx.body = BEIJING;
    } else if (locate == "Shanghai") {
        ctx.body = SHANGHAI;
    } else {
        ctx.body = {
            "errorCode": "NO_RESULT"
        }
    }

})

app.use(router.routes()).use(router.allowedMethods());
app.listen(3002);