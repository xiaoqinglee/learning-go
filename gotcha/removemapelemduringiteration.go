package gotcha

/**

Is it safe to remove selected keys from map within a range loop?

This is safe!
The iteration order over maps is not specified and is not guaranteed
to be the same from one iteration to the next. If map entries that have
not yet been reached are removed during iteration, the corresponding
iteration values will not be produced. If map entries are created during
iteration, that entry may be produced during the iteration or may be
skipped. The choice may vary for each entry created and from one
iteration to the next. If the map is nil, the number of iterations is 0.

for key := range m {
    if key.expired() {
        delete(m, key)
    }
}

*/
