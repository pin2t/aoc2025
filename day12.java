import static java.lang.IO.*;
import static java.lang.Integer.parseInt;

void main() {
    var line = readln();
    int res = 0, n = 0;
    var tiles = new HashMap<Integer, Integer>();
    var re = Pattern.compile("\\d+");
    while (line != null) {
        if (line.length() == 2 && line.charAt(1) == ':') {
            tiles.put(n++, (int) (readln() + readln() + readln()).chars().filter(c -> c == '#').count());
        } else if (line.contains("x")) {
            var nums = re.matcher(line).results().map(r -> parseInt(r.group())).toList();
            var area = nums.get(0) * nums.get(1);
            nums = nums.stream().skip(2).toList();
            var s = 0;
            for (var i = 0; i < nums.size(); i++) {
                s += nums.get(i) * tiles.get(i);
            }
            if (s <= area) {
                res++;
            }
        }
        line = readln();
    }
    println(res);
}