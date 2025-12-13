import static java.lang.Integer.parseInt;
import static java.lang.Math.*;
import static java.lang.System.in;
import static java.lang.System.out;
import static java.util.Comparator.naturalOrder;

record Pos(int x, int y) {}

enum Dir {
    UP(0, -1), DOWN(0, 1), LEFT(-1, 0), RIGHT(1, 0);

    final int dx, dy;

    Dir(int dx, int dy) {
        this.dx = dx;
        this.dy = dy;
    }
}

long area(Pos a, Pos b) {
    return (long) (abs(a.x - b.x) + 1) * (abs(a.y - b.y) + 1);
}

void main() {
    var reds = new BufferedReader(new InputStreamReader(in)).lines()
        .map(l -> l.split(",")).map(p -> new Pos(parseInt(p[0]), parseInt(p[1]))).toList();
    var maxarea = 0L;
    for (int i = 0; i < reds.size() - 1; i++) {
        for (int j = i + 1; j < reds.size(); j++) {
            maxarea = max(maxarea, area(reds.get(i), reds.get(j)));
        }
    }
    out.print(maxarea);
    var xs = new ArrayList<>(reds.stream().map(Pos::x).collect(Collectors.toSet()));
    var ys = new ArrayList<>(reds.stream().map(Pos::y).collect(Collectors.toSet()));
    xs.sort(naturalOrder());
    ys.sort(naturalOrder());
    var virtual = reds.stream().map(p -> new Pos(2 + xs.indexOf(p.x) * 2, 2 + ys.indexOf(p.y) * 2)).toList();
    var border = new HashSet<Pos>();
    for (int i = 0; i < virtual.size(); i++) {
        var a = virtual.get(i);
        var b = virtual.get((i + 1) % virtual.size());
        if (a.x == b.x) {
            for (int y = min(a.y, b.y); y <= max(a.y, b.y); y++) {
                border.add(new Pos(a.x, y));
            }
        } else {
            for (int x = min(a.x, b.x); x <= max(a.x, b.x); x++) {
                border.add(new Pos(x, a.y));
            }
        }
    }
    var maxx = virtual.stream().mapToInt(p -> p.x).max().getAsInt();
    var maxy = virtual.stream().mapToInt(p -> p.y).max().getAsInt();
    var outside = new HashSet<Pos>();
    var q = new LinkedList<Pos>();
    q.offer(new Pos(0, 0));
    while (!q.isEmpty()) {
        var p = q.poll();
        if (!outside.add(p)) { continue; }
        for (var dir : Dir.values()) {
            var next = new Pos(p.x + dir.dx, p.y + dir.dy);
            if (next.x >= 0 && next.x < maxx + 2 && next.y >= 0 && next.y < maxy + 2 && !border.contains(next)) {
                q.offer(next);
            }
        }
    }
    int ia = 0, ib = 1;
    for (int i = 0; i < virtual.size() - 1; i++) {
        for (int j = i + 1; j < virtual.size(); j++) {
            var a = virtual.get(i);
            var b = virtual.get(j);
            var inside = true;
            for (int x = min(a.x, b.x); x <= max(a.x, b.x) && inside; x++) {
                inside = !outside.contains(new Pos(x, min(a.y, b.y))) &&
                         !outside.contains(new Pos(x, max(a.y, b.y)));
            }
            if (!inside) { continue; }
            for (int y = min(a.y, b.y); y <= max(a.y, b.y) && inside; y++) {
                inside = !outside.contains(new Pos(min(a.x, b.x), y)) &&
                         !outside.contains(new Pos(max(a.x, b.x), y));
            }
            if (!inside) { continue; }
            if (area(a, b) > area(virtual.get(ia), virtual.get(ib))) {
                ia = i;
                ib = j;
            }
        }
    }
    out.println(" " + area(reds.get(ia), reds.get(ib)));
}