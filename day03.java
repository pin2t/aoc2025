import static java.lang.IO.*;
import static java.lang.String.format;

void main() {
    var jolts = new long[2];
    var bank = readln();
    while (bank != null) {
        jolts[0] += maxJolts(bank, 2);
        jolts[1] += maxJolts(bank, 12);
        bank = readln();
    }
    println(jolts[0] + " " + jolts[1]);
}

long maxJolts(String bank, int digits) {
    long mx = 0;
    var pos = -1;
    for (int i = 1; i <= digits; i++) {
        pos++;
        for (int j = pos + 1; j < bank.length() - digits + i; j++) {
            if (bank.charAt(j) > bank.charAt(pos)) {
                pos = j;
            }
        }
        mx = mx * 10 + bank.charAt(pos) - '0';
    }
    return mx;
}
