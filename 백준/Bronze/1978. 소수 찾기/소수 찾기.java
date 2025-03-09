import java.io.*;
import java.util.*;

public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(br.readLine());

        List<Integer> nums = new ArrayList<>(Collections.nCopies(n, 0));
        
        StringTokenizer stk = new StringTokenizer(br.readLine());
        while (stk.hasMoreTokens()) {
            nums.add(Integer.parseInt(stk.nextToken()));
        }

        int cnt = 0;

        for(int num:nums){
            if (num > 1) {
                boolean isPrime = true;
                for(int i = 2; i <= Math.sqrt(num); i++){
                    if (num%i == 0){
                        isPrime = false;
                        break;
                    }
                }
                if (isPrime){
                    cnt++;
                }
            }
        }
        System.out.println(cnt);
    }
}