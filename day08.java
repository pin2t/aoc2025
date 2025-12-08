import static java.lang.IO.println;
import static java.lang.Long.parseLong;
import static java.lang.System.in;
import static java.util.Comparator.comparingLong;
import static java.util.Comparator.reverseOrder;

void main() {
    record Pos(long x, long y, long z) {}
    record Pair(Pos a, Pos b) {
        long dist() { return (a.x - b.x) * (a.x - b.x) + (a.y - b.y) * (a.y - b.y) + (a.z - b.z) * (a.z - b.z); }
    }
    var boxes = new BufferedReader(new InputStreamReader(in)).lines().map(l -> {
        var p = l.split(",");
        return new Pos(parseLong(p[0]), parseLong(p[1]), parseLong(p[2]));
    }).toList();
    var pairs = new ArrayList<Pair>();
    for (int i = 0; i < boxes.size() - 1; i++) {
        for (int j = i + 1; j < boxes.size(); j++) {
            pairs.add(new Pair(boxes.get(i), boxes.get(j)));
        }
    }
    pairs.sort(comparingLong(Pair::dist));
    long n = 0, res1 = 0;
    var circuits = new ArrayList<HashSet<Pos>>();
    for (var p : pairs) {
        var ca = circuits.stream().filter(c -> c.contains(p.a)).findFirst();
        var cb = circuits.stream().filter(c -> c.contains(p.b)).findFirst();
        if (ca.isEmpty()) {
            if (cb.isEmpty()) {
                circuits.add(new HashSet<>(List.of(p.a, p.b)));
            } else {
                cb.get().add(p.a);
            }
        } else {
            if (cb.isEmpty()) {
                ca.get().add(p.b);
            } else if (ca.get() != cb.get()) {
                ca.get().addAll(cb.get());
                cb.get().clear();
            }
        }
        circuits.removeIf(HashSet::isEmpty);
        if (++n == 1000) {
            var sizes = new ArrayList<>(circuits.stream().map(HashSet::size).toList());
            sizes.sort(reverseOrder());
            res1 = (long) sizes.get(0) * sizes.get(1) * sizes.get(2);
        }
        if (circuits.getFirst().size() == boxes.size()) {
            println(res1 + " " + p.a.x * p.b.x);
            break;
        }
    }
}