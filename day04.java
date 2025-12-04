import static java.lang.IO.*;

void main() {
    var y = 0;
    var rolls = new HashSet<Roll>();
    var line = readln();
    while (line != null) {
        for (int x = 0; x < line.length(); x++) {
            if (line.charAt(x) == '@') {
                rolls.add(new Roll(x, y));
            }
        }
        line = readln();
        y++;
    }
    var a = accessible(rolls);
    var acc = a.size();
    var total = 0;
    while (!a.isEmpty()) {
        total += a.size();
        rolls.removeAll(a);
        a = accessible(rolls);
    }
    println(acc + " " + total);
}

Set<Roll> accessible(Set<Roll> rolls) {
    var result = new HashSet<Roll>();
    for (var r : rolls) {
        var inter = r.adjacent();
        inter.retainAll(rolls);
        if (inter.size() < 4) {
            result.add(r);
        }
    }
    return result;
}

record Roll(int x, int y) {
    Set<Roll> adjacent() {
        return new HashSet<>(List.of(new Roll(x - 1, y), new Roll(x + 1, y), new Roll(x, y - 1), new Roll(x, y + 1), new Roll(x - 1, y - 1), new Roll(x + 1, y - 1), new Roll(x - 1, y + 1), new Roll(x + 1, y + 1)));
    }
}