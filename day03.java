import static java.lang.String.format;

void main() {
    var jolts = new long[2];
    var bank = IO.readln();
    while (bank != null) {
        jolts[0] += maxJolts(bank, 2);
        jolts[1] += maxJolts(bank, 12);
        bank = IO.readln();
    }
    IO.println(format("%d %d", jolts[0], jolts[1]));
}

long maxJolts(String bank, int digits) {
    long m = 0;
    var dm = -1;
    for (int i = 1; i <= digits; i++) {
        dm++;
        for (int j = dm + 1; j < bank.length() - digits + i; j++) {
            if (bank.charAt(j) > bank.charAt(dm)) {
                dm = j;
            }
        }
        m = m * 10 + bank.charAt(dm) - '0';
    }
    return m;
}