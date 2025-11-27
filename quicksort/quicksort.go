package mergesort

// Partition teilt das Array arr in zwei Teile, basierend auf dem ersten Element
// (dem Pivot). Alle Elemente kleiner oder gleich dem Pivot werden nach links,
// alle größeren nach rechts einsortiert. Gibt die neue Position des Pivotelements zurück.
func Partition(arr []int) int {
	links := []int{}
	rechts := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[0] {
			rechts = append(rechts, arr[i])
		} else {
			links = append(links, arr[i])
		}
	}

	return len(links)
}

// QuickSort sortiert die übergebene Liste mittels des Quick-Sort-Algorithmus.
func QuickSort(arr []int) {

}
