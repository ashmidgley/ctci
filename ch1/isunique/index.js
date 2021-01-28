const index = {};

index.allUnique = (input) => {
    let set = new Set();
    for (var i = 0; i < input.length; i++) {
        const c = input.charAt(i);
        if (set.has(c)) {
            return false;
        }
        set.add(c);
    }
    return true;
}

module.exports = index;