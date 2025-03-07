import java.io.*;
import java.util.*;
import java.util.stream.Collectors;

public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(br.readLine());

        List<Integer> sizes = Arrays.stream(br.readLine().split(" "))
                                        .map(Integer::parseInt)
                                        .collect(Collectors.toList());

        StringTokenizer stk = new StringTokenizer(br.readLine());
        int t = Integer.parseInt(stk.nextToken());
        int p = Integer.parseInt(stk.nextToken());

        int tBundleCount = sizes.stream()
                                    .mapToInt(size -> (size / t) + (size % t != 0 ? 1 : 0))
                                    .sum();

        System.out.println(tBundleCount);
        System.out.println((n/p) + " " + (n%p));
    }
}