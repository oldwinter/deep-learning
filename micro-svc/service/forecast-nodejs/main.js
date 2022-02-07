const Koa = require('koa');
const Router = require('koa-router');
const axios = require('axios')
const app = new Koa();
const router = new Router();

var fs=require('fs');
var file="./mock-data.json";
var result=JSON.parse(fs.readFileSync(file));

const HANGZHOU=result.HANGZHOU 
const BEIJING=result.BEIJING
const SHANGHAI=result.SHANGHAI

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