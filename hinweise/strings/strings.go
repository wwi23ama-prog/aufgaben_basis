package stringfuncs

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {
	result := 0
	for _, char := range s {
		if char == 'A' {
			result++
		}
	}
	return result
}

// Erwartet einen String s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {
	count := 0
	// Hinweis: Gehen Sie wie bei Contains vor. Aber statt des Early-Out erhöhen Sie
	// bei jedem Vorkommen von c einen Zähler, den Sie am Ende liefern.
	// TODO
	return count
}

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	result := ""
	for _, char := range s {
		result += string(char)
		result += string(char)
	}
	return result
}

// Erwartet einen String s und liefert s rückwärts zurück.
func Reverse(s string) string {
	result := ""
	// Hinweis: Schreiben Sie eine Schleife, die von rückwärts über s läuft
	// und in jedem Durchlauf ein Zeichen an result anhängt.
	// TODO
	return result
}

// Erwartet zwei Strings s1 und s2 und prüft, ob der eine der andere umgedreht ist.
func IsReverse(s1, s2 string) bool {
	// Hinweis: Verwenden Sie die Funktion Reverse() und vergleichen Sie die Ergebnisse.
	// TODO
	return false
}

// Erwartet einen String s und prüft, ob s ein Palindrom ist.
// Ein Palindrom ist ein String, der vorwärts und rückwärts gelesen gleich ist.
func IsPalindrome(s string) bool {
	// Hinweis: Verwenden Sie die Funktion IsReverse() oder Reverse().
	// TODO
	return false
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	// Hinweis: Laufen Sie in einer Schleife über s1 und prüfen Sie für jedes Zeichen,
	// ob es in s2 genauso oft vorkommt. Das können Sie mit der Funktion CountChar()
	// erledigen. Falls Sie ein Zeichen finden, das in s1 und s2 unterschiedlich oft
	// vorkommt, können Sie false liefern ("Early Out"), ansonsten liefern Sie am Ende true.
	// TODO
	return true
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	// Hinweis: Verwenden Sie die Funktion strings.ToLower() oder strings.ToUpper()
	// zusammen mit der Funktion IsAnagram().
	// TODO
	return false
}

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	// Hinweis: Prüfen Sie in einer Schleife für jedes Zeichen in s, ob es gleich c ist.
	// Falls ja, können Sie true liefern ("Early Out"), ansonsten liefern Sie am Ende false.
	// TODO
	return false
}

// Erwartet einen String s und einen Buchstaben c.
// Liefert die Position, an der c in s vorkommt.
// Liefert die Länge von s, falls c nicht in s vorkommt.
// Kommt c mehrfach vor, soll die erste Position geliefert werden.
func PositionOf(s string, c byte) int {
	// Hinweis: Gehen Sie wie bei Contains vor, aber führen Sie zusätzlich einen
	// Schleifenzähler mit, dessen Wert Sie im Fall einer Fundstelle von c liefern.
	// TODO
	return len(s)
}

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {
	// Hinweis: Verwenden Sie eine "Slice", um einen Teilstring von s zu gewinnen.
	// Beispiel: s[3:6] würde Ihnen die Zeichen aus s von 3 bis 5 liefern.
	// Vergleichen Sie solche Teilstrings in einer Schleife mit t.
	for i := 0; i < len(s)-len(t)+1; i++ {
		// TODO: Vergleichen Sie das s-Teilstück von i bis i+len(t) mit t.
	}
	return false
}

// Erwartet einen String und prüft, ob er korrekte Klammerpaare enthält.
// D.h. der Eingabestring enthält Klammern '(' und ')'.
// Die Funktion soll nun prüfen, ob der String für jede öffnende Klammer auch eine
// schließende Klammer enthält.
// Außerdem darf es keine schließenden Klammern geben, für die es nicht vorher eine
// passende öffnende Klammer gegeben hat.
// Die Funktion soll true liefern, falls der String korrekt geklammert ist.
func CheckParentheses(s string) bool {
	// Hinweis: Laufen Sie in einer Schleife über s.
	// Erhöhen Sie einen Klammer-Zähler um 1, wenn Sie eine öffnende Klammer sehen,
	// und verringern Sie ihn, wenn Sie eine schließende Klammer sehen.
	// Falls der Zähler jemals negativ wird oder am Ende nicht 0 ist,
	// liefern Sie false.
	counter := 0
	// TODO
	return counter == 0
}

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {
	// Hinweis: Schreiben Sie eine Schleife mit n-1 Durchläufen, die jedes Mal s und sep
	// an das Ergebnis anhängt. Am Ende müssen Sie noch ein weiteres Mal s anhängen.
	result := ""
	// TODO
	return result + s
}

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	// Hinweis: Es ist Code vorgegeben, der die minimale Länge von s1 und s2 bestimmt.
	// Schreiben Sie eine Schleife, die bis zu diesem min läuft und in jedem Durchlauf
	// jeweils ein Zeichen aus s1 und eines aus s2 ans Ergebnis anhängt.
	// Hängen Sie anschließend noch die beiden restlichen Strings an.
	result := ""

	min := len(s1)
	if len(s2) < min {
		min = len(s2)
	}
	// TODO

	return result
}
