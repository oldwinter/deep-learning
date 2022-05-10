package com.huaweicloud.asm.main;
import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;
import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;

import java.io.IOException;
import java.io.OutputStream;
import java.io.UnsupportedEncodingException;
import java.net.InetSocketAddress;
import java.net.InetAddress;
import java.net.URLDecoder;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Main {

    public static void main(String[] arg) throws Exception {
        HttpServer server = HttpServer.create(new InetSocketAddress(InetAddress.getByName("0.0.0.0"),3005), 0);
        server.createContext("/activity", new ActivityHandler());
        server.createContext("/status", new StatusHandler());
        server.start();
    }

    static class ActivityHandler implements HttpHandler{
        @Override
        public void handle(HttpExchange exchange) throws IOException {
            
            String queryString =  exchange.getRequestURI().getQuery();
            Map<String,String> postInfo = formData2Dic(queryString);

    		JSONObject json = new JSONObject();
    		
    		String weather = postInfo.get("weather");
    		int temp = Integer.parseInt(postInfo.get("temp"));
    		
            if(weather.equals("Rain")){
                json.put("sport", "Do some indoor activities");
            }else if(weather.equals("Clear")) {   
                json.put("sport", "Go climbing");
            }else if(weather.equals("Snow")) {
                json.put("sport", "Go skating");
            }else {
                json.put("sport", "Whatever you do :)");
            }

            if(temp > 30 ){
                json.put("dress", "Put on your T-shirt Please");
            }else if(temp > 20 ) {
            	json.put("dress", "Make sure you wear your jacket");
            }else {
                json.put("dress", "Please wear warm clothes");
            }

            exchange.sendResponseHeaders(200, 0);
            OutputStream os = exchange.getResponseBody();
            String response = json.toJSONString();
            System.out.println(response);
            os.write(response.getBytes());
            os.close();
        }
    }
    static class StatusHandler implements HttpHandler{
        @Override
        public void handle(HttpExchange exchange) throws IOException {
            String response = "ok";
            exchange.sendResponseHeaders(200, 0);
            OutputStream os = exchange.getResponseBody();
            os.write(response.getBytes());
            os.close();
        }
    }
    public static Map<String,String> formData2Dic(String formData ) {
        Map<String,String> result = new HashMap<>();
        if(formData== null || formData.trim().length() == 0) {
            return result;
        }
        final String[] items = formData.split("&");
        Arrays.stream(items).forEach(item ->{
            final String[] keyAndVal = item.split("=");
            if( keyAndVal.length == 2) {
                try{
                    final String key = URLDecoder.decode( keyAndVal[0],"utf8");
                    final String val = URLDecoder.decode( keyAndVal[1],"utf8");
                    result.put(key,val);
                }catch (UnsupportedEncodingException e) {}
            }
        });
        return result;
    }
}
