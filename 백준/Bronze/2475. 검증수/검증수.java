import java.io.*;
import java.util.StringTokenizer;
	/*
	*	5개의 숫자를 각각 제곱한 수의 합을 10으로 나눈 나머지
	*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringTokenizer stk = new StringTokenizer(br.readLine());
        double a = Float.parseFloat(stk.nextToken());
        double b = Float.parseFloat(stk.nextToken());
        double c = Float.parseFloat(stk.nextToken());
        double d = Float.parseFloat(stk.nextToken());
        double e = Float.parseFloat(stk.nextToken());

        double sumNum = Math.pow(a, 2) + Math.pow(b, 2) + Math.pow(c, 2) + Math.pow(d, 2) + Math.pow(e, 2);
        int result = (int) sumNum % 10;

        System.out.println(result);
    }
}