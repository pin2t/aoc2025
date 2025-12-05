import static java.lang.IO.*;
import static java.lang.Long.parseLong;
import static java.util.Comparator.comparingLong;

import java.util.ArrayList;

void main() {
    record Range (long first, long last) { } 
    var ranges = new ArrayList<Range>();
    var line = readln();
    while (line != null) {
        if (line.isBlank()) break;
        ranges.add(new Range(parseLong(line.substring(0, line.indexOf('-'))), parseLong(line.substring(line.indexOf('-') + 1))));
        line = readln();
    }
    long n = 0L;
    line = readln();
    while (line != null) {
        var id = parseLong(line);
        if (ranges.stream().anyMatch(r -> r.first <= id && id <= r.last)) n++;
        line = readln();
    }
    ranges.sort(comparingLong(a -> a.first));
    var merged = new ArrayList<Range>();
    for (var r : ranges) {
        if (!merged.isEmpty() && r.first <= merged.getLast().last) {
            merged.set(merged.size() - 1, new Range(merged.getLast().first, Math.max(merged.getLast().last, r.last))); 
        } else {
            merged.add(r);
        }
    }
    println(n + " " + merged.stream().mapToLong(r -> r.last - r.first + 1).sum());
}


