import java.io.*;
import java.util.*;

public class Main{
    // 어떤 수가 "존재하는지" 빠르게 판단해야 하는 문제, value는 사용할 필요 없으므로 HashSet 사용
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        
        int n = Integer.parseInt(br.readLine());
        StringTokenizer stkN = new StringTokenizer(br.readLine());

        HashSet<Integer> a = new HashSet<>();

        for (int i = 0; i < n; i++) {
            a.add(Integer.parseInt(stkN.nextToken()));
        }

        int m = Integer.parseInt(br.readLine());
        StringTokenizer stkM = new StringTokenizer(br.readLine());

        for (int j = 0; j < m; j++) {
            if (a.contains(Integer.parseInt(stkM.nextToken()))) {
                System.out.println(1);
            }else {
                System.out.println(0);
            }


        }
    }
}