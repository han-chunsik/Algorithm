import java.io.*;
import java.util.List;
import java.util.ArrayList;

/*
*		1. 테스트 케이스 개수 입력받기
*		2. 테스트 케이스 만큼 거스름 돈 입력받기
*		3. 큰 단위부터 순서대로 나누어 필요한 동전 개수 얻기
*		4. 0이 될때까지 반복하며 배열에 저장
*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int t = Integer.parseInt(br.readLine());

        List<Integer> coins = List.of(25, 10, 5, 1);
        int[][] coinCount = new int[t][coins.size()]; 

        for (int i = 0; i < t; i++) {
            int amount;
            amount = Integer.parseInt(br.readLine());
            for (int j = 0; j < coins.size(); j++) {
                int coin = coins.get(j);
                coinCount[i][j] = amount / coin;
			    amount %= coin;
            }
        }

        for (int a = 0; a < t; a++) {
            for (int b = 0; b < coins.size(); b++) {
                System.out.print(coinCount[a][b] + " ");
            }
            System.out.println();
        }
    }
}