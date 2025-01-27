import java.io.*;
import java.util.StringTokenizer;

/*
* 10진수 값과 N을 입력받는다.
* 10진수를 N으로 나누어 몫과 나머지를 구한다.
* 나머지를 기록한다.
* 몫이 0이 될 때까지 2단계와 3단계를 반복한다.
* 기록된 나머지를 거꾸로 조합하여 N진수 결과를 얻는다.
*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringTokenizer stk = new StringTokenizer(br.readLine());
        int n = Integer.parseInt(stk.nextToken());
        int b = Integer.parseInt(stk.nextToken());
        
        String nums = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        String result = "";
        int rem;

        while (n > 0) {
            rem = n%b;
            result = nums.charAt(rem) + result;
            n /= b;
        }
        System.out.println(result);
    }
}