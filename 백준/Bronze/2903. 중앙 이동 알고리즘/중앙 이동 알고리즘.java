import java.io.*;

	/*
	*	점의 개수
	*		초기 정사각형: 가로 2개, 세로 2개
	*			2*2 = 4개
	*		1번 변형: 가로 3개, 세로 3개
	*			3*3 = 9개
    *
	*		- 변형 될 때마다 점의 개수는 2의 N제곱 + 1 로 늘어남
	*		- 총 점의 개수는 가로 축에있는 점의 개수 * 세로 축에있는 점의 개수 이므로, 은 결국 2의 N제곱 +1 을 제곱한 값
	*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        float n = Float.parseFloat(br.readLine());

        int d = (int) Math.pow((Math.pow(2, n)+1),2);

        System.out.println(d);
    }
}