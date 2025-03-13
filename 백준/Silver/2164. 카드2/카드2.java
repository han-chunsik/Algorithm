import java.io.*;
import java.util.*;

public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine());

        ArrayDeque<Integer> cards = new ArrayDeque<>();

        for(int i=1;i<n+1;i++){
            cards.addLast(i);
        }

        while (cards.size() != 1) {
            cards.pollFirst();
            cards.addLast(cards.pollFirst());
        }

        System.out.println(cards.getFirst());
    }
}