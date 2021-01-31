let index = {};

index.isPermutation = (a, b) => {
    if(a.length != b.length) {
        return false;
    }

    sa = a.split("").sort().join("");
    sb = b.split("").sort().join("");
    for(var i = 0; i < a.length; i++) {
        if(sa[i] != sb[i]) {
            return false;
        }
    }
    return true;
}

module.exports = index;