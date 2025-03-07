import java.io.*;
import java.util.*;

public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(br.readLine());

        List<Integer> sizes = new ArrayList<>();
        
        StringTokenizer stk = new StringTokenizer(br.readLine());
        while (stk.hasMoreTokens()) {
            sizes.add(Integer.parseInt(stk.nextToken()));
        }

        stk = new StringTokenizer(br.readLine());
        int t = Integer.parseInt(stk.nextToken());
        int p = Integer.parseInt(stk.nextToken());

        int tBundleCount = 0;
        for (int size: sizes) {
            tBundleCount+=size/t;
            if(size%t > 0) {
                tBundleCount++;
            }
        }

        System.out.println(tBundleCount);
        System.out.println((n/p) + " " + (n%p));
    }
}