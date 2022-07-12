package gotcha

/**
不论那种语言(除了rust外), 在迭代一个容器时添加或删除元素都是非常危险的动作,
因为这些动作改变了容器的大小可能会让该容器重新被分配地址, 造成之前地址不可用.

目前阶段在迭代 golang 的 map 时添加或删除元素是安全的, 不代表在未来的实现中仍然是这样.
使用非公开的 api 不会获得质量承诺.

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
