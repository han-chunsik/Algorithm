import java.io.*;
	/*
	*	A, B, C 가 주어질 때,
	*	첫째 줄에는 정수로 계산한 A + B - C
	*	둘째 줄에는 문자열로 계산한 A + B - C
    *
	*	예시)
	*	A = 3
	*	B = 4
	*	C = 5
    *
	*	첫째 줄: (3 + 4) - 5 = 2
	*	둘째 줄: "3" + "4" - "5" = "34" - "5" = 29
	*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int a = Integer.parseInt(br.readLine());
        int b = Integer.parseInt(br.readLine());
        int c = Integer.parseInt(br.readLine());

        int resultInt = (a + b) - c;
        int resultStr = Integer.valueOf(String.valueOf(a) + String.valueOf(b)) - c;

        System.out.println(resultInt);
        System.out.println(resultStr);
    }
}