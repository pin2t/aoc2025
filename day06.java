import static java.lang.IO.println;
import static java.lang.Long.parseLong;
import static java.lang.System.in;

void main() throws IOException {
    var lines = new BufferedReader(new InputStreamReader(in)).readAllLines();
    var ops = lines.getLast();
    lines = lines.subList(0, lines.size() - 1);
    var re = Pattern.compile("\\d+");
    long res = 0, res2 = 0;
    for (int i = 0; i < ops.length(); i++) {
        var ii = i;
        var nums = lines.stream().map(l -> re.matcher(l.substring(ii)).results().findFirst().get().group()).mapToLong(Long::parseLong);
        switch (ops.charAt(i)) {
        case '+': res += nums.sum(); break;
        case '*': res += nums.reduce(1, (a, b) -> a * b); break;
        }
    }
    var nums = new ArrayList<Long>();
    for (int i = lines.getFirst().length() - 1; i >= 0; i--) {
        var s = new StringBuilder();
        for (var l : lines) { s.append(l.charAt(i)); }
        var ss = s.toString().trim();
        if (!ss.isBlank()) {
            nums.add(parseLong(ss));
        }
        switch (i < ops.length() ? ops.charAt(i) : ' ') {
        case '+':
            res2 += nums.stream().reduce(Long::sum).get();
            nums.clear();
            break;
        case '*':
            res2 += nums.stream().mapToLong(n -> n).reduce(1, (a, b) -> a * b);
            nums.clear();
            break;
        }
    }
    println(res + " " + res2);
}