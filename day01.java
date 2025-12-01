void main() {
    int n = 50, stops = 0, clicks = 0;
    var line = IO.readln();
    while (line != null) {
        for (int i = 0; i < Integer.parseInt(line.substring(1)); i++) {
           n = (n + (line.charAt(0) == 'R' ? 1 : 99)) % 100;
           if (n == 0) clicks++;
        }
        if (n == 0) stops++;
        line = IO.readln();
    }
    IO.println(stops + " " + clicks);
}
