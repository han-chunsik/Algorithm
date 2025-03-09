import java.io.*;
import java.util.*;

public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        
        StringTokenizer stk = new StringTokenizer(br.readLine());

        long n = Long.parseLong(stk.nextToken());
        long m = Long.parseLong(stk.nextToken());

        System.out.println(Math.abs(n-m));
    }
}