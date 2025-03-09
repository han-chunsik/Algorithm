import java.io.IOException;
import java.time.*;
import java.time.format.DateTimeFormatter;

public class Main{
    public static void main(String[] args) throws IOException {
        LocalDate today = LocalDate.now();
        DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd");
        String formattedDate = today.format(formatter);
        System.out.println(formattedDate);
    }
}