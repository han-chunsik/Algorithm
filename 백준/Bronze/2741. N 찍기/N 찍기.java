import java.io.*;
	/*
	*	자연수 N이 주어졌을 때, 1부터 N까지 한 줄에 하나씩 출력
	*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine());

        for (int i=1; i<=n; i++){
            System.out.println(i);
        }
    }
}