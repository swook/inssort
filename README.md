# inssort

inssort is a small library made for a very specific purpose. To sort a list
where changes are appended (using insertion sort), and also optionally track if
list entries have been moved.

A possible use-case of this library is:

1. A list of data is kept (eg. top ranked players).
2. New data is appended (good for insertion sort to be used).
3. Know if ordering has changed to reduce database transactions.
4. *[Optionally]* Check changes for only the top N entries.

---

inssort.Sort has the following type signature:

    func Sort(sort.Interface, ...int) bool

The variadic int argument can be provided in the following ways.

1. (None) Sort assuming all entries new.
2. (1 argument) Sort assuming mentioned index is new.
3. (2 arguments) Sort assuming arg1:arg2 as range for new entries.
4. (3 arguments) Same as previous, but limit the checking of movement to the first
   arg3 entries.
