import java.io.*;
import java.util.*;
	/*
	*	주어진 각 변 a, b, c로 직삼각형인지 확인하는 방법
	*	- a제곱 + b제곱 = c제곱(가장 긴 수) 일 경우 직삼각형
	*	입력 마지막이 0 0 0 일 경우 입력 종료
	*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        while (true) {
            StringTokenizer stk = new StringTokenizer(br.readLine());
            int a = Integer.parseInt(stk.nextToken());
            int b = Integer.parseInt(stk.nextToken());
            int c = Integer.parseInt(stk.nextToken());

            if (a+b+c == 0) {
                break;
            }

            double powA = Math.pow(a,2);
            double powB = Math.pow(b,2);
            double powC = Math.pow(c,2);

            String result = "wrong";
            if (a>b && a>c) {
                if (powA-powB-powC ==0){
                    result = "right";
                }
            }

            if (b>a && b>c) {
                if (powB-powA-powC ==0){
                    result = "right";
                }
            }

            if (c>a && c>b) {
                if (powC-powA-powB ==0){
                    result = "right";
                }
            }

            System.out.println(result);
        }
        
        
    }
}