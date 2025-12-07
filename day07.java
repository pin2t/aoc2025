import static java.lang.IO.println;
import static java.lang.IO.readln;
import static java.util.Arrays.stream;

void main() {
    var line = readln();
    var tl = new long[line.length()];
    for (int i = 0; i < line.length(); i++) {
        tl[i] = line.charAt(i) == 'S' ? 1L : 0L;
    }
    line = readln();
    var splits = 0;
    while (line != null) {
        for (int i = 0; i < line.length(); i++) {
            if (line.charAt(i) == '^' && tl[i] > 0L) {
                tl[i - 1] = tl[i - 1] + tl[i];
                tl[i + 1] = tl[i + 1] + tl[i];
                tl[i] = 0L;
                splits++;
            }
        }
        line = readln();
    }
    println(splits + " " + stream(tl).sum());
}