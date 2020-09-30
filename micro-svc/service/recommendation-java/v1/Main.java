import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;


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
            String responseDressing = "";
            String responseSport = "";
            // String postString = IOUtils.toString(exchange.getRequestBody());
            // Map<String,String> postInfo = formData2Dic(postString);
            
            String queryString =  exchange.getRequestURI().getQuery();
            Map<String,String> postInfo = formData2Dic(queryString);

            System.out.printf("queryString: %s",queryString);

            if(postInfo.get("weather").equals("Rain")){
                responseSport = "sport: Do some indoor activities";
            }else if(postInfo.get("weather").equals("Clear")) {
                responseSport = "sport: Go climbing";
            }else if(postInfo.get("weather").equals("Snow")) {
                responseSport = "sport: Go skating";
            }else {
                responseSport = "sport: Whatever you do :)";
            }

            if(Integer.parseInt(postInfo.get("temp")) > 30 ){
                responseDressing = "dress: Put on your T-shirt Please";
            }else if(Integer.parseInt(postInfo.get("temp")) > 20 ) {
                responseDressing = "dress: Make sure you wear your jacket";
            }else {
                responseDressing = "dress: Please wear warm clothes";
            }
            exchange.sendResponseHeaders(200, 0);
            OutputStream os = exchange.getResponseBody();
            String response = responseDressing + ";" + responseSport;
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