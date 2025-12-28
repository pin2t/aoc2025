import com.microsoft.z3.*;
import java.util.regex.Pattern;

import static java.lang.Integer.parseInt;
import static java.lang.System.in;
import static java.lang.System.out;
import static java.util.Arrays.*;
import static java.util.Comparator.comparingInt;

void main() throws IOException {
    var reWire = Pattern.compile("\\(.*?\\)");
    var reMachine = Pattern.compile("^\\[([.#]+)] ([()\\d, ]+) \\{([\\d,]+)}$");
    var reNum = Pattern.compile("\\d+");
    var presses = new int[]{0, 0};
    for (var line : new BufferedReader(new InputStreamReader(in)).readAllLines()) {
        var matcher = reMachine.matcher(line);
        if (!matcher.matches()) { continue; }
        var gr = matcher.group(1);
        var lights = new boolean[gr.length()];
        for (var i = 0; i < lights.length; i++) {
            lights[i] = gr.charAt(i) == '#';
        }
        var mtWires = reWire.matcher(matcher.group(2));
        var buttons = new ArrayList<List<Integer>>();
        for (var w : mtWires.results().toList()) {
            buttons.add(reNum.matcher(w.group()).results().map(r -> parseInt(r.group(0))).toList());
        }
        var jolts = reNum.matcher(matcher.group(3)).results().map(r -> parseInt(r.group(0))).toList();
        var initlights = new boolean[lights.length];
        fill(initlights, false);
        record State(boolean[] lights, int presses) {}
        var q = new PriorityQueue<State>(comparingInt(State::presses));
        q.offer(new State(initlights, 0));
        while (!q.isEmpty()) {
            var st = q.poll();
            if (Arrays.equals(lights, st.lights)) {
                presses[0] += st.presses;
                break;
            }
            for (var b : buttons) {
                var nxlight = copyOf(st.lights, st.lights.length);
                for (var i : b) nxlight[i] = !nxlight[i];
                q.offer(new State(nxlight, st.presses + 1));
            }
        }
        try (var ctx = new Context()) {
            var opt = ctx.mkOptimize();
            var vars = new ArrayList<IntExpr>();
            for (int i = 0; i < buttons.size(); i++) {
                var ni = ctx.mkIntConst("n" + i);
                opt.Assert(ctx.mkGe(ni, ctx.mkInt(0)));
                vars.add(ni);
            }
            for (int i = 0; i < jolts.size(); i++) {
                var ints = new ArrayList<IntExpr>();
                for (int j = 0; j < buttons.size(); j++) {
                    int finalI = i;
                    if (buttons.get(j).stream().anyMatch(out -> out.equals(finalI))) {
                        ints.add(vars.get(j));
                    }
                }
                BoolExpr be;
                if (ints.size() == 1) {
                    be = ctx.mkEq(ints.get(0), ctx.mkInt(jolts.get(i)));
                } else {
                    be = ctx.mkEq(ctx.mkAdd(ints.toArray(new IntExpr[0])), ctx.mkInt(jolts.get(i)));
                }
                opt.Assert(be);
            }
            Optimize.Handle<IntSort> res = opt.MkMinimize(ctx.mkAdd(vars.toArray(new IntExpr[0])));
            if (opt.Check() == Status.SATISFIABLE) {
                presses[1] += parseInt(res.getValue().toString());
            }
        }
    }
    out.println(presses[0] + " " + presses[1]);
}