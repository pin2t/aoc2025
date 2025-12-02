import static java.lang.String.format;
import static java.util.Arrays.stream;

void main() {
    var line = IO.readln();
    var ids = new long[2];
    stream(line.split(",")).map(r -> r.split("-")).forEach(parts -> {
        for (long n = Long.parseLong(parts[0]); n <= Long.parseLong(parts[1]); n++) {
            var sn = Long.toString(n);
            if (sn.substring(0, sn.length() / 2).equals(sn.substring(sn.length() / 2))) {
                ids[0] += n;
            }
            for (int i = 1; i <= sn.length() / 2; i++) {
                if (sn.length() % i != 0) { continue; }
                var j = 0;
                while (j < sn.length() && sn.substring(j, j + i).equals(sn.substring(0, i))) { j += i; }
                if (j == sn.length()) {
                    ids[1] += n;
                    break;
                }
            }
        }
    });
    IO.println(format("%d %d", ids[0], ids[1]));
}