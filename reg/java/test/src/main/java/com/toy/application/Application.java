package com.toy.application;


import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Application {

    static String imgReg = "<p[^<>]*>[^<>]*<img[^<>]*\\ssrc=(?:'|\")?(?:[^\\s'\"<>]+)//s\\.cimg\\.163\\.com/i/([^\\s'\"<>]+?)\\.\\d+x\\d+\\.auto.jpg(?:'|\")?[^<>]*>(.*?)</p>";
    static String tagReg = "<[^<>]*>";
    static String whiteSpaceReg = "[\\s　]*";

    static Pattern imgRegPattern = Pattern.compile(imgReg);
    static Pattern tagRegPattern = Pattern.compile(tagReg);
    static Pattern whiteSpaceRegPattern = Pattern.compile(whiteSpaceReg);

    static int threshold = 10000;
    public static void main(String[] args) {
        run();
        find();
    }

    static String format(String app, String method, long costTime) {
        return String.format(" %s %s cost %d millionSecond. \n", app, method, costTime);
    }

    static void find() {
        Long start = System.currentTimeMillis();
        for(int i=1; i<threshold; i++){
            find01();
        }
        Long end = System.currentTimeMillis();
        System.out.print(format("java", "find", end-start));
    }

    static void find01() {
        String data = read("src/main/resources/data/online_data.txt");
        Matcher matcher = imgRegPattern.matcher(data);
        while( matcher.find() ){

        }
    }

    static void run() {
        Long start = System.currentTimeMillis();
        for(int i=1; i<threshold; i++){
            replace();
        }
        Long end = System.currentTimeMillis();
        System.out.print(format("java", "run", end-start));
    }

    static void replace() {
        String data = read("src/main/resources/data/online_data.txt");
        String tmp = data.replaceAll(imgReg, "");
        tmp = tmp.replaceAll(tagReg, "");
        tmp = tmp.replaceAll(whiteSpaceReg, "");
    }

    static String read(String path) {
        if( path == null || path == "" ){
            return "";
        }
        File file = new File(path);
        if( file == null ){
            return "";
        }

        BufferedReader reader = null;
        try{
            reader = new BufferedReader(new FileReader(file));
            String tmp;
            StringBuilder sb = new StringBuilder();
            while( (tmp = reader.readLine()) != null ){
                sb.append(tmp);
            }
            return sb.toString();
        }catch (Exception e) {
            e.printStackTrace();
        }finally {
            if( reader != null ) {
                try{
                    reader.close();
                }catch ( Exception e ){

                }
            }
        }
        return "";
    }
}
