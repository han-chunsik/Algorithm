import java.io.*;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
	/*
	* "OOXXOXXOOO"의 점수는 1+2+0+0+1+0+0+1+2+3 = 10점
	* OX퀴즈의 결과가 주어졌을 때, 점수를 구하는 프로그램
	*/
public class Main{
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine());

        List<Integer> scoreList = new ArrayList<>(Collections.nCopies(n, 0));
        
        for (int i=0; i<n; i++){
            int cnt = 0;
            String tc = br.readLine();
            for (char digit : tc.toCharArray()) {
                String t = Character.toString(digit);
                if (t.equals("O")){
                    cnt+=1;
                }else{
                    cnt=0;
                }
                scoreList.set(i, scoreList.get(i)+cnt);
            }
        }
        for (int score : scoreList){
            System.out.println(score);
        }
    }
}