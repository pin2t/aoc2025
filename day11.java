import static java.util.Arrays.stream;
import static java.util.Collections.emptyList;
import static java.lang.IO.*;

Map<String, List<String>> conns = new HashMap<>();
Map<String, Long> cache = new HashMap<>();

long paths(String from, String to) {
    if (from.equals(to)) { return 1; }
    if (cache.containsKey(from + to)) { return cache.get(from + to); }
    var result = conns.getOrDefault(from, emptyList())
        .stream().mapToLong(it -> paths(it, to)).sum();
    cache.put(from + to, result);
    return result;
}

void main() {
    var line = readln();
    while (line != null) {
        var parts = line.split(" ");
        conns.put(parts[0].substring(0, parts[0].length() - 1), stream(parts).skip(1).toList());
        line = readln();
    }
    println(paths("you", "out") + " " + paths("svr", "fft") * paths("fft", "dac") * paths("dac", "out"));
}